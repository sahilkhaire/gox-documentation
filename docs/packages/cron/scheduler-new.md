---
title: "Scheduler.New"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "14"
---

<SymbolHeader pkg="cron" title="Scheduler.New" node="node-cron" import-path="github.com/sahilkhaire/gox/cron" />
## Overview

New creates a scheduler and starts its internal runner.

Part of the **`cron`** package — Node.js analog: *node-cron*.

Method on **`Scheduler`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func New() *Scheduler
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const job = cron.schedule('* * * * *', fn);
```

```go [Standard Go]
c := cron.New()
c.AddFunc(spec, fn)
c.Start()
```

```go [gox]
import "github.com/sahilkhaire/gox/cron"

sched := cron.New()
sched.Schedule("* * * * *", fn)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/cron"

sched := cron.New()
sched.Schedule("* * * * *", fn)
```

## Tips

Obtain a `Scheduler` value first (see constructors on the package overview), then call `New`.

## Standard library alternative

Use the standard library directly:

```go
c := cron.New()
c.AddFunc(spec, fn)
c.Start()
```

← [Back to cron package overview](/packages/cron/)
