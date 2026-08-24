package gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/schema"
)

func TestGormJSONArray(t *testing.T) {
	t.Run("constructors", func(t *testing.T) {
		arr, err := NewJSONArray[int]([]int{1, 2, 3})
		assert.NoError(t, err)
		assert.Equal(t, JSONArray[int]{1, 2, 3}, arr)

		arr = MustJSONArray[int]([]int{4, 5})
		assert.Equal(t, JSONArray[int]{4, 5}, arr)

		assert.Panics(t, func() {
			MustJSONArray[int](map[string]int{"invalid": 1})
		})
	})

	t.Run("data_type", func(t *testing.T) {
		var arr JSONArray[int]
		assert.Equal(t, "json", arr.GormDataType())
		assert.Equal(t, "jsonb", arr.GormDBDataType(createMockDB("postgres"), &schema.Field{}))
		assert.Equal(t, "JSON", arr.GormDBDataType(createMockDB("clickhouse"), &schema.Field{}))
	})

	t.Run("value_scan_json", func(t *testing.T) {
		arr := MustJSONArray[int]([]int{1, 2})
		assert.Equal(t, "[1,2]", arr.String())

		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "[1,2]", v)

		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "[1,2]", string(data))

		var got JSONArray[int]
		assert.NoError(t, got.Scan("[3,4]"))
		assert.Equal(t, JSONArray[int]{3, 4}, got)

		assert.NoError(t, got.UnmarshalJSON([]byte("[5]")))
		assert.Equal(t, JSONArray[int]{5}, got)

		assert.NoError(t, got.DecodeValue([]int{6}))
		assert.Equal(t, JSONArray[int]{6}, got)

		assert.NoError(t, got.SetValue([]int{7, 8}))
		assert.Equal(t, JSONArray[int]{6, 7, 8}, got)
	})
}

func TestGormNullableJSONArray(t *testing.T) {
	t.Run("constructors", func(t *testing.T) {
		arr, err := NewNullableJSONArray[int]([]int{1, 2})
		assert.NoError(t, err)
		assert.Equal(t, NullableJSONArray[int]{1, 2}, arr)

		arr = MustNullableJSONArray[int]([]int{3})
		assert.Equal(t, NullableJSONArray[int]{3}, arr)

		assert.Panics(t, func() {
			MustNullableJSONArray[int](map[string]int{"invalid": 1})
		})
	})

	t.Run("data_type", func(t *testing.T) {
		var arr NullableJSONArray[int]
		assert.Equal(t, "json", arr.GormDataType())
		assert.Equal(t, "jsonb", arr.GormDBDataType(createMockDB("postgres"), &schema.Field{}))
		assert.Equal(t, "Nullable(JSON)", arr.GormDBDataType(createMockDB("clickhouse"), &schema.Field{}))
	})

	t.Run("nil_and_values", func(t *testing.T) {
		var arr NullableJSONArray[int]
		assert.Equal(t, "null", arr.String())

		data, err := arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "null", string(data))

		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Nil(t, v)

		assert.NoError(t, arr.Scan(nil))
		assert.Nil(t, arr)

		assert.NoError(t, arr.SetValue([]int{1, 2}))
		assert.Equal(t, NullableJSONArray[int]{1, 2}, arr)

		assert.NoError(t, arr.DecodeValue([]int{3}))
		assert.Equal(t, NullableJSONArray[int]{3}, arr)

		assert.NoError(t, arr.UnmarshalJSON([]byte("[4,5]")))
		assert.Equal(t, NullableJSONArray[int]{4, 5}, arr)
	})
}
