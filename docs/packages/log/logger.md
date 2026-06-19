---
title: "Logger"
package: "log"
import: "github.com/sahilkhaire/gox/log"
gox-doc-version: "11"
---

<SymbolHeader pkg="log" title="Logger" node="winston, pino" import-path="github.com/sahilkhaire/gox/log" />
## Overview

Logger wraps slog.Logger with familiar leveled helpers.

Part of the **`log`** package — Node.js analog: *winston, pino*.

`Logger` is a type exported by gox. Methods on this type are documented separately.

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
// Typical winston, pino pattern in Node.js
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

## Example

```go
import "github.com/sahilkhaire/gox/log"

_ = log.Logger
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/log/set-default">SetDefault</a>
</div>

← [Back to log package overview](/packages/log/)
