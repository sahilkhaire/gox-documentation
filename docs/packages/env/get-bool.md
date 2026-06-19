---
title: "GetBool"
package: "env"
import: "github.com/sahilkhaire/gox/env"
gox-doc-version: "14"
---

<SymbolHeader pkg="env" title="GetBool" node="dotenv / process.env" import-path="github.com/sahilkhaire/gox/env" />
## Overview

GetBool parses a bool environment variable.

Part of the **`env`** package — Node.js analog: *dotenv / process.env*.

## Signature

<div class="signature-block">

```go
func GetBool(key string, def bool) (bool, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical dotenv / process.env pattern in Node.js
```

```go [Standard Go]
v := os.Getenv("KEY")
if v == "" {
    v = "default"
}
```

```go [gox]
import "github.com/sahilkhaire/gox/env"

debug, err := env.GetBool("DEBUG", false)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/env"

debug, err := env.GetBool("DEBUG", false)
```

## Tips

Call `Load()` once at startup, then use typed getters instead of parsing strings manually.

## Standard library alternative

Use the standard library directly:

```go
v := os.Getenv("KEY")
if v == "" {
    v = "default"
}
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/env/get">Get</a><a class="related-chip" href="/packages/env/get-duration">GetDuration</a><a class="related-chip" href="/packages/env/get-int">GetInt</a>
</div>

← [Back to env package overview](/packages/env/)
