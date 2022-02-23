# gosql

Library of standart sql collections an types like:

 * Char
 * NumberArrays generic (Ordered, Nullable) Numbers only (int, float, ...)
 * StringArray
 * JSON generic with any type of value `structs`, `scalars`, ...

New version supports generics from `go1.8` to make easear using of data processing.

```go
import (
  "github.com/geniusrabbit/gosql/v2"
)

type Model struct {
  ID uint64
  Title string

  Configuration gosql.JSON[Config]
  Tags gosql.StringArray

  DebugMetricQuartiles gosql.NullableOrderedNumberArray[float64]
}
```

# License MIT

[LICENSE](LICENSE)
