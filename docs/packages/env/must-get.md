---
title: "MustGet"
package: "env"
import: "github.com/sahilkhaire/gox/env"
gox-doc-version: "10"
---

<SymbolHeader pkg="env" title="MustGet" node="dotenv / process.env" import-path="github.com/sahilkhaire/gox/env" />
## Overview

MustGet returns the value or panics if missing.

## Signature

<div class="signature-block">

```go
func MustGet(key string) string
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
_ = env.MustGet(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/env/get">Get</a><a class="related-chip" href="/packages/env/get-bool">GetBool</a><a class="related-chip" href="/packages/env/get-duration">GetDuration</a>
</div>

← [Back to env package overview](/packages/env/)
