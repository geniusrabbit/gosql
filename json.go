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
type NullableJSON struct {
	value []byte
}

// NewNullableJSON field object
func NewNullableJSON(v ...interface{}) (jobj *NullableJSON, err error) {
	jobj = &NullableJSON{}
	if len(v) > 0 && v[0] != nil {
		err = jobj.SetValue(v[0])
	}
	return
}

// String value
func (f *NullableJSON) String() string {
	if f == nil || f.value == nil {
		return "null"
	}
	return string(f.value)
}

// SetValue of json
func (f *NullableJSON) SetValue(value interface{}) (err error) {
	f.value, err = json.Marshal(value)
	return
}

// Bytes body
func (f NullableJSON) Bytes() []byte {
	return f.value
}

// Length of the body
func (f NullableJSON) Length() int {
	return len(f.value)
}

// Value implements the driver.Valuer interface, json field interface
func (f NullableJSON) Value() (_ driver.Value, err error) {
	var v []byte
	if v, err := f.MarshalJSON(); nil == err && nil != v {
		return string(v), nil
	}
	return v, err
}

// Scan implements the driver.Valuer interface, json field interface
func (f *NullableJSON) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		value = []byte(v)
	case []byte:
	case nil:
		f.value = nil
		return nil
	default:
		return ErrInvalidScan
	}

	var (
		trg interface{}
		err = json.Unmarshal(value.([]byte), &trg)
	)

	if err == nil {
		f.value = append([]byte{}, value.([]byte)...)
	}
	return err
}

// MarshalJSON implements the json.Marshaler
func (f NullableJSON) MarshalJSON() ([]byte, error) {
	if f.value == nil {
		return []byte("null"), nil
	}
	return f.value, nil
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableJSON) UnmarshalJSON(b []byte) error {
	if len(f.value) > 0 {
		f.value = f.value[:0]
	}
	f.value = append(f.value, b...)
	return nil
}

// DecodeValue implements the gocast.Decoder
func (f *NullableJSON) DecodeValue(v interface{}) (err error) {
	switch val := v.(type) {
	case []byte:
		err = f.UnmarshalJSON(val)
	case string:
		err = f.UnmarshalJSON([]byte(val))
	default:
		err = f.SetValue(v)
	}
	return
}

// UnmarshalTo object
func (f *NullableJSON) UnmarshalTo(v interface{}) (err error) {
	if len(f.value) < 1 {
		return
	}
	return json.Unmarshal(f.value, v)
}

///////////////////////////////////////////////////////////////////////////////

// JSON field
type JSON NullableJSON

// NewJSON field object
func NewJSON(v ...interface{}) (*JSON, error) {
	jobj, err := NewNullableJSON(v...)
	return (*JSON)(jobj), err
}

// Value implements the driver.Valuer interface, json field interface
func (f JSON) Value() (driver.Value, error) {
	if f.value == nil {
		return nil, ErrNullValueNotAllowed
	}
	return NullableJSON(f).MarshalJSON()
}

// Scan implements the driver.Valuer interface, json field interface
func (f *JSON) Scan(value interface{}) error {
	if f == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableJSON)(f).Scan(value)
}

// MarshalJSON Implement json.Marshaler
func (f JSON) MarshalJSON() ([]byte, error) {
	if f.value == nil {
		return []byte("{}"), nil
	}
	return NullableJSON(f).MarshalJSON()
}

// UnmarshalJSON Implement json.Unmarshaller
func (f *JSON) UnmarshalJSON(b []byte) error {
	if b == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableJSON)(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *JSON) DecodeValue(v interface{}) (err error) {
	if v == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableJSON)(f).DecodeValue(v)
}
