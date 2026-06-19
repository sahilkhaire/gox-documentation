---
title: "Client.New"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "14"
---

<SymbolHeader pkg="queue" title="Client.New" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

New creates a queue client connected to redisAddr (host:port).

Part of the **`queue`** package — Node.js analog: *bull*.

Method on **`Client`** — call it on a value of that type after constructing or receiving one from a constructor.

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

## Example

```go
import "github.com/sahilkhaire/gox/queue"

client, err := queue.New("localhost:6379")
```

## Tips

Obtain a `Client` value first (see constructors on the package overview), then call `New`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to queue package overview](/packages/queue/)
