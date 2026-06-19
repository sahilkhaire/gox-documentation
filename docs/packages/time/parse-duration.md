---
title: "ParseDuration"
package: "time"
import: "github.com/sahilkhaire/gox/time"
node: "ms('2d')"
gox-doc-version: "11"
---

<SymbolHeader pkg="time" title="ParseDuration" node="ms('2d')" import-path="github.com/sahilkhaire/gox/time" />
## Overview

If you are coming from Node.js, the closest pattern is **`ms('2d')`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/time"

timex.ParseDuration("2d")
```

## Tips

Import `github.com/sahilkhaire/gox/time` and call `ParseDuration` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/time/add">Add</a><a class="related-chip" href="/packages/time/diff">Diff</a><a class="related-chip" href="/packages/time/end-of-day">EndOfDay</a>
</div>

← [Back to time package overview](/packages/time/)
