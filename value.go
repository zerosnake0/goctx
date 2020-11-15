package goctx

import (
	"context"
	"sync/atomic"
	"unsafe"

	"github.com/zerosnake0/gounsafe"
)

// ValueFunc takes a context and a key, returns the value and the parent of the context (if exists)
type ValueFunc func(ctx context.Context, key interface{}) (value interface{}, parent context.Context)

var (
	valueFuncMap = map[uintptr]ValueFunc{}
)

func loadMap() (unsafe.Pointer, map[uintptr]ValueFunc) {
	ptr := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&valueFuncMap)))
	m := *(*map[uintptr]ValueFunc)(unsafe.Pointer(&ptr))
	return ptr, m
}

// RegisterValueFunc registers a value function
func RegisterValueFunc(ctx context.Context, f ValueFunc) {
	for {
		ptr, m := loadMap()
		newM := map[uintptr]ValueFunc{}
		for k, v := range m {
			newM[k] = v
		}
		newM[(*gounsafe.Iface)(unsafe.Pointer(&ctx)).Itab.RType] = f
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&valueFuncMap)),
			ptr, *(*unsafe.Pointer)(unsafe.Pointer(&newM))) {
			return
		}
	}
}

// Value can replace the `context.Context.Value` call
func Value(ctx context.Context, key interface{}) (v interface{}) {
	_, m := loadMap()
	for {
		if ctx == nil {
			return
		}
		ifc := (*gounsafe.Iface)(unsafe.Pointer(&ctx))
		switch ifc.Itab.RType {
		// the rtype value will be zero if initialization failed
		case cancelCtxRtype:
			ctx = (*cancelCtx)((*gounsafe.Iface)(unsafe.Pointer(&ctx)).Data).Context
		case emptyCtxRtype:
			return nil
		case timerCtxRtype:
			ctx = (*timerCtx)((*gounsafe.Iface)(unsafe.Pointer(&ctx)).Data).Context
		case valueCtxRtype:
			ptr := (*valueCtx)((*gounsafe.Iface)(unsafe.Pointer(&ctx)).Data)
			if key == ptr.key {
				return ptr.value
			}
			ctx = ptr.Context
		default:
			f := m[ifc.Itab.RType]
			if f == nil {
				return ctx.Value(key)
			}
			v, ctx = f(ctx, key)
			if v != nil {
				return
			}
		}
	}
}
