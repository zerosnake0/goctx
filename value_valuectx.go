package goctx

import (
	"context"
	"reflect"
	"unsafe"

	"github.com/zerosnake0/gounsafe"
)

var (
	valueCtxRtype gounsafe.RType
)

type valueCtx struct {
	context.Context
	key   interface{}
	value interface{}
}

func valueCtxFunc(ctx context.Context, key interface{}) (value interface{}, parent context.Context) {
	ptr := (*valueCtx)((*gounsafe.Iface)(unsafe.Pointer(&ctx)).Data)
	if key == ptr.key {
		return ptr.value, ptr.Context
	}
	return nil, ptr.Context
}

func init() {
	ctx := context.WithValue(context.TODO(), valueCtx{}, "")

	typ := reflect.TypeOf(ctx)
	if typ.Kind() != reflect.Ptr {
		return
	}
	typ = typ.Elem()
	if typ.String() != "context.valueCtx" {
		return
	}
	if !checkContextField(typ, "Context", 0) {
		return
	}
	ourType := reflect.TypeOf(valueCtx{})
	if !checkInterfaceField(typ, "key", ourType.Field(1).Offset) {
		return
	}
	if !checkInterfaceField(typ, "val", ourType.Field(2).Offset) {
		return
	}

	RegisterValueFunc(ctx, valueCtxFunc)
	valueCtxRtype = (*gounsafe.Iface)(unsafe.Pointer(&ctx)).Itab.RType
}
