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

func init() {
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	typ := reflect.TypeOf(ctx)
	if typ.Kind() != reflect.Ptr {
		return
	}
	typ = typ.Elem()
	if typ.String() != "context.cancelCtx" {
		return
	}
	if !checkContextField(typ, "Context", 0) {
		return
	}
	RegisterValueFunc(ctx, cancelCtxFunc)
	cancelCtxRtype = (*gounsafe.Iface)(unsafe.Pointer(&ctx)).Itab.RType

}
