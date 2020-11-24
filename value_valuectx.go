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

func checkValueCtx(o interface{}) bool {
	if o == nil {
		return false
	}
	typ := reflect.TypeOf(o)
	if typ.Kind() != reflect.Ptr {
		return false
	}
	typ = typ.Elem()
	if typ.String() != "context.valueCtx" {
		return false
	}
	ourType := reflect.TypeOf(valueCtx{})
	return checkInterfaceField(typ, "key", ourType.Field(1).Offset) &&
		checkInterfaceField(typ, "val", ourType.Field(2).Offset) &&
		checkContextField(typ, "Context", 0)
}

func init() {
	ctx := context.WithValue(context.TODO(), valueCtx{}, "")
	if checkValueCtx(ctx) {
		RegisterValueFunc(ctx, valueCtxFunc)
		valueCtxRtype = (*gounsafe.Iface)(unsafe.Pointer(&ctx)).Itab.RType
	}
}
