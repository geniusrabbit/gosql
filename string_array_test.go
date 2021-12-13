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
