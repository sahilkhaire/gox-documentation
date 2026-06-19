---
title: "Emitter"
package: "event"
import: "github.com/sahilkhaire/gox/event"
gox-doc-version: "10"
---

<SymbolHeader pkg="event" title="Emitter" node="EventEmitter" import-path="github.com/sahilkhaire/gox/event" />
## Overview

Emitter dispatches named events to registered handlers.

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
// See package overview
```

```go [Standard Go]
// sync.Mutex + map[string][]func() or channels
```

```go [gox]
import "github.com/sahilkhaire/gox/event"

_ = event.Emitter
```

:::

← [Back to event package overview](/packages/event/)
