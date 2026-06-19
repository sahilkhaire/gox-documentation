---
title: "JobID"
package: "cron"
import: "github.com/sahilkhaire/gox/cron"
gox-doc-version: "10"
---

<SymbolHeader pkg="cron" title="JobID" node="node-cron" import-path="github.com/sahilkhaire/gox/cron" />
## Overview

JobID identifies a scheduled job.

## Signature

<div class="signature-block">

```go
type JobID int64
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

_ = cron.JobID
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/cron/parse">Parse</a>
</div>

← [Back to cron package overview](/packages/cron/)
