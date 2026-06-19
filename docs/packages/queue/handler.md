---
title: "Handler"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "11"
---

<SymbolHeader pkg="queue" title="Handler" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

Handler processes a task payload.

Part of the **`queue`** package — Node.js analog: *bull*.

`Handler` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Handler func(ctx context.Context, payload []byte) error
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

_ = queue.Handler
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/queue"

_ = queue.Handler
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to queue package overview](/packages/queue/)
