---
title: "Dirname"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.dirname(p)"
gox-doc-version: "10"
---

<SymbolHeader pkg="path" title="Dirname" node="path.dirname(p)" import-path="github.com/sahilkhaire/gox/path" />
## Overview

Dirname returns the directory (path.dirname).

**Node.js equivalent:** `path.dirname(p)`

## Signature

<div class="signature-block">

```go
func Dirname(p string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
path.dirname('/foo/bar.txt');
```

```go [Standard Go]
dir := filepath.Dir("/foo/bar.txt")
```

```go [gox]
import "github.com/sahilkhaire/gox/path"

dir := path.Dirname("/foo/bar.txt")
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/path/basename">Basename</a><a class="related-chip" href="/packages/path/extname">Extname</a><a class="related-chip" href="/packages/path/is-abs">IsAbs</a>
</div>

← [Back to path package overview](/packages/path/)
