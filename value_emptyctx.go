package goctx

import (
	"context"
	"reflect"
	"unsafe"

	"github.com/zerosnake0/gounsafe"
)

var (
	emptyCtxRtype gounsafe.RType
)

func emptyCtxFunc(ctx context.Context, key interface{}) (value interface{}, parent context.Context) {
	return nil, nil
}

func checkEmptyCtx(o interface{}) bool {
	if o == nil {
		return false
	}
	typ := reflect.TypeOf(o)
	if typ.Kind() != reflect.Ptr {
		return false
	}
	typ = typ.Elem()
	return typ.String() == "context.emptyCtx"
}

func init() {
	ctx := context.TODO()
	if checkEmptyCtx(ctx) {
		RegisterValueFunc(ctx, emptyCtxFunc)
		emptyCtxRtype = (*gounsafe.Iface)(unsafe.Pointer(&ctx)).Itab.RType
	}
}
