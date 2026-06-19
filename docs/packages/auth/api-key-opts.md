---
title: "APIKeyOpts"
package: "auth"
import: "github.com/sahilkhaire/gox/auth"
gox-doc-version: "10"
---

<SymbolHeader pkg="auth" title="APIKeyOpts" node="passport" import-path="github.com/sahilkhaire/gox/auth" />
## Overview

APIKeyOpts configures API key middleware.

## Signature

<div class="signature-block">

```go
type APIKeyOpts struct {
	Keys     map[string]bool
	Validate func(key string) bool
}
```

</div>

## Compare: Node.js · Standard Go · gox

::: code-group

```js [Node.js]
// See package overview
```

```go [Standard Go]
// net/http middleware checking Authorization header
```

```go [gox]
import "github.com/sahilkhaire/gox/auth"

_ = auth.APIKeyOpts
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/auth/api-key">APIKey</a><a class="related-chip" href="/packages/auth/basic">Basic</a><a class="related-chip" href="/packages/auth/bearer">Bearer</a>
</div>

← [Back to auth package overview](/packages/auth/)
