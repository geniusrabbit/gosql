package gorm

import (
	"context"
	"database/sql/driver"

	"github.com/geniusrabbit/gosql/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// Char field type declaration with GORM type methods
type Char gosql.Char

// GormDataType gorm common data type
func (Char) GormDataType() string {
	return "char"
}

// GormDBDataType gorm db data type
func (c Char) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql", "mariadb", "postgres", "sqlserver":
		return "char"
	case "sqlite", "sqlite3":
		return "text"
	case "ydb", "clickhouse":
		return "Int32"
	}
	return ""
}

// Value implements the driver.Valuer interface, char field
func (f Char) Value() (driver.Value, error) {
	return (gosql.Char)(f).Value()
}

// GormValue gorm expr for char field
func (f Char) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	switch db.Dialector.Name() {
	case "ydb", "clickhouse":
		return clause.Expr{SQL: "CAST(? AS Int32)", Vars: []any{int32(f)}}
	}
	return clause.Expr{SQL: "?", Vars: []any{f}}
}

// Scan implements the sql.Scanner interface, char field
func (f *Char) Scan(value any) error {
	return (*gosql.Char)(f).Scan(value)
}

// MarshalJSON implements the json.Marshaler
func (f Char) MarshalJSON() ([]byte, error) {
	return (gosql.Char)(f).MarshalJSON()
}

// UnmarshalJSON implements the json.Unmarshaller
func (f *Char) UnmarshalJSON(b []byte) (err error) {
	return (*gosql.Char)(f).UnmarshalJSON(b)
}
