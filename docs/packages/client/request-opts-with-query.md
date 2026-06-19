---
title: "RequestOpts.WithQuery"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "10"
---

<SymbolHeader pkg="client" title="RequestOpts.WithQuery" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

WithQuery returns RequestOpts with query values.

## Signature

<div class="signature-block">

```go
func WithQuery(q url.Values) *RequestOpts
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
v.WithQuery(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/client/request-opts-json-body">RequestOpts.JSONBody</a><a class="related-chip" href="/packages/client/request-opts-with-headers">RequestOpts.WithHeaders</a><a class="related-chip" href="/packages/client/request-opts-with-method">RequestOpts.WithMethod</a>
</div>

← [Back to client package overview](/packages/client/)
