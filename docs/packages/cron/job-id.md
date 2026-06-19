---
title: "JobID"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: node-cron</span><span class="api-badge import">github.com/sahilkhaire/gox/cron</span></div>
# JobID

## Overview

JobID identifies a scheduled job.

## Signature

```go
type JobID int64
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

_ = cron.JobID
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Parse](/packages/cron/parse)

← [Back to cron package overview](/packages/cron/)
