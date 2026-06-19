---
title: "Version"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
gox-doc-version: "14"
---

<SymbolHeader pkg="semver" title="Version" node="semver" import-path="github.com/sahilkhaire/gox/semver" />
## Overview

Version is a semantic version.

Part of the **`semver`** package — Node.js analog: *semver*.

`Version` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type Version struct {
	// contains filtered or unexported fields
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical semver pattern in Node.js
```

```go [Standard Go]
// github.com/Masterminds/semver/v3
```

```go [gox]
import "github.com/sahilkhaire/gox/semver"

v, err := semver.Parse("1.2.3")
if err != nil {
    return err
}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/semver"

v, err := semver.Parse("1.2.3")
if err != nil {
    return err
}
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/semver/compare">Compare</a><a class="related-chip" href="/packages/semver/inc">Inc</a><a class="related-chip" href="/packages/semver/satisfies">Satisfies</a>
</div>

← [Back to semver package overview](/packages/semver/)
