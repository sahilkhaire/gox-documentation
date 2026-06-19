---
title: "Load"
package: "env"
import: "github.com/sahilkhaire/gox/env"
node: "require('dotenv').config()"
gox-doc-version: "11"
---

<SymbolHeader pkg="env" title="Load" node="require('dotenv').config()" import-path="github.com/sahilkhaire/gox/env" />
## Overview

Load parses a dotenv file and sets variables in the environment and override layer.

If you are coming from Node.js, the closest pattern is **`require('dotenv').config()`**.

## Signature

<div class="signature-block">

```go
func Load(path string) error
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
require('dotenv').config();
```

```go [Standard Go]
// use os.Getenv after loading .env manually or via a config lib
```

```go [gox]
import "github.com/sahilkhaire/gox/env"

if err := env.Load(); err != nil {
    log.Fatal(err)
}
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/env"

dir := t.TempDir()
path := filepath.Join(dir, ".env")
Set("FOO", "")
Set("ONLY_TEST", "1")
n, err := GetInt("MISSING", 7)
Set("N", "42")
```

## Tips

Call `Load()` once at startup, then use typed getters instead of parsing strings manually.

## Standard library alternative

Use `os.Getenv` and `os.Setenv` from the standard library, or a config library like viper.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/env/get">Get</a><a class="related-chip" href="/packages/env/get-bool">GetBool</a><a class="related-chip" href="/packages/env/get-duration">GetDuration</a>
</div>

← [Back to env package overview](/packages/env/)
