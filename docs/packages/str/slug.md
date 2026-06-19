---
title: "Slug"
package: "str"
import: "github.com/sahilkhaire/gox/str"
node: "slugify(s)"
gox-doc-version: "10"
---

<SymbolHeader pkg="str" title="Slug" node="slugify(s)" import-path="github.com/sahilkhaire/gox/str" />
## Overview

Converts arbitrary text into a URL-safe slug — strips punctuation, lowercases, and replaces spaces with hyphens.

## Signature

<div class="signature-block">

```go
func Slug(s string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
slugify('Hello World');
```

```go [Standard Go]
s := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(title), " ", "-"))
```

```go [gox]
import "github.com/sahilkhaire/gox/str"

slug := str.Slug("Hello World!")
```

:::

## Tips

Use before writing user-generated titles to URL paths.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/str/camel">Camel</a><a class="related-chip" href="/packages/str/capitalize">Capitalize</a><a class="related-chip" href="/packages/str/is-blank">IsBlank</a>
</div>

← [Back to str package overview](/packages/str/)
