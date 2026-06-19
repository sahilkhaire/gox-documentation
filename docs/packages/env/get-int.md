---
title: "GetInt"
package: "env"
import: "github.com/sahilkhaire/gox/env"
gox-doc-version: "10"
---

<SymbolHeader pkg="env" title="GetInt" node="dotenv / process.env" import-path="github.com/sahilkhaire/gox/env" />
## Overview

GetInt parses an int environment variable.

## Signature

<div class="signature-block">

```go
func GetInt(key string, def int) (int, error)
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
_ = env.GetInt(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/env/get">Get</a><a class="related-chip" href="/packages/env/get-bool">GetBool</a><a class="related-chip" href="/packages/env/get-duration">GetDuration</a>
</div>

← [Back to env package overview](/packages/env/)
