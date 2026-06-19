---
title: "Write"
package: "csv"
import: "github.com/sahilkhaire/gox/csv"
gox-doc-version: "11"
---

<SymbolHeader pkg="csv" title="Write" node="papaparse" import-path="github.com/sahilkhaire/gox/csv" />
## Overview

Write writes rows to w using columns for field order (header row).

Part of the **`csv`** package — Node.js analog: *papaparse*.

## Signature

<div class="signature-block">

```go
func Write(ctx context.Context, w io.Writer, rows []map[string]string, columns []string) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical papaparse pattern in Node.js
```

```go [Standard Go]
r := csv.NewReader(f)
records, err := r.ReadAll()
```

```go [gox]
import "github.com/sahilkhaire/gox/csv"

// csv
_ = csv.Write(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/csv"

// csv
_ = csv.Write(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/csv` and call `Write` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/csv/parse-string">ParseString</a><a class="related-chip" href="/packages/csv/read">Read</a><a class="related-chip" href="/packages/csv/read-all">ReadAll</a>
</div>

← [Back to csv package overview](/packages/csv/)
