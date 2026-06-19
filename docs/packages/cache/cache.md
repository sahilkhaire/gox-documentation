---
title: "Cache"
package: "cache"
import: "github.com/sahilkhaire/gox/cache"
gox-doc-version: "10"
---

<SymbolHeader pkg="cache" title="Cache" node="lru-cache" import-path="github.com/sahilkhaire/gox/cache" />
## Overview

Cache is a thread-safe LRU cache with per-key TTL.

## Signature

<div class="signature-block">

```go
type Cache struct {
	// contains filtered or unexported fields
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// sync.Map or lru cache library
```

```go [gox]
import "github.com/sahilkhaire/gox/cache"

_ = cache.Cache
```

:::

← [Back to cache package overview](/packages/cache/)
