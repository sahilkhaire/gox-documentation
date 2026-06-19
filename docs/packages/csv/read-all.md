---
title: "ReadAll"
package: "csv"
import: "github.com/sahilkhaire/gox/csv"
gox-doc-version: "10"
---

<SymbolHeader pkg="csv" title="ReadAll" node="papaparse" import-path="github.com/sahilkhaire/gox/csv" />
## Overview

ReadAll reads a CSV file at path.

## Signature

<div class="signature-block">

```go
func ReadAll(ctx context.Context, path string) ([]map[string]string, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/csv/parse-string">ParseString</a><a class="related-chip" href="/packages/csv/read">Read</a><a class="related-chip" href="/packages/csv/write">Write</a>
</div>

← [Back to csv package overview](/packages/csv/)
