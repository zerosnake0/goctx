package goctx

import (
	"context"
	"reflect"
	"unsafe"

	"github.com/zerosnake0/gounsafe"
)

var (
	cancelCtxRtype gounsafe.RType
)

type cancelCtx struct {
	context.Context
}

func cancelCtxFunc(ctx context.Context, key interface{}) (value interface{}, parent context.Context) {
	return nil, (*cancelCtx)((*gounsafe.Iface)(unsafe.Pointer(&ctx)).Data).Context
}

func checkCancelCtx(o interface{}) bool {
	if o == nil {
		return false
	}
	typ := reflect.TypeOf(o)
	if typ.Kind() != reflect.Ptr {
		return false
	}
	typ = typ.Elem()
	if typ.String() != "context.cancelCtx" {
		return false
	}
	return checkContextField(typ, "Context", 0)
}

func init() {
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()
	if checkCancelCtx(ctx) {
		RegisterValueFunc(ctx, cancelCtxFunc)
		cancelCtxRtype = (*gounsafe.Iface)(unsafe.Pointer(&ctx)).Itab.RType
	}
}
