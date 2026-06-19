---
title: "Version.Parse"
package: "semver"
import: "github.com/sahilkhaire/gox/semver"
node: "semver.parse(v)"
gox-doc-version: "11"
---

<SymbolHeader pkg="semver" title="Version.Parse" node="semver.parse(v)" import-path="github.com/sahilkhaire/gox/semver" />
## Overview

Parse parses a semver string.

If you are coming from Node.js, the closest pattern is **`semver.parse(v)`**.

Method on **`Version`** — call it on a value of that type after constructing or receiving one from a constructor.

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

## Example

```go
import "github.com/sahilkhaire/gox/semver"

v, err := semver.Parse("1.2.3")
```

## Tips

Obtain a `Version` value first (see constructors on the package overview), then call `Parse`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to semver package overview](/packages/semver/)
