---
title: "Logger"
package: "log"
import: "github.com/sahilkhaire/gox/log"
gox-doc-version: "10"
---

<SymbolHeader pkg="log" title="Logger" node="winston, pino" import-path="github.com/sahilkhaire/gox/log" />
## Overview

Logger wraps slog.Logger with familiar leveled helpers.

## Signature

<div class="signature-block">

```go
type Logger struct {
	*slog.Logger
}
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

_ = log.Logger
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/log/set-default">SetDefault</a>
</div>

← [Back to log package overview](/packages/log/)
