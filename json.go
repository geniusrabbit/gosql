//
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 – 2018, 2022, 2025
//

package gosql

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
)

// JSON field
type JSON[T any] struct {
	Data T
}

// NewJSON creates new JSON object
func NewJSON[T any](val any) (*JSON[T], error) {
	var obj JSON[T]
	if err := obj.SetValue(val); err != nil {
		return nil, err
	}
	return &obj, nil
}

// MustJSON creates new JSON object
func MustJSON[T any](val any) *JSON[T] {
	obj, err := NewJSON[T](val)
	if err != nil {
		panic(err)
	}
	return obj
}

// String value
func (f *JSON[T]) String() string {
	if f == nil {
		return "{}"
	}
	data, _ := f.MarshalJSON()
	return string(data)
}

// SetValue of json
func (f *JSON[T]) SetValue(value any) error {
	switch vl := value.(type) {
	case T:
		f.Data = vl
		return nil
	case *T:
		f.Data = *vl
		return nil
	case nil:
		return ErrNullValueNotAllowed
	}

	switch vl := value.(type) {
	case string:
		return f.UnmarshalJSON([]byte(vl))
	case []byte:
		return f.UnmarshalJSON(vl)
	default:
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return f.UnmarshalJSON(data)
	}
}

// Value implements the driver.Valuer interface, json field interface
func (f JSON[T]) Value() (driver.Value, error) {
	v, err := f.MarshalJSON()
	if err == nil && v != nil {
		return string(v), nil
	}
	return nil, err
}

// Scan implements the driver.Valuer interface, json field interface
func (f *JSON[T]) Scan(value any) error {
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
	return f.UnmarshalJSON(data)
}

// MarshalJSON implements the json.Marshaler
func (f JSON[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Data)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *JSON[T]) UnmarshalJSON(data []byte) error {
	f.Data = *new(T)
	if len(data) == 0 {
		return nil
	}
	if data = bytes.TrimSpace(data); len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, &f.Data)
}

// DecodeValue implements the gocast.Decoder
func (f *JSON[T]) DecodeValue(v any) error {
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
