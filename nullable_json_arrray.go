package gosql

import "database/sql/driver"

// NullableJSONArray object
type NullableJSONArray[T any] JSONArray[T]

// NewNullableJSONArray creates new NullableJSONArray object
func NewNullableJSONArray[T any](val any) (NullableJSONArray[T], error) {
	arr, err := NewJSONArray[T](val)
	return NullableJSONArray[T](arr), err
}

// MustNullableJSONArray creates new NullableJSONArray object
func MustNullableJSONArray[T any](val any) NullableJSONArray[T] {
	obj := MustJSONArray[T](val)
	return NullableJSONArray[T](obj)
}

// String value
func (f NullableJSONArray[T]) String() string {
	if f == nil {
		return "null"
	}
	return JSONArray[T](f).String()
}

// SetValue of json
func (f *NullableJSONArray[T]) SetValue(value any) error {
	if value == nil {
		*f = nil
		return nil
	}
	return (*JSONArray[T])(f).SetValue(value)
}

// Value of the object
func (f NullableJSONArray[T]) Value() (driver.Value, error) {
	if len(f) == 0 {
		return nil, nil
	}
	return JSONArray[T](f).Value()
}

// Scan value from database
func (f *NullableJSONArray[T]) Scan(value interface{}) error {
	if value == nil {
		*f = nil
		return nil
	}
	return (*JSONArray[T])(f).Scan(value)
}

// MarshalJSON data
func (f NullableJSONArray[T]) MarshalJSON() ([]byte, error) {
	if f == nil {
		return []byte("null"), nil
	}
	return JSONArray[T](f).MarshalJSON()
}

// UnmarshalJSON data
func (f *NullableJSONArray[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		*f = nil
		return nil
	}
	return (*JSONArray[T])(f).UnmarshalJSON(data)
}

// DecodeValue of the object
func (f *NullableJSONArray[T]) DecodeValue(v any) error {
	if f == nil {
		return nil
	}
	return (*JSONArray[T])(f).DecodeValue(v)
}
