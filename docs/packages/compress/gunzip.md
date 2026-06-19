---
title: "Gunzip"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
node: "gunzipSync"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: gunzipSync</span><span class="api-badge import">github.com/sahilkhaire/gox/compress</span></div>
# Gunzip

## Overview

Maps the Node.js pattern `gunzipSync` to gox `compress.Gunzip(data)`.

**Node.js equivalent:** `gunzipSync`

## Signature

```go
func Gunzip(data []byte) ([]byte, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
gunzipSync
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/compress"

compress.Gunzip(data)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Gzip](/packages/compress/gzip)
- [GzipReader](/packages/compress/gzip-reader)
- [GzipWriter](/packages/compress/gzip-writer)

← [Back to compress package overview](/packages/compress/)
