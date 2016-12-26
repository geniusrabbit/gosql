//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package gosql

import (
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/lib/pq/hstore"
)

// Hstore implementation
type Hstore hstore.Hstore

// Map values for hstore
func (h Hstore) Data() map[string]sql.NullString {
	return h.Map
}

// Get value by key
func (h Hstore) Get(key string) (string, bool) {
	if v, ok := h.Map[key]; ok {
		return v.String, true
	}
	return "", false
}

// Set value for key
func (h *Hstore) Set(key, value string) {
	if nil == h.Map {
		h.Map = make(map[string]sql.NullString)
	}
	h.Map[key] = sql.NullString{value, true}
}

// Unset value by key
func (h *Hstore) Unset(key string) {
	if _, ok := h.Get(key); ok {
		delete(h.Map, key)
	}
}

// MarshalJSON method
func (h Hstore) MarshalJSON() ([]byte, error) {
	parts := make([]string, 0, len(h.Map))
	for key, val := range h.Hstore.Map {
		parts = append(parts, key+"=>"+val.String)
	}
	return json.Marshal(strings.Join(parts, ","))
}

// String method implementation of Stringer
func (h Hstore) String() string {
	if v, err := h.Value(); err == nil {
		if value, ok := v.([]byte); ok {
			return string(value)
		}
	}
	return ""
}
