---
title: "Cmd"
package: "exec"
import: "github.com/sahilkhaire/gox/exec"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: child_process</span><span class="api-badge import">github.com/sahilkhaire/gox/exec</span></div>
# Cmd

## Overview

Cmd wraps os/exec.Cmd with context cancellation.

## Signature

```go
type Cmd struct {
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
import "github.com/sahilkhaire/gox/exec"

_ = exec.Cmd
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Exec](/packages/exec/exec)

← [Back to exec package overview](/packages/exec/)
