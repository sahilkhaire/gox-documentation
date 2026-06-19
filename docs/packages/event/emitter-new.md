---
title: "Emitter.New"
package: "event"
import: "github.com/sahilkhaire/gox/event"
gox-doc-version: "11"
---

<SymbolHeader pkg="event" title="Emitter.New" node="EventEmitter" import-path="github.com/sahilkhaire/gox/event" />
## Overview

New returns an empty Emitter.

Part of the **`event`** package — Node.js analog: *EventEmitter*.

Method on **`Emitter`** — call it on a value of that type after constructing or receiving one from a constructor.

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

## Example

```go
import "github.com/sahilkhaire/gox/event"

em := event.New()
```

## Tips

Obtain a `Emitter` value first (see constructors on the package overview), then call `New`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to event package overview](/packages/event/)
