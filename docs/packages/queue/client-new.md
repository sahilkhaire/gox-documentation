---
title: "Client.New"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "10"
---

<SymbolHeader pkg="queue" title="Client.New" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

New creates a queue client connected to redisAddr (host:port).

## Signature

<div class="signature-block">

```go
func New(redisAddr string) *Client
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const queue = new Queue('tasks');
```

```go [Standard Go]
// hibiken/asynq client + worker
```

```go [gox]
import "github.com/sahilkhaire/gox/queue"

client, err := queue.New("localhost:6379")
```

:::

← [Back to queue package overview](/packages/queue/)
