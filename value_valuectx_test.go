package goctx

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_valueCtxFunc(t *testing.T) {
	key, value, ctx := prepare(0)

	gotV, parent := valueCtxFunc(ctx, key)
	require.Equal(t, value, gotV)
	require.Equal(t, context.TODO(), parent)

	key2 := uuid.New().String()
	require.NotEqual(t, key, key2)

	gotV, parent = valueCtxFunc(ctx, key2)
	require.Equal(t, nil, gotV)
	require.Equal(t, context.TODO(), parent)
}

func Test_checkValueCtx(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		require.False(t, checkValueCtx(nil))
	})
	t.Run("wrong kind", func(t *testing.T) {
		require.False(t, checkValueCtx(0))
	})
	t.Run("wrong type name", func(t *testing.T) {
		type valueCtx struct{}
		require.False(t, checkValueCtx(&valueCtx{}))
	})
}
