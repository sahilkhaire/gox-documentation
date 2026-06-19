---
title: "Client"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: axios, fetch</span><span class="api-badge import">github.com/sahilkhaire/gox/client</span></div>
# Client

## Overview

Client is a reusable HTTP client with defaults (axios instance).

## Signature

```go
type Client struct {
	HTTP           *http.Client
	BaseURL        string
	DefaultHeaders http.Header
	RetryOn5xx     bool
}
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/client"

_ = client.Client
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [SetDefaultClient](/packages/client/set-default-client)

← [Back to client package overview](/packages/client/)
