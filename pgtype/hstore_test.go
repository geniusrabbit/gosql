package pgtype

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHstoreJSON(t *testing.T) {
	var (
		store1 Hstore
		store2 Hstore
	)
	store1.Set("key1", "val1")
	store1.Set("key2", "val2")
	store1.Set("key3", "val3")
	data, err := json.Marshal(&store1)
	assert.NoError(t, err, "marshal")
	err = json.Unmarshal(data, &store2)
	assert.NoError(t, err, "unmarshal")
	assert.True(t, reflect.DeepEqual(store1.Data(), store2.Data()),
		"compare encode/decode result")
}
