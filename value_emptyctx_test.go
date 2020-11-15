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
