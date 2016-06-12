//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package gosql

import (
	"database/sql/driver"
	"sort"

	"github.com/cznic/sortutil"
)

///////////////////////////////////////////////////////////////////////////////

// NullableUintArray type of field
type NullableUintArray []uint

// Value implements the driver.Valuer interface, []int field
func (f NullableUintArray) Value() (driver.Value, error) {
	if nil == f {
		return nil, nil
	}
	return encodeUintArray('{', '}', f).String(), nil
}

// Scan implements the driver.Valuer interface, []int field
func (f *NullableUintArray) Scan(value interface{}) error {
	if res, err := decodeUintArray(value); nil == err {
		*f = NullableUintArray(res)
	} else {
		return err
	}
	return nil
}

// MarshalJSON implements the json.Marshaler
func (f NullableUintArray) MarshalJSON() ([]byte, error) {
	if nil == f {
		return []byte("null"), nil
	}
	return encodeUintArray('[', ']', f).Bytes(), nil
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableUintArray) UnmarshalJSON(b []byte) error {
	res, err := decodeUintArray(b)
	*f = NullableUintArray(res)
	return err
}

// DecodeValue implements the gocast.Decoder
func (f *NullableUintArray) DecodeValue(v interface{}) error {
	switch val := v.(type) {
	case []uint:
		*f = NullableUintArray(val)
	case NullableUintArray:
		*f = val
	case UintArray:
		*f = NullableUintArray(val)
	case []byte, string:
		list, err := decodeUintArray(v)
		if nil == err {
			*f = NullableUintArray(list)
		}
		return err
	default:
		return ErrInvalidDecodeValue
	}
	return nil
}

// Sort ints array
func (f NullableUintArray) Sort() {
	sortutil.UintSlice(f).Sort()
}

// IndexOf array value
func (f NullableUintArray) IndexOf(v uint) int {
	if nil != f {
		for i, vl := range f {
			if vl == v {
				return i
			}
		}
	}
	return -1
}

///////////////////////////////////////////////////////////////////////////////

// NullableOrderedUintArray type of field
type NullableOrderedUintArray NullableUintArray

// Scan implements the driver.Valuer interface, []int field
func (f *NullableOrderedUintArray) Scan(value interface{}) (err error) {
	if err = (*NullableUintArray)(f).Scan(value); nil == err {
		NullableUintArray(*f).Sort()
	}
	return
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableOrderedUintArray) UnmarshalJSON(b []byte) (err error) {
	if err = (*NullableUintArray)(f).UnmarshalJSON(b); nil == err {
		NullableUintArray(*f).Sort()
	}
	return
}

// DecodeValue implements the gocast.Decoder
func (f *NullableOrderedUintArray) DecodeValue(v interface{}) (err error) {
	if err = (*NullableUintArray)(f).DecodeValue(v); nil == err {
		NullableUintArray(*f).Sort()
	}
	return
}

// IndexOf array value
func (f NullableOrderedUintArray) IndexOf(v uint) int {
	if nil != f {
		i := sort.Search(len(f), func(i int) bool { return f[i] >= v })
		if i >= 0 && i < len(f) && f[i] == v {
			return i
		}
	}
	return -1
}

///////////////////////////////////////////////////////////////////////////////

// UintArray type of field
type UintArray NullableUintArray

// Scan implements the driver.Valuer interface, []int field
func (f *UintArray) Scan(value interface{}) error {
	if nil == value {
		return ErrNullValueNotAllowed
	}
	return (*NullableUintArray)(f).Scan(value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *UintArray) UnmarshalJSON(b []byte) error {
	if nil == b {
		return ErrNullValueNotAllowed
	}
	return (*NullableUintArray)(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *UintArray) DecodeValue(v interface{}) error {
	if nil == v {
		return ErrNullValueNotAllowed
	}
	return (*NullableUintArray)(f).DecodeValue(v)
}

///////////////////////////////////////////////////////////////////////////////

// OrderedUintArray type of field
type OrderedUintArray NullableUintArray

// Scan implements the driver.Valuer interface, []int field
func (f *OrderedUintArray) Scan(value interface{}) (err error) {
	if nil == value {
		return ErrNullValueNotAllowed
	}
	if err = (*NullableUintArray)(f).Scan(value); nil == err {
		NullableUintArray(*f).Sort()
	}
	return
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *OrderedUintArray) UnmarshalJSON(b []byte) (err error) {
	if nil == b {
		return ErrNullValueNotAllowed
	}
	if err = (*NullableUintArray)(f).UnmarshalJSON(b); nil == err {
		NullableUintArray(*f).Sort()
	}
	return err
}

// DecodeValue implements the gocast.Decoder
func (f *OrderedUintArray) DecodeValue(v interface{}) (err error) {
	if nil == v {
		return ErrNullValueNotAllowed
	}
	if err = (*NullableUintArray)(f).DecodeValue(v); nil == err {
		NullableUintArray(*f).Sort()
	}
	return
}
