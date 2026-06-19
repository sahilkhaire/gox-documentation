---
title: "Parse"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "11"
---

<SymbolHeader pkg="cron" title="Parse" node="node-cron" import-path="github.com/sahilkhaire/gox/cron" />
## Overview

Parse validates a cron expression (with optional seconds field).

Part of the **`cron`** package — Node.js analog: *node-cron*.

## Signature

<div class="signature-block">

```go
func Parse(expr string) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical node-cron pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/cron"

// cron
_ = cron.Parse(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/cron` and call `Parse` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to cron package overview](/packages/cron/)
