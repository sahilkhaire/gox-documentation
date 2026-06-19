---
title: "RequestOpts.WithQuery"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "11"
---

<SymbolHeader pkg="client" title="RequestOpts.WithQuery" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

WithQuery returns RequestOpts with query values.

Part of the **`client`** package — Node.js analog: *axios, fetch*.

Method on **`RequestOpts`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func WithQuery(q url.Values) *RequestOpts
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical axios, fetch pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/client"

var v RequestOpts
v.WithQuery(/* args */)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/client/request-opts-json-body">RequestOpts.JSONBody</a><a class="related-chip" href="/packages/client/request-opts-with-headers">RequestOpts.WithHeaders</a><a class="related-chip" href="/packages/client/request-opts-with-method">RequestOpts.WithMethod</a>
</div>

← [Back to client package overview](/packages/client/)
