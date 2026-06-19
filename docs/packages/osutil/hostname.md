---
title: "Hostname"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.hostname()"
gox-doc-version: "11"
---

<SymbolHeader pkg="osutil" title="Hostname" node="os.hostname()" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

Hostname returns the host name.

If you are coming from Node.js, the closest pattern is **`os.hostname()`**.

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

## Example

```go
import "github.com/sahilkhaire/gox/osutil"

name, err := osutil.Hostname()
```

## Tips

Import `github.com/sahilkhaire/gox/osutil` and call `Hostname` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/arch">Arch</a><a class="related-chip" href="/packages/osutil/cp-us">CPUs</a><a class="related-chip" href="/packages/osutil/homedir">Homedir</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
