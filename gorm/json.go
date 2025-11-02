package gorm

import (
	"github.com/geniusrabbit/gosql/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// JSON field type declaration with GORM type methods
type JSON[T any] struct {
	gosql.JSON[T]
}

// GormDataType gorm common data type
func (JSON[T]) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (j JSON[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return jsonGormDBDataType(db, false)
}

// Data returns the underlying data
func (j *JSON[T]) Data() T {
	return j.JSON.Data
}

// NullableJSON field type declaration with GORM type methods
type NullableJSON[T any] struct {
	gosql.NullableJSON[T]
}

// GormDataType gorm common data type
func (NullableJSON[T]) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (j NullableJSON[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return jsonGormDBDataType(db, true)
}

// Data returns the underlying data pointer
func (j *NullableJSON[T]) Data() *T {
	return j.NullableJSON.Data
}

func jsonGormDBDataType(db *gorm.DB, nullable bool) string {
	switch db.Dialector.Name() {
	case "mysql", "mariadb":
		return "json"
	case "postgres":
		return "jsonb"
	case "sqlite", "sqlite3":
		return "text"
	case "sqlserver":
		return "nvarchar(max)"
	case "ydb":
		if nullable {
			return "Optional<JSONDocument>"
		}
		return "JSONDocument"
	case "clickhouse":
		if nullable {
			return "Nullable(JSON)"
		}
		return "JSON"
	}
	return ""
}
