---
title: "Exec"
package: "exec"
import: "github.com/sahilkhaire/gox/exec"
node: "exec('ls -la')"
gox-doc-version: "14"
---

<SymbolHeader pkg="exec" title="Exec" node="exec('ls -la')" import-path="github.com/sahilkhaire/gox/exec" />
## Overview

Exec runs a shell-less command string split on whitespace (child_process.exec).

If you are coming from Node.js, the closest pattern is **`exec('ls -la')`**.

## Signature

<div class="signature-block">

```go
func Exec(ctx context.Context, command string) ([]byte, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
exec('ls -la')
```

```go [Standard Go]
cmd := exec.CommandContext(ctx, name, args...)
out, err := cmd.CombinedOutput()
```

```go [gox]
import "github.com/sahilkhaire/gox/exec"

exec.Exec(ctx, command)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/exec"

exec.Exec(ctx, command)
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
