---
title: "ParseFile"
package: "json"
import: "github.com/sahilkhaire/gox/json"
gox-doc-version: "14"
---

<SymbolHeader pkg="json" title="ParseFile" node="JSON.parse/stringify" import-path="github.com/sahilkhaire/gox/json" />
## Overview

ParseFile reads path and unmarshals into v.

Part of the **`json`** package — Node.js analog: *JSON.parse/stringify*.

## Signature

<div class="signature-block">

```go
func ParseFile(ctx context.Context, path string, v any) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
JSON.parse(fs.readFileSync('cfg.json','utf8'));
```

```go [Standard Go]
b, err := os.ReadFile(path)
err = json.Unmarshal(b, &cfg)
```

```go [gox]
import "github.com/sahilkhaire/gox/json"

dir := t.TempDir()
path := filepath.Join(dir, "data.json")
var m map[string]int
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/json"

dir := t.TempDir()
path := filepath.Join(dir, "data.json")
var m map[string]int
```

## Tips

Import `github.com/sahilkhaire/gox/json` and call `ParseFile` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
b, err := os.ReadFile(path)
err = json.Unmarshal(b, &cfg)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/json/must-parse">MustParse</a><a class="related-chip" href="/packages/json/must-stringify">MustStringify</a><a class="related-chip" href="/packages/json/parse">Parse</a>
</div>

← [Back to json package overview](/packages/json/)
