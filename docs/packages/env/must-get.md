---
title: "MustGet"
package: "env"
import: "github.com/sahilkhaire/gox/env"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: dotenv / process.env</span><span class="api-badge import">github.com/sahilkhaire/gox/env</span></div>
# MustGet

## Overview

MustGet returns the value or panics if missing.

## Signature

```go
func MustGet(key string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
v := os.Getenv("KEY")
if v == "" {
    v = "default"
}
```

```go [gox]
import "github.com/sahilkhaire/gox/env"

// env
_ = env.MustGet(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Get](/packages/env/get)
- [GetBool](/packages/env/get-bool)
- [GetDuration](/packages/env/get-duration)

← [Back to env package overview](/packages/env/)
