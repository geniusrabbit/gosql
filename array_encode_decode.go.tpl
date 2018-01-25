package gosql

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/cheekybits/genny/generic"
)

type GenType generic.Type

// GenTypeArrayDecode decodes array of type GenType
func GenTypeArrayDecode(data interface{}) (result []GenType, err error) {
	var (
		arr string
		t   GenType
	)
	switch data.(type) {
	case []byte:
		arr = string(data.([]byte))
	case string:
		arr = data.(string)
	case []GenType:
		return data.([]GenType), nil
	case nil:
		return nil, nil
	default:
		return nil, ErrInvalidScan
	}

	if "null" == arr || "NULL" == arr {
		return nil, nil
	}

	result = make([]GenType, 0)
	if arr != "{}" && arr != "[]" && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}[]"), ","); vals != nil {
			for _, cid := range vals {
				var br = false

				switch interface{}(t).(type) {
				case int, int8, int16, int32, int64:
					var v int64
					if v, err = strconv.ParseInt(strings.Trim(cid, "'\""), 10, 64); err == nil {
						result = append(result, GenType(v))
					} else {
						br = true
					}
				case uint, uint8, uint16, uint32, uint64:
					var v uint64
					if v, err = strconv.ParseUint(strings.Trim(cid, "'\""), 10, 64); nil == err {
						result = append(result, GenType(v))
					} else {
						br = true
					}
				case float32, float64:
					var v float64
					if v, err = strconv.ParseFloat(strings.Trim(cid, "'\""), 64); nil == err {
						result = append(result, GenType(v))
					} else {
						br = true
					}
				}

				if br {
					break
				}

			} // end for
		}
	}
	return
}

// GenTypeArrayEncode encodes array of type GenType
func GenTypeArrayEncode(begin, end byte, arr []GenType) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}

			switch interface{}(v).(type) {
			case int, int8, int16, int32, int64:
				buff.WriteString(strconv.Itoa(int(v)))
			case uint, uint8, uint16, uint32, uint64:
				buff.WriteString(strconv.FormatUint(uint64(v), 10))
			case float32, float64:
				buff.WriteString(strconv.FormatFloat(float64(v), 'G', -1, 64))
			}

		}
	}

	buff.WriteByte(end)
	return &buff
}
