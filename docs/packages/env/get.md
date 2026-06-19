---
title: "Get"
package: "env"
import: "github.com/sahilkhaire/gox/env"
node: "process.env.KEY"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: process.env.KEY</span><span class="api-badge import">github.com/sahilkhaire/gox/env</span></div>
# Get

## Overview

Get returns the value for key (override, then os.Getenv).

**Node.js equivalent:** `process.env.KEY`

## Signature

```go
func Get(key string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const port = process.env.PORT || '8080';
```

```go [Standard Go]
port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
```

```go [gox]
import "github.com/sahilkhaire/gox/env"

port := env.Get("PORT", "8080")
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [GetBool](/packages/env/get-bool)
- [GetDuration](/packages/env/get-duration)
- [GetInt](/packages/env/get-int)

← [Back to env package overview](/packages/env/)
