---
title: "MustParse"
package: "json"
import: "github.com/sahilkhaire/gox/json"
gox-doc-version: "11"
---

<SymbolHeader pkg="json" title="MustParse" node="JSON.parse/stringify" import-path="github.com/sahilkhaire/gox/json" />
## Overview

MustParse unmarshals or panics.

Part of the **`json`** package — Node.js analog: *JSON.parse/stringify*.

## Signature

<div class="signature-block">

```go
func MustParse(data []byte, v any)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical JSON.parse/stringify pattern in Node.js
```

```go [Standard Go]
err := json.Unmarshal([]byte(raw), &v)
```

```go [gox]
import "github.com/sahilkhaire/gox/json"

// json
_ = json.MustParse(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/json"

// json
_ = json.MustParse(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/json` and call `MustParse` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/json/must-stringify">MustStringify</a><a class="related-chip" href="/packages/json/parse">Parse</a><a class="related-chip" href="/packages/json/parse-file">ParseFile</a>
</div>

← [Back to json package overview](/packages/json/)
