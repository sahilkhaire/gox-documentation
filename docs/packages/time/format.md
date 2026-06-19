---
title: "Format"
package: "time"
import: "github.com/sahilkhaire/gox/time"
node: "moment().format()"
gox-doc-version: "10"
---

<SymbolHeader pkg="time" title="Format" node="moment().format()" import-path="github.com/sahilkhaire/gox/time" />
## Overview

**Node.js equivalent:** `moment().format()`

## Signature

<div class="signature-block">

```go
func Format(t stdtime.Time, layout string) string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
moment().format()
```

```go [Standard Go]
t := time.Now()
t.Format(time.RFC3339)
```

```go [gox]
import "github.com/sahilkhaire/gox/time"

timex.Format(t, layout)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/time/add">Add</a><a class="related-chip" href="/packages/time/diff">Diff</a><a class="related-chip" href="/packages/time/end-of-day">EndOfDay</a>
</div>

← [Back to time package overview](/packages/time/)
