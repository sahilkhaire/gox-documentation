---
title: "TempDir"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.tmpdir()"
gox-doc-version: "14"
---

<SymbolHeader pkg="osutil" title="TempDir" node="os.tmpdir()" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

TempDir returns the default temp directory.

If you are coming from Node.js, the closest pattern is **`os.tmpdir()`**.

## Signature

<div class="signature-block">

```go
func TempDir() string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
os.tmpdir()
```

```go [Standard Go]
val, err := os.TempDir()
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

osutil.TempDir()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/osutil"

osutil.TempDir()
```

## Tips

Import `github.com/sahilkhaire/gox/osutil` and call `TempDir` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
val, err := os.TempDir()
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/arch">Arch</a><a class="related-chip" href="/packages/osutil/cp-us">CPUs</a><a class="related-chip" href="/packages/osutil/homedir">Homedir</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
