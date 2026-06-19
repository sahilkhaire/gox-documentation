---
title: "Join"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.join(...)"
gox-doc-version: "14"
---

<SymbolHeader pkg="path" title="Join" node="path.join(...)" import-path="github.com/sahilkhaire/gox/path" />
## Overview

Join joins path segments (path.join).

If you are coming from Node.js, the closest pattern is **`path.join(...)`**.

## Signature

<div class="signature-block">

```go
func Join(elem ...string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
path.join('a', 'b', 'c');
```

```go [Standard Go]
filepath.Join("a", "b", "c")
```

```go [gox]
import "github.com/sahilkhaire/gox/path"

p := path.Join("a", "b", "c")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/path"

p := path.Join("a", "b", "c")
```

## Tips

Uses OS-specific separators via filepath underneath.

## Standard library alternative

Use the standard library directly:

```go
filepath.Join("a", "b", "c")
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/path/basename">Basename</a><a class="related-chip" href="/packages/path/dirname">Dirname</a><a class="related-chip" href="/packages/path/extname">Extname</a>
</div>

← [Back to path package overview](/packages/path/)
