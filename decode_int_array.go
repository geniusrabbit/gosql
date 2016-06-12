//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package gosql

import (
	"bytes"
	"strconv"
	"strings"
)

func decodeIntArray(data interface{}) (result []int, err error) {
	var arr string
	switch data.(type) {
	case []byte:
		arr = string(data.([]byte))
		break
	case string:
		arr = data.(string)
		break
	case nil:
		return nil, nil
	default:
		return nil, ErrInvalidScan
	}

	if "null" == arr || "NULL" == arr {
		return nil, nil
	}

	result = make([]int, 0)
	if "{}" != arr && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}"), ","); nil != vals {
			for _, cid := range vals {
				var v int64
				if v, err = strconv.ParseInt(strings.Trim(cid, "'\""), 10, 64); nil == err {
					result = append(result, int(v))
				} else {
					break
				}
			} // end for
		}
	}
	return
}

func encodeIntArray(begin, end byte, arr []int) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}
			buff.WriteString(strconv.Itoa(v))
		}
	}

	buff.WriteByte(end)
	return &buff
}

///////////////////////////////////////////////////////////////////////////////

func decodeInt8Array(data interface{}) (result []int8, err error) {
	var arr string
	switch data.(type) {
	case []byte:
		arr = string(data.([]byte))
		break
	case string:
		arr = data.(string)
		break
	case nil:
		return nil, nil
		break
	default:
		return nil, ErrInvalidScan
	}

	if "null" == arr || "NULL" == arr {
		return nil, nil
	}

	result = make([]int8, 0)
	if "{}" != arr && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}"), ","); nil != vals {
			for _, cid := range vals {
				var v int64
				if v, err = strconv.ParseInt(strings.Trim(cid, "'\""), 10, 64); nil == err {
					result = append(result, int8(v))
				} else {
					break
				}
			} // end for
		}
	}
	return
}

func encodeInt8Array(begin, end byte, arr []int8) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}
			buff.WriteString(strconv.Itoa(int(v)))
		}
	}

	buff.WriteByte(end)
	return &buff
}

///////////////////////////////////////////////////////////////////////////////

func decodeInt16Array(data interface{}) (result []int16, err error) {
	var arr string
	switch data.(type) {
	case []byte:
		arr = string(data.([]byte))
		break
	case string:
		arr = data.(string)
		break
	case nil:
		return nil, nil
		break
	default:
		return nil, ErrInvalidScan
	}

	if "null" == arr || "NULL" == arr {
		return nil, nil
	}

	result = make([]int16, 0)
	if "{}" != arr && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}"), ","); nil != vals {
			for _, cid := range vals {
				var v int64
				if v, err = strconv.ParseInt(strings.Trim(cid, "'\""), 10, 64); nil == err {
					result = append(result, int16(v))
				} else {
					break
				}
			} // end for
		}
	}
	return
}

func encodeInt16Array(begin, end byte, arr []int32) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}
			buff.WriteString(strconv.Itoa(int(v)))
		}
	}

	buff.WriteByte(end)
	return &buff
}

///////////////////////////////////////////////////////////////////////////////

func decodeInt32Array(data interface{}) (result []int32, err error) {
	var arr string
	switch data.(type) {
	case []byte:
		arr = string(data.([]byte))
		break
	case string:
		arr = data.(string)
		break
	case nil:
		return nil, nil
		break
	default:
		return nil, ErrInvalidScan
	}

	if "null" == arr || "NULL" == arr {
		return nil, nil
	}

	result = make([]int32, 0)
	if "{}" != arr && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}"), ","); nil != vals {
			for _, cid := range vals {
				var v int64
				if v, err = strconv.ParseInt(strings.Trim(cid, "'\""), 10, 64); nil == err {
					result = append(result, int32(v))
				} else {
					break
				}
			} // end for
		}
	}
	return
}

func encodeInt32Array(begin, end byte, arr []int32) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}
			buff.WriteString(strconv.Itoa(int(v)))
		}
	}

	buff.WriteByte(end)
	return &buff
}

///////////////////////////////////////////////////////////////////////////////

func decodeInt64Array(data interface{}) (result []int64, err error) {
	var arr string
	switch data.(type) {
	case []byte:
		arr = string(data.([]byte))
		break
	case string:
		arr = data.(string)
		break
	case nil:
		return nil, nil
		break
	default:
		return nil, ErrInvalidScan
	}

	if "null" == arr || "NULL" == arr {
		return nil, nil
	}

	result = make([]int64, 0)
	if "{}" != arr && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}"), ","); nil != vals {
			for _, cid := range vals {
				var v int64
				if v, err = strconv.ParseInt(strings.Trim(cid, "'\""), 10, 64); nil == err {
					result = append(result, v)
				} else {
					break
				}
			} // end for
		}
	}
	return
}

func encodeInt64Array(begin, end byte, arr []int64) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}
			buff.WriteString(strconv.FormatInt(v, 10))
		}
	}

	buff.WriteByte(end)
	return &buff
}
