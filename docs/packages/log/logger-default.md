---
title: "Logger.Default"
package: "log"
import: "github.com/sahilkhaire/gox/log"
gox-doc-version: "10"
---

<SymbolHeader pkg="log" title="Logger.Default" node="winston, pino" import-path="github.com/sahilkhaire/gox/log" />
## Overview

Default returns the package-level default logger.

## Signature

<div class="signature-block">

```go
func Default() *Logger
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("msg", "key", val)
```

```go [gox]
import "github.com/sahilkhaire/gox/log"

var v Logger
v.Default(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/log/logger-new">Logger.New</a><a class="related-chip" href="/packages/log/logger-new-with-level">Logger.NewWithLevel</a>
</div>

← [Back to log package overview](/packages/log/)
