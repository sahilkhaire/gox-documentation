---
title: "MustParse"
package: "json"
import: "github.com/sahilkhaire/gox/json"
gox-doc-version: "10"
---

<SymbolHeader pkg="json" title="MustParse" node="JSON.parse/stringify" import-path="github.com/sahilkhaire/gox/json" />
## Overview

MustParse unmarshals or panics.

## Signature

<div class="signature-block">

```go
func MustParse(data []byte, v any)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/json/must-stringify">MustStringify</a><a class="related-chip" href="/packages/json/parse">Parse</a><a class="related-chip" href="/packages/json/parse-file">ParseFile</a>
</div>

← [Back to json package overview](/packages/json/)
