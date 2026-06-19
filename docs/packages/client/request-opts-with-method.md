---
title: "RequestOpts.WithMethod"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: axios, fetch</span><span class="api-badge import">github.com/sahilkhaire/gox/client</span></div>
# RequestOpts.WithMethod

## Overview

WithMethod sets method on opts (helper for Fetch).

## Signature

```go
func WithMethod(method string) *RequestOpts
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

var v RequestOpts
v.WithMethod(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [RequestOpts.JSONBody](/packages/client/request-opts-json-body)
- [RequestOpts.WithHeaders](/packages/client/request-opts-with-headers)
- [RequestOpts.WithQuery](/packages/client/request-opts-with-query)

← [Back to client package overview](/packages/client/)
