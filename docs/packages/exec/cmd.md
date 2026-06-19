---
title: "Cmd"
package: "exec"
import: "github.com/sahilkhaire/gox/exec"
gox-doc-version: "14"
---

<SymbolHeader pkg="exec" title="Cmd" node="child_process" import-path="github.com/sahilkhaire/gox/exec" />
## Overview

Cmd wraps os/exec.Cmd with context cancellation.

Part of the **`exec`** package — Node.js analog: *child_process*.

`Cmd` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Cmd struct {
	// contains filtered or unexported fields
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical child_process pattern in Node.js
```

```go [Standard Go]
cmd := exec.CommandContext(ctx, name, args...)
out, err := cmd.CombinedOutput()
```

```go [gox]
import "github.com/sahilkhaire/gox/exec"

cmd := exec.Command(ctx, "git", "status")
out, err := cmd.Output()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/exec"

cmd := exec.Command(ctx, "git", "status")
out, err := cmd.Output()
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use the standard library directly:

```go
cmd := exec.CommandContext(ctx, name, args...)
out, err := cmd.CombinedOutput()
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/exec/exec">Exec</a>
</div>

← [Back to exec package overview](/packages/exec/)
