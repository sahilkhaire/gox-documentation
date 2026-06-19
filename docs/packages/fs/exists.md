---
title: "Exists"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.exists(path)"
gox-doc-version: "11"
---

<SymbolHeader pkg="fs" title="Exists" node="fs.exists(path)" import-path="github.com/sahilkhaire/gox/fs" />
## Overview

Exists reports whether path exists.

If you are coming from Node.js, the closest pattern is **`fs.exists(path)`**.

## Signature

<div class="signature-block">

```go
func Exists(ctx context.Context, path string) (bool, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
await fs.promises.access('path');
```

```go [Standard Go]
_, err := os.Stat(path)
exists := err == nil
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

ok, err := fs.Exists(ctx, "path")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/fs"

ok, err := fs.Exists(ctx, "path")
```

## Tips

Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly.

## Standard library alternative

Use `os.ReadFile`, `os.WriteFile`, etc. Pass `context.Context` manually for cancellation.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/fs/copy">Copy</a><a class="related-chip" href="/packages/fs/mkdir">Mkdir</a><a class="related-chip" href="/packages/fs/read-dir">ReadDir</a>
</div>

← [Back to fs package overview](/packages/fs/)
