//
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 – 2018, 2022
//

package gosql

import (
	"database/sql/driver"
	"encoding/json"
)

// JSON field
type JSON[T any] struct {
	value T
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
		f.value = vl
	case *T:
		f.value = *vl
	case nil:
		return ErrNullValueNotAllowed
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
func (f JSON[T]) Value() (_ driver.Value, err error) {
	if v, err := f.MarshalJSON(); err == nil && v != nil {
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
	case nil:
		return ErrNullValueNotAllowed
	default:
		return ErrInvalidScan
	}
	return json.Unmarshal(data, &f.value)
}

// MarshalJSON implements the json.Marshaler
func (f JSON[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *JSON[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, f.value)
}

// DecodeValue implements the gocast.Decoder
func (f *JSON[T]) DecodeValue(v any) error {
	switch val := v.(type) {
	case []byte:
		return f.UnmarshalJSON(val)
	case string:
		return f.UnmarshalJSON([]byte(val))
	default:
		return f.SetValue(v)
	}
}
