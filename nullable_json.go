//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 – 2018
//

package gosql

import (
	"database/sql/driver"
	"encoding/json"
)

// NullableJSON field
type NullableJSON[T any] struct {
	value *T
}

// String value
func (f *NullableJSON[T]) String() string {
	if f == nil || f.value == nil {
		return "null"
	}
	data, _ := f.MarshalJSON()
	return string(data)
}

// SetValue of json
func (f *NullableJSON[T]) SetValue(value any) error {
	switch vl := value.(type) {
	case T:
		if f.value == nil {
			f.value = new(T)
		}
		*f.value = vl
	case *T:
		if f.value == nil {
			f.value = new(T)
		}
		*f.value = *vl
	case nil:
		f.value = nil
	case string:
		return f.UnmarshalJSON([]byte(vl))
	case []byte:
		return f.UnmarshalJSON(vl)
	default:
		return ErrInvalidDecodeValue
	}
	return nil
}

// Value implements the driver.Valuer interface, json field interface
func (f NullableJSON[T]) Value() (_ driver.Value, err error) {
	if v, err := f.MarshalJSON(); err == nil && v != nil {
		return string(v), nil
	}
	return nil, err
}

// Scan implements the driver.Valuer interface, json field interface
func (f *NullableJSON[T]) Scan(value any) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	case nil:
		f.value = nil
		return nil
	default:
		return ErrInvalidScan
	}
	if f.value == nil {
		f.value = new(T)
	}
	return json.Unmarshal(data, f.value)
}

// MarshalJSON implements the json.Marshaler
func (f NullableJSON[T]) MarshalJSON() ([]byte, error) {
	if f.value == nil {
		return []byte("null"), nil
	}
	return json.Marshal(f.value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableJSON[T]) UnmarshalJSON(data []byte) error {
	if f.value == nil {
		f.value = new(T)
	}
	return json.Unmarshal(data, f.value)
}

// DecodeValue implements the gocast.Decoder
func (f *NullableJSON[T]) DecodeValue(v any) error {
	switch val := v.(type) {
	case []byte:
		return f.UnmarshalJSON(val)
	case string:
		return f.UnmarshalJSON([]byte(val))
	default:
		return f.SetValue(v)
	}
}
