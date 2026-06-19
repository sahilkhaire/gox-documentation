---
title: "SetDefault"
package: "log"
import: "github.com/sahilkhaire/gox/log"
gox-doc-version: "11"
---

<SymbolHeader pkg="log" title="SetDefault" node="winston, pino" import-path="github.com/sahilkhaire/gox/log" />
## Overview

SetDefault replaces the package default logger.

Part of the **`log`** package — Node.js analog: *winston, pino*.

## Signature

<div class="signature-block">

```go
func SetDefault(l *Logger)
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical winston, pino pattern in Node.js
```

```go [Standard Go]
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("msg", "key", val)
```

```go [gox]
import "github.com/sahilkhaire/gox/log"

// log
_ = log.SetDefault(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/log"

// log
_ = log.SetDefault(/* args */)
```

## Tips

Import `github.com/sahilkhaire/gox/log` and call `SetDefault` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

← [Back to log package overview](/packages/log/)
