# Getting Started

**gox** is a Node-friendly Go toolkit. If you have shipped Express apps, used lodash daily, or configured dotenv and axios — you already know most of the surface area. gox maps those patterns onto typed, idiomatic Go.

<HomeStats />

## Install

```bash
go get github.com/sahilkhaire/gox
```

gox is **not a framework**. Each capability lives in its own import path — pull in only what your service needs:

```go
import (
    "github.com/sahilkhaire/gox/cond"
    "github.com/sahilkhaire/gox/slice"
)

label := cond.If(age >= 18, "adult", "minor")
adults := slice.Filter(users, func(u User) bool { return u.Age >= 18 })
```

::: info Import paths
Every package is `github.com/sahilkhaire/gox/<name>`. The time helpers live at `gox/time` but use the identifier **`timex`** to avoid clashing with the standard library.
:::

## Your npm stack → gox

| You used in Node.js | gox package | Docs |
|---------------------|-------------|------|
| `express`, `cors`, `helmet` | [`http`](/packages/http/) | [HTTP guide](/guide/http) |
| `axios`, `fetch` | [`client`](/packages/client/) | [client.Fetch](/packages/client/fetch) |
| `lodash`, `Array.*` | [`slice`](/packages/slice/), [`maputil`](/packages/maputil/) | [slice.Map](/packages/slice/map) |
| `dotenv` | [`env`](/packages/env/) | [env.Load](/packages/env/load) |
| `zod`, `joi` | [`validate`](/packages/validate/) | [validate.Object](/packages/validate/object) |
| `knex`, `pg` | [`db`](/packages/db/) | [Migration: Knex](/migration/knex) |
| `ioredis` | [`redis`](/packages/redis/) | [redis.New](/packages/redis/new) |
| `mongoose` | [`mongo`](/packages/mongo/) | [mongo.Connect](/packages/mongo/connect) |
| `jsonwebtoken` | [`jwt`](/packages/jwt/) | [jwt.Sign](/packages/jwt/sign) |
| `passport` | [`auth`](/packages/auth/) | [Migration: Passport](/migration/passport) |
| `bull` | [`queue`](/packages/queue/) | [queue.New](/packages/queue/new) |
| `a ? b : c`, `??` | [`cond`](/packages/cond/) | [cond.If](/packages/cond/if) |

See the full [Cheat Sheet](/reference/cheatsheet) for every mapping.

## Three-way docs on every API

Each reference page follows the same premium layout:

1. **Signature** — full Go declaration with generics
2. **Overview** — what it does and when to reach for it
3. **Compare** — tabbed **Node.js · Standard Go · gox** examples
4. **Tips** — migration notes, performance, escape hatches

Example — [`slice.Map`](/packages/slice/map):

::: code-group

```js [Node.js]
const names = users.map(u => u.name);
```

```go [Standard Go]
names := make([]string, len(users))
for i, u := range users {
    names[i] = u.Name
}
```

```go [gox]
names := slice.Map(users, func(u User) string { return u.Name })
```

:::

## Recommended learning path

1. **[Architecture](/guide/architecture)** — module layout, context-first I/O, dependency policy
2. **Pick a package group** — [Utilities](/packages/cond/) → [Web](/packages/http/) → [Data](/packages/db/)
3. **[Migration guide](/migration/express)** for your heaviest npm dependency
4. **Browse package APIs** — each function has its own page with Node.js vs Go comparisons and examples

::: warning Must* helpers
Some packages expose `MustOpen`, `MustParse`, `MustValidate`, etc. These panic on failure — same trade-off as Node's `JSON.parse` vs `try/catch`. Prefer explicit error handling in library code.
:::

## Next steps

- [HTTP Guide](/guide/http) — Express → gox/http in one place
- [Package reference](/packages/cond/) — all 38 packages, 267+ APIs
- [Express migration](/migration/express) · [axios](/migration/axios) · [Knex](/migration/knex)

## License

MIT
