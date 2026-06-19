---
title: "EncodeQuery"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "querystring.stringify(obj)"
gox-doc-version: "11"
---

<SymbolHeader pkg="url" title="EncodeQuery" node="querystring.stringify(obj)" import-path="github.com/sahilkhaire/gox/url" />
## Overview

EncodeQuery encodes values (querystring.stringify).

If you are coming from Node.js, the closest pattern is **`querystring.stringify(obj)`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/url"

url.EncodeQuery(obj)
```

## Tips

Import `github.com/sahilkhaire/gox/url` and call `EncodeQuery` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/url/format">Format</a><a class="related-chip" href="/packages/url/parse-query">ParseQuery</a><a class="related-chip" href="/packages/url/resolve">Resolve</a>
</div>

← [Back to url package overview](/packages/url/)
