---
title: "Satisfies"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
node: "semver.satisfies(v, range)"
gox-doc-version: "14"
---

<SymbolHeader pkg="semver" title="Satisfies" node="semver.satisfies(v, range)" import-path="github.com/sahilkhaire/gox/semver" />
## Overview

Satisfies reports whether version matches constraint (npm range syntax).

If you are coming from Node.js, the closest pattern is **`semver.satisfies(v, range)`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/semver"

semver.Satisfies(v, constraint)
```

## Tips

Import `github.com/sahilkhaire/gox/semver` and call `Satisfies` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/semver/compare">Compare</a><a class="related-chip" href="/packages/semver/inc">Inc</a>
</div>

← [Back to semver package overview](/packages/semver/)
