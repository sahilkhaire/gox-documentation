---
title: "URL.Parse"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "new URL(str)"
gox-doc-version: "10"
---

<SymbolHeader pkg="url" title="URL.Parse" node="new URL(str)" import-path="github.com/sahilkhaire/gox/url" />
## Overview

Parse parses a URL string (url.parse).

**Node.js equivalent:** `new URL(str)`

## Signature

<div class="signature-block">

```go
func Parse(raw string) (*URL, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
new URL('https://example.com/path');
```

```go [Standard Go]
u, err := url.Parse(rawURL)
```

```go [gox]
import "github.com/sahilkhaire/gox/url"

u, err := url.Parse("https://example.com/path")
```

:::

← [Back to url package overview](/packages/url/)
