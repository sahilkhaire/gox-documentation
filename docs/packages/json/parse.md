---
title: "Parse"
package: "json"
import: "github.com/sahilkhaire/gox/json"
node: "JSON.parse(str)"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: JSON.parse(str)</span><span class="api-badge import">github.com/sahilkhaire/gox/json</span></div>
# Parse

## Overview

Parse unmarshals data into v (JSON.parse).

**Node.js equivalent:** `JSON.parse(str)`

## Signature

```go
func Parse(data []byte, v any) error
```

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const obj = JSON.parse(str);
```

```go [Standard Go]
json.Unmarshal([]byte(str), &obj)
```

```go [gox]
import "github.com/sahilkhaire/gox/json"

obj, err := json.Parse[MyType](str)
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
- [ParseFile](/packages/json/parse-file)

← [Back to json package overview](/packages/json/)
