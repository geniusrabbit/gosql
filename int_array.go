//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 – 2017, 2020
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
	if f == nil {
		return nil, nil
	}
	return IntArrayEncode('{', '}', f).String(), nil
}

// Scan implements the driver.Valuer interface, []int field
func (f *NullableIntArray) Scan(value interface{}) error {
	if res, err := IntArrayDecode(value); err == nil {
		*f = NullableIntArray(res)
	} else {
		return err
	}
	return nil
}

// MarshalJSON implements the json.Marshaler
func (f NullableIntArray) MarshalJSON() ([]byte, error) {
	if f == nil {
		return []byte("null"), nil
	}
	return IntArrayEncode('[', ']', f).Bytes(), nil
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableIntArray) UnmarshalJSON(b []byte) error {
	res, err := IntArrayDecode(b)
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
		list, err := IntArrayDecode(v)
		if err == nil {
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

// Len of array
func (f NullableIntArray) Len() int {
	return len(f)
}

// IndexOf array value
func (f NullableIntArray) IndexOf(v int) int {
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
func (f NullableIntArray) OneOf(vals []int) bool {
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
func (f NullableIntArray) Ordered() NullableOrderedIntArray {
	f.Sort()
	return NullableOrderedIntArray(f)
}

// Filter current array and create filtered copy
func (f NullableIntArray) Filter(fn func(v int) (int, bool)) (resp NullableIntArray) {
	for _, v := range f {
		if nv, ok := fn(v); ok {
			resp = append(resp, nv)
		}
	}
	return
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
	if err = (*NullableIntArray)(f).Scan(value); err == nil {
		NullableIntArray(*f).Sort()
	}
	return
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *NullableOrderedIntArray) UnmarshalJSON(b []byte) (err error) {
	if err = (*NullableIntArray)(f).UnmarshalJSON(b); err == nil {
		NullableIntArray(*f).Sort()
	}
	return err
}

// DecodeValue implements the gocast.Decoder
func (f *NullableOrderedIntArray) DecodeValue(v interface{}) (err error) {
	if err = (*NullableIntArray)(f).DecodeValue(v); err == nil {
		NullableIntArray(*f).Sort()
	}
	return
}

// Sort ints array
func (f NullableOrderedIntArray) Sort() {
	(NullableIntArray)(f).Sort()
}

// Len of array
func (f NullableOrderedIntArray) Len() int {
	return len(f)
}

// IndexOf array value
func (f NullableOrderedIntArray) IndexOf(v int) int {
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
func (f NullableOrderedIntArray) OneOf(vals []int) bool {
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
func (f NullableOrderedIntArray) Filter(fn func(v int) (int, bool)) (resp NullableOrderedIntArray) {
	for _, v := range f {
		if nv, ok := fn(v); ok {
			resp = append(resp, nv)
		}
	}
	return
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
	if value == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableIntArray)(f).Scan(value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *IntArray) UnmarshalJSON(b []byte) error {
	if b == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableIntArray)(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *IntArray) DecodeValue(v interface{}) error {
	if v == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableIntArray)(f).DecodeValue(v)
}

// Sort ints array
func (f IntArray) Sort() {
	(NullableIntArray)(f).Sort()
}

// Len of array
func (f IntArray) Len() int {
	return len(f)
}

// IndexOf array value
func (f IntArray) IndexOf(v int) int {
	return NullableIntArray(f).IndexOf(v)
}

// OneOf value in array
func (f IntArray) OneOf(vals []int) bool {
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
func (f IntArray) Ordered() OrderedIntArray {
	f.Sort()
	return OrderedIntArray(f)
}

// Filter current array and create filtered copy
func (f IntArray) Filter(fn func(v int) (int, bool)) (resp IntArray) {
	for _, v := range f {
		if nv, ok := fn(v); ok {
			resp = append(resp, nv)
		}
	}
	return
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
	if value == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableOrderedIntArray)(f).Scan(value)
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *OrderedIntArray) UnmarshalJSON(b []byte) (err error) {
	if b == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableOrderedIntArray)(f).UnmarshalJSON(b)
}

// DecodeValue implements the gocast.Decoder
func (f *OrderedIntArray) DecodeValue(v interface{}) error {
	if v == nil {
		return ErrNullValueNotAllowed
	}
	return (*NullableOrderedIntArray)(f).DecodeValue(v)
}

// Sort ints array
func (f OrderedIntArray) Sort() {
	(NullableOrderedIntArray)(f).Sort()
}

// Len of array
func (f OrderedIntArray) Len() int {
	return len(f)
}

// IndexOf array value
func (f OrderedIntArray) IndexOf(v int) int {
	return (NullableOrderedIntArray)(f).IndexOf(v)
}

// OneOf value in array
func (f OrderedIntArray) OneOf(vals []int) bool {
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
func (f OrderedIntArray) Filter(fn func(v int) (int, bool)) (resp OrderedIntArray) {
	for _, v := range f {
		if nv, ok := fn(v); ok {
			resp = append(resp, nv)
		}
	}
	return
}
