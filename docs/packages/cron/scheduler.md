---
title: "Scheduler"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "11"
---

<SymbolHeader pkg="cron" title="Scheduler" node="node-cron" import-path="github.com/sahilkhaire/gox/cron" />
## Overview

Scheduler runs cron jobs in the background.

Part of the **`cron`** package — Node.js analog: *node-cron*.

`Scheduler` is a type exported by gox. Methods on this type are documented separately.

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
// Typical node-cron pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/cron"

_ = cron.Scheduler
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/cron/parse">Parse</a>
</div>

← [Back to cron package overview](/packages/cron/)
