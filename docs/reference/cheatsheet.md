# Node.js → gox Cheat Sheet

Quick lookup: npm module → `gox` import path.

## Language & utilities

| npm / Node.js | gox import | Key API |
|---------------|------------|---------|
| `a ? b : c` | `gox/cond` | `cond.If(cond, a, b)` |
| `a ?? b` | `gox/cond` | `cond.Coalesce(a, b)` |
| `lodash` | `gox/slice`, `gox/maputil`, `gox/str` | `slice.Map`, `maputil.Pick` |
| `Promise.all` | `gox/async` | `async.All(ctx, fns...)` |
| `setTimeout` | `gox/async` | `async.After(ctx, d, fn)` |
| `events` | `gox/event` | `em.On`, `em.Emit` |
| `Buffer` | `gox/buffer` | `buffer.From`, `buffer.Concat` |

## File system & OS

| npm / Node.js | gox import | Key API |
|---------------|------------|---------|
| `fs.promises` | `gox/fs` | `fs.ReadFile(ctx, path)` |
| `path` | `gox/path` | `path.Join`, `path.Resolve` |
| `dotenv` | `gox/env` | `env.Load()`, `env.Get("KEY", def)` |
| `child_process` | `gox/exec` | `exec.Command(ctx, "ls")` |
| `os` | `gox/osutil` | `osutil.Hostname()`, `osutil.Homedir()` |
| `archiver` | `gox/archive` | `archive.ZipCreate`, `ZipExtract` |
| `zlib` | `gox/compress` | `compress.Gzip`, `Gunzip` |
| `stream` | `gox/stream` | `stream.Pipe`, `Transform` |

## Web

| npm / Node.js | gox import | Key API |
|---------------|------------|---------|
| `express` | `gox/http` | `http.New()`, `app.Get`, `app.Use` |
| `cors`, `helmet`, `morgan` | `gox/http` | `http.CORS`, `Security`, `Logger` |
| `multer` | `gox/http` | `http.ParseMultipart` |
| `express-session` | `gox/http` | `http.SessionMiddleware` |
| `express-rate-limit` | `gox/http` | `http.RateLimit` |
| `axios`, `fetch` | `gox/client` | `client.Get`, `client.Post` |
| `ws` | `gox/ws` | `ws.HandleWS`, `ws.Dial` |
| `http-errors` | `gox/err` | `err.NotFound`, `err.BadRequest` |
| `url`, `querystring` | `gox/url` | `url.Parse`, `url.ParseQuery` |

## Data

| npm / Node.js | gox import | Key API |
|---------------|------------|---------|
| `zod`, `joi` | `gox/validate` | `validate.Validate(&v)` |
| `JSON.parse` | `gox/json` | `json.Parse`, `json.Stringify` |
| `knex`, `pg` | `gox/db` | `db.From("t").Where(...).First` |
| `mongoose` | `gox/mongo` | `coll.FindOne`, `mongo.Eq` |
| `ioredis` | `gox/redis` | `rdb.Get`, `rdb.Set` |
| `papaparse` | `gox/csv` | `csv.Read`, `csv.Write` |

## Security

| npm / Node.js | gox import | Key API |
|---------------|------------|---------|
| `bcrypt`, `crypto` | `gox/crypto` | `crypto.HashPassword`, `SHA256` |
| `jsonwebtoken` | `gox/jwt` | `jwt.Sign`, `jwt.Verify` |
| `uuid`, `nanoid` | `gox/id` | `id.NewUUID`, `id.NewNanoID` |
| `passport` | `gox/auth` | `auth.Bearer`, `auth.APIKey` |

## Infrastructure

| npm / Node.js | gox import | Key API |
|---------------|------------|---------|
| `winston`, `pino` | `gox/log` | `log.Info`, `log.With` |
| `node-cache` | `gox/cache` | `cache.Set`, `cache.Get` |
| `node-cron` | `gox/cron` | `cron.Schedule`, `cron.Every` |
| `bull`, `bullmq` | `gox/queue` | `queue.Enqueue`, `Worker` |
| `nodemailer` | `gox/mail` | `mail.Send` |
| `moment`, `dayjs`, `ms` | `gox/time` | `timex.Format`, `timex.ParseDuration` |
| `semver` | `gox/semver` | `semver.Parse`, `semver.Satisfies` |

## Ecosystem guides (no gox package)

| npm / Node.js | Go approach | Guide |
|---------------|-------------|-------|
| `apollo-server`, `graphql` | gqlgen | [graphql.md](./migration/graphql.md) |
| `@grpc/grpc-js` | `google.golang.org/grpc` | [grpc.md](./migration/grpc.md) |
| `kafkajs` | `segmentio/kafka-go` | [kafka.md](./migration/kafka.md) |
| `@opentelemetry/*` | `go.opentelemetry.io/otel` | [opentelemetry.md](./migration/opentelemetry.md) |
| `jest`, `vitest` | `testing`, `httptest` | [testing.md](./migration/testing.md) |
| `socket.io` | `gox/ws` + patterns | [ws.md](./migration/ws.md) |

## Install

```bash
go get github.com/sahilkhaire/gox
```

Import only what you need — each package is independent:

```go
import (
    "github.com/sahilkhaire/gox/http"
    "github.com/sahilkhaire/gox/db"
)
```
