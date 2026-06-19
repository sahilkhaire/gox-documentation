---
title: "ReadAll"
package: "csv"
import: "github.com/sahilkhaire/gox/csv"
gox-doc-version: "11"
---

<SymbolHeader pkg="csv" title="ReadAll" node="papaparse" import-path="github.com/sahilkhaire/gox/csv" />
## Overview

ReadAll reads a CSV file at path.

Part of the **`csv`** package — Node.js analog: *papaparse*.

## Signature

<div class="signature-block">

```go
func ReadAll(ctx context.Context, path string) ([]map[string]string, error)
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
_ = csv.ReadAll(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/csv"

// csv
_ = csv.ReadAll(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/csv` and call `ReadAll` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/csv/parse-string">ParseString</a><a class="related-chip" href="/packages/csv/read">Read</a><a class="related-chip" href="/packages/csv/write">Write</a>
</div>

← [Back to csv package overview](/packages/csv/)
