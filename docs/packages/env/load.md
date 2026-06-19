---
title: "Load"
package: "env"
import: "github.com/sahilkhaire/gox/env"
node: "require('dotenv').config()"
gox-doc-version: "7"
---

<div class="api-meta"><span class="api-badge node">Node: require('dotenv').config()</span><span class="api-badge import">github.com/sahilkhaire/gox/env</span></div>
# Load

## Overview

Load parses a dotenv file and sets variables in the environment and override layer.

**Node.js equivalent:** `require('dotenv').config()`

## Signature

```go
func Load(path string) error
```

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

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

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

::: tip When to use gox
- Familiar API if you are migrating from Node.js
- Typed generics and explicit error handling (idiomatic Go underneath)
- Consistent naming across the gox toolkit
:::

## Related APIs

- [Get](/packages/env/get)
- [GetBool](/packages/env/get-bool)
- [GetDuration](/packages/env/get-duration)

← [Back to env package overview](/packages/env/)
