---
title: "Parse"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "14"
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

spec, err := cron.Parse("0 * * * *")
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/cron"

spec, err := cron.Parse("0 * * * *")
```

## Tips

Import `github.com/sahilkhaire/gox/cron` and call `Parse` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
c := cron.New()
c.AddFunc(spec, fn)
c.Start()
```

← [Back to cron package overview](/packages/cron/)
