# GoSQL Types Collection

[![Build Status](https://github.com/geniusrabbit/gosql/workflows/run%20tests/badge.svg)](https://github.com/geniusrabbit/gosql/actions?workflow=run%20tests)
[![Go Report Card](https://goreportcard.com/badge/github.com/geniusrabbit/gosql)](https://goreportcard.com/report/github.com/geniusrabbit/gosql)
[![GoDoc](https://godoc.org/github.com/geniusrabbit/gosql?status.svg)](https://godoc.org/github.com/geniusrabbit/gosql)
[![Coverage Status](https://coveralls.io/repos/github/geniusrabbit/gosql/badge.svg)](https://coveralls.io/github/geniusrabbit/gosql)

A comprehensive Go library providing SQL-compatible types and collections with full JSON marshaling/unmarshaling support and database integration.

## Features

### Core Types

- **Char** - Single character type with SQL and JSON support
- **Duration** - Extended duration type with custom parsing (ns, us, ms, s, m, h, d, w)
- **JSON** - Generic JSON type for any value (structs, scalars, arrays)
- **StringArray** - Array of strings with PostgreSQL-compatible formatting
- **NumberArray** - Generic numeric arrays supporting integers and floats
- **NullableJSON** - JSON type with nullable support

### Array Types

All array types support:

- **Ordered** variants for sorted arrays
- **Nullable** variants that can be null
- PostgreSQL array format parsing and generation
- JSON marshaling/unmarshaling
- SQL scanning and value generation

### ORM Integration

Full GORM support is provided via the `gorm` subpackage with:

- Database-specific type mapping (MySQL, PostgreSQL, SQLite, YDB, ClickHouse)
- Custom value expressions for different SQL dialects
- Proper migration support

## Installation

```bash
go get github.com/geniusrabbit/gosql/v2
```

For GORM integration:

```bash
go get github.com/geniusrabbit/gosql/gorm
```

## Usage

### Basic Types

```go
import (
  "github.com/geniusrabbit/gosql/v2"
)

type Model struct {
  ID uint64
  Title string
  Status gosql.Char
  
  // JSON configurations
  Configuration gosql.JSON[Config]
  Metadata gosql.NullableJSON[map[string]any]
  
  // Arrays
  Tags gosql.StringArray
  Scores gosql.OrderedNumberArray[float64]
  Metrics gosql.NullableOrderedNumberArray[int]
  
  // Duration with custom parsing
  Timeout gosql.Duration
}

// Usage examples
model := Model{
  Status: gosql.Char('A'),
  Configuration: gosql.MustJSON(Config{Debug: true}),
  Tags: gosql.StringArray{"tag1", "tag2", "tag3"},
  Scores: gosql.OrderedNumberArray[float64]{9.5, 8.7, 9.1},
  Timeout: gosql.Duration(5 * time.Minute),
}
```

### ORM Usage

```go
import (
  "github.com/geniusrabbit/gosql/gorm"
)

type User struct {
  ID uint64
  Name string
  Status gorm.Char
  Settings gorm.JSON[UserSettings]
  Tags gorm.StringArray
}

// The GORM types automatically handle:
// - Database-specific SQL generation
// - Type casting for different databases
// - Proper migration schemas
```

### Supported Databases

The library provides optimized support for:

- **PostgreSQL** - Native array and JSON support
- **MySQL/MariaDB** - JSON and text-based arrays
- **SQLite** - Text-based storage with JSON parsing
- **SQL Server** - JSON and varchar arrays
- **YDB** - Specialized type casting
- **ClickHouse** - Optimized for analytics workloads

### Testing

The library includes comprehensive tests covering:

- Core type functionality (scan, value, JSON marshaling)
- Database-specific behavior testing
- GORM integration testing with mock dialectors
- Error handling and edge cases
- Unicode and special character support

Run tests:

```bash
# Core library tests
go test -v

# GORM integration tests
cd gorm && go test -v
```

## Contributing

Contributions are welcome! The library follows standard Go conventions and includes:

- Comprehensive test coverage
- Proper error handling
- Database compatibility testing
- JSON marshaling/unmarshaling validation

## License

MIT License - see [LICENSE](LICENSE) for details.
