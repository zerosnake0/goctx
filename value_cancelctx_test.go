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
