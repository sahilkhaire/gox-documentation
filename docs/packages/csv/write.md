---
title: "Write"
package: "csv"
import: "github.com/sahilkhaire/gox/csv"
gox-doc-version: "10"
---

<SymbolHeader pkg="csv" title="Write" node="papaparse" import-path="github.com/sahilkhaire/gox/csv" />
## Overview

Write writes rows to w using columns for field order (header row).

## Signature

<div class="signature-block">

```go
func Write(ctx context.Context, w io.Writer, rows []map[string]string, columns []string) error
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
_ = csv.Write(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/csv/parse-string">ParseString</a><a class="related-chip" href="/packages/csv/read">Read</a><a class="related-chip" href="/packages/csv/read-all">ReadAll</a>
</div>

← [Back to csv package overview](/packages/csv/)
