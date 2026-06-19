---
title: "Resolve"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.resolve(...)"
gox-doc-version: "10"
---

<SymbolHeader pkg="path" title="Resolve" node="path.resolve(...)" import-path="github.com/sahilkhaire/gox/path" />
## Overview

Resolve resolves to an absolute path (path.resolve).

**Node.js equivalent:** `path.resolve(...)`

## Signature

<div class="signature-block">

```go
func Resolve(elem ...string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
path.resolve('/foo', 'bar');
```

```go [Standard Go]
abs, err := filepath.Abs(filepath.Join(base, rel))
```

```go [gox]
import "github.com/sahilkhaire/gox/path"

p := path.Resolve("/foo", "bar")
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/path/basename">Basename</a><a class="related-chip" href="/packages/path/dirname">Dirname</a><a class="related-chip" href="/packages/path/extname">Extname</a>
</div>

← [Back to path package overview](/packages/path/)
