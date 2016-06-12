//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
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

///////////////////////////////////////////////////////////////////////////////

// NullableOrderedIntArray type of field
type NullableOrderedIntArray NullableIntArray

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

///////////////////////////////////////////////////////////////////////////////

// OrderedIntArray type of field
type OrderedIntArray NullableIntArray

// Scan implements the driver.Valuer interface, []int field
func (f *OrderedIntArray) Scan(value interface{}) (err error) {
	if nil == value {
		return ErrNullValueNotAllowed
	}
	if err = (*NullableIntArray)(f).Scan(value); nil == err {
		NullableIntArray(*f).Sort()
	}
	return
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *OrderedIntArray) UnmarshalJSON(b []byte) (err error) {
	if nil == b {
		return ErrNullValueNotAllowed
	}
	if err = (*NullableIntArray)(f).UnmarshalJSON(b); nil == err {
		NullableIntArray(*f).Sort()
	}
	return err
}
