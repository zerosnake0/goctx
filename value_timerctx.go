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

func init() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	cancel()

	typ := reflect.TypeOf(ctx)
	if typ.Kind() != reflect.Ptr {
		return
	}
	typ = typ.Elem()
	if typ.String() != "context.timerCtx" {
		return
	}
	if !checkCancelContextField(typ, "cancelCtx", 0) {
		return
	}

	RegisterValueFunc(ctx, timerCtxFunc)
	timerCtxRtype = (*gounsafe.Iface)(unsafe.Pointer(&ctx)).Itab.RType
}
