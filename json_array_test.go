package gosql

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONArray(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		js, err := NewJSONArray[int]([]int{1, 2, 3})
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, 1, js[0])
		assert.Equal(t, 2, js[1])
		assert.Equal(t, 3, js[2])

		assert.Error(t, js.SetValue([]string{"invalid"}))

		assert.Panics(t, func() {
			MustJSONArray[int](map[string]int{"invalid": 1})
		})
	})

	t.Run("json_empty_parse", func(t *testing.T) {
		var js JSONArray[int]
		assert.NoError(t, js.Scan([]byte{}), "scan empty byte")
		assert.NoError(t, js.Scan([]byte(nil)), "scan nil byte")
		assert.NoError(t, js.Scan([]byte(" ")), "scan empty string")
	})

	t.Run("json_encode", func(t *testing.T) {
		js, err := NewJSONArray[int]([]int{1, 2, 3})
		if !assert.NoError(t, err) {
			return
		}
		data, err := js.MarshalJSON()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, "[1,2,3]", string(data))
		assert.Equal(t, "[1,2,3]", js.String())
		v, _ := js.Value()
		assert.Equal(t, "[1,2,3]", v)
	})

	t.Run("json_encode_empty", func(t *testing.T) {
		var js JSONArray[int]
		data, err := js.MarshalJSON()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, "[]", string(data))
		assert.Equal(t, "[]", js.String())
		v, _ := js.Value()
		assert.Equal(t, "[]", v)
	})

	t.Run("set_value", func(t *testing.T) {
		var js JSONArray[int]
		assert.NoError(t, js.SetValue([]int{4, 5}))
		assert.Equal(t, JSONArray[int]{4, 5}, js)

		src := []int{6}
		assert.NoError(t, js.SetValue(&src))
		assert.Equal(t, JSONArray[int]{4, 5, 6}, js)

		assert.Equal(t, ErrNullValueNotAllowed, js.SetValue(nil))
		assert.NoError(t, js.SetValue("[7,8]"))
		assert.Equal(t, JSONArray[int]{7, 8}, js)

		assert.NoError(t, js.SetValue([]byte("[9]")))
		assert.Equal(t, JSONArray[int]{9}, js)

		assert.NoError(t, js.SetValue(json.RawMessage("[10]")))
		assert.Equal(t, JSONArray[int]{10}, js)

		assert.Equal(t, ErrInvalidSetValue, js.SetValue(1))
	})

	t.Run("scan", func(t *testing.T) {
		var js JSONArray[int]
		assert.NoError(t, js.Scan("[1,2]"))
		assert.Equal(t, JSONArray[int]{1, 2}, js)

		assert.NoError(t, js.Scan(json.RawMessage("[3]")))
		assert.Equal(t, JSONArray[int]{3}, js)

		assert.Equal(t, ErrNullValueNotAllowed, js.Scan(nil))
		assert.Equal(t, ErrInvalidScan, js.Scan(1))
	})

	t.Run("decode_value", func(t *testing.T) {
		var js JSONArray[int]
		assert.NoError(t, js.DecodeValue([]int{1, 2}))
		assert.Equal(t, JSONArray[int]{1, 2}, js)

		src := []int{3, 4}
		assert.NoError(t, js.DecodeValue(&src))
		assert.Equal(t, JSONArray[int]{3, 4}, js)

		assert.NoError(t, js.DecodeValue(JSONArray[int]{5}))
		assert.Equal(t, JSONArray[int]{5}, js)

		assert.NoError(t, js.DecodeValue("[6,7]"))
		assert.Equal(t, JSONArray[int]{6, 7}, js)

		assert.NoError(t, js.DecodeValue([]byte("[8]")))
		assert.Equal(t, JSONArray[int]{8}, js)

		assert.NoError(t, js.DecodeValue(json.RawMessage("[9]")))
		assert.Equal(t, JSONArray[int]{9}, js)
	})

	t.Run("unmarshal_invalid", func(t *testing.T) {
		var js JSONArray[int]
		assert.Error(t, js.UnmarshalJSON([]byte(`{"a":1}`)))
	})
}
