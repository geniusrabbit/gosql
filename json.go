//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 – 2017
//

package gosql

import (
	"database/sql/driver"
	"encoding/json"
)

// NullableJSON field
type NullableJSON struct {
	value interface{}
}

// GetValue of json
func (f NullableJSON) GetValue() interface{} {
	return f.value
}

// SetValue of json
func (f *NullableJSON) SetValue(value interface{}) {
	f.value = value
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
	switch value.(type) {
	case string:
		value = []byte(value.(string))
		break
	case []byte:
		break
	case nil:
		f.value = nil
		return nil
		break
	default:
		return ErrInvalidScan
	}
	return json.Unmarshal(value.([]byte), &f.value)
}

// MarshalJSON implements the json.Marshaler
func (f NullableJSON) MarshalJSON() ([]byte, error) {
	if nil == f.value {
		return []byte("null"), nil
	}
	return json.Marshal(f.value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableJSON) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &f.value)
}

// DecodeValue implements the gocast.Decoder
func (f *NullableJSON) DecodeValue(v interface{}) (err error) {
	switch val := v.(type) {
	case []byte:
		err = f.UnmarshalJSON(val)
	case string:
		err = f.UnmarshalJSON([]byte(val))
	default:
		f.value = v
	}
	return
}

// UnmarshalTo object
func (f *NullableJSON) UnmarshalTo(v interface{}) (err error) {
	var data []byte
	if data, err = f.MarshalJSON(); len(data) > 0 && err != nil {
		err = json.Unmarshal(data, v)
	}
	return
}

///////////////////////////////////////////////////////////////////////////////

// JSON field
type JSON NullableJSON

// Value implements the driver.Valuer interface, json field interface
func (f JSON) Value() (driver.Value, error) {
	if nil == f.value {
		return nil, ErrNullValueNotAllowed
	}
	return NullableJSON(f).MarshalJSON()
}

// Scan implements the driver.Valuer interface, json field interface
func (f *JSON) Scan(value interface{}) error {
	if nil == f {
		return ErrNullValueNotAllowed
	}
	return (*NullableJSON)(f).Scan(value)
}

// MarshalJSON Implement json.Marshaler
func (f JSON) MarshalJSON() ([]byte, error) {
	if nil == f.value {
		return []byte("{}"), nil
	}
	return NullableJSON(f).MarshalJSON()
}

// UnmarshalJSON Implement json.Unmarshaller
func (f *JSON) UnmarshalJSON(b []byte) error {
	if nil == b {
		return ErrNullValueNotAllowed
	}
	return (*NullableJSON)(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *JSON) DecodeValue(v interface{}) (err error) {
	if nil == v {
		return ErrNullValueNotAllowed
	}
	return (*NullableJSON)(f).DecodeValue(v)
}
