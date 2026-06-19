---
title: "URL"
package: "url"
import: "github.com/sahilkhaire/gox/url"
gox-doc-version: "10"
---

<SymbolHeader pkg="url" title="URL" node="url, querystring" import-path="github.com/sahilkhaire/gox/url" />
## Overview

URL wraps net/url.URL with Node-friendly helpers.

## Signature

<div class="signature-block">

```go
type URL struct {
	*url.URL
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

_ = url.URL
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/url"

u, err := Parse("https://example.com/a?x=1")
res, err := Resolve("https://example.com/foo/", "bar")
vals, err := ParseQuery("a=1&b=2")
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/url/encode-query">EncodeQuery</a><a class="related-chip" href="/packages/url/format">Format</a><a class="related-chip" href="/packages/url/parse-query">ParseQuery</a>
</div>

← [Back to url package overview](/packages/url/)
