---
title: "Cmd.Command"
package: "exec"
import: "github.com/sahilkhaire/gox/exec"
node: "spawn(cmd, args)"
gox-doc-version: "10"
---

<SymbolHeader pkg="exec" title="Cmd.Command" node="spawn(cmd, args)" import-path="github.com/sahilkhaire/gox/exec" />
## Overview

Command builds a command (child_process.spawn / execFile).

**Node.js equivalent:** `spawn(cmd, args)`

## Signature

<div class="signature-block">

```go
func Command(ctx context.Context, name string, args ...string) *Cmd
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
spawn('ls', ['-la']);
```

```go [Standard Go]
cmd := exec.CommandContext(ctx, name, args...)
out, err := cmd.CombinedOutput()
```

```go [gox]
import "github.com/sahilkhaire/gox/exec"

cmd := exec.Command(ctx, "ls", "-la")
```

:::

← [Back to exec package overview](/packages/exec/)
