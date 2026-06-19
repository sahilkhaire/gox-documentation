---
title: "Worker.NewWorker"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "14"
---

<SymbolHeader pkg="queue" title="Worker.NewWorker" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

NewWorker creates a worker that dispatches to mux.

Part of the **`queue`** package — Node.js analog: *bull*.

Method on **`Worker`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func NewWorker(redisAddr string, mux *ServeMux) *Worker
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical bull pattern in Node.js
```

```go [Standard Go]
// hibiken/asynq client + worker
```

```go [gox]
import "github.com/sahilkhaire/gox/queue"

worker := queue.NewWorker("localhost:6379", mux)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/queue"

worker := queue.NewWorker("localhost:6379", mux)
```

## Tips

Obtain a `Worker` value first (see constructors on the package overview), then call `NewWorker`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to queue package overview](/packages/queue/)
