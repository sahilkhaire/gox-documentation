---
title: "ServeMux.NewServeMux"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "14"
---

<SymbolHeader pkg="queue" title="ServeMux.NewServeMux" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

NewServeMux returns an empty task router.

Part of the **`queue`** package — Node.js analog: *bull*.

Method on **`ServeMux`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func NewServeMux() *ServeMux
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
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/queue"

mux := queue.NewServeMux()
```

## Tips

Obtain a `ServeMux` value first (see constructors on the package overview), then call `NewServeMux`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to queue package overview](/packages/queue/)
