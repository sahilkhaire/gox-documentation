---
title: "Extname"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.extname(p)"
gox-doc-version: "14"
---

<SymbolHeader pkg="path" title="Extname" node="path.extname(p)" import-path="github.com/sahilkhaire/gox/path" />
## Overview

Extname returns the extension including dot (path.extname).

If you are coming from Node.js, the closest pattern is **`path.extname(p)`**.

## Signature

<div class="signature-block">

```go
func Extname(p string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
path.extname('file.txt');
```

```go [Standard Go]
ext := filepath.Ext("file.txt")
```

```go [gox]
import "github.com/sahilkhaire/gox/path"

ext := path.Extname("file.txt")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/path"

ext := path.Extname("file.txt")
```

## Tips

Import `github.com/sahilkhaire/gox/path` and call `Extname` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
ext := filepath.Ext("file.txt")
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/path/basename">Basename</a><a class="related-chip" href="/packages/path/dirname">Dirname</a><a class="related-chip" href="/packages/path/is-abs">IsAbs</a>
</div>

← [Back to path package overview](/packages/path/)
