---
title: "MustStringify"
package: "json"
import: "github.com/sahilkhaire/gox/json"
gox-doc-version: "11"
---

<SymbolHeader pkg="json" title="MustStringify" node="JSON.parse/stringify" import-path="github.com/sahilkhaire/gox/json" />
## Overview

MustStringify marshals or panics.

Part of the **`json`** package — Node.js analog: *JSON.parse/stringify*.

## Signature

<div class="signature-block">

```go
func MustStringify(v any) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical JSON.parse/stringify pattern in Node.js
```

```go [Standard Go]
b, err := json.Marshal(v)
```

```go [gox]
import "github.com/sahilkhaire/gox/json"

// json
_ = json.MustStringify(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/json"

// json
_ = json.MustStringify(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/json` and call `MustStringify` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/json/must-parse">MustParse</a><a class="related-chip" href="/packages/json/parse">Parse</a><a class="related-chip" href="/packages/json/parse-file">ParseFile</a>
</div>

← [Back to json package overview](/packages/json/)
