---
title: "Add"
package: "time"
import: "github.com/sahilkhaire/gox/time"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: moment, dayjs (package timex)</span><span class="api-badge import">github.com/sahilkhaire/gox/time</span></div>
# Add

## Overview

## Signature

```go
func Add(t stdtime.Time, d stdtime.Duration) stdtime.Time
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
import "github.com/sahilkhaire/gox/time"

// timex
_ = timex.Add(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Diff](/packages/time/diff)
- [EndOfDay](/packages/time/end-of-day)
- [EndOfMonth](/packages/time/end-of-month)

← [Back to time package overview](/packages/time/)
