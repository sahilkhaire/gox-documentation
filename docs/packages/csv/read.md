---
title: "Read"
package: "csv"
import: "github.com/sahilkhaire/gox/csv"
gox-doc-version: "11"
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

// csv
_ = csv.Read(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/csv"

// csv
_ = csv.Read(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/csv` and call `Read` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/csv/parse-string">ParseString</a><a class="related-chip" href="/packages/csv/read-all">ReadAll</a><a class="related-chip" href="/packages/csv/write">Write</a>
</div>

← [Back to csv package overview](/packages/csv/)
