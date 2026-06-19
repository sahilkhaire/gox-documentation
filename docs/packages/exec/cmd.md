---
title: "Cmd"
package: "exec"
import: "github.com/sahilkhaire/gox/exec"
gox-doc-version: "10"
---

<SymbolHeader pkg="exec" title="Cmd" node="child_process" import-path="github.com/sahilkhaire/gox/exec" />
## Overview

Cmd wraps os/exec.Cmd with context cancellation.

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
// See package overview
```

```go [Standard Go]
cmd := exec.CommandContext(ctx, name, args...)
out, err := cmd.CombinedOutput()
```

```go [gox]
import "github.com/sahilkhaire/gox/exec"

_ = exec.Cmd
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/exec/exec">Exec</a>
</div>

← [Back to exec package overview](/packages/exec/)
