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
