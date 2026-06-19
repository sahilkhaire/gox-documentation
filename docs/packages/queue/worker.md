---
title: "Worker"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "14"
---

<SymbolHeader pkg="queue" title="Worker" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

Worker processes tasks from Redis.

Part of the **`queue`** package — Node.js analog: *bull*.

`Worker` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Worker struct {
	// contains filtered or unexported fields
}
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

mux := queue.NewServeMux()
mux.HandleFunc("email", handleEmail)
worker := queue.NewWorker("localhost:6379", mux)
worker.Run()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/queue"

mux := queue.NewServeMux()
mux.HandleFunc("email", handleEmail)
worker := queue.NewWorker("localhost:6379", mux)
worker.Run()
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to queue package overview](/packages/queue/)
