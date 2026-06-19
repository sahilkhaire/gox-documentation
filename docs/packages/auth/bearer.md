---
title: "Bearer"
package: "auth"
import: "github.com/sahilkhaire/gox/auth"
gox-doc-version: "10"
---

<SymbolHeader pkg="auth" title="Bearer" node="passport" import-path="github.com/sahilkhaire/gox/auth" />
## Overview

Bearer validates Authorization: Bearer &lt;jwt&gt; and stores claims on the request context.

## Signature

<div class="signature-block">

```go
func Bearer(secret []byte, opts *BearerOptions) goxhttp.Middleware
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
_ = auth.Bearer(/* args */)
```

:::

## Related APIs

<div class="related-chips">
<a class="related-chip" href="/packages/auth/api-key">APIKey</a><a class="related-chip" href="/packages/auth/basic">Basic</a><a class="related-chip" href="/packages/auth/get-api-key">GetAPIKey</a>
</div>

← [Back to auth package overview](/packages/auth/)
