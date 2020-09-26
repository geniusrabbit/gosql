//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 – 2017, 2020
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
	if f == nil {
		return nil, nil
	}
	return UintArrayEncode('{', '}', f).String(), nil
}

// Scan implements the driver.Valuer interface, []int field
func (f *NullableUintArray) Scan(value interface{}) error {
	if res, err := UintArrayDecode(value); err == nil {
		*f = NullableUintArray(res)
	} else {
		return err
	}
	return nil
}

// MarshalJSON implements the json.Marshaler
func (f NullableUintArray) MarshalJSON() ([]byte, error) {
	if f == nil {
		return []byte("null"), nil
	}
	return UintArrayEncode('[', ']', f).Bytes(), nil
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableUintArray) UnmarshalJSON(b []byte) error {
	res, err := UintArrayDecode(b)
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
		list, err := UintArrayDecode(v)
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

// Len of array
func (f NullableUintArray) Len() int {
	return len(f)
}

// IndexOf array value
func (f NullableUintArray) IndexOf(v uint) int {
	if f == nil {
		return -1
	}
	for i, vl := range f {
		if vl == v {
			return i
		}
	}
	return -1
}

// OneOf value in array
func (f NullableUintArray) OneOf(vals []uint) bool {
	if len(f) < 1 || len(vals) < 1 {
		return false
	}
	for _, v := range vals {
		if f.IndexOf(v) != -1 {
			return true
		}
	}
	return false
}

// Ordered object
func (f NullableUintArray) Ordered() NullableOrderedUintArray {
	f.Sort()
	return NullableOrderedUintArray(f)
}

// Filter current array and create filtered copy
func (f NullableUintArray) Filter(fn func(v uint) (uint, bool)) (resp NullableUintArray) {
	for _, v := range f {
		if nv, ok := fn(v); ok {
			resp = append(resp, nv)
		}
	}
	return
}

///////////////////////////////////////////////////////////////////////////////

// NullableOrderedUintArray type of field
type NullableOrderedUintArray NullableUintArray

// Value implements the driver.Valuer interface, []int field
func (f NullableOrderedUintArray) Value() (driver.Value, error) {
	return NullableUintArray(f).Value()
}

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

// Sort ints array
func (f NullableOrderedUintArray) Sort() {
	(NullableUintArray)(f).Sort()
}

// Len of array
func (f NullableOrderedUintArray) Len() int {
	return len(f)
}

// IndexOf array value
func (f NullableOrderedUintArray) IndexOf(v uint) int {
	if f == nil {
		return -1
	}
	i := sort.Search(len(f), func(i int) bool { return f[i] >= v })
	if i >= 0 && i < len(f) && f[i] == v {
		return i
	}
	return -1
}

// OneOf value in array
func (f NullableOrderedUintArray) OneOf(vals []uint) bool {
	if len(f) < 1 || len(vals) < 1 {
		return false
	}
	for _, v := range vals {
		if f.IndexOf(v) != -1 {
			return true
		}
	}
	return false
}

// Filter current array and create filtered copy
func (f NullableOrderedUintArray) Filter(fn func(v uint) (uint, bool)) (resp NullableOrderedUintArray) {
	for _, v := range f {
		if nv, ok := fn(v); ok {
			resp = append(resp, nv)
		}
	}
	return
}

///////////////////////////////////////////////////////////////////////////////

// UintArray type of field
type UintArray NullableUintArray

// Value implements the driver.Valuer interface, []int field
func (f UintArray) Value() (driver.Value, error) {
	if f == nil {
		return "{}", nil
	}
	return NullableUintArray(f).Value()
}

// Scan implements the driver.Valuer interface, []int field
func (f *UintArray) Scan(value interface{}) error {
	if value == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableUintArray)(f).Scan(value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *UintArray) UnmarshalJSON(b []byte) error {
	if b == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableUintArray)(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *UintArray) DecodeValue(v interface{}) error {
	if v == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableUintArray)(f).DecodeValue(v)
}

// Sort ints array
func (f UintArray) Sort() {
	(NullableUintArray)(f).Sort()
}

// Len of array
func (f UintArray) Len() int {
	return len(f)
}

// IndexOf array value
func (f UintArray) IndexOf(v uint) int {
	return NullableUintArray(f).IndexOf(v)
}

// OneOf value in array
func (f UintArray) OneOf(vals []uint) bool {
	if len(f) < 1 || len(vals) < 1 {
		return false
	}
	for _, v := range vals {
		if f.IndexOf(v) != -1 {
			return true
		}
	}
	return false
}

// Ordered object
func (f UintArray) Ordered() OrderedUintArray {
	f.Sort()
	return OrderedUintArray(f)
}

// Filter current array and create filtered copy
func (f UintArray) Filter(fn func(v uint) (uint, bool)) (resp UintArray) {
	for _, v := range f {
		if nv, ok := fn(v); ok {
			resp = append(resp, nv)
		}
	}
	return
}

///////////////////////////////////////////////////////////////////////////////

// OrderedUintArray type of field
type OrderedUintArray NullableOrderedUintArray

// Value implements the driver.Valuer interface, []int field
func (f OrderedUintArray) Value() (driver.Value, error) {
	return UintArray(f).Value()
}

// Scan implements the driver.Valuer interface, []int field
func (f *OrderedUintArray) Scan(value interface{}) error {
	if value == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableOrderedUintArray)(f).Scan(value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *OrderedUintArray) UnmarshalJSON(b []byte) (err error) {
	if b == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableOrderedUintArray)(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *OrderedUintArray) DecodeValue(v interface{}) error {
	if v == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableOrderedUintArray)(f).DecodeValue(v)
}

// Sort ints array
func (f OrderedUintArray) Sort() {
	(NullableOrderedUintArray)(f).Sort()
}

// Len of array
func (f OrderedUintArray) Len() int {
	return len(f)
}

// IndexOf array value
func (f OrderedUintArray) IndexOf(v uint) int {
	return (NullableOrderedUintArray)(f).IndexOf(v)
}

// OneOf value in array
func (f OrderedUintArray) OneOf(vals []uint) bool {
	if len(f) < 1 || len(vals) < 1 {
		return false
	}
	for _, v := range vals {
		if f.IndexOf(v) != -1 {
			return true
		}
	}
	return false
}

// Filter current array and create filtered copy
func (f OrderedUintArray) Filter(fn func(v uint) (uint, bool)) (resp OrderedUintArray) {
	for _, v := range f {
		if nv, ok := fn(v); ok {
			resp = append(resp, nv)
		}
	}
	return
}
