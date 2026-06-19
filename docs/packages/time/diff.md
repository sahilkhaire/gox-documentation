---
title: "Diff"
package: "time"
import: "github.com/sahilkhaire/gox/time"
gox-doc-version: "11"
---

<SymbolHeader pkg="time" title="Diff" node="moment, dayjs (package timex)" import-path="github.com/sahilkhaire/gox/time" />
## Overview

Part of the **`time`** package — Node.js analog: *moment, dayjs (package timex)*.

## Signature

<div class="signature-block">

```go
func Diff(a, b stdtime.Time) stdtime.Duration
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical moment, dayjs (package timex) pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/time"

// timex
_ = timex.Diff(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/time` and call `Diff` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/time/add">Add</a><a class="related-chip" href="/packages/time/end-of-day">EndOfDay</a><a class="related-chip" href="/packages/time/end-of-month">EndOfMonth</a>
</div>

← [Back to time package overview](/packages/time/)
