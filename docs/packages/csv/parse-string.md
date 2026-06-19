---
title: "ParseString"
package: "csv"
import: "github.com/sahilkhaire/gox/csv"
gox-doc-version: "14"
---

<SymbolHeader pkg="csv" title="ParseString" node="papaparse" import-path="github.com/sahilkhaire/gox/csv" />
## Overview

ParseString parses CSV text with a header row.

Part of the **`csv`** package — Node.js analog: *papaparse*.

## Signature

<div class="signature-block">

```go
func ParseString(s string) ([]map[string]string, error)
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

rows, err := csv.ParseString(raw, csv.ReadOptions{Header: true})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/csv"

rows, err := csv.ParseString(raw, csv.ReadOptions{Header: true})
```

## Tips

Import `github.com/sahilkhaire/gox/csv` and call `ParseString` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
r := csv.NewReader(f)
records, err := r.ReadAll()
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/csv/read">Read</a><a class="related-chip" href="/packages/csv/read-all">ReadAll</a><a class="related-chip" href="/packages/csv/write">Write</a>
</div>

← [Back to csv package overview](/packages/csv/)
