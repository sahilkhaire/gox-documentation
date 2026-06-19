---
title: "Worker"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "10"
---

<SymbolHeader pkg="queue" title="Worker" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

Worker processes tasks from Redis.

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
// See package overview
```

```go [Standard Go]
// hibiken/asynq client + worker
```

```go [gox]
import "github.com/sahilkhaire/gox/queue"

_ = queue.Worker
```

:::

← [Back to queue package overview](/packages/queue/)
