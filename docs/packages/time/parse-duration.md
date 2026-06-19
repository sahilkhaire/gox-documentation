---
title: "ParseDuration"
package: "time"
import: "github.com/sahilkhaire/gox/time"
node: "ms('2d')"
gox-doc-version: "10"
---

<SymbolHeader pkg="time" title="ParseDuration" node="ms('2d')" import-path="github.com/sahilkhaire/gox/time" />
## Overview

**Node.js equivalent:** `ms('2d')`

## Signature

<div class="signature-block">

```go
func ParseDuration(s string) (stdtime.Duration, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
ms('2d')
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/time"

timex.ParseDuration("2d")
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/time/add">Add</a><a class="related-chip" href="/packages/time/diff">Diff</a><a class="related-chip" href="/packages/time/end-of-day">EndOfDay</a>
</div>

← [Back to time package overview](/packages/time/)
