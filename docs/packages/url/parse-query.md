---
title: "ParseQuery"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "querystring.parse(qs)"
gox-doc-version: "14"
---

<SymbolHeader pkg="url" title="ParseQuery" node="querystring.parse(qs)" import-path="github.com/sahilkhaire/gox/url" />
## Overview

ParseQuery parses a query string (querystring.parse).

If you are coming from Node.js, the closest pattern is **`querystring.parse(qs)`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/url"

vals, err := url.ParseQuery("a=1&b=2")
```

## Tips

Import `github.com/sahilkhaire/gox/url` and call `ParseQuery` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
vals, err := url.ParseQuery(q)
// or url.Values.Encode()
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/url/encode-query">EncodeQuery</a><a class="related-chip" href="/packages/url/format">Format</a><a class="related-chip" href="/packages/url/resolve">Resolve</a>
</div>

← [Back to url package overview](/packages/url/)
