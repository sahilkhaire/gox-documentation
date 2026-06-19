---
title: "ReadFile"
package: "fs"
import: "github.com/sahilkhaire/gox/fs"
node: "fs.readFile(path)"
gox-doc-version: "11"
---

<SymbolHeader pkg="fs" title="ReadFile" node="fs.readFile(path)" import-path="github.com/sahilkhaire/gox/fs" />
## Overview

Reads an entire file with context cancellation — fs.promises.readFile with Go context support.

If you are coming from Node.js, the closest pattern is **`fs.readFile(path)`**.

## Signature

<div class="signature-block">

```go
func ReadFile(ctx context.Context, path string) ([]byte, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const data = await fs.promises.readFile('file.txt');
```

```go [Standard Go]
data, err := os.ReadFile("file.txt")
```

```go [gox]
import "github.com/sahilkhaire/gox/fs"

data, err := fs.ReadFile(ctx, "file.txt")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/fs"

data, err := fs.ReadFile(ctx, "file.txt")
```

## Tips

Prefer context-aware I/O so requests can cancel long reads.

## Standard library alternative

Use `os.ReadFile`, `os.WriteFile`, etc. Pass `context.Context` manually for cancellation.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/fs/copy">Copy</a><a class="related-chip" href="/packages/fs/exists">Exists</a><a class="related-chip" href="/packages/fs/mkdir">Mkdir</a>
</div>

← [Back to fs package overview](/packages/fs/)
