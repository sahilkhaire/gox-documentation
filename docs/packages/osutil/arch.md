---
title: "Arch"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "process.arch"
gox-doc-version: "10"
---

<SymbolHeader pkg="osutil" title="Arch" node="process.arch" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

Arch returns the architecture (runtime.GOARCH).

**Node.js equivalent:** `process.arch`

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/cp-us">CPUs</a><a class="related-chip" href="/packages/osutil/homedir">Homedir</a><a class="related-chip" href="/packages/osutil/hostname">Hostname</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
