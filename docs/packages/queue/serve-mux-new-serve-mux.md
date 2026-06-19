---
title: "ServeMux.NewServeMux"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "10"
---

<SymbolHeader pkg="queue" title="ServeMux.NewServeMux" node="bull" import-path="github.com/sahilkhaire/gox/queue" />
## Overview

NewServeMux returns an empty task router.

## Signature

<div class="signature-block">

```go
func NewServeMux() *ServeMux
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

var v ServeMux
v.NewServeMux(/* args */)
```

:::

← [Back to queue package overview](/packages/queue/)
