---
title: "IsAbs"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.isAbsolute(p)"
gox-doc-version: "10"
---

<SymbolHeader pkg="path" title="IsAbs" node="path.isAbsolute(p)" import-path="github.com/sahilkhaire/gox/path" />
## Overview

IsAbs reports whether p is absolute (path.isAbsolute).

**Node.js equivalent:** `path.isAbsolute(p)`

## Signature

<div class="signature-block">

```go
func IsAbs(p string) bool
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
path.isAbsolute('/foo');
```

```go [Standard Go]
ok := filepath.IsAbs(p)
```

```go [gox]
import "github.com/sahilkhaire/gox/path"

ok := path.IsAbs(p)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/path/basename">Basename</a><a class="related-chip" href="/packages/path/dirname">Dirname</a><a class="related-chip" href="/packages/path/extname">Extname</a>
</div>

← [Back to path package overview](/packages/path/)
