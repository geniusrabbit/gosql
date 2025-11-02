package gorm

import (
	"database/sql/driver"

	"github.com/geniusrabbit/gosql/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// NullableJSONArray field type declaration with GORM type methods
type NullableJSONArray[T any] gosql.NullableJSONArray[T]

// NewNullableJSONArray creates new NullableJSONArray object
func NewNullableJSONArray[T any](val any) (NullableJSONArray[T], error) {
	arr, err := gosql.NewNullableJSONArray[T](val)
	if err != nil {
		return nil, err
	}
	return NullableJSONArray[T](arr), nil
}

// MustNullableJSONArray creates new NullableJSONArray object
func MustNullableJSONArray[T any](val any) NullableJSONArray[T] {
	obj, err := NewNullableJSONArray[T](val)
	if err != nil {
		panic(err)
	}
	return obj
}

// GormDataType gorm common data type
func (NullableJSONArray[T]) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (j NullableJSONArray[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return jsonGormDBDataType(db, true)
}

// String value
func (f NullableJSONArray[T]) String() string {
	return (gosql.NullableJSONArray[T](f)).String()
}

// SetValue of json
func (f *NullableJSONArray[T]) SetValue(value any) error {
	return (*gosql.NullableJSONArray[T])(f).SetValue(value)
}

// Value implements the driver.Valuer interface, []T field
func (f NullableJSONArray[T]) Value() (driver.Value, error) {
	return (gosql.NullableJSONArray[T](f)).Value()
}

// Scan implements the driver.Valuer interface, []T field
func (f *NullableJSONArray[T]) Scan(value any) error {
	return (*gosql.NullableJSONArray[T])(f).Scan(value)
}

// MarshalJSON implements the json.Marshaler
func (f NullableJSONArray[T]) MarshalJSON() ([]byte, error) {
	return (gosql.NullableJSONArray[T](f)).MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableJSONArray[T]) UnmarshalJSON(b []byte) error {
	return (*gosql.NullableJSONArray[T])(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *NullableJSONArray[T]) DecodeValue(v any) error {
	return (*gosql.NullableJSONArray[T])(f).DecodeValue(v)
}
