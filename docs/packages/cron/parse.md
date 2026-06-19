---
title: "Parse"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: node-cron</span><span class="api-badge import">github.com/sahilkhaire/gox/cron</span></div>
# Parse

## Overview

Parse validates a cron expression (with optional seconds field).

## Signature

```go
func Parse(expr string) error
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
import "github.com/sahilkhaire/gox/cron"

// cron
_ = cron.Parse(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to cron package overview](/packages/cron/)
