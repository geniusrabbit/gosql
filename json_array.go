package gosql

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
)

type JSONArray[T any] []T

// NewJSONArray creates new JSONArray object
func NewJSONArray[T any](val any) (JSONArray[T], error) {
	var obj JSONArray[T]
	if err := obj.SetValue(val); err != nil {
		return nil, err
	}
	return obj, nil
}

// MustJSONArray creates new JSONArray object
func MustJSONArray[T any](val any) JSONArray[T] {
	obj, err := NewJSONArray[T](val)
	if err != nil {
		panic(err)
	}
	return obj
}

// String value
func (f JSONArray[T]) String() string {
	data, _ := f.MarshalJSON()
	return string(data)
}

// SetValue of json
func (f *JSONArray[T]) SetValue(value any) error {
	switch vl := value.(type) {
	case []T:
		*f = append(*f, vl...)
	case *[]T:
		*f = append(*f, *vl...)
	case nil:
		return ErrNullValueNotAllowed
	default:
		switch vl := value.(type) {
		case string:
			return f.UnmarshalJSON([]byte(vl))
		case []byte:
			return f.UnmarshalJSON(vl)
		case json.RawMessage:
			return f.UnmarshalJSON(vl)
		default:
			return ErrInvalidSetValue
		}
	}
	return nil
}

// Value implements the driver.Valuer interface, []T field
func (f JSONArray[T]) Value() (driver.Value, error) {
	v, err := f.MarshalJSON()
	if err == nil && len(v) > 1 {
		return string(v), nil
	}
	return nil, err
}

// Scan implements the driver.Valuer interface, []T field
func (f *JSONArray[T]) Scan(value any) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	case json.RawMessage:
		data = v
	case nil:
		return ErrNullValueNotAllowed
	default:
		return ErrInvalidScan
	}
	if data = bytes.TrimSpace(data); len(data) == 0 {
		*f = nil
		return nil
	}
	return f.UnmarshalJSON(data)
}

// MarshalJSON implements the json.Marshaler
func (f JSONArray[T]) MarshalJSON() ([]byte, error) {
	if len(f) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal([]T(f))
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *JSONArray[T]) UnmarshalJSON(b []byte) error {
	var res []T
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*f = append((*f)[:0], res...)
	return nil
}

// DecodeValue implements the gocast.Decoder
func (f *JSONArray[T]) DecodeValue(v any) error {
	switch val := v.(type) {
	case []T:
		*f = val
	case *[]T:
		*f = *val
	case JSONArray[T]:
		*f = val
	default:
		switch val := v.(type) {
		case []byte:
			return f.UnmarshalJSON(val)
		case string:
			return f.UnmarshalJSON([]byte(val))
		case json.RawMessage:
			return f.UnmarshalJSON(val)
		default:
			return f.SetValue(v)
		}
	}
	return nil
}
