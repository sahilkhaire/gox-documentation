---
title: "Logger.NewWithLevel"
package: "log"
import: "github.com/sahilkhaire/gox/log"
gox-doc-version: "10"
---

<SymbolHeader pkg="log" title="Logger.NewWithLevel" node="winston, pino" import-path="github.com/sahilkhaire/gox/log" />
## Overview

NewWithLevel returns a Logger at the given level.

## Signature

<div class="signature-block">

```go
func NewWithLevel(level slog.Level) *Logger
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
v.NewWithLevel(/* args */)
```

:::

## Example from tests

Extracted from the gox test suite — runnable patterns used in CI:

```go
import "github.com/sahilkhaire/gox/log"

l := NewWithLevel(slog.LevelError)
```

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/log/logger-default">Logger.Default</a><a class="related-chip" href="/packages/log/logger-new">Logger.New</a>
</div>

← [Back to log package overview](/packages/log/)
