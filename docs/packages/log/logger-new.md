---
title: "Logger.New"
package: "log"
import: "github.com/sahilkhaire/gox/log"
node: "winston.createLogger()"
gox-doc-version: "10"
---

<SymbolHeader pkg="log" title="Logger.New" node="winston.createLogger()" import-path="github.com/sahilkhaire/gox/log" />
## Overview

New returns a Logger writing JSON to stderr at info level.

**Node.js equivalent:** `winston.createLogger()`

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

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/log/logger-default">Logger.Default</a><a class="related-chip" href="/packages/log/logger-new-with-level">Logger.NewWithLevel</a>
</div>

← [Back to log package overview](/packages/log/)
