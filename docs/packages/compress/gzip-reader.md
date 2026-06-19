---
title: "GzipReader"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: zlib</span><span class="api-badge import">github.com/sahilkhaire/gox/compress</span></div>
# GzipReader

## Overview

GzipReader wraps a gzip reader over src.

## Signature

```go
func GzipReader(src io.Reader) (*gzip.Reader, error)
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
import "github.com/sahilkhaire/gox/compress"

// compress
_ = compress.GzipReader(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Gunzip](/packages/compress/gunzip)
- [Gzip](/packages/compress/gzip)
- [GzipWriter](/packages/compress/gzip-writer)

← [Back to compress package overview](/packages/compress/)
