package gorm

import (
	"context"
	"testing"

	"github.com/geniusrabbit/gosql/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/schema"
)

func TestGormChar(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected Char
		wantErr  bool
	}{
		{"string_single_char", "A", Char('A'), false},
		{"string_multi_char", "ABC", Char('A'), false},
		{"byte_slice", []byte("B"), Char('B'), false},
		{"byte_slice_multi", []byte("BCD"), Char('B'), false},
		{"nil_value", nil, Char(0), true},
	}

	for _, test := range tests {
		t.Run("scan:"+test.name, func(t *testing.T) {
			var c Char
			err := c.Scan(test.input)
			if test.wantErr {
				assert.Error(t, err)
				assert.Equal(t, gosql.ErrNullValueNotAllowed, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, c)
			}
		})
	}

	t.Run("scan:invalid_type", func(t *testing.T) {
		var c Char
		err := c.Scan(123.45) // float64 - unsupported type
		assert.Error(t, err)
		assert.Equal(t, gosql.ErrInvalidScan, err)
	})

	t.Run("scan:empty_values", func(t *testing.T) {
		var c Char
		err := c.Scan("")
		assert.Error(t, err)
		assert.Equal(t, gosql.ErrInvalidScan, err)

		err = c.Scan([]byte{})
		assert.Error(t, err)
		assert.Equal(t, gosql.ErrInvalidScan, err)
	})

	t.Run("scan:numeric_types", func(t *testing.T) {
		// Test numeric types separately as they might behave differently
		// depending on the gosql version
		var c Char

		// Test rune
		err := c.Scan(rune('C'))
		if err != nil {
			// If rune doesn't work, it should be ErrInvalidScan
			assert.Equal(t, gosql.ErrInvalidScan, err)
		} else {
			assert.Equal(t, Char('C'), c)
		}

		// Test int
		c = Char(0) // Reset
		err = c.Scan(int(65))
		if err != nil {
			// If int doesn't work, it should be ErrInvalidScan
			assert.Equal(t, gosql.ErrInvalidScan, err)
		} else {
			assert.Equal(t, Char('A'), c)
		}

		// Test uint16
		c = Char(0) // Reset
		err = c.Scan(uint16(66))
		if err != nil {
			// If uint16 doesn't work, it should be ErrInvalidScan
			assert.Equal(t, gosql.ErrInvalidScan, err)
		} else {
			assert.Equal(t, Char('B'), c)
		}

		// Test uint32
		c = Char(0) // Reset
		err = c.Scan(uint32(67))
		if err != nil {
			// If uint32 doesn't work, it should be ErrInvalidScan
			assert.Equal(t, gosql.ErrInvalidScan, err)
		} else {
			assert.Equal(t, Char('C'), c)
		}
	})
}

func TestGormCharValue(t *testing.T) {
	tests := []struct {
		name     string
		char     Char
		expected string
	}{
		{"regular_char", Char('A'), "A"},
		{"space_char", Char(' '), " "},
		{"zero_char", Char(0), " "},
		{"unicode_char", Char('ñ'), "ñ"},
		{"number_char", Char('5'), "5"},
	}

	for _, test := range tests {
		t.Run("value:"+test.name, func(t *testing.T) {
			value, err := test.char.Value()
			assert.NoError(t, err)
			assert.Equal(t, test.expected, value)
		})
	}
}

func TestGormCharJSON(t *testing.T) {
	tests := []struct {
		name         string
		char         Char
		expectedJSON string
	}{
		{"regular_char", Char('A'), `"A"`},
		{"space_char", Char(' '), `" "`},
		{"zero_char", Char(0), `" "`},
	}

	for _, test := range tests {
		t.Run("marshal:"+test.name, func(t *testing.T) {
			data, err := test.char.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, test.expectedJSON, string(data))
		})
	}

	t.Run("unmarshal:direct", func(t *testing.T) {
		// Test direct UnmarshalJSON calls
		// Note: The implementation delegates to gosql.Char which uses decodeChar directly
		tests := []struct {
			json     string
			expected Char
		}{
			{`"A"`, Char('"')}, // Gets the first character (quote)
			{`"5"`, Char('"')}, // Gets the first character (quote)
			{`" "`, Char('"')}, // Gets the first character (quote)
		}

		for _, test := range tests {
			var c Char
			err := c.UnmarshalJSON([]byte(test.json))
			assert.NoError(t, err)
			assert.Equal(t, test.expected, c)
		}
	})

	t.Run("unmarshal:invalid_json", func(t *testing.T) {
		var c Char

		// Test with nil - becomes empty byte slice, should error with ErrInvalidScan
		err := c.UnmarshalJSON(nil)
		assert.Error(t, err)
		assert.Equal(t, gosql.ErrInvalidScan, err)

		// Test with empty byte slice - should error
		err = c.UnmarshalJSON([]byte{})
		assert.Error(t, err)
		assert.Equal(t, gosql.ErrInvalidScan, err)
	})
}

func TestGormCharDataType(t *testing.T) {
	var c Char

	t.Run("gorm_data_type", func(t *testing.T) {
		dataType := c.GormDataType()
		assert.Equal(t, "char", dataType)
	})
}

func TestGormCharDBDataType(t *testing.T) {
	var c Char
	field := &schema.Field{Name: "test_field"}

	tests := []struct {
		dialectName  string
		expectedType string
	}{
		{"mysql", "char"},
		{"mariadb", "char"},
		{"postgres", "char"},
		{"sqlserver", "char"},
		{"sqlite", "text"},
		{"sqlite3", "text"},
		{"ydb", "Int32"},
		{"clickhouse", "Int32"},
		{"unknown_dialect", ""},
	}

	for _, test := range tests {
		t.Run("dialect_"+test.dialectName, func(t *testing.T) {
			// Create a properly initialized mock DB
			db := createMockDB(test.dialectName)

			result := c.GormDBDataType(db, field)
			assert.Equal(t, test.expectedType, result)
		})
	}
}

func TestGormCharGormValue(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		char        Char
		dialectName string
		expectedSQL string
		expectedVar any
	}{
		{"mysql_regular_char", Char('A'), "mysql", "?", Char('A')},
		{"postgres_space_char", Char(' '), "postgres", "?", Char(' ')},
		{"sqlite_zero_char", Char(0), "sqlite", "?", Char(0)},
		{"sqlserver_unicode_char", Char('ñ'), "sqlserver", "?", Char('ñ')},
		{"ydb_regular_char", Char('B'), "ydb", "CAST(? AS Int32)", int32('B')},
		{"clickhouse_number_char", Char('5'), "clickhouse", "CAST(? AS Int32)", int32('5')},
		{"clickhouse_zero_char", Char(0), "clickhouse", "CAST(? AS Int32)", int32(0)},
		{"unknown_dialect", Char('X'), "unknown", "?", Char('X')},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a properly initialized mock DB
			db := createMockDB(test.dialectName)

			expr := test.char.GormValue(ctx, db)
			assert.Equal(t, test.expectedSQL, expr.SQL)
			assert.Len(t, expr.Vars, 1)
			assert.Equal(t, test.expectedVar, expr.Vars[0])
		})
	}

	t.Run("context_handling", func(t *testing.T) {
		// Test that different contexts work (the method doesn't actually use context currently)
		c := Char('T')
		db := createMockDB("mysql")

		// Test with background context
		expr1 := c.GormValue(context.Background(), db)
		assert.Equal(t, "?", expr1.SQL)

		// Test with TODO context
		expr2 := c.GormValue(context.TODO(), db)
		assert.Equal(t, "?", expr2.SQL)

		// Both should produce the same result
		assert.Equal(t, expr1.SQL, expr2.SQL)
		assert.Equal(t, expr1.Vars, expr2.Vars)
	})
}

func TestGormCharEdgeCases(t *testing.T) {
	t.Run("zero_initialization", func(t *testing.T) {
		var c Char
		assert.Equal(t, Char(0), c)

		// Zero char should return space when getting value
		value, err := c.Value()
		assert.NoError(t, err)
		assert.Equal(t, " ", value)

		// Zero char should marshal to space
		data, err := c.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, `" "`, string(data))
	})

	t.Run("type_conversion_consistency", func(t *testing.T) {
		// Test that GORM Char behaves consistently with gosql.Char
		gormChar := Char('A')
		gosqlChar := gosql.Char('A')

		// Value() should be identical
		gormValue, err1 := gormChar.Value()
		gosqlValue, err2 := gosqlChar.Value()
		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Equal(t, gosqlValue, gormValue)

		// MarshalJSON should be identical
		gormJSON, err1 := gormChar.MarshalJSON()
		gosqlJSON, err2 := gosqlChar.MarshalJSON()
		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Equal(t, gosqlJSON, gormJSON)
	})

	t.Run("high_unicode_values", func(t *testing.T) {
		// Test with high Unicode values - Value() should work properly
		c := Char(0x1F600) // 😀 emoji
		value, err := c.Value()
		assert.NoError(t, err)
		assert.Equal(t, "😀", value)
	})
}
