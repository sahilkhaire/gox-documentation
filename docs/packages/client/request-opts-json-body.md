---
title: "RequestOpts.JSONBody"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "14"
---

<SymbolHeader pkg="client" title="RequestOpts.JSONBody" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

JSONBody returns RequestOpts with a JSON body.

Part of the **`client`** package — Node.js analog: *axios, fetch*.

Method on **`RequestOpts`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func JSONBody(v any) *RequestOpts
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

opts := client.RequestOpts{}.JSONBody(map[string]string{"name": "Ada"})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/client"

opts := client.RequestOpts{}.JSONBody(map[string]string{"name": "Ada"})
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use the standard library directly:

```go
resp, err := http.NewRequestWithContext(ctx, method, url, body)
client.Do(req)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/client/request-opts-with-headers">RequestOpts.WithHeaders</a><a class="related-chip" href="/packages/client/request-opts-with-method">RequestOpts.WithMethod</a><a class="related-chip" href="/packages/client/request-opts-with-query">RequestOpts.WithQuery</a>
</div>

← [Back to client package overview](/packages/client/)
