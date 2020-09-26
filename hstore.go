//
// @project GeniusRabbit
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 – 2017, 2020
//

package gosql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/lib/pq/hstore"
)

// Hstore implementation
type Hstore hstore.Hstore

// Data values for hstore
func (h Hstore) Data() map[string]sql.NullString {
	return h.Map
}

// GetBool value from store
func (h Hstore) GetBool(key string) bool {
	if v, ok := h.Get(key); ok {
		ok, _ = strconv.ParseBool(v)
		return ok
	}
	return false
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
	if h.Map == nil {
		h.Map = make(map[string]sql.NullString)
	}
	h.Map[key] = sql.NullString{String: value, Valid: true}
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
	for key, val := range h.Map {
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

// Scan implements the Scanner interface.
//
// Note h.Map is reallocated before the scan to clear existing values. If the
// hstore column's database value is NULL, then h.Map is set to nil instead.
func (h *Hstore) Scan(value interface{}) error {
	return (*hstore.Hstore)(h).Scan(value)
}

// Value implements the driver Valuer interface. Note if h.Map is nil, the
// database column value will be set to NULL.
func (h Hstore) Value() (driver.Value, error) {
	return (hstore.Hstore)(h).Value()
}
