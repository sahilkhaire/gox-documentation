---
title: "Stringify"
package: "json"
import: "github.com/sahilkhaire/gox/json"
node: "JSON.stringify(obj)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: JSON.stringify(obj)</span><span class="api-badge import">github.com/sahilkhaire/gox/json</span></div>
# Stringify

## Overview

Stringify marshals v to a string (JSON.stringify).

**Node.js equivalent:** `JSON.stringify(obj)`

## Signature

```go
func Stringify(v any) (string, error)
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const str = JSON.stringify(obj);
```

```go [Standard Go]
b, err := json.Marshal(obj)
```

```go [gox]
import "github.com/sahilkhaire/gox/json"

str, err := json.Stringify(obj)
```

:::

## Tips

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [MustParse](/packages/json/must-parse)
- [MustStringify](/packages/json/must-stringify)
- [Parse](/packages/json/parse)

← [Back to json package overview](/packages/json/)
