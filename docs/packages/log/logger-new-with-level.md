---
title: "Logger.NewWithLevel"
package: "log"
import: "github.com/sahilkhaire/gox/log"
gox-doc-version: "11"
---

<SymbolHeader pkg="log" title="Logger.NewWithLevel" node="winston, pino" import-path="github.com/sahilkhaire/gox/log" />
## Overview

NewWithLevel returns a Logger at the given level.

Part of the **`log`** package — Node.js analog: *winston, pino*.

Method on **`Logger`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func NewWithLevel(level slog.Level) *Logger
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

var v Logger
v.NewWithLevel(/* args */)
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/log"

l := NewWithLevel(slog.LevelError)
```

## Tips

Obtain a `Logger` value first (see constructors on the package overview), then call `NewWithLevel`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/log/logger-default">Logger.Default</a><a class="related-chip" href="/packages/log/logger-new">Logger.New</a>
</div>

← [Back to log package overview](/packages/log/)
