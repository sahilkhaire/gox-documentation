---
title: "URL.Parse"
package: "url"
import: "github.com/sahilkhaire/gox/url"
node: "new URL(str)"
gox-doc-version: "11"
---

<SymbolHeader pkg="url" title="URL.Parse" node="new URL(str)" import-path="github.com/sahilkhaire/gox/url" />
## Overview

Parse parses a URL string (url.parse).

If you are coming from Node.js, the closest pattern is **`new URL(str)`**.

Method on **`URL`** — call it on a value of that type after constructing or receiving one from a constructor.

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

## Example

```go
import "github.com/sahilkhaire/gox/url"

u, err := url.Parse("https://example.com/path")
```

## Tips

Obtain a `URL` value first (see constructors on the package overview), then call `Parse`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to url package overview](/packages/url/)
