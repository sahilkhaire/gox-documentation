---
title: "TarExtract"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: archiver</span><span class="api-badge import">github.com/sahilkhaire/gox/archive</span></div>
# TarExtract

## Overview

TarExtract extracts a gzip-compressed tar archive into destDir.

## Signature

```go
func TarExtract(data []byte, destDir string) error
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
import "github.com/sahilkhaire/gox/archive"

// archive
_ = archive.TarExtract(/* args */)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [TarCreate](/packages/archive/tar-create)
- [ZipCreate](/packages/archive/zip-create)
- [ZipCreateEntries](/packages/archive/zip-create-entries)

← [Back to archive package overview](/packages/archive/)
