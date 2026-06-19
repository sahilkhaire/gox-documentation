---
title: "Response"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "14"
---

<SymbolHeader pkg="client" title="Response" node="axios, fetch" import-path="github.com/sahilkhaire/gox/client" />
## Overview

Response wraps an HTTP response.

Part of the **`client`** package — Node.js analog: *axios, fetch*.

`Response` is a type exported by gox. Methods on this type are documented separately.

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
// Typical axios, fetch pattern in Node.js
```

```go [Standard Go]
resp, err := http.NewRequestWithContext(ctx, method, url, body)
client.Do(req)
```

```go [gox]
import "github.com/sahilkhaire/gox/client"

if resp.StatusCode() != 200 {
    return fmt.Errorf("unexpected status %d", resp.StatusCode())
}
var body User
if err := resp.JSON(&body); err != nil {
    return err
}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/client"

if resp.StatusCode() != 200 {
    return fmt.Errorf("unexpected status %d", resp.StatusCode())
}
var body User
if err := resp.JSON(&body); err != nil {
    return err
}
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
<a class="related-chip" href="/packages/client/set-default-client">SetDefaultClient</a>
</div>

← [Back to client package overview](/packages/client/)
