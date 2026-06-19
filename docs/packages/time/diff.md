---
title: "Diff"
package: "time"
import: "github.com/sahilkhaire/gox/time"
gox-doc-version: "10"
---

<SymbolHeader pkg="time" title="Diff" node="moment, dayjs (package timex)" import-path="github.com/sahilkhaire/gox/time" />
## Overview

## Signature

<div class="signature-block">

```go
func Diff(a, b stdtime.Time) stdtime.Duration
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
t := time.Now()
t.Format(time.RFC3339)
```

```go [gox]
import "github.com/sahilkhaire/gox/time"

// timex
_ = timex.Diff(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/time/add">Add</a><a class="related-chip" href="/packages/time/end-of-day">EndOfDay</a><a class="related-chip" href="/packages/time/end-of-month">EndOfMonth</a>
</div>

← [Back to time package overview](/packages/time/)
