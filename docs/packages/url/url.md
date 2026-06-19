---
title: "URL"
package: "url"
import: "github.com/sahilkhaire/gox/url"
gox-doc-version: "14"
---

<SymbolHeader pkg="url" title="URL" node="url, querystring" import-path="github.com/sahilkhaire/gox/url" />
## Overview

URL wraps net/url.URL with Node-friendly helpers.

Part of the **`url`** package — Node.js analog: *url, querystring*.

`URL` is a type exported by gox. Methods on this type are documented separately.

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
// Typical url, querystring pattern in Node.js
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

u, err := Parse("https://example.com/a?x=1")
res, err := Resolve("https://example.com/foo/", "bar")
vals, err := ParseQuery("a=1&b=2")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/url"

u, err := Parse("https://example.com/a?x=1")
res, err := Resolve("https://example.com/foo/", "bar")
vals, err := ParseQuery("a=1&b=2")
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/url/encode-query">EncodeQuery</a><a class="related-chip" href="/packages/url/format">Format</a><a class="related-chip" href="/packages/url/parse-query">ParseQuery</a>
</div>

← [Back to url package overview](/packages/url/)
