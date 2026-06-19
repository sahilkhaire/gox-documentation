---
title: "DefaultUpgrader"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "11"
---

<SymbolHeader pkg="ws" title="DefaultUpgrader" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

DefaultUpgrader allows all origins (suitable for tests; tighten in production).

Part of the **`ws`** package — Node.js analog: *ws*.

Use this constant or variable directly — no function call required.

## Signature

<div class="signature-block">

```go
var DefaultUpgrader
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical ws pattern in Node.js
```

```go [Standard Go]
// gorilla/websocket Upgrader or Dialer
```

```go [gox]
import "github.com/sahilkhaire/gox/ws"

_ = ws.DefaultUpgrader
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/ws"

_ = ws.DefaultUpgrader
```

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to ws package overview](/packages/ws/)
