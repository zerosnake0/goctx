package goctx

import (
	"context"
	"reflect"
)

func checkField(typ reflect.Type, name string, offset uintptr, exp reflect.Type) bool {
	field, ok := typ.FieldByName(name)
	if !ok {
		return false
	}
	if field.Offset != offset {
		return false
	}
	if field.Type != exp {
		return false
	}
	return true
}

func checkContextField(typ reflect.Type, name string, offset uintptr) bool {
	return checkField(typ, name, offset, reflect.TypeOf((*context.Context)(nil)).Elem())
}

func checkInterfaceField(typ reflect.Type, name string, offset uintptr) bool {
	return checkField(typ, name, offset, reflect.TypeOf((*interface{})(nil)).Elem())
}

func checkCancelContextField(typ reflect.Type, name string, offset uintptr) bool {
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()
	return checkField(typ, name, offset, reflect.TypeOf(ctx).Elem())
}
