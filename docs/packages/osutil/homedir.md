---
title: "Homedir"
package: "osutil"
import: "github.com/sahilkhaire/gox/osutil"
node: "os.homedir()"
gox-doc-version: "14"
---

<SymbolHeader pkg="osutil" title="Homedir" node="os.homedir()" import-path="github.com/sahilkhaire/gox/osutil" />
## Overview

Homedir returns the current user's home directory.

If you are coming from Node.js, the closest pattern is **`os.homedir()`**.

## Signature

<div class="signature-block">

```go
func Homedir() (string, error)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
os.homedir()
```

```go [Standard Go]
val, err := os.UserHomeDir()
```

```go [gox]
import "github.com/sahilkhaire/gox/osutil"

osutil.Homedir()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/osutil"

osutil.Homedir()
```

## Tips

Import `github.com/sahilkhaire/gox/osutil` and call `Homedir` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
val, err := os.UserHomeDir()
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/osutil/arch">Arch</a><a class="related-chip" href="/packages/osutil/cp-us">CPUs</a><a class="related-chip" href="/packages/osutil/hostname">Hostname</a>
</div>

← [Back to osutil package overview](/packages/osutil/)
