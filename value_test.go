package goctx

import (
	"context"
	"fmt"
	"testing"
	"time"
	"unsafe"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/zerosnake0/gounsafe"
)

func prepare(count int) (_, _ interface{}, _ context.Context) {
	key := uuid.New().String()
	value := uuid.New().String()
	ctx := context.WithValue(context.TODO(), key, value)

	for i := 0; i < count; i++ {
		var cancel func()
		switch i % 3 {
		case 0:
			ctx = context.WithValue(ctx, " "+uuid.New().String(), " "+uuid.New().String())
		case 1:
			ctx, cancel = context.WithTimeout(ctx, time.Second)
			cancel()
		default:
			ctx, cancel = context.WithCancel(ctx)
			cancel()
		}
	}
	return key, value, ctx
}

type testCtx struct {
	context.Context
}

func (t testCtx) Value(key interface{}) interface{} {
	if key == t {
		return t
	}
	return nil
}

var _ context.Context = testCtx{}

func TestRegisterValueFunc(t *testing.T) {
	ctx := testCtx{}
	require.Equal(t, ctx, Value(ctx, ctx))
	require.Equal(t, nil, Value(ctx, 0))

	RegisterValueFunc(testCtx{}, func(ctx context.Context, key interface{}) (value interface{}, parent context.Context) {
		if key == ctx {
			return ctx, nil
		}
		return nil, nil
	})
	require.Equal(t, ctx, Value(ctx, ctx))
	require.Equal(t, nil, Value(ctx, 0))
}

func TestValue(t *testing.T) {
	t.Run("nested", func(t *testing.T) {
		key, value, ctx := prepare(20)
		require.Equal(t, value, ctx.Value(key))
		require.Equal(t, value, Value(ctx, key))

		anotherKey := -1
		require.NotEqual(t, anotherKey, key)
		require.Equal(t, nil, ctx.Value(anotherKey))
		require.Equal(t, nil, Value(ctx, anotherKey))
	})
	t.Run("custom context", func(t *testing.T) {
		ctx := testCtx{}
		v := Value(ctx, ctx)
		require.Equal(t, ctx, v)
	})
}

// pure map implementation
func value2(ctx context.Context, key interface{}) (v interface{}) {
	_, m := loadMap()
	for {
		if ctx == nil {
			return
		}
		ifc := (*gounsafe.Iface)(unsafe.Pointer(&ctx))
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

func BenchmarkValue(b *testing.B) {
	const (
		std     = "std"
		goctx   = "goctx"
		pureMap = "pure map"
	)

	for _, count := range []int{
		5, 10, 20,
	} {
		b.Run(fmt.Sprint(count), func(b *testing.B) {
			key, value, ctx := prepare(count)
			require.Equal(b, value, ctx.Value(key))
			require.Equal(b, value, Value(ctx, key))

			require.Equal(b, nil, ctx.Value(-1))
			require.Equal(b, nil, Value(ctx, -1))

			b.Run("non parallel", func(b *testing.B) {
				b.Run(std, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						ctx.Value(key)
					}
				})
				b.Run(goctx, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						Value(ctx, key)
					}
				})
				b.Run(pureMap, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						value2(ctx, key)
					}
				})
			})
			b.Run("parallel", func(b *testing.B) {
				b.Run(std, func(b *testing.B) {
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							ctx.Value(key)
						}
					})
				})
				b.Run(goctx, func(b *testing.B) {
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							Value(ctx, key)
						}
					})
				})
				b.Run(pureMap, func(b *testing.B) {
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							value2(ctx, key)
						}
					})
				})
			})
		})
	}
}
