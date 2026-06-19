---
title: "Join"
package: "path"
import: "github.com/sahilkhaire/gox/path"
node: "path.join(...)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: path.join(...)</span><span class="api-badge import">github.com/sahilkhaire/gox/path</span></div>
# Join

## Overview

Join joins path segments (path.join).

**Node.js equivalent:** `path.join(...)`

## Signature

```go
func Join(elem ...string) string
```

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

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Basename](/packages/path/basename)
- [Dirname](/packages/path/dirname)
- [Extname](/packages/path/extname)

← [Back to path package overview](/packages/path/)
