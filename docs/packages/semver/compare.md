---
title: "Compare"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
node: "semver.compare(a, b)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: semver.compare(a, b)</span><span class="api-badge import">github.com/sahilkhaire/gox/semver</span></div>
# Compare

## Overview

Maps the Node.js pattern `semver.compare(a, b)` to gox `semver.Compare(a, b)`.

**Node.js equivalent:** `semver.compare(a, b)`

## Signature

```go
func Compare(a, b string) int
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
semver.compare(a, b)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/semver"

semver.Compare(a, b)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Inc](/packages/semver/inc)
- [Satisfies](/packages/semver/satisfies)

← [Back to semver package overview](/packages/semver/)
