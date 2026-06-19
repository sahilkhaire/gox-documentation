---
title: "FileEntry"
package: "archive"
import: "github.com/sahilkhaire/gox/archive"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: archiver</span><span class="api-badge import">github.com/sahilkhaire/gox/archive</span></div>
# FileEntry

## Overview

FileEntry is a named file for archive creation.

## Signature

```go
type FileEntry struct {
	Name string
	Data []byte
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
import "github.com/sahilkhaire/gox/archive"

_ = archive.FileEntry
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
- [ZipCreate](/packages/archive/zip-create)

← [Back to archive package overview](/packages/archive/)
