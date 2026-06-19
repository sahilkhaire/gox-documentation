---
title: "Format"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "url.format(obj)"
gox-doc-version: "10"
---

<SymbolHeader pkg="url" title="Format" node="url.format(obj)" import-path="github.com/sahilkhaire/gox/url" />
## Overview

Format returns the serialized URL (url.format).

**Node.js equivalent:** `url.format(obj)`

## Signature

<div class="signature-block">

```go
func Format(u *URL) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
url.format(obj)
```

```go [Standard Go]
u := &url.URL{Scheme: "https", Host: host, Path: path}
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

url.Format(u)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/url/encode-query">EncodeQuery</a><a class="related-chip" href="/packages/url/parse-query">ParseQuery</a><a class="related-chip" href="/packages/url/resolve">Resolve</a>
</div>

← [Back to url package overview](/packages/url/)
