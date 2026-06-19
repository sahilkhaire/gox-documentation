---
title: "GzipWriter"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: zlib</span><span class="api-badge import">github.com/sahilkhaire/gox/compress</span></div>
# GzipWriter

## Overview

GzipWriter wraps dst with a gzip writer; caller must Close the writer.

## Signature

```go
func GzipWriter(dst io.Writer) *gzip.Writer
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
_ = compress.GzipWriter(/* args */)
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
- [GzipReader](/packages/compress/gzip-reader)

← [Back to compress package overview](/packages/compress/)
