---
title: "ParseQuery"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "querystring.parse(qs)"
gox-doc-version: "10"
---

<SymbolHeader pkg="url" title="ParseQuery" node="querystring.parse(qs)" import-path="github.com/sahilkhaire/gox/url" />
## Overview

ParseQuery parses a query string (querystring.parse).

**Node.js equivalent:** `querystring.parse(qs)`

## Signature

<div class="signature-block">

```go
func ParseQuery(query string) (url.Values, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
querystring.parse('a=1&b=2');
```

```go [Standard Go]
vals, err := url.ParseQuery(q)
// or url.Values.Encode()
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

vals, err := url.ParseQuery("a=1&b=2")
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/url/encode-query">EncodeQuery</a><a class="related-chip" href="/packages/url/format">Format</a><a class="related-chip" href="/packages/url/resolve">Resolve</a>
</div>

← [Back to url package overview](/packages/url/)
