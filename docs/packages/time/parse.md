---
title: "Parse"
package: "time"
import: "github.com/sahilkhaire/gox/time"
gox-doc-version: "10"
---

<SymbolHeader pkg="time" title="Parse" node="moment, dayjs (package timex)" import-path="github.com/sahilkhaire/gox/time" />
## Overview

## Signature

<div class="signature-block">

```go
func Parse(layout, value string) (stdtime.Time, error)
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
_ = timex.Parse(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/time/add">Add</a><a class="related-chip" href="/packages/time/diff">Diff</a><a class="related-chip" href="/packages/time/end-of-day">EndOfDay</a>
</div>

← [Back to time package overview](/packages/time/)
