package gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/schema"
)

func TestGormJSON(t *testing.T) {
	t.Run("constructors", func(t *testing.T) {
		js, err := NewJSON(map[string]string{"k": "v"})
		assert.NoError(t, err)
		assert.Equal(t, "v", js.Data()["k"])

		js = MustJSON(map[string]string{"k": "w"})
		assert.Equal(t, "w", js.Data()["k"])
	})

	t.Run("data_type", func(t *testing.T) {
		var js JSON[int]
		assert.Equal(t, "json", js.GormDataType())
	})

	t.Run("db_data_type", func(t *testing.T) {
		var js JSON[int]
		field := &schema.Field{Name: "data"}
		tests := []struct {
			dialect  string
			expected string
		}{
			{"mysql", "json"},
			{"mariadb", "json"},
			{"postgres", "jsonb"},
			{"sqlite", "text"},
			{"sqlite3", "text"},
			{"sqlserver", "nvarchar(max)"},
			{"ydb", "JsonDocument"},
			{"clickhouse", "JSON"},
			{"unknown", ""},
		}
		for _, test := range tests {
			t.Run(test.dialect, func(t *testing.T) {
				assert.Equal(t, test.expected, js.GormDBDataType(createMockDB(test.dialect), field))
			})
		}
	})
}

func TestGormNullableJSON(t *testing.T) {
	t.Run("constructors", func(t *testing.T) {
		val := map[string]string{"k": "v"}
		js, err := NewNullableJSON(&val)
		assert.NoError(t, err)
		assert.Equal(t, "v", (*js.Data())["k"])

		js = MustNullableJSON(&val)
		assert.Equal(t, "v", (*js.Data())["k"])
	})

	t.Run("data_type", func(t *testing.T) {
		var js NullableJSON[int]
		assert.Equal(t, "json", js.GormDataType())
	})

	t.Run("db_data_type", func(t *testing.T) {
		var js NullableJSON[int]
		field := &schema.Field{Name: "data"}
		assert.Equal(t, "jsonb", js.GormDBDataType(createMockDB("postgres"), field))
		assert.Equal(t, "Nullable(JSON)", js.GormDBDataType(createMockDB("clickhouse"), field))
		assert.Equal(t, "json", js.GormDBDataType(createMockDB("mysql"), field))
	})
}
