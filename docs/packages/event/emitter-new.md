---
title: "Emitter.New"
package: "event"
import: "github.com/sahilkhaire/gox/event"
gox-doc-version: "10"
---

<SymbolHeader pkg="event" title="Emitter.New" node="EventEmitter" import-path="github.com/sahilkhaire/gox/event" />
## Overview

New returns an empty Emitter.

## Signature

<div class="signature-block">

```go
func New() *Emitter
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const emitter = new EventEmitter();
```

```go [Standard Go]
// sync.Mutex + map[string][]func() or channels
```

```go [gox]
import "github.com/sahilkhaire/gox/event"

em := event.New()
```

:::

← [Back to event package overview](/packages/event/)
