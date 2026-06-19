---
title: "Version"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: semver</span><span class="api-badge import">github.com/sahilkhaire/gox/semver</span></div>
# Version

## Overview

Version is a semantic version.

## Signature

```go
type Version struct {
	// contains filtered or unexported fields
}
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
import "github.com/sahilkhaire/gox/semver"

_ = semver.Version
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
- [Satisfies](/packages/semver/satisfies)

← [Back to semver package overview](/packages/semver/)
