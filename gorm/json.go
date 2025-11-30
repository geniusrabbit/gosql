package gorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/geniusrabbit/gosql/v2"
)

// JSON field type declaration with GORM type methods
type JSON[T any] struct {
	gosql.JSON[T]
}

// NewJSON creates new JSON object
// This function is used to create a new JSON object with the provided data.
// It wraps the gosql.NewJSON function to return a GORM-compatible JSON type.
// If an error occurs during the creation, it returns nil and the error.
func NewJSON[T any](data T) (*JSON[T], error) {
	jdata, err := gosql.NewJSON[T](data)
	if err != nil {
		return nil, err
	}
	return &JSON[T]{*jdata}, nil
}

// MustJSON creates new JSON object
func MustJSON[T any](data T) *JSON[T] {
	return &JSON[T]{*gosql.MustJSON[T](data)}
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

// NewNullableJSON creates new NullableJSON object
// This function is used to create a new NullableJSON object with the provided data.
// It wraps the gosql.NewNullableJSON function to return a GORM-compatible NullableJSON type.
// If an error occurs during the creation, it returns nil and the error.
func NewNullableJSON[T any](data *T) (*NullableJSON[T], error) {
	jdata, err := gosql.NewNullableJSON[T](data)
	if err != nil {
		return nil, err
	}
	return &NullableJSON[T]{*jdata}, nil
}

// MustNullableJSON creates new NullableJSON object
func MustNullableJSON[T any](data *T) *NullableJSON[T] {
	return &NullableJSON[T]{*gosql.MustNullableJSON[T](data)}
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
		return "JsonDocument"
	case "clickhouse":
		if nullable {
			return "Nullable(JSON)"
		}
		return "JSON"
	}
	return ""
}
