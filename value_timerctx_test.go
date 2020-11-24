package goctx

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_timerCtxFunc(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	cancel()

	key := uuid.New().String()
	gotV, parent := timerCtxFunc(ctx, key)
	require.Equal(t, nil, gotV)
	require.Equal(t, context.TODO(), parent)
}

func Test_checkTimerCtx(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		require.False(t, checkTimerCtx(nil))
	})
	t.Run("wrong kind", func(t *testing.T) {
		require.False(t, checkTimerCtx(0))
	})
	t.Run("wrong type name", func(t *testing.T) {
		type timerCtx struct{}
		require.False(t, checkTimerCtx(&timerCtx{}))
	})
}
