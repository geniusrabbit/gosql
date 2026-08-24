package gosql

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullableNumberArrayValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var arr NullableNumberArray[int]
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("empty", func(t *testing.T) {
		arr := NullableNumberArray[int]{}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{}", v)
	})

	t.Run("values", func(t *testing.T) {
		arr := NullableNumberArray[int]{1, 2, 3}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{1,2,3}", v)
	})
}

func TestNullableNumberArrayScan(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		arr := NullableNumberArray[int]{1}
		err := arr.Scan(nil)
		assert.NoError(t, err)
		assert.Nil(t, arr)
	})

	t.Run("sql_string", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.Scan("{1,2,3}")
		assert.NoError(t, err)
		assert.Equal(t, NullableNumberArray[int]{1, 2, 3}, arr)
	})

	t.Run("sql_bytes", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.Scan([]byte("{4,5}"))
		assert.NoError(t, err)
		assert.Equal(t, NullableNumberArray[int]{4, 5}, arr)
	})

	t.Run("null_string", func(t *testing.T) {
		arr := NullableNumberArray[int]{1}
		err := arr.Scan("null")
		assert.NoError(t, err)
		assert.Nil(t, arr)
	})

	t.Run("empty_braces", func(t *testing.T) {
		arr := NullableNumberArray[int]{1}
		err := arr.Scan("{}")
		assert.NoError(t, err)
		assert.Equal(t, NullableNumberArray[int]{}, arr)
	})

	t.Run("slice", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.Scan([]int{7, 8})
		assert.NoError(t, err)
		assert.Equal(t, NullableNumberArray[int]{7, 8}, arr)
	})

	t.Run("invalid_type", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.Scan(true)
		assert.Equal(t, ErrInvalidScan, err)
	})
}

func TestNullableNumberArrayJSON(t *testing.T) {
	t.Run("marshal_nil", func(t *testing.T) {
		var arr NullableNumberArray[int]
		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "null", string(data))
	})

	t.Run("marshal_empty", func(t *testing.T) {
		arr := NullableNumberArray[int]{}
		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "[]", string(data))
	})

	t.Run("marshal_values", func(t *testing.T) {
		arr := NullableNumberArray[int]{1, 2, 3}
		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "[1,2,3]", string(data))
	})

	t.Run("unmarshal_null", func(t *testing.T) {
		arr := NullableNumberArray[int]{1}
		err := arr.UnmarshalJSON([]byte("null"))
		assert.NoError(t, err)
		assert.Nil(t, arr)
	})

	t.Run("unmarshal_values", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.UnmarshalJSON([]byte("[1,2,3]"))
		assert.NoError(t, err)
		assert.Equal(t, NullableNumberArray[int]{1, 2, 3}, arr)
	})

	t.Run("roundtrip", func(t *testing.T) {
		arr1 := NullableNumberArray[int]{1, 2, 3}
		data, err := json.Marshal(arr1)
		assert.NoError(t, err)

		var arr2 NullableNumberArray[int]
		err = json.Unmarshal(data, &arr2)
		assert.NoError(t, err)
		assert.Equal(t, arr1, arr2)
	})

	t.Run("roundtrip_nil", func(t *testing.T) {
		var arr1 NullableNumberArray[int]
		data, err := json.Marshal(arr1)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(data))

		var arr2 NullableNumberArray[int]
		err = json.Unmarshal(data, &arr2)
		assert.NoError(t, err)
		assert.Nil(t, arr2)
	})
}

func TestNullableNumberArrayDecodeValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		arr := NullableNumberArray[int]{1}
		err := arr.DecodeValue(nil)
		assert.NoError(t, err)
		assert.Nil(t, arr)
	})

	t.Run("slice", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.DecodeValue([]int{1, 2})
		assert.NoError(t, err)
		assert.Equal(t, NullableNumberArray[int]{1, 2}, arr)
	})

	t.Run("nullable_number_array", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.DecodeValue(NullableNumberArray[int]{3, 4})
		assert.NoError(t, err)
		assert.Equal(t, NullableNumberArray[int]{3, 4}, arr)
	})

	t.Run("number_array", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.DecodeValue(NumberArray[int]{5, 6})
		assert.NoError(t, err)
		assert.Equal(t, NullableNumberArray[int]{5, 6}, arr)
	})

	t.Run("json_bytes", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.DecodeValue([]byte("[7,8]"))
		assert.NoError(t, err)
		assert.Equal(t, NullableNumberArray[int]{7, 8}, arr)
	})

	t.Run("invalid_type", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.DecodeValue(1.5)
		assert.Equal(t, ErrInvalidDecodeValue, err)
	})

	t.Run("invalid_string", func(t *testing.T) {
		var arr NullableNumberArray[int]
		err := arr.DecodeValue("[1,abc]")
		assert.Error(t, err)
	})
}

func TestNullableNumberArrayHelpers(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		arr := NullableNumberArray[int]{3, 1, 2}
		got := arr.Sort()
		assert.Equal(t, NullableNumberArray[int]{1, 2, 3}, arr)
		assert.Equal(t, arr, got)
		assert.Equal(t, 3, arr.Len())
		assert.True(t, arr.Less(0, 2))
	})

	t.Run("index_of", func(t *testing.T) {
		arr := NullableNumberArray[int]{1, 3, 5}
		assert.Equal(t, 0, arr.IndexOf(1))
		assert.Equal(t, 2, arr.IndexOf(5))
		assert.Equal(t, -1, arr.IndexOf(2))

		var nilArr NullableNumberArray[int]
		assert.Equal(t, -1, nilArr.IndexOf(1))
	})

	t.Run("one_of", func(t *testing.T) {
		arr := NullableNumberArray[int]{1, 3, 5}
		assert.True(t, arr.OneOf([]int{5}))
		assert.False(t, arr.OneOf([]int{2}))
		assert.False(t, arr.OneOf(nil))
		assert.False(t, NullableNumberArray[int]{}.OneOf([]int{1}))
	})

	t.Run("ordered", func(t *testing.T) {
		arr := NullableNumberArray[int]{3, 1, 2}
		got := arr.Ordered()
		assert.Equal(t, NullableNumberArray[int]{1, 2, 3}, arr)
		assert.Equal(t, NullableOrderedNumberArray[int]{1, 2, 3}, got)
	})

	t.Run("filter", func(t *testing.T) {
		arr := NullableNumberArray[int]{1, 2, 3, 4}
		got := arr.Filter(func(v int) bool { return v > 2 })
		assert.Equal(t, NullableNumberArray[int]{3, 4}, got)
	})

	t.Run("map", func(t *testing.T) {
		arr := NullableNumberArray[int]{1, 2, 3}
		got := arr.Map(func(v int) (int, bool) { return v + 1, true })
		assert.Equal(t, NullableNumberArray[int]{2, 3, 4}, got)
	})
}
