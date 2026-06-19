---
title: "Version.Parse"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
node: "semver.parse(v)"
gox-doc-version: "10"
---

<SymbolHeader pkg="semver" title="Version.Parse" node="semver.parse(v)" import-path="github.com/sahilkhaire/gox/semver" />
## Overview

Parse parses a semver string.

**Node.js equivalent:** `semver.parse(v)`

## Signature

<div class="signature-block">

```go
func Parse(v string) (*Version, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
semver.parse('1.2.3');
```

```go [Standard Go]
// github.com/Masterminds/semver/v3
```

```go [gox]
import "github.com/sahilkhaire/gox/semver"

v, err := semver.Parse("1.2.3")
```

:::

← [Back to semver package overview](/packages/semver/)
