---
title: "Gzip"
package: "compress"
import: "github.com/sahilkhaire/gox/compress"
node: "gzipSync"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: gzipSync</span><span class="api-badge import">github.com/sahilkhaire/gox/compress</span></div>
# Gzip

## Overview

Gzip compresses data with gzip.

**Node.js equivalent:** `gzipSync`

## Signature

```go
func Gzip(data []byte) ([]byte, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
zlib.gzipSync(data);
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches (e.g. db.SQL, redis.RDB).
```

```go [gox]
import "github.com/sahilkhaire/gox/compress"

compressed, err := compress.Gzip(data)
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
- [GzipReader](/packages/compress/gzip-reader)
- [GzipWriter](/packages/compress/gzip-writer)

← [Back to compress package overview](/packages/compress/)
