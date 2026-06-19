---
title: "SetDefault"
package: "log"
import: "github.com/sahilkhaire/gox/log"
gox-doc-version: "14"
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

log.SetDefault(log.New())
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/log"

log.SetDefault(log.New())
```

## Tips

Import `github.com/sahilkhaire/gox/log` and call `SetDefault` directly. See the comparison below for the standard library equivalent.

## Standard library alternative

Use the standard library directly:

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("msg", "key", val)
```

← [Back to log package overview](/packages/log/)
