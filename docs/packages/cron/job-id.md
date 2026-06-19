---
title: "JobID"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "14"
---

<SymbolHeader pkg="cron" title="JobID" node="node-cron" import-path="github.com/sahilkhaire/gox/cron" />
## Overview

JobID identifies a scheduled job.

Part of the **`cron`** package — Node.js analog: *node-cron*.

`JobID` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type JobID int64
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

var id cron.JobID = "nightly-backup"
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/cron"

var id cron.JobID = "nightly-backup"
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

Use the standard library directly:

```go
c := cron.New()
c.AddFunc(spec, fn)
c.Start()
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/cron/parse">Parse</a>
</div>

← [Back to cron package overview](/packages/cron/)
