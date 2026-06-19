---
title: "Emitter"
package: "event"
import: "github.com/sahilkhaire/gox/event"
gox-doc-version: "14"
---

<SymbolHeader pkg="event" title="Emitter" node="EventEmitter" import-path="github.com/sahilkhaire/gox/event" />
## Overview

Emitter dispatches named events to registered handlers.

Part of the **`event`** package — Node.js analog: *EventEmitter*.

`Emitter` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Emitter struct {
	// contains filtered or unexported fields
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical EventEmitter pattern in Node.js
```

```go [Standard Go]
// sync.Mutex + map[string][]func() or channels
```

```go [gox]
import "github.com/sahilkhaire/gox/event"

emitter := event.NewEmitter()
emitter.On("ready", func(args ...any) {
    fmt.Println("ready", args)
})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/event"

emitter := event.NewEmitter()
emitter.On("ready", func(args ...any) {
    fmt.Println("ready", args)
})
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to event package overview](/packages/event/)
