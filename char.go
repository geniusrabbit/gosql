//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package gosql

import "database/sql/driver"

// Char type of field
type Char rune

// Value implements the driver.Valuer interface, char field
func (f Char) Value() (driver.Value, error) {
	if 0 == f {
		return " ", nil
	}
	return string(f), nil
}

// Scan implements the sql.Scanner interface, char field
func (f *Char) Scan(value interface{}) (err error) {
	*f, err = decodeChar(value)
	return
}

// MarshalJSON implements the json.Marshaler
func (f Char) MarshalJSON() ([]byte, error) {
	if 0 == f {
		return []byte("\" \""), nil
	}
	return []byte{'"', byte(f), '"'}, nil
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *Char) UnmarshalJSON(b []byte) (err error) {
	*f, err = decodeChar(b)
	return
}

///////////////////////////////////////////////////////////////////////////////
/// Helpers
///////////////////////////////////////////////////////////////////////////////

func decodeChar(value interface{}) (Char, error) {
	if nil == value {
		return Char(0), ErrNullValueNotAllowed
	}

	switch v := value.(type) {
	case []byte:
		if len(v) > 0 {
			return Char(v[0]), nil
		}
		break
	case string:
		if len(v) > 0 {
			return Char(v[0]), nil
		}
		break
	}
	return Char(0), ErrInvalidScan
}
