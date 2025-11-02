package gorm

import (
	"database/sql/driver"

	"github.com/geniusrabbit/gosql/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// JSONArray field type declaration with GORM type methods
type JSONArray[T any] gosql.JSONArray[T]

// NewJSONArray creates new JSONArray object
func NewJSONArray[T any](val any) (JSONArray[T], error) {
	arr, err := gosql.NewJSONArray[T](val)
	if err != nil {
		return nil, err
	}
	return JSONArray[T](arr), nil
}

// MustJSONArray creates new JSONArray object
func MustJSONArray[T any](val any) JSONArray[T] {
	obj, err := NewJSONArray[T](val)
	if err != nil {
		panic(err)
	}
	return obj
}

// GormDataType gorm common data type
func (JSONArray[T]) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (j JSONArray[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return jsonGormDBDataType(db, false)
}

// String value
func (f JSONArray[T]) String() string {
	return (gosql.JSONArray[T](f)).String()
}

// SetValue of json
func (f *JSONArray[T]) SetValue(value any) error {
	return (*gosql.JSONArray[T])(f).SetValue(value)
}

// Value implements the driver.Valuer interface, []T field
func (f JSONArray[T]) Value() (driver.Value, error) {
	return (gosql.JSONArray[T](f)).Value()
}

// Scan implements the driver.Valuer interface, []T field
func (f *JSONArray[T]) Scan(value any) error {
	return (*gosql.JSONArray[T])(f).Scan(value)
}

// MarshalJSON implements the json.Marshaler
func (f JSONArray[T]) MarshalJSON() ([]byte, error) {
	return (gosql.JSONArray[T](f)).MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *JSONArray[T]) UnmarshalJSON(b []byte) error {
	return (*gosql.JSONArray[T])(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *JSONArray[T]) DecodeValue(v any) error {
	return (*gosql.JSONArray[T])(f).DecodeValue(v)
}
