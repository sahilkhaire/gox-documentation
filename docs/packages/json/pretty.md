---
title: "Pretty"
package: "json"
import: "github.com/sahilkhaire/gox/json"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: JSON.parse/stringify</span><span class="api-badge import">github.com/sahilkhaire/gox/json</span></div>
# Pretty

## Overview

Pretty returns indented JSON.

## Signature

```go
func Pretty(v any) (string, error)
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
import "github.com/sahilkhaire/gox/json"

// json
_ = json.Pretty(/* args */)
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/json"

s, err := goxjson.Pretty(map[string]int{"a": 1})
```

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [MustParse](/packages/json/must-parse)
- [MustStringify](/packages/json/must-stringify)
- [Parse](/packages/json/parse)

← [Back to json package overview](/packages/json/)
