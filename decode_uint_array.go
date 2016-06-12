//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package gosql

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func decodeUintArray(data interface{}) (result []uint, err error) {
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

	result = make([]uint, 0)
	if "{}" != arr && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}"), ","); nil != vals {
			fmt.Println(">>>>>>>>", vals)
			for _, cid := range vals {
				var v uint64
				if v, err = strconv.ParseUint(strings.Trim(cid, "'\""), 10, 64); nil == err {
					result = append(result, uint(v))
				} else {
					break
				}
			} // end for
		}
	}
	return
}

func encodeUintArray(begin, end byte, arr []uint) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}
			buff.WriteString(strconv.FormatUint(uint64(v), 10))
		}
	}

	buff.WriteByte(end)
	return &buff
}

///////////////////////////////////////////////////////////////////////////////

func decodeUint8Array(data interface{}) (result []uint8, err error) {
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

	result = make([]uint8, 0)
	if "{}" != arr && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}"), ","); nil != vals {
			for _, cid := range vals {
				var v uint64
				if v, err = strconv.ParseUint(strings.Trim(cid, "'\""), 10, 64); nil == err {
					result = append(result, uint8(v))
				} else {
					break
				}
			} // end for
		}
	}
	return
}

func encodeUint8Array(begin, end byte, arr []int8) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr && len(arr) > 0 {
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

func decodeUint16Array(data interface{}) (result []uint16, err error) {
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

	result = make([]uint16, 0)
	if "{}" != arr && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}"), ","); nil != vals {
			for _, cid := range vals {
				var v uint64
				if v, err = strconv.ParseUint(strings.Trim(cid, "'\""), 10, 64); nil == err {
					result = append(result, uint16(v))
				} else {
					break
				}
			} // end for
		}
	}
	return
}

func encodeUint16Array(begin, end byte, arr []uint32) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}
			buff.WriteString(strconv.FormatUint(uint64(v), 10))
		}
	}

	buff.WriteByte(end)
	return &buff
}

///////////////////////////////////////////////////////////////////////////////

func decodeUint32Array(data interface{}) (result []uint32, err error) {
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

	result = make([]uint32, 0)
	if "{}" != arr && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}"), ","); nil != vals {
			for _, cid := range vals {
				var v uint64
				if v, err = strconv.ParseUint(strings.Trim(cid, "'\""), 10, 64); nil == err {
					result = append(result, uint32(v))
				} else {
					break
				}
			} // end for
		}
	}
	return
}

func encodeUint32Array(begin, end byte, arr []int32) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}
			buff.WriteString(strconv.FormatUint(uint64(v), 10))
		}
	}

	buff.WriteByte(end)
	return &buff
}

///////////////////////////////////////////////////////////////////////////////

func decodeUint64Array(data interface{}) (result []uint64, err error) {
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

	result = make([]uint64, 0)
	if "{}" != arr && len(arr) > 0 {
		if vals := strings.Split(strings.Trim(arr, "{}"), ","); nil != vals {
			for _, cid := range vals {
				var v uint64
				if v, err = strconv.ParseUint(strings.Trim(cid, "'\""), 10, 64); nil == err {
					result = append(result, v)
				} else {
					break
				}
			} // end for
		}
	}
	return
}

func encodeUint64Array(begin, end byte, arr []int64) *bytes.Buffer {
	var buff bytes.Buffer
	buff.WriteByte(begin)

	if nil != arr {
		for i, v := range arr {
			if i > 0 {
				buff.WriteByte(',')
			}
			buff.WriteString(strconv.FormatUint(uint64(v), 10))
		}
	}

	buff.WriteByte(end)
	return &buff
}
