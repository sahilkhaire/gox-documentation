---
title: "WriteFile"
package: "json"
import: "github.com/sahilkhaire/gox/json"
gox-doc-version: "10"
---

<SymbolHeader pkg="json" title="WriteFile" node="JSON.parse/stringify" import-path="github.com/sahilkhaire/gox/json" />
## Overview

WriteFile marshals v and writes to path (helper for examples).

## Signature

<div class="signature-block">

```go
func WriteFile(ctx context.Context, path string, v any, perm os.FileMode) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
fs.writeFileSync('out.json', JSON.stringify(obj));
```

```go [Standard Go]
b, err := json.Marshal(obj)
err = os.WriteFile(path, b, 0644)
```

```go [gox]
import "github.com/sahilkhaire/gox/json"

err := json.WriteFile(ctx, "out.json", obj)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/json/must-parse">MustParse</a><a class="related-chip" href="/packages/json/must-stringify">MustStringify</a><a class="related-chip" href="/packages/json/parse">Parse</a>
</div>

← [Back to json package overview](/packages/json/)
