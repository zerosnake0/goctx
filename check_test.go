package goctx

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkField(t *testing.T) {
	typ := reflect.TypeOf(struct {
		A int
	}{})
	fieldType := reflect.TypeOf(1)
	t.Run("name fail", func(t *testing.T) {
		require.False(t, checkField(typ, "B", 0, fieldType))
	})
	t.Run("offset fail", func(t *testing.T) {
		require.False(t, checkField(typ, "A", 1, fieldType))
	})
	t.Run("field type fail", func(t *testing.T) {
		require.False(t, checkField(typ, "A", 0, reflect.TypeOf("")))
	})
	t.Run("succ", func(t *testing.T) {
		require.True(t, checkField(typ, "A", 0, fieldType))
	})
}
