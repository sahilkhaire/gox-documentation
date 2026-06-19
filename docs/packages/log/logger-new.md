---
title: "Logger.New"
package: "log"
import: "github.com/sahilkhaire/gox/log"
node: "winston.createLogger()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: winston.createLogger()</span><span class="api-badge import">github.com/sahilkhaire/gox/log</span></div>
# Logger.New

## Overview

New returns a Logger writing JSON to stderr at info level.

**Node.js equivalent:** `winston.createLogger()`

## Signature

```go
func New() *Logger
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
winston.createLogger();
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/log"

logger := log.New()
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Logger.Default](/packages/log/logger-default)
- [Logger.NewWithLevel](/packages/log/logger-new-with-level)

← [Back to log package overview](/packages/log/)
