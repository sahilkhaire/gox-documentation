---
title: "WSHandler"
package: "http"
import: "github.com/sahilkhaire/gox/http"
gox-doc-version: "10"
---

<SymbolHeader pkg="http" title="WSHandler" node="express, cors, helmet, morgan" import-path="github.com/sahilkhaire/gox/http" />
## Overview

WSHandler handles a WebSocket connection in an Express-style handler.

## Signature

<div class="signature-block">

```go
type WSHandler func(*Ctx, *goxws.Conn) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
func handler(w http.ResponseWriter, r *http.Request) {
    // chi or net/http
}
```

```go [gox]
import "github.com/sahilkhaire/gox/http"

_ = http.WSHandler
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/http/save-uploaded-file">SaveUploadedFile</a>
</div>

← [Back to http package overview](/packages/http/)
