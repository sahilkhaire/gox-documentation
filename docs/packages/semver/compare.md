---
title: "Compare"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
node: "semver.compare(a, b)"
gox-doc-version: "10"
---

<SymbolHeader pkg="semver" title="Compare" node="semver.compare(a, b)" import-path="github.com/sahilkhaire/gox/semver" />
## Overview

Compare returns -1 if a &lt; b, 0 if equal, 1 if a &gt; b.

**Node.js equivalent:** `semver.compare(a, b)`

## Signature

<div class="signature-block">

```go
func Compare(a, b string) int
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
semver.compare(a, b)
```

```go [Standard Go]
// github.com/Masterminds/semver/v3
```

```go [gox]
import "github.com/sahilkhaire/gox/semver"

semver.Compare(a, b)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/semver/inc">Inc</a><a class="related-chip" href="/packages/semver/satisfies">Satisfies</a>
</div>

← [Back to semver package overview](/packages/semver/)
