---
title: "CPUs"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.cpus().length"
gox-doc-version: "10"
---

<SymbolHeader pkg="osutil" title="CPUs" node="os.cpus().length" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

CPUs returns the number of logical CPUs.

**Node.js equivalent:** `os.cpus().length`

## Signature

<div class="signature-block">

```go
func CPUs() int
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
os.cpus().length
```

```go [Standard Go]
// Use the underlying stdlib or driver directly.
// See package overview for escape hatches.
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

osutil.CPUs()
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/arch">Arch</a><a class="related-chip" href="/packages/osutil/homedir">Homedir</a><a class="related-chip" href="/packages/osutil/hostname">Hostname</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
