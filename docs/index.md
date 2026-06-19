---
layout: home

hero:
  name: gox
  text: Write Go like you know Node
  tagline: A premium toolkit for JavaScript developers — Express-style HTTP, lodash collections, Zod validation, Knex queries, and 38 independent packages. Idiomatic Go underneath.
  image:
    src: /hero.svg
    alt: gox — Node-friendly Go toolkit
  actions:
    - theme: brand
      text: Get Started →
      link: /guide/getting-started
    - theme: alt
      text: Browse packages
      link: /packages/cond/
    - theme: alt
      text: npm → gox cheat sheet
      link: /reference/cheatsheet

features:
  - icon: 🚀
    title: Express-style HTTP
    details: Routes, middleware, CORS, Helmet-style security, sessions, multipart uploads, SSE, rate limits, and WebSocket upgrades — with familiar handler signatures.
  - icon: 🧩
    title: lodash & Array helpers
    details: Generic Map, Filter, Reduce, GroupBy, object Pick/Omit/Merge, and Set operations — typed with Go generics, zero reflection magic.
  - icon: 🗄️
    title: Full data layer
    details: Zod-like validation, Knex-style SQL, Redis, MongoDB, JSON helpers, Bull-like queues, and LRU caching — import only what you need.
  - icon: 🔐
    title: Security & identity
    details: bcrypt, JWT sign/verify, UUID/nanoid, and Passport-style Bearer, API key, and Basic auth middleware.
  - icon: 📦
    title: Independent packages
    details: No monolithic framework — each package is a separate import path with minimal coupling and documented escape hatches to underlying drivers.
  - icon: 📖
    title: Side-by-side docs
    details: Every API page shows Node.js, standard Go, and gox examples so you always know what you are replacing and why.
---

<HomeStats />

## Explore packages

<TierGrid />

## Side-by-side: the gox difference

Every API reference includes a three-way comparison. Here is a taste:

::: code-group

```js [Node.js]
const adults = users
  .filter(u => u.age >= 18)
  .map(u => ({ ...u, label: u.age >= 21 ? 'legal' : 'minor' }));
```

```go [Standard Go]
var adults []User
for _, u := range users {
    if u.Age < 18 {
        continue
    }
    label := "minor"
    if u.Age >= 21 {
        label = "legal"
    }
    adults = append(adults, User{ /* copy fields */, Label: label })
}
```

```go [gox]
import (
    "github.com/sahilkhaire/gox/cond"
    "github.com/sahilkhaire/gox/slice"
)

adults := slice.Map(
    slice.Filter(users, func(u User) bool { return u.Age >= 18 }),
    func(u User) User {
        u.Label = cond.If(u.Age >= 21, "legal", "minor")
        return u
    },
)
```

:::

## Install in seconds

```bash
go get github.com/sahilkhaire/gox
```

::: tip New from Node.js?
Start with the [Getting Started guide](/guide/getting-started), then jump to the package you used most in npm — [http](/packages/http/), [client](/packages/client/), or [slice](/packages/slice/).
:::
