---
title: "Exec"
package: "exec"
import: "github.com/sahilkhaire/gox/exec"
node: "exec('ls -la')"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: exec('ls -la')</span><span class="api-badge import">github.com/sahilkhaire/gox/exec</span></div>
# Exec

## Overview

Maps the Node.js pattern `exec('ls -la')` to gox `exec.Exec(ctx, command)`.

**Node.js equivalent:** `exec('ls -la')`

## Signature

```go
func Exec(ctx context.Context, command string) ([]byte, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
exec('ls -la')
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/exec"

exec.Exec(ctx, command)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to exec package overview](/packages/exec/)
