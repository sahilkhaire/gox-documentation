---
title: "Cache.New"
package: "cache"
import: "github.com/sahilkhaire/gox/cache"
node: "new LRUCache({ max })"
gox-doc-version: "14"
---

<SymbolHeader pkg="cache" title="Cache.New" node="new LRUCache({ max })" import-path="github.com/sahilkhaire/gox/cache" />
## Overview

New creates a cache that evicts least-recently-used entries when size exceeds maxSize.
Node: new LRUCache({ max: maxSize })

If you are coming from Node.js, the closest pattern is **`new LRUCache({ max })`**.

Method on **`Cache`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func New(maxSize int) *Cache
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const cache = new LRUCache({ max: 100 });
```

```go [Standard Go]
// sync.Map or lru cache library
```

```go [gox]
import "github.com/sahilkhaire/gox/cache"

c := cache.New(100)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/cache"

c := cache.New(100)
```

## Tips

Obtain a `Cache` value first (see constructors on the package overview), then call `New`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to cache package overview](/packages/cache/)
