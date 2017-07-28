//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package gosql

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"strings"
)

// NullableStringArray implementation
type NullableStringArray []string

// Join array to string
func (f NullableStringArray) Join(sep string) string {
	if nil == f {
		return ""
	}
	return strings.Join(f, sep)
}

// SetArray value
func (f *NullableStringArray) SetArray(arr []string) *NullableStringArray {
	*f = arr
	return f
}

// Value implements the driver.Valuer interface, []string field
func (f NullableStringArray) Value() (driver.Value, error) {
	if nil == f {
		return nil, nil
	}
	return encodeNullableStringArray('{', '}', byte(0), f).String(), nil
}

// Scan implements the driver.Valuer interface, []string field
func (f *NullableStringArray) Scan(value interface{}) error {
	switch value.(type) {
	case []byte:
		*f = decodeNullableStringArray(string(value.([]byte)))
		break
	case string:
		*f = decodeNullableStringArray(value.(string))
		break
	case nil:
		*f = nil
		break
	}
	return nil
}

// MarshalJSON implements the json.Marshaler
func (f NullableStringArray) MarshalJSON() ([]byte, error) {
	if nil == f {
		return []byte("null"), nil
	}
	return json.Marshal([]string(f))
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableStringArray) UnmarshalJSON(b []byte) error {
	if *f = decodeNullableStringArray(string(b)); nil == *f {
		*f = []string{}
	}
	return nil
}

// DecodeValue implements the gocast.Decoder
func (f *NullableStringArray) DecodeValue(v interface{}) error {
	switch val := v.(type) {
	case []string:
		*f = NullableStringArray(val)
	case NullableStringArray:
		*f = val
	case StringArray:
		*f = NullableStringArray(val)
	case []byte:
		*f = NullableStringArray(strings.Split(string(val), ","))
	case string:
		*f = NullableStringArray(strings.Split(val, ","))
	default:
		return ErrInvalidDecodeValue
	}
	return nil
}

// IndexOf array value
func (f NullableStringArray) IndexOf(v string) int {
	if f != nil {
		for i, vl := range f {
			if vl == v {
				return i
			}
		}
	}
	return -1
}

///////////////////////////////////////////////////////////////////////////////

// StringArray implementation
type StringArray NullableStringArray

// Join array to string
func (f StringArray) Join(sep string) string {
	return NullableStringArray(f).Join(sep)
}

// SetArray value
func (f *StringArray) SetArray(arr []string) *StringArray {
	if nil == arr {
		arr = []string{}
	}
	*f = arr
	return f
}

// Value implements the driver.Valuer interface, []string field
func (f StringArray) Value() (driver.Value, error) {
	if nil == f {
		return "{}", ErrNullValueNotAllowed
	}
	return NullableStringArray(f).Value()
}

// Scan implements the driver.Valuer interface, []string field
func (f *StringArray) Scan(value interface{}) error {
	if nil == value {
		return ErrNullValueNotAllowed
	}
	return (*NullableStringArray)(f).Scan(value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *StringArray) UnmarshalJSON(b []byte) error {
	if nil == b {
		return ErrNullValueNotAllowed
	}
	return (*NullableStringArray)(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *StringArray) DecodeValue(v interface{}) error {
	if nil == v {
		return ErrNullValueNotAllowed
	}
	return (*NullableStringArray)(f).DecodeValue(v)
}

// IndexOf array value
func (f StringArray) IndexOf(v string) int {
	return (NullableStringArray)(f).IndexOf(v)
}

///////////////////////////////////////////////////////////////////////////////
/// Helpers
///////////////////////////////////////////////////////////////////////////////

func decodeNullableStringArray(arr string) []string {
	if "null" == arr || "NULL" == arr {
		return nil
	}
	if "{}" == arr {
		return []string{}
	}
	return strings.Split(strings.Trim(arr, "{}"), ",")
}

func encodeNullableStringArray(begin, end, border byte, arr []string) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}
			if byte(0) != border {
				buff.WriteByte(border)
			}
			buff.WriteString(v)
			if byte(0) != border {
				buff.WriteByte(border)
			}
		}
	}

	buff.WriteByte(end)
	return &buff
}
