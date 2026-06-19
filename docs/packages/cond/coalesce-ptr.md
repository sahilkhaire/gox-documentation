---
title: "CoalescePtr"
package: "cond"
import: "github.com/sahilkhaire/gox/cond"
node: "ptr ?? fallback"
gox-doc-version: "10"
---

<SymbolHeader pkg="cond" title="CoalescePtr" node="ptr ?? fallback" import-path="github.com/sahilkhaire/gox/cond" />
## Overview

CoalescePtr returns the first non-nil pointer's value, or zero if all nil.

**Node.js equivalent:** `ptr ?? fallback`

## Signature

<div class="signature-block">

```go
func CoalescePtr[T any](vals ...*T) T
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
const name = user?.name ?? 'guest';
```

```go [Standard Go]
name := "guest"
if user != nil {
    name = user.Name
}
```

```go [gox]
import "github.com/sahilkhaire/gox/cond"

name := cond.CoalescePtr(&user.Name, "guest")
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/cond"

a, b := 1, 2
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/cond/coalesce">Coalesce</a><a class="related-chip" href="/packages/cond/coalesce-fn">CoalesceFn</a><a class="related-chip" href="/packages/cond/if">If</a>
</div>

← [Back to cond package overview](/packages/cond/)
