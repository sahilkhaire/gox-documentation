---
title: "Basename"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.basename(p)"
gox-doc-version: "14"
---

<SymbolHeader pkg="path" title="Basename" node="path.basename(p)" import-path="github.com/sahilkhaire/gox/path" />
## Overview

Basename returns the last element (path.basename).

If you are coming from Node.js, the closest pattern is **`path.basename(p)`**.

## Signature

<div class="signature-block">

```go
func Basename(p string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
path.basename('/foo/bar.txt');
```

```go [Standard Go]
name := filepath.Base("/foo/bar.txt")
```

```go [gox]
import "github.com/sahilkhaire/gox/path"

name := path.Basename("/foo/bar.txt")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/path"

name := path.Basename("/foo/bar.txt")
```

## Tips

Import `github.com/sahilkhaire/gox/path` and call `Basename` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
name := filepath.Base("/foo/bar.txt")
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/path/dirname">Dirname</a><a class="related-chip" href="/packages/path/extname">Extname</a><a class="related-chip" href="/packages/path/is-abs">IsAbs</a>
</div>

← [Back to path package overview](/packages/path/)
