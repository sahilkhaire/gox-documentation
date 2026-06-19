---
title: "Format"
package: "time"
import: "github.com/sahilkhaire/gox/time"
node: "moment().format()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: moment().format()</span><span class="api-badge import">github.com/sahilkhaire/gox/time</span></div>
# Format

## Overview

Maps the Node.js pattern `moment().format()` to gox `timex.Format(t, layout)`.

**Node.js equivalent:** `moment().format()`

## Signature

```go
func Format(t stdtime.Time, layout string) string
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
moment().format()
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/time"

timex.Format(t, layout)
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
