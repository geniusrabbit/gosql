package gosql

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberArrayValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var arr NumberArray[int]
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{}", v)
	})

	t.Run("empty", func(t *testing.T) {
		arr := NumberArray[int]{}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{}", v)
	})

	t.Run("values", func(t *testing.T) {
		arr := NumberArray[int]{1, 2, 3}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{1,2,3}", v)
	})

	t.Run("uint", func(t *testing.T) {
		arr := NumberArray[uint]{1, 2}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{1,2}", v)
	})
}

func TestNumberArrayScan(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.Scan(nil)
		assert.Equal(t, ErrNullValueNotAllowed, err)
	})

	t.Run("sql_string", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.Scan("{1,2,3}")
		assert.NoError(t, err)
		assert.Equal(t, NumberArray[int]{1, 2, 3}, arr)
	})

	t.Run("sql_bytes", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.Scan([]byte("{4,5}"))
		assert.NoError(t, err)
		assert.Equal(t, NumberArray[int]{4, 5}, arr)
	})

	t.Run("empty_braces", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.Scan("{}")
		assert.NoError(t, err)
		assert.Equal(t, NumberArray[int]{}, arr)
	})

	t.Run("invalid_type", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.Scan(123)
		assert.Equal(t, ErrInvalidScan, err)
	})

	t.Run("invalid_number", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.Scan("{1,abc}")
		assert.Error(t, err)
	})
}

func TestNumberArrayJSON(t *testing.T) {
	t.Run("marshal_nil", func(t *testing.T) {
		var arr NumberArray[int]
		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "[]", string(data))
	})

	t.Run("marshal_values", func(t *testing.T) {
		arr := NumberArray[int]{1, 2, 3}
		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "[1,2,3]", string(data))
	})

	t.Run("unmarshal_nil_bytes", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.UnmarshalJSON(nil)
		assert.Equal(t, ErrNullValueNotAllowed, err)
	})

	t.Run("unmarshal_values", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.UnmarshalJSON([]byte("[1,2,3]"))
		assert.NoError(t, err)
		assert.Equal(t, NumberArray[int]{1, 2, 3}, arr)
	})

	t.Run("roundtrip", func(t *testing.T) {
		arr1 := NumberArray[int]{1, 2, 3}
		data, err := json.Marshal(arr1)
		assert.NoError(t, err)

		var arr2 NumberArray[int]
		err = json.Unmarshal(data, &arr2)
		assert.NoError(t, err)
		assert.Equal(t, arr1, arr2)
	})
}

func TestNumberArrayDecodeValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.DecodeValue(nil)
		assert.Equal(t, ErrNullValueNotAllowed, err)
	})

	t.Run("slice", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.DecodeValue([]int{1, 2})
		assert.NoError(t, err)
		assert.Equal(t, NumberArray[int]{1, 2}, arr)
	})

	t.Run("json_string", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.DecodeValue("[3,4]")
		assert.NoError(t, err)
		assert.Equal(t, NumberArray[int]{3, 4}, arr)
	})

	t.Run("invalid_type", func(t *testing.T) {
		var arr NumberArray[int]
		err := arr.DecodeValue(123)
		assert.Equal(t, ErrInvalidDecodeValue, err)
	})
}

func TestNumberArrayHelpers(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		arr := NumberArray[int]{3, 1, 2}
		got := arr.Sort()
		assert.Equal(t, NumberArray[int]{1, 2, 3}, arr)
		assert.Equal(t, arr, got)
		assert.Equal(t, 3, arr.Len())
	})

	t.Run("index_of", func(t *testing.T) {
		arr := NumberArray[int]{1, 3, 5}
		assert.Equal(t, 1, arr.IndexOf(3))
		assert.Equal(t, -1, arr.IndexOf(2))
	})

	t.Run("one_of", func(t *testing.T) {
		arr := NumberArray[int]{1, 3, 5}
		assert.True(t, arr.OneOf([]int{2, 3}))
		assert.False(t, arr.OneOf([]int{2, 4}))
		assert.False(t, arr.OneOf(nil))
	})

	t.Run("ordered", func(t *testing.T) {
		arr := NumberArray[int]{3, 1, 2}
		got := arr.Ordered()
		assert.Equal(t, NumberArray[int]{1, 2, 3}, arr)
		assert.Equal(t, OrderedNumberArray[int]{1, 2, 3}, got)
	})

	t.Run("filter", func(t *testing.T) {
		arr := NumberArray[int]{1, 2, 3, 4}
		got := arr.Filter(func(v int) bool { return v%2 == 0 })
		assert.Equal(t, NumberArray[int]{2, 4}, got)
		assert.Equal(t, NumberArray[int]{1, 2, 3, 4}, arr)
	})

	t.Run("map", func(t *testing.T) {
		arr := NumberArray[int]{1, 2, 3}
		got := arr.Map(func(v int) (int, bool) {
			return v * 2, v != 2
		})
		assert.Equal(t, NumberArray[int]{2, 6}, got)
	})
}
