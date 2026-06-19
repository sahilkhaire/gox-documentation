---
title: "Set"
package: "env"
import: "github.com/sahilkhaire/gox/env"
gox-doc-version: "14"
---

<SymbolHeader pkg="env" title="Set" node="dotenv / process.env" import-path="github.com/sahilkhaire/gox/env" />
## Overview

Set sets an override value (for tests).

Part of the **`env`** package — Node.js analog: *dotenv / process.env*.

## Signature

<div class="signature-block">

```go
func Set(key, value string)
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

env.Set("APP_ENV", "development")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/env"

env.Set("APP_ENV", "development")
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
<a class="related-chip" href="/packages/env/get">Get</a><a class="related-chip" href="/packages/env/get-bool">GetBool</a><a class="related-chip" href="/packages/env/get-duration">GetDuration</a>
</div>

← [Back to env package overview](/packages/env/)
