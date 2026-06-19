---
title: "ZipCreate"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: archiver</span><span class="api-badge import">github.com/sahilkhaire/gox/archive</span></div>
# ZipCreate

## Overview

ZipCreate builds a zip archive from a map of path -&gt; contents.

## Signature

```go
func ZipCreate(files map[string][]byte) ([]byte, error)
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
_ = archive.ZipCreate(/* args */)
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
- [TarExtract](/packages/archive/tar-extract)
- [ZipCreateEntries](/packages/archive/zip-create-entries)

← [Back to archive package overview](/packages/archive/)
