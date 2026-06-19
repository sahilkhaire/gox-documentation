---
title: "Satisfies"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
node: "semver.satisfies(v, range)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: semver.satisfies(v, range)</span><span class="api-badge import">github.com/sahilkhaire/gox/semver</span></div>
# Satisfies

## Overview

Maps the Node.js pattern `semver.satisfies(v, range)` to gox `semver.Satisfies(v, constraint)`.

**Node.js equivalent:** `semver.satisfies(v, range)`

## Signature

```go
func Satisfies(version, constraint string) bool
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
semver.satisfies(v, range)
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/semver"

semver.Satisfies(v, constraint)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Compare](/packages/semver/compare)
- [Inc](/packages/semver/inc)

← [Back to semver package overview](/packages/semver/)
