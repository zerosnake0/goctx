package goctx

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_cancelCtxFunc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	cancel()

	key := uuid.New().String()
	gotV, parent := cancelCtxFunc(ctx, key)
	require.Equal(t, nil, gotV)
	require.Equal(t, context.TODO(), parent)
}

func Test_checkCancelCtx(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		require.False(t, checkCancelCtx(nil))
	})
	t.Run("wrong kind", func(t *testing.T) {
		require.False(t, checkCancelCtx(0))
	})
	t.Run("wrong type name", func(t *testing.T) {
		type cancelCtx struct{}
		require.False(t, checkCancelCtx(&cancelCtx{}))
	})
}
