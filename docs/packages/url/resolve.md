---
title: "Resolve"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "url.resolve(base, rel)"
gox-doc-version: "10"
---

<SymbolHeader pkg="url" title="Resolve" node="url.resolve(base, rel)" import-path="github.com/sahilkhaire/gox/url" />
## Overview

Resolve resolves relative against base (url.resolve).

**Node.js equivalent:** `url.resolve(base, rel)`

## Signature

<div class="signature-block">

```go
func Resolve(base, relative string) (string, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
url.resolve(base, rel)
```

```go [Standard Go]
u := &url.URL{Scheme: "https", Host: host, Path: path}
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

url.Resolve(base, rel)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/url/encode-query">EncodeQuery</a><a class="related-chip" href="/packages/url/format">Format</a><a class="related-chip" href="/packages/url/parse-query">ParseQuery</a>
</div>

← [Back to url package overview](/packages/url/)
