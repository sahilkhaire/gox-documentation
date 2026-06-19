---
title: "Parse"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "10"
---

<SymbolHeader pkg="cron" title="Parse" node="node-cron" import-path="github.com/sahilkhaire/gox/cron" />
## Overview

Parse validates a cron expression (with optional seconds field).

## Signature

<div class="signature-block">

```go
func Parse(expr string) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
c := cron.New()
c.AddFunc(spec, fn)
c.Start()
```

```go [gox]
import "github.com/sahilkhaire/gox/cron"

// cron
_ = cron.Parse(/* args */)
```

:::

← [Back to cron package overview](/packages/cron/)
