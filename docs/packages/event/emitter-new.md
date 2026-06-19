---
title: "Emitter.New"
package: "event"
import: "github.com/sahilkhaire/gox/event"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: EventEmitter</span><span class="api-badge import">github.com/sahilkhaire/gox/event</span></div>
# Emitter.New

## Overview

New returns an empty Emitter.

## Signature

```go
func New() *Emitter
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const emitter = new EventEmitter();
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/event"

em := event.New()
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to event package overview](/packages/event/)
