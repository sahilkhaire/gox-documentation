---
title: "EnqueueOpts"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "10"
---

<SymbolHeader pkg="queue" title="EnqueueOpts" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

EnqueueOpts configures task enqueue behavior.

## Signature

<div class="signature-block">

```go
type EnqueueOpts struct {
	MaxRetry int
	Queue    string
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

_ = queue.EnqueueOpts
```

:::

← [Back to queue package overview](/packages/queue/)
