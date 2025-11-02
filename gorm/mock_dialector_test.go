package gorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// mockDialector is a minimal mock implementation of gorm.Dialector for testing
type mockDialector struct {
	name string
}

func (m *mockDialector) Name() string {
	return m.name
}

func (m *mockDialector) Initialize(*gorm.DB) error {
	return nil
}

func (m *mockDialector) Migrator(db *gorm.DB) gorm.Migrator {
	return nil
}

func (m *mockDialector) DataTypeOf(*schema.Field) string {
	return ""
}

func (m *mockDialector) DefaultValueOf(*schema.Field) clause.Expression {
	return clause.Expr{}
}

func (m *mockDialector) BindVarTo(clause.Writer, *gorm.Statement, any) {
}

func (m *mockDialector) QuoteTo(clause.Writer, string) {
}

func (m *mockDialector) Explain(string, ...any) string {
	return ""
}

// createMockDB creates a properly initialized mock GORM DB for testing
func createMockDB(dialectName string) *gorm.DB {
	db := &gorm.DB{
		Config: &gorm.Config{},
	}
	db.Dialector = &mockDialector{name: dialectName}
	return db
}
