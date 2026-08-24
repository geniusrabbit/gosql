package gosql

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringArrayJSON(t *testing.T) {
	var arr1 = StringArray{"val1", "val2", "val3"}
	data, err := json.Marshal(arr1)
	assert.NoError(t, err, "marshal array")
	var arr2 StringArray
	err = json.Unmarshal(data, &arr2)
	assert.NoError(t, err, "unmarshal array")
	assert.True(t, reflect.DeepEqual(arr1, arr2),
		"compare encode/decode result")
}

func TestStringArraySQL(t *testing.T) {
	var arr StringArray
	err := arr.Scan("{10000,10000,10000,10000}")
	assert.NoError(t, err, "scan array")
	assert.ElementsMatch(t, arr, []string{"10000", "10000", "10000", "10000"},
		"compare scan result")

	arr = arr[:0]
	sqlStringArray := `{"breakfast", "consulting", "bar-""#1"""}`
	err = arr.Scan(sqlStringArray)
	assert.NoError(t, err, "scan array")
	assert.ElementsMatch(t, arr, []string{"breakfast", "consulting", `bar-"#1"`},
		"compare scan result")
	sqlVal, err := arr.Value()
	assert.NoError(t, err, "encode array")
	assert.Equal(t, strings.ReplaceAll(sqlStringArray, " ", ""), sqlVal)
}

func TestStringArrayValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var arr StringArray
		v, err := arr.Value()
		assert.Equal(t, ErrNullValueNotAllowed, err)
		assert.Equal(t, "{}", v)
	})

	t.Run("empty", func(t *testing.T) {
		arr := StringArray{}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{}", v)
	})
}

func TestStringArrayScanNil(t *testing.T) {
	var arr StringArray
	err := arr.Scan(nil)
	assert.Equal(t, ErrNullValueNotAllowed, err)
}

func TestStringArrayUnmarshalJSONNil(t *testing.T) {
	var arr StringArray
	err := arr.UnmarshalJSON(nil)
	assert.Equal(t, ErrNullValueNotAllowed, err)

	err = arr.UnmarshalJSON([]byte(`{`))
	assert.Error(t, err)
}

func TestStringArrayHelpers(t *testing.T) {
	t.Run("join", func(t *testing.T) {
		arr := StringArray{"a", "b", "c"}
		assert.Equal(t, "a,b,c", arr.Join(","))
		assert.Equal(t, "", StringArray(nil).Join(","))
	})

	t.Run("set_array", func(t *testing.T) {
		var arr StringArray
		got := arr.SetArray([]string{"x", "y"})
		assert.Equal(t, StringArray{"x", "y"}, arr)
		assert.Equal(t, &arr, got)

		arr.SetArray(nil)
		assert.Equal(t, StringArray{}, arr)
	})

	t.Run("len_index_oneof", func(t *testing.T) {
		arr := StringArray{"a", "b", "c"}
		assert.Equal(t, 3, arr.Len())
		assert.Equal(t, 1, arr.IndexOf("b"))
		assert.Equal(t, -1, arr.IndexOf("z"))
		assert.True(t, arr.OneOf([]string{"z", "b"}))
		assert.False(t, arr.OneOf([]string{"z"}))
		assert.False(t, arr.OneOf(nil))
	})
}

func TestStringArrayDecodeValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var arr StringArray
		err := arr.DecodeValue(nil)
		assert.Equal(t, ErrNullValueNotAllowed, err)
	})

	t.Run("slice", func(t *testing.T) {
		var arr StringArray
		err := arr.DecodeValue([]string{"a", "b"})
		assert.NoError(t, err)
		assert.Equal(t, StringArray{"a", "b"}, arr)
	})

	t.Run("comma_string", func(t *testing.T) {
		var arr StringArray
		err := arr.DecodeValue("a,b,c")
		assert.NoError(t, err)
		assert.Equal(t, StringArray{"a", "b", "c"}, arr)
	})

	t.Run("invalid_type", func(t *testing.T) {
		var arr StringArray
		err := arr.DecodeValue(123)
		assert.Equal(t, ErrInvalidDecodeValue, err)
	})
}

func TestNullableStringArray(t *testing.T) {
	t.Run("value_nil", func(t *testing.T) {
		var arr NullableStringArray
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Nil(t, v)
	})

	t.Run("value_empty", func(t *testing.T) {
		arr := NullableStringArray{}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, "{}", v)
	})

	t.Run("value_items", func(t *testing.T) {
		arr := NullableStringArray{"a", "b"}
		v, err := arr.Value()
		assert.NoError(t, err)
		assert.Equal(t, `{"a","b"}`, v)
	})

	t.Run("scan_nil", func(t *testing.T) {
		arr := NullableStringArray{"a"}
		err := arr.Scan(nil)
		assert.NoError(t, err)
		assert.Nil(t, arr)
	})

	t.Run("scan_null_string", func(t *testing.T) {
		arr := NullableStringArray{"a"}
		err := arr.Scan("null")
		assert.NoError(t, err)
		assert.Nil(t, arr)
	})

	t.Run("scan_empty", func(t *testing.T) {
		var arr NullableStringArray
		err := arr.Scan("{}")
		assert.NoError(t, err)
		assert.Equal(t, NullableStringArray{}, arr)
	})

	t.Run("scan_bytes", func(t *testing.T) {
		var arr NullableStringArray
		err := arr.Scan([]byte("{x,y}"))
		assert.NoError(t, err)
		assert.Equal(t, NullableStringArray{"x", "y"}, arr)
	})

	t.Run("scan_empty_quoted", func(t *testing.T) {
		var arr NullableStringArray
		err := arr.Scan(`{""}`)
		assert.NoError(t, err)
		assert.Equal(t, NullableStringArray{""}, arr)
	})

	t.Run("json", func(t *testing.T) {
		var nilArr NullableStringArray
		data, err := nilArr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "null", string(data))

		arr := NullableStringArray{"a", "b"}
		data, err = arr.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, `["a","b"]`, string(data))

		var got NullableStringArray
		err = got.UnmarshalJSON(data)
		assert.NoError(t, err)
		assert.Equal(t, arr, got)
	})

	t.Run("decode_value", func(t *testing.T) {
		var arr NullableStringArray
		assert.NoError(t, arr.DecodeValue([]string{"a"}))
		assert.Equal(t, NullableStringArray{"a"}, arr)

		assert.NoError(t, arr.DecodeValue(NullableStringArray{"b"}))
		assert.Equal(t, NullableStringArray{"b"}, arr)

		assert.NoError(t, arr.DecodeValue(StringArray{"c"}))
		assert.Equal(t, NullableStringArray{"c"}, arr)

		assert.NoError(t, arr.DecodeValue([]byte("d,e")))
		assert.Equal(t, NullableStringArray{"d", "e"}, arr)

		assert.NoError(t, arr.DecodeValue("f,g"))
		assert.Equal(t, NullableStringArray{"f", "g"}, arr)

		assert.Equal(t, ErrInvalidDecodeValue, arr.DecodeValue(1))
	})

	t.Run("helpers", func(t *testing.T) {
		var nilArr NullableStringArray
		assert.Equal(t, "", nilArr.Join(","))
		assert.Equal(t, 0, nilArr.Len())
		assert.Equal(t, -1, nilArr.IndexOf("a"))

		arr := NullableStringArray{"a", "b"}
		assert.Equal(t, "a|b", arr.Join("|"))
		assert.Equal(t, 2, arr.Len())
		assert.Equal(t, 0, arr.IndexOf("a"))
		assert.True(t, arr.OneOf([]string{"b"}))
		assert.False(t, arr.OneOf([]string{"z"}))
		assert.False(t, arr.OneOf(nil))

		arr.SetArray([]string{"x"})
		assert.Equal(t, NullableStringArray{"x"}, arr)
	})
}
