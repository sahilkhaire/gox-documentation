---
title: "SetDefault"
package: "log"
import: "github.com/sahilkhaire/gox/log"
gox-doc-version: "10"
---

<SymbolHeader pkg="log" title="SetDefault" node="winston, pino" import-path="github.com/sahilkhaire/gox/log" />
## Overview

SetDefault replaces the package default logger.

## Signature

<div class="signature-block">

```go
func SetDefault(l *Logger)
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

// log
_ = log.SetDefault(/* args */)
```

:::

← [Back to log package overview](/packages/log/)
