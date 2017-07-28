//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 – 2017
//

package gosql

import (
	"database/sql/driver"
	"sort"
)

///////////////////////////////////////////////////////////////////////////////

// NullableIntArray type of field
type NullableIntArray []int

// Value implements the driver.Valuer interface, []int field
func (f NullableIntArray) Value() (driver.Value, error) {
	if nil == f {
		return nil, nil
	}
	return encodeIntArray('{', '}', f).String(), nil
}

// Scan implements the driver.Valuer interface, []int field
func (f *NullableIntArray) Scan(value interface{}) error {
	if res, err := decodeIntArray(value); nil == err {
		*f = NullableIntArray(res)
	} else {
		return err
	}
	return nil
}

// MarshalJSON implements the json.Marshaler
func (f NullableIntArray) MarshalJSON() ([]byte, error) {
	if nil == f {
		return []byte("null"), nil
	}
	return encodeIntArray('[', ']', f).Bytes(), nil
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableIntArray) UnmarshalJSON(b []byte) error {
	res, err := decodeIntArray(b)
	*f = NullableIntArray(res)
	return err
}

// DecodeValue implements the gocast.Decoder
func (f *NullableIntArray) DecodeValue(v interface{}) error {
	switch val := v.(type) {
	case []int:
		*f = NullableIntArray(val)
	case NullableIntArray:
		*f = val
	case IntArray:
		*f = NullableIntArray(val)
	case []byte, string:
		list, err := decodeIntArray(v)
		if nil == err {
			*f = NullableIntArray(list)
		}
		return err
	default:
		return ErrInvalidDecodeValue
	}
	return nil
}

// Sort ints array
func (f NullableIntArray) Sort() {
	sort.Ints(f)
}

// IndexOf array value
func (f NullableIntArray) IndexOf(v int) int {
	if nil != f {
		for i, vl := range f {
			if vl == v {
				return i
			}
		}
	}
	return -1
}

// Ordered object
func (f NullableIntArray) Ordered() NullableOrderedIntArray {
	f.Sort()
	return NullableOrderedIntArray(f)
}

///////////////////////////////////////////////////////////////////////////////

// NullableOrderedIntArray type of field
type NullableOrderedIntArray NullableIntArray

// Value implements the driver.Valuer interface, []int field
func (f NullableOrderedIntArray) Value() (driver.Value, error) {
	return NullableIntArray(f).Value()
}

// Scan implements the driver.Valuer interface, []int field
func (f *NullableOrderedIntArray) Scan(value interface{}) (err error) {
	if err = (*NullableIntArray)(f).Scan(value); nil == err {
		NullableIntArray(*f).Sort()
	}
	return
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableOrderedIntArray) UnmarshalJSON(b []byte) (err error) {
	if err = (*NullableIntArray)(f).UnmarshalJSON(b); nil == err {
		NullableIntArray(*f).Sort()
	}
	return err
}

// DecodeValue implements the gocast.Decoder
func (f *NullableOrderedIntArray) DecodeValue(v interface{}) (err error) {
	if err = (*NullableIntArray)(f).DecodeValue(v); nil == err {
		NullableIntArray(*f).Sort()
	}
	return
}

// Sort ints array
func (f NullableOrderedIntArray) Sort() {
	(NullableIntArray)(f).Sort()
}

// IndexOf array value
func (f NullableOrderedIntArray) IndexOf(v int) int {
	if nil != f {
		i := sort.Search(len(f), func(i int) bool { return f[i] >= v })
		if i >= 0 && i < len(f) && f[i] == v {
			return i
		}
	}
	return -1
}

///////////////////////////////////////////////////////////////////////////////

// IntArray type of field
type IntArray NullableIntArray

// Value implements the driver.Valuer interface, []int field
func (f IntArray) Value() (driver.Value, error) {
	if f == nil {
		return "{}", nil
	}
	return NullableIntArray(f).Value()
}

// Scan implements the driver.Valuer interface, []int field
func (f *IntArray) Scan(value interface{}) error {
	if nil == value {
		return ErrNullValueNotAllowed
	}
	return (*NullableIntArray)(f).Scan(value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *IntArray) UnmarshalJSON(b []byte) error {
	if nil == b {
		return ErrNullValueNotAllowed
	}
	return (*NullableIntArray)(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *IntArray) DecodeValue(v interface{}) error {
	if nil == v {
		return ErrNullValueNotAllowed
	}
	return (*NullableIntArray)(f).DecodeValue(v)
}

// Sort ints array
func (f IntArray) Sort() {
	(NullableIntArray)(f).Sort()
}

// IndexOf array value
func (f IntArray) IndexOf(v int) int {
	return NullableIntArray(f).IndexOf(v)
}

// Ordered object
func (f IntArray) Ordered() OrderedIntArray {
	f.Sort()
	return OrderedIntArray(f)
}

///////////////////////////////////////////////////////////////////////////////

// OrderedIntArray type of field
type OrderedIntArray NullableOrderedIntArray

// Value implements the driver.Valuer interface, []int field
func (f OrderedIntArray) Value() (driver.Value, error) {
	return IntArray(f).Value()
}

// Scan implements the driver.Valuer interface, []int field
func (f *OrderedIntArray) Scan(value interface{}) (err error) {
	if nil == value {
		return ErrNullValueNotAllowed
	}
	return (*NullableOrderedIntArray)(f).Scan(value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *OrderedIntArray) UnmarshalJSON(b []byte) (err error) {
	if nil == b {
		return ErrNullValueNotAllowed
	}
	return (*NullableOrderedIntArray)(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *OrderedIntArray) DecodeValue(v interface{}) error {
	if nil == v {
		return ErrNullValueNotAllowed
	}
	return (*NullableOrderedIntArray)(f).DecodeValue(v)
}

// Sort ints array
func (f OrderedIntArray) Sort() {
	(NullableOrderedIntArray)(f).Sort()
}

// IndexOf array value
func (f OrderedIntArray) IndexOf(v int) int {
	return (NullableOrderedIntArray)(f).IndexOf(v)
}
