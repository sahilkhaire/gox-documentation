---
title: "Platform"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "process.platform"
gox-doc-version: "11"
---

<SymbolHeader pkg="osutil" title="Platform" node="process.platform" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

Platform returns the OS (runtime.GOOS).

If you are coming from Node.js, the closest pattern is **`process.platform`**.

## Signature

<div class="signature-block">

```go
func Platform() string
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
process.platform
```

```go [Standard Go]
val, err := os.Platform()
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

osutil.Platform()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/osutil"

osutil.Platform()
```

## Tips

Import `github.com/sahilkhaire/gox/osutil` and call `Platform` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/arch">Arch</a><a class="related-chip" href="/packages/osutil/cp-us">CPUs</a><a class="related-chip" href="/packages/osutil/homedir">Homedir</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
