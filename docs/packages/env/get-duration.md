---
title: "GetDuration"
package: "env"
import: "github.com/sahilkhaire/gox/env"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: dotenv / process.env</span><span class="api-badge import">github.com/sahilkhaire/gox/env</span></div>
# GetDuration

## Overview

GetDuration parses a duration environment variable.

## Signature

```go
func GetDuration(key string, def time.Duration) (time.Duration, error)
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
_ = env.GetDuration(/* args */)
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
- [GetInt](/packages/env/get-int)

← [Back to env package overview](/packages/env/)
