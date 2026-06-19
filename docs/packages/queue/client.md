---
title: "Client"
package: "queue"
import: "github.com/sahilkhaire/gox/queue"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: bull</span><span class="api-badge import">github.com/sahilkhaire/gox/queue</span></div>
# Client

## Overview

Client enqueues background tasks.

## Signature

```go
type Client struct {
	// contains filtered or unexported fields
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
import "github.com/sahilkhaire/gox/queue"

_ = queue.Client
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/queue"

mr, err := miniredis.Run()
defer mr.Close()
c := queue.New(mr.Addr())
defer c.Close()
mux := queue.NewServeMux()
mux.HandleFunc("email:send", func(ctx context.Context, payload []byte) error {
	return nil
})
```

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

← [Back to queue package overview](/packages/queue/)
