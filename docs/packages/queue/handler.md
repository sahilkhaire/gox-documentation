---
title: "Handler"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "10"
---

<SymbolHeader pkg="queue" title="Handler" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

Handler processes a task payload.

## Signature

<div class="signature-block">

```go
type Handler func(ctx context.Context, payload []byte) error
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

_ = queue.Handler
```

:::

← [Back to queue package overview](/packages/queue/)
