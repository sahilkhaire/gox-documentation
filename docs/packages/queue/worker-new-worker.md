---
title: "Worker.NewWorker"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "10"
---

<SymbolHeader pkg="queue" title="Worker.NewWorker" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

NewWorker creates a worker that dispatches to mux.

## Signature

<div class="signature-block">

```go
func NewWorker(redisAddr string, mux *ServeMux) *Worker
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

var v Worker
v.NewWorker(/* args */)
```

:::

← [Back to queue package overview](/packages/queue/)
