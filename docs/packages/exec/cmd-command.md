---
title: "Cmd.Command"
package: "exec"
import: "github.com/sahilkhaire/gox/exec"
node: "spawn(cmd, args)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: spawn(cmd, args)</span><span class="api-badge import">github.com/sahilkhaire/gox/exec</span></div>
# Cmd.Command

## Overview

Command builds a command (child_process.spawn / execFile).

**Node.js equivalent:** `spawn(cmd, args)`

## Signature

```go
func Command(ctx context.Context, name string, args ...string) *Cmd
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
spawn('ls', ['-la']);
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/exec"

cmd := exec.Command(ctx, "ls", "-la")
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to exec package overview](/packages/exec/)
