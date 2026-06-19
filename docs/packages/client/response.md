---
title: "Response"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "10"
---

<SymbolHeader pkg="client" title="Response" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

Response wraps an HTTP response.

## Signature

<div class="signature-block">

```go
type Response struct {
	*http.Response
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
resp, err := http.NewRequestWithContext(ctx, method, url, body)
client.Do(req)
```

```go [gox]
import "github.com/sahilkhaire/gox/client"

_ = client.Response
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/client/set-default-client">SetDefaultClient</a>
</div>

← [Back to client package overview](/packages/client/)
