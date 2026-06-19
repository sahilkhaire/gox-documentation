---
title: "Get"
package: "env"
import: "github.com/sahilkhaire/gox/env"
node: "process.env.KEY"
gox-doc-version: "14"
---

<SymbolHeader pkg="env" title="Get" node="process.env.KEY" import-path="github.com/sahilkhaire/gox/env" />
## Overview

Get returns the value for key (override, then os.Getenv).

If you are coming from Node.js, the closest pattern is **`process.env.KEY`**.

## Signature

<div class="signature-block">

```go
func Get(key string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const port = process.env.PORT || '8080';
```

```go [Standard Go]
port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
```

```go [gox]
import "github.com/sahilkhaire/gox/env"

port := env.Get("PORT")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/env"

port := env.Get("PORT")
```

## Tips

Call `Load()` once at startup, then use typed getters instead of parsing strings manually.

## Standard library alternative

Use the standard library directly:

```go
port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/env/get-bool">GetBool</a><a class="related-chip" href="/packages/env/get-duration">GetDuration</a><a class="related-chip" href="/packages/env/get-int">GetInt</a>
</div>

← [Back to env package overview](/packages/env/)
