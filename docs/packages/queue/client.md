---
title: "Client"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "10"
---

<SymbolHeader pkg="queue" title="Client" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

Client enqueues background tasks.

## Signature

<div class="signature-block">

```go
type Client struct {
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

_ = queue.Client
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/queue"

mr, err := miniredis.Run()
defer mr.Close()
c := queue.New(mr.Addr())
defer c.Close()
mux := queue.NewServeMux()
mux.HandleFunc("email:send", func(ctx context.Context, payload []byte) error {
	return nil
})
```

← [Back to queue package overview](/packages/queue/)
