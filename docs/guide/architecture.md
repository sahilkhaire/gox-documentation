# gox Architecture

## Module

`github.com/sahilkhaire/gox` — Go 1.22+

## Folder rules

1. **Public packages** live at the module root (`http/`, `slice/`, …). Import path = directory name.
2. **Nesting** only when a domain outgrows one package (e.g. `http/middleware/`).
3. **`internal/`** holds private code shared by 2+ packages. Not importable externally.
4. **`examples/`** — one subdirectory per shipped feature; no empty placeholders.
5. **`docs/`** — plans, architecture, API mapping. Root keeps only `README.md`.
6. **No `pkg/` or `cmd/`** unless explicitly needed later.

## Conventions

| Rule | Detail |
|------|--------|
| Context | I/O functions take `context.Context` as the first argument |
| Errors | Return errors; `Must*` helpers may panic (documented) |
| Generics | Collection helpers use type parameters |
| Tests | Stdlib `testing` only; table-driven tests |
| Imports | No cycles between public packages |
| Node docs | Exported symbols include `// Node: ...` where applicable |

## Dependency policy

| Category | Policy |
|----------|--------|
| Utilities & I/O | stdlib only |
| Web, data, security, infrastructure | Thin wrappers over proven libraries; justify in this file |

### Approved dependencies

| Package | Dependency | Reason |
|---------|------------|--------|
| `gox/http` | `github.com/go-chi/chi/v5` | stdlib-compatible router |
| `gox/validate` | `github.com/go-playground/validator/v10` | validation engine |
| `gox/db` | `github.com/jmoiron/sqlx` | SQL ergonomics |
| `gox/redis` | `github.com/redis/go-redis/v9` | Redis client |
| `gox/mongo` | `go.mongodb.org/mongo-driver` | MongoDB client |
| `gox/jwt` | `github.com/golang-jwt/jwt/v5` | JWT |
| `gox/cron` | `github.com/robfig/cron/v3` | Cron scheduling |
| `gox/queue` | `github.com/hibiken/asynq` | Redis task queue (Bull-like) |
| `gox/semver` | `github.com/Masterminds/semver/v3` | Semantic versioning |
| `gox/ws` | `github.com/gorilla/websocket` | WebSocket upgrade and framing |
| `gox/http` (rate limit) | `golang.org/x/time/rate` | Token-bucket rate limiting |

Packages using stdlib only include `gox/archive`, `gox/csv`, `gox/exec`, `gox/osutil`, `gox/cache`, `gox/event`, `gox/mail` (`net/smtp`), `gox/compress` (`compress/gzip`), `gox/stream` (`io`), and `gox/time` (package `timex`). `gox/http` upload, session, and SSE features use stdlib multipart and cookies.

New dependencies require a row in this table before merge.

## Adding a new package

1. Create `/<pkg>/` with `doc.go`, implementation, tests
2. Add section to `docs/node-mapping.md`
3. Mark **shipped** in `docs/feature-catalog.md`
4. Add `examples/<pkg>/` if the API is non-trivial

### Test-only dependencies

| Package | Dependency | Reason |
|---------|------------|--------|
| `gox/db` tests | `modernc.org/sqlite` | In-memory SQLite without CGO |
| `gox/redis` tests | `github.com/alicebob/miniredis/v2` | Embedded Redis for unit tests |
| `gox/queue` tests | `github.com/alicebob/miniredis/v2` | Embedded Redis for asynq client/worker tests |
| `gox/crypto` | `golang.org/x/crypto/bcrypt` | Password hashing |
| `gox/id` | `github.com/google/uuid` | UUID generation |
