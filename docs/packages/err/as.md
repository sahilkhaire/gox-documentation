---
title: "As"
package: "err"
import: "github.com/sahilkhaire/gox/err"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: http-errors</span><span class="api-badge import">github.com/sahilkhaire/gox/err</span></div>
# As

## Overview

As finds the first error in the chain assignable to target.

## Signature

```go
func As(err error, target any) bool
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
import "github.com/sahilkhaire/gox/err"

// err
_ = err.As(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Is](/packages/err/is)

← [Back to err package overview](/packages/err/)
