package gosql

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	tests := []struct {
		value  string
		target Duration
	}{
		{"1ns", Nanosecond},
		{"1us", Microsecond},
		{"1ms", Millisecond},
		{"1s", Second},
		{"1m", Minute},
		{"1h", Hour},
		{"1d", Duration(24 * time.Hour)},
		{"1w", Duration(7 * 24 * time.Hour)},
	}

	for _, test := range tests {
		t.Run("parse:"+test.value, func(t *testing.T) {
			d, err := ParseDuration(test.value)
			if assert.NoError(t, err) {
				assert.Equal(t, test.target, d)
			}
		})
		t.Run("scan:"+test.value, func(t *testing.T) {
			var d Duration
			if err := d.Scan(test.value); assert.NoError(t, err) {
				assert.Equal(t, test.target, d)
			}
			if err := d.Scan([]byte(test.value)); assert.NoError(t, err) {
				assert.Equal(t, test.target, d)
			}
			if err := d.Scan(int64(d)); assert.NoError(t, err) {
				assert.Equal(t, test.target, d)
			}
		})
	}

	t.Run("scan:invalid", func(t *testing.T) {
		var d Duration
		assert.Error(t, d.Scan("invalid"))
		assert.Error(t, d.Scan([]byte("invalid")))
		assert.Error(t, d.Scan("w-error"))
		assert.Error(t, d.Scan(any(nil)))
	})

	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "1ns", Nanosecond.String())
		assert.Equal(t, "1µs", Microsecond.String())
	})

	t.Run("values", func(t *testing.T) {
		assert.Equal(t, int64(1), Nanosecond.Nanoseconds())
		assert.Equal(t, int64(1), Microsecond.Microseconds())
		assert.Equal(t, int64(1), Millisecond.Milliseconds())
		assert.Equal(t, float64(1), Second.Seconds())
		assert.Equal(t, float64(1), Minute.Minutes())
		assert.Equal(t, float64(1), Hour.Hours())
		assert.Equal(t, Hour, Hour.Truncate(Minute))
		assert.Equal(t, Hour, Hour.Round(Minute))
		assert.Equal(t, Hour, Hour.Abs())

		v, err := Hour.Value()
		assert.NoError(t, err)
		assert.Equal(t, "1h0m0s", v)
	})

	t.Run("json", func(t *testing.T) {
		var d Duration
		assert.NoError(t, d.UnmarshalJSON([]byte(`"1h"`)))
		assert.Equal(t, Hour, d)

		v, err := d.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(`"1h0m0s"`), v)

		assert.Error(t, d.UnmarshalJSON([]byte(`1h`)))
		assert.Error(t, d.UnmarshalJSON([]byte(``)))
	})
}
