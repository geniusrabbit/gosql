package gosql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		js, err := NewJSON[map[string]string](map[string]any{
			"name":   "Yoda",
			"number": "1",
		})
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, "Yoda", js.Data["name"])
		assert.Equal(t, "1", js.Data["number"])

		assert.Error(t, js.SetValue(map[string]any{"invalid": 1}))

		assert.Panics(t, func() {
			MustJSON[map[string]string](nil)
		})
	})

	t.Run("json_empty_parse", func(t *testing.T) {
		var js JSON[map[string]string]
		assert.NoError(t, js.Scan([]byte{}), "scan empty byte")
		assert.NoError(t, js.Scan([]byte(" ")), "scan empty string")
	})
}
