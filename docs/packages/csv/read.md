---
title: "Read"
package: "csv"
import: "github.com/sahilkhaire/gox/csv"
gox-doc-version: "14"
---

<SymbolHeader pkg="csv" title="Read" node="papaparse" import-path="github.com/sahilkhaire/gox/csv" />
## Overview

Read parses CSV from r; the first row is used as column keys.

Part of the **`csv`** package — Node.js analog: *papaparse*.

## Signature

<div class="signature-block">

```go
func Read(ctx context.Context, r io.Reader) ([]map[string]string, error)
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

rows, err := csv.Read(ctx, r, csv.ReadOptions{Header: true})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/csv"

rows, err := csv.Read(ctx, r, csv.ReadOptions{Header: true})
```

## Tips

Import `github.com/sahilkhaire/gox/csv` and call `Read` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
r := csv.NewReader(f)
records, err := r.ReadAll()
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/csv/parse-string">ParseString</a><a class="related-chip" href="/packages/csv/read-all">ReadAll</a><a class="related-chip" href="/packages/csv/write">Write</a>
</div>

← [Back to csv package overview](/packages/csv/)
