---
title: "Scheduler.New"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: node-cron</span><span class="api-badge import">github.com/sahilkhaire/gox/cron</span></div>
# Scheduler.New

## Overview

New creates a scheduler and starts its internal runner.

## Signature

```go
func New() *Scheduler
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const job = cron.schedule('* * * * *', fn);
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/cron"

sched := cron.New()
sched.Schedule("* * * * *", fn)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to cron package overview](/packages/cron/)
