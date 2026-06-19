---
title: "Stat"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.stat(path)"
gox-doc-version: "11"
---

<SymbolHeader pkg="fs" title="Stat" node="fs.stat(path)" import-path="github.com/sahilkhaire/gox/fs" />
## Overview

Stat returns file info (fs.promises.stat).

If you are coming from Node.js, the closest pattern is **`fs.stat(path)`**.

## Signature

<div class="signature-block">

```go
func Stat(ctx context.Context, path string) (os.FileInfo, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.stat(path)
```

```go [Standard Go]
// os / io helpers — e.g. os.stat
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.Stat(ctx, path)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/fs"

fs.Stat(ctx, path)
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use `os.ReadFile`, `os.WriteFile`, etc. Pass `context.Context` manually for cancellation.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/fs/copy">Copy</a><a class="related-chip" href="/packages/fs/exists">Exists</a><a class="related-chip" href="/packages/fs/mkdir">Mkdir</a>
</div>

← [Back to fs package overview](/packages/fs/)
