---
title: "Inc"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
node: "semver.inc(v, 'minor')"
gox-doc-version: "14"
---

<SymbolHeader pkg="semver" title="Inc" node="semver.inc(v, 'minor')" import-path="github.com/sahilkhaire/gox/semver" />
## Overview

Inc increments part (major, minor, patch) and returns the new version string.

If you are coming from Node.js, the closest pattern is **`semver.inc(v, 'minor')`**.

## Signature

<div class="signature-block">

```go
func Inc(version, part string) (string, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
semver.inc(v, 'minor')
```

```go [Standard Go]
// github.com/Masterminds/semver/v3
```

```go [gox]
import "github.com/sahilkhaire/gox/semver"

semver.Inc(v, part)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/semver"

semver.Inc(v, part)
```

## Tips

Import `github.com/sahilkhaire/gox/semver` and call `Inc` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/semver/compare">Compare</a><a class="related-chip" href="/packages/semver/satisfies">Satisfies</a>
</div>

← [Back to semver package overview](/packages/semver/)
