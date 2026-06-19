---
title: "IsAbs"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.isAbsolute(p)"
gox-doc-version: "11"
---

<SymbolHeader pkg="path" title="IsAbs" node="path.isAbsolute(p)" import-path="github.com/sahilkhaire/gox/path" />
## Overview

IsAbs reports whether p is absolute (path.isAbsolute).

If you are coming from Node.js, the closest pattern is **`path.isAbsolute(p)`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/path"

ok := path.IsAbs(p)
```

## Tips

Import `github.com/sahilkhaire/gox/path` and call `IsAbs` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/path/basename">Basename</a><a class="related-chip" href="/packages/path/dirname">Dirname</a><a class="related-chip" href="/packages/path/extname">Extname</a>
</div>

← [Back to path package overview](/packages/path/)
