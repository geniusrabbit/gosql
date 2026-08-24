package gosql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayNumberDecode(t *testing.T) {
	t.Run("int_sql", func(t *testing.T) {
		got, err := ArrayNumberDecode[int]("{1,2,3}", '{', '}')
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, got)
	})

	t.Run("int_json", func(t *testing.T) {
		got, err := ArrayNumberDecode[int]("[1,2,3]", '[', ']')
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, got)
	})

	t.Run("quoted_numbers", func(t *testing.T) {
		got, err := ArrayNumberDecode[int](`{"1","2"}`, '{', '}')
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2}, got)
	})

	t.Run("uint", func(t *testing.T) {
		got, err := ArrayNumberDecode[uint]("{10,20}", '{', '}')
		assert.NoError(t, err)
		assert.Equal(t, []uint{10, 20}, got)
	})

	t.Run("float64", func(t *testing.T) {
		got, err := ArrayNumberDecode[float64]("{1.5,2.25}", '{', '}')
		assert.NoError(t, err)
		assert.Equal(t, []float64{1.5, 2.25}, got)
	})

	t.Run("bytes", func(t *testing.T) {
		got, err := ArrayNumberDecode[int]([]byte("{7,8}"), '{', '}')
		assert.NoError(t, err)
		assert.Equal(t, []int{7, 8}, got)
	})

	t.Run("slice", func(t *testing.T) {
		src := []int{1, 2}
		got, err := ArrayNumberDecode[int](src, '{', '}')
		assert.NoError(t, err)
		assert.Equal(t, src, got)
	})

	t.Run("nil", func(t *testing.T) {
		got, err := ArrayNumberDecode[int](nil, '{', '}')
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("null_string", func(t *testing.T) {
		got, err := ArrayNumberDecode[int]("null", '{', '}')
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("empty", func(t *testing.T) {
		got, err := ArrayNumberDecode[int]("{}", '{', '}')
		assert.NoError(t, err)
		assert.Equal(t, []int{}, got)

		got, err = ArrayNumberDecode[int]("[]", '[', ']')
		assert.NoError(t, err)
		assert.Equal(t, []int{}, got)

		got, err = ArrayNumberDecode[int]("", '{', '}')
		assert.NoError(t, err)
		assert.Equal(t, []int{}, got)
	})

	t.Run("invalid_type", func(t *testing.T) {
		_, err := ArrayNumberDecode[int](123, '{', '}')
		assert.Equal(t, ErrInvalidScan, err)
	})

	t.Run("invalid_int", func(t *testing.T) {
		_, err := ArrayNumberDecode[int]("{1,x}", '{', '}')
		assert.Error(t, err)
	})

	t.Run("invalid_uint", func(t *testing.T) {
		_, err := ArrayNumberDecode[uint]("{-1}", '{', '}')
		assert.Error(t, err)
	})

	t.Run("invalid_float", func(t *testing.T) {
		_, err := ArrayNumberDecode[float64]("{1.2.3}", '{', '}')
		assert.Error(t, err)
	})
}

func TestArrayNumberEncode(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, "{1,2,3}", ArrayNumberEncode('{', '}', []int{1, 2, 3}).String())
		assert.Equal(t, "[1,2,3]", string(ArrayNumberEncode('[', ']', []int{1, 2, 3}).Bytes()))
	})

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, "{}", ArrayNumberEncode('{', '}', []int{}).String())
	})

	t.Run("uint", func(t *testing.T) {
		assert.Equal(t, "{1,2}", ArrayNumberEncode('{', '}', []uint{1, 2}).String())
	})

	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, "{1.5,2.25}", ArrayNumberEncode('{', '}', []float64{1.5, 2.25}).String())
	})
}
