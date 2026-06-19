---
title: "APIKey"
package: "auth"
import: "github.com/sahilkhaire/gox/auth"
gox-doc-version: "10"
---

<SymbolHeader pkg="auth" title="APIKey" node="passport" import-path="github.com/sahilkhaire/gox/auth" />
## Overview

APIKey validates a header API key and stores it on the request context.

## Signature

<div class="signature-block">

```go
func APIKey(header string, opts APIKeyOpts) goxhttp.Middleware
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

// auth
_ = auth.APIKey(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/auth/basic">Basic</a><a class="related-chip" href="/packages/auth/bearer">Bearer</a><a class="related-chip" href="/packages/auth/get-api-key">GetAPIKey</a>
</div>

← [Back to auth package overview](/packages/auth/)
