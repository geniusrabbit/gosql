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

	t.Run("data_or", func(t *testing.T) {
		js, _ := NewNullableJSON[map[string]string](map[string]any{
			"name":   "Yoda",
			"number": "1",
		})
		assert.Equal(t, "Yoda", js.DataOr(map[string]string{})["name"])
		assert.Equal(t, "1", js.DataOr(map[string]string{})["number"])
		assert.Equal(t, "Yoda", js.DataOr(map[string]string{"name": "Luke"})["name"])
		assert.Equal(t, "1", js.DataOr(map[string]string{"number": "2"})["number"])

		null, _ := NewNullableJSON[map[string]string](nil)
		assert.Equal(t, "Luke", null.DataOr(map[string]string{"name": "Luke"})["name"])
		assert.Equal(t, "2", null.DataOr(map[string]string{"number": "2"})["number"])
		assert.Nil(t, null.DataOr(nil))
	})

	t.Run("set_value", func(t *testing.T) {
		var js NullableJSON[map[string]string]
		assert.NoError(t, js.SetValue(map[string]string{"k": "v"}))
		assert.Equal(t, "v", (*js.Data)["k"])

		val := map[string]string{"k": "w"}
		assert.NoError(t, js.SetValue(&val))
		assert.Equal(t, "w", (*js.Data)["k"])

		assert.NoError(t, js.SetValue(nil))
		assert.Nil(t, js.Data)

		assert.NoError(t, js.SetValue(`{"k":"x"}`))
		assert.Equal(t, "x", (*js.Data)["k"])

		assert.NoError(t, js.SetValue([]byte(`{"k":"y"}`)))
		assert.Equal(t, "y", (*js.Data)["k"])
	})

	t.Run("string_and_value", func(t *testing.T) {
		n := 7
		js := MustNullableJSON[int](&n)
		assert.Equal(t, "7", js.String())

		v, err := js.Value()
		assert.NoError(t, err)
		assert.Equal(t, "7", v)

		var empty NullableJSON[int]
		assert.Equal(t, "null", empty.String())
		v, err = empty.Value()
		assert.NoError(t, err)
		assert.Equal(t, "null", v)

		var nilJS *NullableJSON[int]
		assert.Equal(t, "null", nilJS.String())
	})

	t.Run("scan", func(t *testing.T) {
		var js NullableJSON[int]
		assert.NoError(t, js.Scan("12"))
		assert.Equal(t, 12, *js.Data)

		assert.NoError(t, js.Scan([]byte("13")))
		assert.Equal(t, 13, *js.Data)

		assert.NoError(t, js.Scan(nil))
		assert.Nil(t, js.Data)

		assert.Equal(t, ErrInvalidScan, js.Scan(1))
	})

	t.Run("decode_value", func(t *testing.T) {
		var js NullableJSON[int]
		assert.NoError(t, js.DecodeValue("14"))
		assert.Equal(t, 14, *js.Data)

		assert.NoError(t, js.DecodeValue([]byte("15")))
		assert.Equal(t, 15, *js.Data)

		assert.NoError(t, js.DecodeValue(16))
		assert.Equal(t, 16, *js.Data)
	})

	t.Run("json_roundtrip", func(t *testing.T) {
		n := 9
		js := MustNullableJSON[int](&n)
		data, err := js.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "9", string(data))

		var got NullableJSON[int]
		assert.NoError(t, got.UnmarshalJSON(data))
		assert.Equal(t, 9, *got.Data)

		var empty NullableJSON[int]
		data, err = empty.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "null", string(data))
	})
}
