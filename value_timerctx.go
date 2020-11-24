package goctx

import (
	"context"
	"reflect"
	"time"
	"unsafe"

	"github.com/zerosnake0/gounsafe"
)

var (
	timerCtxRtype gounsafe.RType
)

type timerCtx struct {
	cancelCtx
}

func timerCtxFunc(ctx context.Context, key interface{}) (value interface{}, parent context.Context) {
	return nil, (*timerCtx)((*gounsafe.Iface)(unsafe.Pointer(&ctx)).Data).Context
}

func checkTimerCtx(o interface{}) bool {
	if o == nil {
		return false
	}
	typ := reflect.TypeOf(o)
	if typ.Kind() != reflect.Ptr {
		return false
	}
	typ = typ.Elem()
	if typ.String() != "context.timerCtx" {
		return false
	}
	return checkCancelContextField(typ, "cancelCtx", 0)
}

func init() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	cancel()
	if checkTimerCtx(ctx) {
		RegisterValueFunc(ctx, timerCtxFunc)
		timerCtxRtype = (*gounsafe.Iface)(unsafe.Pointer(&ctx)).Itab.RType
	}
}
