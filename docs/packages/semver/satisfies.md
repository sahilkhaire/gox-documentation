---
title: "Satisfies"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
node: "semver.satisfies(v, range)"
gox-doc-version: "10"
---

<SymbolHeader pkg="semver" title="Satisfies" node="semver.satisfies(v, range)" import-path="github.com/sahilkhaire/gox/semver" />
## Overview

Satisfies reports whether version matches constraint (npm range syntax).

**Node.js equivalent:** `semver.satisfies(v, range)`

## Signature

<div class="signature-block">

```go
func Satisfies(version, constraint string) bool
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
semver.satisfies(v, range)
```

```go [Standard Go]
// github.com/Masterminds/semver/v3
```

```go [gox]
import "github.com/sahilkhaire/gox/semver"

semver.Satisfies(v, constraint)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/semver/compare">Compare</a><a class="related-chip" href="/packages/semver/inc">Inc</a>
</div>

← [Back to semver package overview](/packages/semver/)
