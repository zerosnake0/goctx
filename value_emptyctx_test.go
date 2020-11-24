package goctx

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_emptyCtxFunc(t *testing.T) {
	ctx := context.TODO()

	key := uuid.New().String()
	gotV, parent := emptyCtxFunc(ctx, key)
	require.Equal(t, nil, gotV)
	require.Equal(t, nil, parent)
}

func Test_checkEmptyCtx(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		require.False(t, checkEmptyCtx(nil))
	})
	t.Run("wrong kind", func(t *testing.T) {
		require.False(t, checkEmptyCtx(0))
	})
	t.Run("wrong type name", func(t *testing.T) {
		type emptyCtx struct{}
		require.False(t, checkEmptyCtx(&emptyCtx{}))
	})
}