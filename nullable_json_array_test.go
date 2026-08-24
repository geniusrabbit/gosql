package gosql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullableJSONArray(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		js, err := NewNullableJSONArray[int]([]int{1, 2, 3})
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, 1, js[0])
		assert.Equal(t, 2, js[1])
		assert.Equal(t, 3, js[2])

		assert.Error(t, js.SetValue([]string{"invalid"}))

		assert.Panics(t, func() {
			MustNullableJSONArray[int](map[string]int{"invalid": 1})
		})
	})

	t.Run("json_empty_parse", func(t *testing.T) {
		var js NullableJSONArray[int]
		assert.NoError(t, js.Scan([]byte{}), "scan empty byte")
		assert.NoError(t, js.Scan([]byte(nil)), "scan nil byte")
		assert.NoError(t, js.Scan([]byte(" ")), "scan empty string")
	})

	t.Run("json_encode", func(t *testing.T) {
		js, err := NewNullableJSONArray[int]([]int{1, 2, 3})
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
		var js NullableJSONArray[int]
		data, err := js.MarshalJSON()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, "null", string(data))
		assert.Equal(t, "null", js.String())
		v, _ := js.Value()
		assert.Nil(t, v)
	})

	t.Run("scan", func(t *testing.T) {
		var js NullableJSONArray[int]
		if assert.NoError(t, js.Scan([]byte("[1,2,3]"))) {
			assert.ElementsMatch(t, []int{1, 2, 3}, js)
		}
	})

	t.Run("scan_nil", func(t *testing.T) {
		js := NullableJSONArray[int]{1}
		assert.NoError(t, js.Scan(nil))
		assert.Nil(t, js)
	})

	t.Run("set_value_nil", func(t *testing.T) {
		js := NullableJSONArray[int]{1}
		assert.NoError(t, js.SetValue(nil))
		assert.Nil(t, js)
	})

	t.Run("unmarshal_empty", func(t *testing.T) {
		js := NullableJSONArray[int]{1}
		assert.NoError(t, js.UnmarshalJSON(nil))
		assert.Nil(t, js)
	})

	t.Run("decode_value", func(t *testing.T) {
		var js NullableJSONArray[int]
		assert.NoError(t, js.DecodeValue([]int{1, 2}))
		assert.Equal(t, NullableJSONArray[int]{1, 2}, js)

		var nilPtr *NullableJSONArray[int]
		assert.NoError(t, nilPtr.DecodeValue([]int{1}))
	})

	t.Run("string_nil", func(t *testing.T) {
		var js NullableJSONArray[int]
		assert.Equal(t, "null", js.String())
	})
}
