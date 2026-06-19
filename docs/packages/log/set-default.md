---
title: "SetDefault"
package: "log"
import: "github.com/sahilkhaire/gox/log"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: winston, pino</span><span class="api-badge import">github.com/sahilkhaire/gox/log</span></div>
# SetDefault

## Overview

SetDefault replaces the package default logger.

## Signature

```go
func SetDefault(l *Logger)
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
import "github.com/sahilkhaire/gox/log"

// log
_ = log.SetDefault(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to log package overview](/packages/log/)
