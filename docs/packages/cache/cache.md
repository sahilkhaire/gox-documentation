---
title: "Cache"
package: "cache"
import: "github.com/sahilkhaire/gox/cache"
gox-doc-version: "14"
---

<SymbolHeader pkg="cache" title="Cache" node="lru-cache" import-path="github.com/sahilkhaire/gox/cache" />
## Overview

Cache is a thread-safe LRU cache with per-key TTL.

Part of the **`cache`** package — Node.js analog: *lru-cache*.

`Cache` is a type exported by gox. Methods on this type are documented separately.

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
// Typical lru-cache pattern in Node.js
```

```go [Standard Go]
// sync.Map or lru cache library
```

```go [gox]
import "github.com/sahilkhaire/gox/cache"

c := cache.New(500)
c.Set("user:1", user, time.Minute)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/cache"

c := cache.New(500)
c.Set("user:1", user, time.Minute)
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to cache package overview](/packages/cache/)
