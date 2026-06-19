---
title: "Parse"
package: "json"
import: "github.com/sahilkhaire/gox/json"
node: "JSON.parse(str)"
gox-doc-version: "10"
---

<SymbolHeader pkg="json" title="Parse" node="JSON.parse(str)" import-path="github.com/sahilkhaire/gox/json" />
## Overview

Parse JSON into a typed value using generics — like JSON.parse but with compile-time type safety.

## Signature

<div class="signature-block">

```go
func Parse(data []byte, v any) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const obj = JSON.parse(str);
```

```go [Standard Go]
err := json.Unmarshal([]byte(str), &obj)
```

```go [gox]
import "github.com/sahilkhaire/gox/json"

obj, err := json.Parse[MyType](str)
```

:::

## Tips

Use MustParse when invalid JSON should panic in init or tests.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/json/must-parse">MustParse</a><a class="related-chip" href="/packages/json/must-stringify">MustStringify</a><a class="related-chip" href="/packages/json/parse-file">ParseFile</a>
</div>

← [Back to json package overview](/packages/json/)
