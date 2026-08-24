package gosql

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedNumberArrayValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{}", v)
	})

	t.Run("empty", func(t *testing.T) {
		arr := OrderedNumberArray[int]{}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{}", v)
	})

	t.Run("values", func(t *testing.T) {
		arr := OrderedNumberArray[int]{1, 2, 3}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{1,2,3}", v)
	})
}

func TestOrderedNumberArrayScan(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		err := arr.Scan(nil)
		assert.Equal(t, ErrNullValueNotAllowed, err)
	})

	t.Run("sql_string", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		err := arr.Scan("{1,3,5}")
		assert.NoError(t, err)
		assert.Equal(t, OrderedNumberArray[int]{1, 3, 5}, arr)
	})

	t.Run("sql_bytes", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		err := arr.Scan([]byte("{2,4}"))
		assert.NoError(t, err)
		assert.Equal(t, OrderedNumberArray[int]{2, 4}, arr)
	})

	t.Run("invalid_type", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		err := arr.Scan(1.2)
		assert.Equal(t, ErrInvalidScan, err)
	})
}

func TestOrderedNumberArrayJSON(t *testing.T) {
	t.Run("marshal_nil", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "[]", string(data))
	})

	t.Run("marshal_values", func(t *testing.T) {
		arr := OrderedNumberArray[int]{1, 2, 3}
		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "[1,2,3]", string(data))
	})

	t.Run("unmarshal_nil_bytes", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		err := arr.UnmarshalJSON(nil)
		assert.Equal(t, ErrNullValueNotAllowed, err)
	})

	t.Run("unmarshal_values", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		err := arr.UnmarshalJSON([]byte("[1,2,3]"))
		assert.NoError(t, err)
		assert.Equal(t, OrderedNumberArray[int]{1, 2, 3}, arr)
	})

	t.Run("roundtrip", func(t *testing.T) {
		arr1 := OrderedNumberArray[int]{1, 2, 3}
		data, err := json.Marshal(arr1)
		assert.NoError(t, err)

		var arr2 OrderedNumberArray[int]
		err = json.Unmarshal(data, &arr2)
		assert.NoError(t, err)
		assert.Equal(t, arr1, arr2)
	})
}

func TestOrderedNumberArrayDecodeValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		err := arr.DecodeValue(nil)
		assert.Equal(t, ErrNullValueNotAllowed, err)
	})

	t.Run("slice", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		err := arr.DecodeValue([]int{1, 2})
		assert.NoError(t, err)
		assert.Equal(t, OrderedNumberArray[int]{1, 2}, arr)
	})

	t.Run("json_string", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		err := arr.DecodeValue("[3,4]")
		assert.NoError(t, err)
		assert.Equal(t, OrderedNumberArray[int]{3, 4}, arr)
	})

	t.Run("invalid_type", func(t *testing.T) {
		var arr OrderedNumberArray[int]
		err := arr.DecodeValue(true)
		assert.Equal(t, ErrInvalidDecodeValue, err)
	})
}

func TestOrderedNumberArrayHelpers(t *testing.T) {
	t.Run("sort", func(t *testing.T) {
		arr := OrderedNumberArray[int]{3, 1, 2}
		got := arr.Sort()
		assert.Equal(t, OrderedNumberArray[int]{1, 2, 3}, arr)
		assert.Equal(t, arr, got)
		assert.Equal(t, 3, arr.Len())
		assert.True(t, arr.Less(0, 1))
		arr.Swap(0, 2)
		assert.Equal(t, OrderedNumberArray[int]{3, 2, 1}, arr)
	})

	t.Run("index_of", func(t *testing.T) {
		arr := OrderedNumberArray[int]{1, 3, 5, 7}
		assert.Equal(t, 0, arr.IndexOf(1))
		assert.Equal(t, 2, arr.IndexOf(5))
		assert.Equal(t, -1, arr.IndexOf(4))
		assert.Equal(t, -1, OrderedNumberArray[int]{}.IndexOf(1))
	})

	t.Run("one_of", func(t *testing.T) {
		arr := OrderedNumberArray[int]{1, 3, 5}
		assert.True(t, arr.OneOf([]int{9, 3}))
		assert.False(t, arr.OneOf([]int{2, 4}))
		assert.False(t, arr.OneOf(nil))
		assert.False(t, OrderedNumberArray[int]{}.OneOf([]int{1}))
	})

	t.Run("filter", func(t *testing.T) {
		arr := OrderedNumberArray[int]{1, 2, 3, 4}
		got := arr.Filter(func(v int) bool { return v >= 3 })
		assert.Equal(t, OrderedNumberArray[int]{3, 4}, got)
	})

	t.Run("map", func(t *testing.T) {
		arr := OrderedNumberArray[int]{1, 2, 3}
		got := arr.Map(func(v int) (int, bool) { return v * 3, v != 2 })
		assert.Equal(t, OrderedNumberArray[int]{3, 9}, got)
	})
}
