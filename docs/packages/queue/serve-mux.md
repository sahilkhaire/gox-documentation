---
title: "ServeMux"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "10"
---

<SymbolHeader pkg="queue" title="ServeMux" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

ServeMux routes tasks by type name.

## Signature

<div class="signature-block">

```go
type ServeMux struct {
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
// hibiken/asynq client + worker
```

```go [gox]
import "github.com/sahilkhaire/gox/queue"

_ = queue.ServeMux
```

:::

← [Back to queue package overview](/packages/queue/)
