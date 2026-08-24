package gosql

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullableOrderedNumberArrayValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("empty", func(t *testing.T) {
		arr := NullableOrderedNumberArray[int]{}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{}", v)
	})

	t.Run("values", func(t *testing.T) {
		arr := NullableOrderedNumberArray[int]{1, 2, 3}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{1,2,3}", v)
	})

	t.Run("float64", func(t *testing.T) {
		arr := NullableOrderedNumberArray[float64]{1.5, 2.25}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{1.5,2.25}", v)
	})
}

func TestNullableOrderedNumberArrayScan(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		arr := NullableOrderedNumberArray[int]{1}
		err := arr.Scan(nil)
		assert.NoError(t, err)
		assert.Nil(t, arr)
	})

	t.Run("sql_string", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.Scan("{1,2,3}")
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{1, 2, 3}, arr)
	})

	t.Run("sql_bytes", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.Scan([]byte("{4,5,6}"))
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{4, 5, 6}, arr)
	})

	t.Run("empty_braces", func(t *testing.T) {
		arr := NullableOrderedNumberArray[int]{1}
		err := arr.Scan("{}")
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{}, arr)
	})

	t.Run("null_string", func(t *testing.T) {
		arr := NullableOrderedNumberArray[int]{1}
		err := arr.Scan("null")
		assert.NoError(t, err)
		assert.Nil(t, arr)
	})

	t.Run("slice", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.Scan([]int{7, 8})
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{7, 8}, arr)
	})

	t.Run("invalid_type", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.Scan(123)
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidScan, err)
	})

	t.Run("invalid_number", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.Scan("{1,abc}")
		assert.Error(t, err)
	})
}

func TestNullableOrderedNumberArrayJSON(t *testing.T) {
	t.Run("marshal_nil", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "null", string(data))
	})

	t.Run("marshal_empty", func(t *testing.T) {
		arr := NullableOrderedNumberArray[int]{}
		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "[]", string(data))
	})

	t.Run("marshal_values", func(t *testing.T) {
		arr := NullableOrderedNumberArray[int]{1, 2, 3}
		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "[1,2,3]", string(data))
	})

	t.Run("unmarshal_nil_bytes", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.UnmarshalJSON(nil)
		assert.Error(t, err)
		assert.Equal(t, ErrNullValueNotAllowed, err)
	})

	t.Run("unmarshal_null", func(t *testing.T) {
		arr := NullableOrderedNumberArray[int]{1}
		err := arr.UnmarshalJSON([]byte("null"))
		assert.NoError(t, err)
		assert.Nil(t, arr)
	})

	t.Run("unmarshal_empty", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.UnmarshalJSON([]byte("[]"))
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{}, arr)
	})

	t.Run("unmarshal_values", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.UnmarshalJSON([]byte("[1,2,3]"))
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{1, 2, 3}, arr)
	})

	t.Run("roundtrip", func(t *testing.T) {
		arr1 := NullableOrderedNumberArray[int]{1, 2, 3}
		data, err := json.Marshal(arr1)
		assert.NoError(t, err)

		var arr2 NullableOrderedNumberArray[int]
		err = json.Unmarshal(data, &arr2)
		assert.NoError(t, err)
		assert.Equal(t, arr1, arr2)
	})

	t.Run("roundtrip_nil", func(t *testing.T) {
		var arr1 NullableOrderedNumberArray[int]
		data, err := json.Marshal(arr1)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(data))

		var arr2 NullableOrderedNumberArray[int]
		err = json.Unmarshal(data, &arr2)
		assert.NoError(t, err)
		assert.Nil(t, arr2)
	})

	t.Run("unmarshal_invalid", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.UnmarshalJSON([]byte("[1,abc]"))
		assert.Error(t, err)
	})
}

func TestNullableOrderedNumberArrayDecodeValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		arr := NullableOrderedNumberArray[int]{1}
		err := arr.DecodeValue(nil)
		assert.NoError(t, err)
		assert.Nil(t, arr)
	})

	t.Run("slice", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.DecodeValue([]int{1, 2, 3})
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{1, 2, 3}, arr)
	})

	t.Run("nullable_number_array", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.DecodeValue(NullableNumberArray[int]{4, 5})
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{4, 5}, arr)
	})

	t.Run("number_array", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.DecodeValue(NumberArray[int]{6, 7})
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{6, 7}, arr)
	})

	t.Run("json_string", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.DecodeValue("[8,9]")
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{8, 9}, arr)
	})

	t.Run("json_bytes", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.DecodeValue([]byte("[10,11]"))
		assert.NoError(t, err)
		assert.Equal(t, NullableOrderedNumberArray[int]{10, 11}, arr)
	})

	t.Run("invalid_type", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.DecodeValue(123)
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidDecodeValue, err)
	})

	t.Run("invalid_string", func(t *testing.T) {
		var arr NullableOrderedNumberArray[int]
		err := arr.DecodeValue("[1,abc]")
		assert.Error(t, err)
	})
}

func TestNullableOrderedNumberArraySort(t *testing.T) {
	arr := NullableOrderedNumberArray[int]{3, 1, 2}
	got := arr.Sort()
	assert.Equal(t, NullableOrderedNumberArray[int]{1, 2, 3}, arr)
	assert.Equal(t, arr, got)
	assert.Equal(t, 3, arr.Len())
	assert.True(t, arr.Less(0, 2))
	assert.False(t, arr.Less(2, 0))

	arr.Swap(0, 2)
	assert.Equal(t, NullableOrderedNumberArray[int]{3, 2, 1}, arr)
}

func TestNullableOrderedNumberArrayIndexOf(t *testing.T) {
	arr := NullableOrderedNumberArray[int]{1, 3, 5, 7}

	assert.Equal(t, 0, arr.IndexOf(1))
	assert.Equal(t, 1, arr.IndexOf(3))
	assert.Equal(t, 3, arr.IndexOf(7))
	assert.Equal(t, -1, arr.IndexOf(2))
	assert.Equal(t, -1, arr.IndexOf(0))
	assert.Equal(t, -1, arr.IndexOf(8))

	var empty NullableOrderedNumberArray[int]
	assert.Equal(t, -1, empty.IndexOf(1))
}

func TestNullableOrderedNumberArrayOneOf(t *testing.T) {
	arr := NullableOrderedNumberArray[int]{1, 3, 5}

	assert.True(t, arr.OneOf([]int{5, 9}))
	assert.True(t, arr.OneOf([]int{1}))
	assert.False(t, arr.OneOf([]int{2, 4}))
	assert.False(t, arr.OneOf(nil))
	assert.False(t, arr.OneOf([]int{}))

	var empty NullableOrderedNumberArray[int]
	assert.False(t, empty.OneOf([]int{1}))
	assert.False(t, NullableOrderedNumberArray[int]{}.OneOf([]int{1}))
}

func TestNullableOrderedNumberArrayFilter(t *testing.T) {
	arr := NullableOrderedNumberArray[int]{1, 2, 3, 4, 5}
	got := arr.Filter(func(v int) bool { return v%2 == 0 })
	assert.Equal(t, NullableOrderedNumberArray[int]{2, 4}, got)
	assert.Equal(t, NullableOrderedNumberArray[int]{1, 2, 3, 4, 5}, arr)

	empty := arr.Filter(func(v int) bool { return v > 10 })
	assert.Equal(t, NullableOrderedNumberArray[int]{}, empty)

	var nilArr NullableOrderedNumberArray[int]
	assert.Equal(t, NullableOrderedNumberArray[int]{}, nilArr.Filter(func(v int) bool { return true }))
}

func TestNullableOrderedNumberArrayMap(t *testing.T) {
	arr := NullableOrderedNumberArray[int]{1, 2, 3, 4}
	got := arr.Map(func(v int) (int, bool) {
		if v%2 != 0 {
			return 0, false
		}
		return v * 10, true
	})
	assert.Equal(t, NullableOrderedNumberArray[int]{20, 40}, got)
	assert.Equal(t, NullableOrderedNumberArray[int]{1, 2, 3, 4}, arr)

	empty := arr.Map(func(v int) (int, bool) { return 0, false })
	assert.Equal(t, NullableOrderedNumberArray[int]{}, empty)
}
