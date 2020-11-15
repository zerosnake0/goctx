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

func init() {
	ctx := context.TODO()

	typ := reflect.TypeOf(ctx)
	if typ.Kind() != reflect.Ptr {
		return
	}
	typ = typ.Elem()
	if typ.String() != "context.emptyCtx" {
		return
	}

	RegisterValueFunc(ctx, emptyCtxFunc)
	emptyCtxRtype = (*gounsafe.Iface)(unsafe.Pointer(&ctx)).Itab.RType
}
