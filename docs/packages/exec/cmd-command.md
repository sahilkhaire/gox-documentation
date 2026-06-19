---
title: "Cmd.Command"
package: "exec"
import: "github.com/sahilkhaire/gox/exec"
node: "spawn(cmd, args)"
gox-doc-version: "14"
---

<SymbolHeader pkg="exec" title="Cmd.Command" node="spawn(cmd, args)" import-path="github.com/sahilkhaire/gox/exec" />
## Overview

Command builds a command (child_process.spawn / execFile).

If you are coming from Node.js, the closest pattern is **`spawn(cmd, args)`**.

Method on **`Cmd`** — call it on a value of that type after constructing or receiving one from a constructor.

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

## Example

```go
import "github.com/sahilkhaire/gox/exec"

cmd := exec.Command(ctx, "ls", "-la")
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use the standard library directly:

```go
cmd := exec.CommandContext(ctx, name, args...)
out, err := cmd.CombinedOutput()
```

← [Back to exec package overview](/packages/exec/)
