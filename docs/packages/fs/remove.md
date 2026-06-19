---
title: "Remove"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.rm(path)"
gox-doc-version: "14"
---

<SymbolHeader pkg="fs" title="Remove" node="fs.rm(path)" import-path="github.com/sahilkhaire/gox/fs" />
## Overview

Remove deletes path (fs.promises.rm).

If you are coming from Node.js, the closest pattern is **`fs.rm(path)`**.

## Signature

<div class="signature-block">

```go
func Remove(ctx context.Context, path string) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.rm(path)
```

```go [Standard Go]
// os / io helpers — e.g. os.remove
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

fs.Remove(ctx, path)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/fs"

fs.Remove(ctx, path)
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
