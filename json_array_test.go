package gosql

import (
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
}
