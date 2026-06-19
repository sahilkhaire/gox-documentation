---
title: "cond Cookbook"
package: "cond"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: ternary ? :, nullish ??</span><span class="api-badge import">github.com/sahilkhaire/gox/cond</span></div>
# cond Cookbook

Ternary and nullish-coalescing patterns from JavaScript, expressed in typed Go.

## Pass / fail label

```go
import "github.com/sahilkhaire/gox/cond"

label := cond.If(score >= 60, "pass", "fail")
```

## Lazy branches (avoid expensive work)

```go
// Only the chosen branch runs — unlike cond.If, both sides are NOT evaluated upfront.
v := cond.IfLazy(useCache, loadFromCache, loadFromDB)
```

## Nullish coalescing chain

```go
name := cond.Coalesce(maybeName, fallbackName, "guest")
```

## Safe pointer fallback

```go
display := cond.CoalescePtr(user.Nickname, user.Name, "Anonymous")
```

← [Back to cond overview](/packages/cond/)
