---
title: "RequestOpts.WithHeaders"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "10"
---

<SymbolHeader pkg="client" title="RequestOpts.WithHeaders" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

WithHeaders returns RequestOpts with headers.

## Signature

<div class="signature-block">

```go
func WithHeaders(h http.Header) *RequestOpts
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

var v RequestOpts
v.WithHeaders(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/client/request-opts-json-body">RequestOpts.JSONBody</a><a class="related-chip" href="/packages/client/request-opts-with-method">RequestOpts.WithMethod</a><a class="related-chip" href="/packages/client/request-opts-with-query">RequestOpts.WithQuery</a>
</div>

← [Back to client package overview](/packages/client/)
