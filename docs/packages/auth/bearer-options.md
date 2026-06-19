---
title: "BearerOptions"
package: "auth"
import: "github.com/sahilkhaire/gox/auth"
gox-doc-version: "11"
---

<SymbolHeader pkg="auth" title="BearerOptions" node="passport" import-path="github.com/sahilkhaire/gox/auth" />
## Overview

BearerOptions configures JWT bearer middleware.

Part of the **`auth`** package — Node.js analog: *passport*.

`BearerOptions` is a type exported by gox. Methods on this type are documented separately.

## Signature

<div class="signature-block">

```go
type BearerOptions struct{}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// Typical passport pattern in Node.js
```

```go [Standard Go]
// net/http middleware checking Authorization header
```

```go [gox]
import "github.com/sahilkhaire/gox/auth"

_ = auth.BearerOptions
```

:::

## Example

```go
import "github.com/sahilkhaire/gox/auth"

_ = auth.BearerOptions
```

## Tips

Browse methods on this type in the sidebar for handler-style APIs and options structs.

## Standard library alternative

gox wraps the Go standard library or a trusted dependency with Node-familiar naming. You can use the underlying library directly — see the package overview for escape hatches.

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/auth/api-key">APIKey</a><a class="related-chip" href="/packages/auth/basic">Basic</a><a class="related-chip" href="/packages/auth/bearer">Bearer</a>
</div>

← [Back to auth package overview](/packages/auth/)
