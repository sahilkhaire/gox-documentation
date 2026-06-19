---
title: "Pretty"
package: "json"
import: "github.com/sahilkhaire/gox/json"
gox-doc-version: "14"
---

<SymbolHeader pkg="json" title="Pretty" node="JSON.parse/stringify" import-path="github.com/sahilkhaire/gox/json" />
## Overview

Pretty returns indented JSON.

Part of the **`json`** package — Node.js analog: *JSON.parse/stringify*.

## Signature

<div class="signature-block">

```go
func Pretty(v any) (string, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
JSON.stringify(obj, null, 2);
```

```go [Standard Go]
b, err := json.MarshalIndent(obj, "", "  ")
```

```go [gox]
import "github.com/sahilkhaire/gox/json"

s, err := goxjson.Pretty(map[string]int{"a": 1})
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/json"

s, err := goxjson.Pretty(map[string]int{"a": 1})
```

## Tips

Import `github.com/sahilkhaire/gox/json` and call `Pretty` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
b, err := json.MarshalIndent(obj, "", "  ")
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/json/must-parse">MustParse</a><a class="related-chip" href="/packages/json/must-stringify">MustStringify</a><a class="related-chip" href="/packages/json/parse">Parse</a>
</div>

← [Back to json package overview](/packages/json/)
