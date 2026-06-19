---
title: "ParseString"
package: "csv"
import: "github.com/sahilkhaire/gox/csv"
gox-doc-version: "10"
---

<SymbolHeader pkg="csv" title="ParseString" node="papaparse" import-path="github.com/sahilkhaire/gox/csv" />
## Overview

ParseString parses CSV text with a header row.

## Signature

<div class="signature-block">

```go
func ParseString(s string) ([]map[string]string, error)
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
_ = csv.ParseString(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/csv/read">Read</a><a class="related-chip" href="/packages/csv/read-all">ReadAll</a><a class="related-chip" href="/packages/csv/write">Write</a>
</div>

← [Back to csv package overview](/packages/csv/)
