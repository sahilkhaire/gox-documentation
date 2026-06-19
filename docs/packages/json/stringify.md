---
title: "Stringify"
package: "json"
import: "github.com/sahilkhaire/gox/json"
node: "JSON.stringify(obj)"
gox-doc-version: "10"
---

<SymbolHeader pkg="json" title="Stringify" node="JSON.stringify(obj)" import-path="github.com/sahilkhaire/gox/json" />
## Overview

Stringify marshals v to a string (JSON.stringify).

**Node.js equivalent:** `JSON.stringify(obj)`

## Signature

<div class="signature-block">

```go
func Stringify(v any) (string, error)
```

</div>

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/json/must-parse">MustParse</a><a class="related-chip" href="/packages/json/must-stringify">MustStringify</a><a class="related-chip" href="/packages/json/parse">Parse</a>
</div>

← [Back to json package overview](/packages/json/)
