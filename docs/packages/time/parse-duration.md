---
title: "ParseDuration"
package: "time"
import: "github.com/sahilkhaire/gox/time"
node: "ms('2d')"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: ms('2d')</span><span class="api-badge import">github.com/sahilkhaire/gox/time</span></div>
# ParseDuration

## Overview

Maps the Node.js pattern `ms('2d')` to gox `timex.ParseDuration("2d")`.

**Node.js equivalent:** `ms('2d')`

## Signature

```go
func ParseDuration(s string) (stdtime.Duration, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
ms('2d')
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/time"

timex.ParseDuration("2d")
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Add](/packages/time/add)
- [Diff](/packages/time/diff)
- [EndOfDay](/packages/time/end-of-day)

← [Back to time package overview](/packages/time/)
