---
title: "Hostname"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.hostname()"
gox-doc-version: "10"
---

<SymbolHeader pkg="osutil" title="Hostname" node="os.hostname()" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

Hostname returns the host name.

**Node.js equivalent:** `os.hostname()`

## Signature

<div class="signature-block">

```go
func Hostname() (string, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
os.hostname();
```

```go [Standard Go]
val, err := os.Hostname()
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

name, err := osutil.Hostname()
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/arch">Arch</a><a class="related-chip" href="/packages/osutil/cp-us">CPUs</a><a class="related-chip" href="/packages/osutil/homedir">Homedir</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
