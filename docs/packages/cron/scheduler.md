---
title: "Scheduler"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "10"
---

<SymbolHeader pkg="cron" title="Scheduler" node="node-cron" import-path="github.com/sahilkhaire/gox/cron" />
## Overview

Scheduler runs cron jobs in the background.

## Signature

<div class="signature-block">

```go
type Scheduler struct {
	// contains filtered or unexported fields
}
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

_ = cron.Scheduler
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/cron/parse">Parse</a>
</div>

← [Back to cron package overview](/packages/cron/)
