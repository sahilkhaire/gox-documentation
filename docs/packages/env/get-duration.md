---
title: "GetDuration"
package: "env"
import: "github.com/sahilkhaire/gox/env"
gox-doc-version: "10"
---

<SymbolHeader pkg="env" title="GetDuration" node="dotenv / process.env" import-path="github.com/sahilkhaire/gox/env" />
## Overview

GetDuration parses a duration environment variable.

## Signature

<div class="signature-block">

```go
func GetDuration(key string, def time.Duration) (time.Duration, error)
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
_ = env.GetDuration(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/env/get">Get</a><a class="related-chip" href="/packages/env/get-bool">GetBool</a><a class="related-chip" href="/packages/env/get-int">GetInt</a>
</div>

← [Back to env package overview](/packages/env/)
