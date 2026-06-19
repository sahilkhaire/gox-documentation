---
title: "Cache.New"
package: "cache"
import: "github.com/sahilkhaire/gox/cache"
node: "new LRUCache({ max })"
gox-doc-version: "10"
---

<SymbolHeader pkg="cache" title="Cache.New" node="new LRUCache({ max })" import-path="github.com/sahilkhaire/gox/cache" />
## Overview

New creates a cache that evicts least-recently-used entries when size exceeds maxSize.

**Node.js equivalent:** `new LRUCache({ max })`

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

← [Back to cache package overview](/packages/cache/)
