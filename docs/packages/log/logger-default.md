---
title: "Logger.Default"
package: "log"
import: "github.com/sahilkhaire/gox/log"
gox-doc-version: "14"
---

<SymbolHeader pkg="log" title="Logger.Default" node="winston, pino" import-path="github.com/sahilkhaire/gox/log" />
## Overview

Default returns the package-level default logger.

Part of the **`log`** package — Node.js analog: *winston, pino*.

Method on **`Logger`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func Default() *Logger
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

logger := log.Default()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/log"

logger := log.Default()
```

## Tips

Obtain a `Logger` value first (see constructors on the package overview), then call `Default`.

## Standard library alternative

Use the standard library directly:

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("msg", "key", val)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/log/logger-new">Logger.New</a><a class="related-chip" href="/packages/log/logger-new-with-level">Logger.NewWithLevel</a>
</div>

← [Back to log package overview](/packages/log/)
