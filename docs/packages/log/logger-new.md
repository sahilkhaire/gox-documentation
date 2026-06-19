---
title: "Logger.New"
package: "log"
import: "github.com/sahilkhaire/gox/log"
node: "winston.createLogger()"
gox-doc-version: "11"
---

<SymbolHeader pkg="log" title="Logger.New" node="winston.createLogger()" import-path="github.com/sahilkhaire/gox/log" />
## Overview

New returns a Logger writing JSON to stderr at info level.

If you are coming from Node.js, the closest pattern is **`winston.createLogger()`**.

Method on **`Logger`** — call it on a value of that type after constructing or receiving one from a constructor.

## Signature

<div class="signature-block">

```go
func New() *Logger
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
winston.createLogger();
```

```go [Standard Go]
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("msg", "key", val)
```

```go [gox]
import "github.com/sahilkhaire/gox/log"

logger := log.New()
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/log"

logger := log.New()
```

## Tips

Obtain a `Logger` value first (see constructors on the package overview), then call `New`.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/log/logger-default">Logger.Default</a><a class="related-chip" href="/packages/log/logger-new-with-level">Logger.NewWithLevel</a>
</div>

← [Back to log package overview](/packages/log/)
