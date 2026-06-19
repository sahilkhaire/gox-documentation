---
title: "Scheduler.New"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "10"
---

<SymbolHeader pkg="cron" title="Scheduler.New" node="node-cron" import-path="github.com/sahilkhaire/gox/cron" />
## Overview

New creates a scheduler and starts its internal runner.

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

← [Back to cron package overview](/packages/cron/)
