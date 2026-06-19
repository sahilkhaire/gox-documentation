---
title: "DefaultUpgrader"
package: "ws"
import: "github.com/sahilkhaire/gox/ws"
gox-doc-version: "10"
---

<SymbolHeader pkg="ws" title="DefaultUpgrader" node="ws" import-path="github.com/sahilkhaire/gox/ws" />
## Overview

DefaultUpgrader allows all origins (suitable for tests; tighten in production).

## Signature

<div class="signature-block">

```go
var DefaultUpgrader
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// gorilla/websocket Upgrader or Dialer
```

```go [gox]
import "github.com/sahilkhaire/gox/ws"

_ = ws.DefaultUpgrader
```

:::

← [Back to ws package overview](/packages/ws/)
