---
title: "StartOfDay"
package: "time"
import: "github.com/sahilkhaire/gox/time"
gox-doc-version: "14"
---

<SymbolHeader pkg="time" title="StartOfDay" node="moment, dayjs (package timex)" import-path="github.com/sahilkhaire/gox/time" />
## Overview

Part of the **`time`** package — Node.js analog: *moment, dayjs (package timex)*.

## Signature

<div class="signature-block">

```go
func StartOfDay(t stdtime.Time) stdtime.Time
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical moment, dayjs (package timex) pattern in Node.js
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/time"

start := timex.StartOfDay(time.Now())
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/time"

start := timex.StartOfDay(time.Now())
```

## Tips

Import `github.com/sahilkhaire/gox/time` and call `StartOfDay` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/time/add">Add</a><a class="related-chip" href="/packages/time/diff">Diff</a><a class="related-chip" href="/packages/time/end-of-day">EndOfDay</a>
</div>

← [Back to time package overview](/packages/time/)
