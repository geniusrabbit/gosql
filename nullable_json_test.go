package gosql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullableJSON(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		js, err := NewNullableJSON[map[string]string](map[string]any{
			"name":   "Yoda",
			"number": "1",
		})
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, "Yoda", (*js.Data)["name"])
		assert.Equal(t, "1", (*js.Data)["number"])

		assert.Error(t, js.SetValue(map[string]any{"invalid": 1}))

		assert.Panics(t, func() {
			MustNullableJSON[map[string]string]([]int{1, 2, 3})
		})
	})

	t.Run("json_empty_parse", func(t *testing.T) {
		var js NullableJSON[map[string]string]
		assert.NoError(t, js.Scan([]byte{}), "scan empty byte")
		assert.NoError(t, js.Scan([]byte(nil)), "scan nil byte")
		assert.NoError(t, js.Scan([]byte(" ")), "scan empty string")
	})
}
