---
title: "MustStringify"
package: "json"
import: "github.com/sahilkhaire/gox/json"
gox-doc-version: "10"
---

<SymbolHeader pkg="json" title="MustStringify" node="JSON.parse/stringify" import-path="github.com/sahilkhaire/gox/json" />
## Overview

MustStringify marshals or panics.

## Signature

<div class="signature-block">

```go
func MustStringify(v any) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/json/must-parse">MustParse</a><a class="related-chip" href="/packages/json/parse">Parse</a><a class="related-chip" href="/packages/json/parse-file">ParseFile</a>
</div>

← [Back to json package overview](/packages/json/)
