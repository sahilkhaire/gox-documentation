---
title: "RequestOpts"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "11"
---

<SymbolHeader pkg="client" title="RequestOpts" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

RequestOpts configures a single request.

Part of the **`client`** package — Node.js analog: *axios, fetch*.

`RequestOpts` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type RequestOpts struct {
	Method  string
	URL     string
	Body    io.Reader
	Headers http.Header
	Query   url.Values
	JSON    any
}
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

_ = client.RequestOpts
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/client"

_ = client.RequestOpts
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/client/set-default-client">SetDefaultClient</a>
</div>

← [Back to client package overview](/packages/client/)
