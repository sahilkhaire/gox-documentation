---
title: "EncodeQuery"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "querystring.stringify(obj)"
gox-doc-version: "10"
---

<SymbolHeader pkg="url" title="EncodeQuery" node="querystring.stringify(obj)" import-path="github.com/sahilkhaire/gox/url" />
## Overview

EncodeQuery encodes values (querystring.stringify).

**Node.js equivalent:** `querystring.stringify(obj)`

## Signature

<div class="signature-block">

```go
func EncodeQuery(v url.Values) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
querystring.stringify(obj)
```

```go [Standard Go]
vals, err := url.ParseQuery(q)
// or url.Values.Encode()
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

url.EncodeQuery(obj)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/url/format">Format</a><a class="related-chip" href="/packages/url/parse-query">ParseQuery</a><a class="related-chip" href="/packages/url/resolve">Resolve</a>
</div>

← [Back to url package overview](/packages/url/)
