---
title: "RequestOpts.JSONBody"
package: "client"
import: "github.com/sahilkhaire/gox/client"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: axios, fetch</span><span class="api-badge import">github.com/sahilkhaire/gox/client</span></div>
# RequestOpts.JSONBody

## Overview

JSONBody returns RequestOpts with a JSON body.

## Signature

```go
func JSONBody(v any) *RequestOpts
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
v.JSONBody(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [RequestOpts.WithHeaders](/packages/client/request-opts-with-headers)
- [RequestOpts.WithMethod](/packages/client/request-opts-with-method)
- [RequestOpts.WithQuery](/packages/client/request-opts-with-query)

← [Back to client package overview](/packages/client/)
