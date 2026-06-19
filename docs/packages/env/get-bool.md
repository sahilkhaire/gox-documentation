---
title: "GetBool"
package: "env"
import: "github.com/sahilkhaire/gox/env"
gox-doc-version: "10"
---

<SymbolHeader pkg="env" title="GetBool" node="dotenv / process.env" import-path="github.com/sahilkhaire/gox/env" />
## Overview

GetBool parses a bool environment variable.

## Signature

<div class="signature-block">

```go
func GetBool(key string, def bool) (bool, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
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
_ = env.GetBool(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/env/get">Get</a><a class="related-chip" href="/packages/env/get-duration">GetDuration</a><a class="related-chip" href="/packages/env/get-int">GetInt</a>
</div>

← [Back to env package overview](/packages/env/)
