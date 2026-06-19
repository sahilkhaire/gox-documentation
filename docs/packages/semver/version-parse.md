---
title: "Version.Parse"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
node: "semver.parse(v)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: semver.parse(v)</span><span class="api-badge import">github.com/sahilkhaire/gox/semver</span></div>
# Version.Parse

## Overview

Parse parses a semver string.

**Node.js equivalent:** `semver.parse(v)`

## Signature

```go
func Parse(v string) (*Version, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
semver.parse('1.2.3');
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/semver"

v, err := semver.Parse("1.2.3")
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to semver package overview](/packages/semver/)
