package gosql

import (
	"encoding/json"
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

	t.Run("set_value_typed", func(t *testing.T) {
		var js JSON[map[string]string]
		assert.NoError(t, js.SetValue(map[string]string{"k": "v"}))
		assert.Equal(t, "v", js.Data["k"])

		val := map[string]string{"k": "w"}
		assert.NoError(t, js.SetValue(&val))
		assert.Equal(t, "w", js.Data["k"])

		assert.NoError(t, js.SetValue(`{"k":"x"}`))
		assert.Equal(t, "x", js.Data["k"])

		assert.NoError(t, js.SetValue([]byte(`{"k":"y"}`)))
		assert.Equal(t, "y", js.Data["k"])

		assert.Equal(t, ErrNullValueNotAllowed, js.SetValue(nil))
	})

	t.Run("string_and_value", func(t *testing.T) {
		js := MustJSON[map[string]int](map[string]int{"n": 1})
		assert.Equal(t, `{"n":1}`, js.String())

		v, err := js.Value()
		assert.NoError(t, err)
		assert.Equal(t, `{"n":1}`, v)

		var nilJS *JSON[int]
		assert.Equal(t, "{}", nilJS.String())
	})

	t.Run("scan", func(t *testing.T) {
		var js JSON[map[string]int]
		assert.NoError(t, js.Scan(`{"n":2}`))
		assert.Equal(t, 2, js.Data["n"])

		assert.NoError(t, js.Scan([]byte(`{"n":3}`)))
		assert.Equal(t, 3, js.Data["n"])

		assert.NoError(t, js.Scan(json.RawMessage(`{"n":4}`)))
		assert.Equal(t, 4, js.Data["n"])

		assert.Equal(t, ErrNullValueNotAllowed, js.Scan(nil))
		assert.Equal(t, ErrInvalidScan, js.Scan(1))
	})

	t.Run("decode_value", func(t *testing.T) {
		var js JSON[map[string]int]
		assert.NoError(t, js.DecodeValue(`{"n":5}`))
		assert.Equal(t, 5, js.Data["n"])

		assert.NoError(t, js.DecodeValue([]byte(`{"n":6}`)))
		assert.Equal(t, 6, js.Data["n"])

		assert.NoError(t, js.DecodeValue(json.RawMessage(`{"n":7}`)))
		assert.Equal(t, 7, js.Data["n"])

		assert.NoError(t, js.DecodeValue(map[string]int{"n": 8}))
		assert.Equal(t, 8, js.Data["n"])
	})

	t.Run("json_roundtrip", func(t *testing.T) {
		js := MustJSON[map[string]string](map[string]string{"a": "b"})
		data, err := js.MarshalJSON()
		assert.NoError(t, err)

		var got JSON[map[string]string]
		assert.NoError(t, got.UnmarshalJSON(data))
		assert.Equal(t, "b", got.Data["a"])
	})
}
