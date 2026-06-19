---
title: "Arch"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "process.arch"
gox-doc-version: "14"
---

<SymbolHeader pkg="osutil" title="Arch" node="process.arch" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

Arch returns the architecture (runtime.GOARCH).

If you are coming from Node.js, the closest pattern is **`process.arch`**.

## Signature

<div class="signature-block">

```go
func Arch() string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
process.arch
```

```go [Standard Go]
val, err := os.Arch()
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

osutil.Arch()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/osutil"

osutil.Arch()
```

## Tips

Import `github.com/sahilkhaire/gox/osutil` and call `Arch` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
val, err := os.Arch()
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/cp-us">CPUs</a><a class="related-chip" href="/packages/osutil/homedir">Homedir</a><a class="related-chip" href="/packages/osutil/hostname">Hostname</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
