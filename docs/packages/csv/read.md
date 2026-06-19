---
title: "Read"
package: "csv"
import: "github.com/sahilkhaire/gox/csv"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: papaparse</span><span class="api-badge import">github.com/sahilkhaire/gox/csv</span></div>
# Read

## Overview

Read parses CSV from r; the first row is used as column keys.

## Signature

```go
func Read(ctx context.Context, r io.Reader) ([]map[string]string, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/csv"

// csv
_ = csv.Read(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [ParseString](/packages/csv/parse-string)
- [ReadAll](/packages/csv/read-all)
- [Write](/packages/csv/write)

← [Back to csv package overview](/packages/csv/)
