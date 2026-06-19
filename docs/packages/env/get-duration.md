---
title: "GetDuration"
package: "env"
import: "github.com/sahilkhaire/gox/env"
gox-doc-version: "11"
---

<SymbolHeader pkg="env" title="GetDuration" node="dotenv / process.env" import-path="github.com/sahilkhaire/gox/env" />
## Overview

GetDuration parses a duration environment variable.

Part of the **`env`** package — Node.js analog: *dotenv / process.env*.

## Signature

<div class="signature-block">

```go
func GetDuration(key string, def time.Duration) (time.Duration, error)
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

// env
_ = env.GetDuration(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/env"

// env
_ = env.GetDuration(/* args */)
```

## Tips

Call `Load()` once at startup, then use typed getters instead of parsing strings manually.

## Standard library alternative

Use `os.Getenv` and `os.Setenv` from the standard library, or a config library like viper.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/env/get">Get</a><a class="related-chip" href="/packages/env/get-bool">GetBool</a><a class="related-chip" href="/packages/env/get-int">GetInt</a>
</div>

← [Back to env package overview](/packages/env/)
