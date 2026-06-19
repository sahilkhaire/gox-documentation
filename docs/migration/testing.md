# Jest / Vitest → Go testing

gox does not ship a test framework — Go's stdlib `testing` package is the standard. This guide maps common Jest patterns.

## Basics

| Jest | Go |
|------|-----|
| `describe('suite', () => { ... })` | `func TestSuite(t *testing.T) { ... }` |
| `it('should work', () => { ... })` | Subtest: `t.Run("should work", func(t *testing.T) { ... })` |
| `expect(x).toBe(y)` | `if x != y { t.Fatalf("got %v want %v", x, y) }` |
| `expect(x).toEqual(y)` | `reflect.DeepEqual` or `cmp.Diff` |
| `beforeEach(fn)` | Setup at start of `t.Run` or `TestMain` |
| `afterEach(fn)` | `defer` in test function |

```go
func TestUser(t *testing.T) {
    tests := []struct {
        name string
        age  int
        want string
    }{
        {"adult", 20, "adult"},
        {"minor", 15, "minor"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := cond.If(tt.age >= 18, "adult", "minor")
            if got != tt.want {
                t.Fatalf("got %q want %q", got, tt.want)
            }
        })
    }
}
```

## Mocking

| Jest | Go |
|------|-----|
| `jest.fn()` | Function fields on structs (interfaces) |
| `jest.mock('module')` | Define interface; inject mock implementation |
| `nock('http')` | `net/http/httptest` |

### HTTP mocking (supertest / nock)

```go
import (
    "net/http/httptest"
    goxhttp "github.com/sahilkhaire/gox/http"
)

func TestGetUser(t *testing.T) {
    app := goxhttp.New()
    app.Get("/users/1", func(c *goxhttp.Ctx) error {
        return c.JSON(200, map[string]string{"id": "1"})
    })

    req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
    rec := httptest.NewRecorder()
    app.ServeHTTP(rec, req)

    if rec.Code != 200 {
        t.Fatalf("status %d", rec.Code)
    }
}
```

### Client mocking (axios)

Use `httptest.NewServer` with `gox/client`:

```go
srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`{"ok":true}`))
}))
defer srv.Close()

c := client.New()
c.BaseURL = srv.URL
resp, err := c.Get(context.Background(), "/", nil)
```

## Async tests

| Jest | Go |
|------|-----|
| `async/await` | No await — call function, check error |
| `Promise.all` | `gox/async.All(ctx, fns...)` |
| `done()` callback | Not needed; test ends when function returns |

```go
func TestAsyncAll(t *testing.T) {
    ctx := context.Background()
    results, err := async.All(ctx,
        func() (int, error) { return 1, nil },
        func() (int, error) { return 2, nil },
    )
    if err != nil || len(results) != 2 {
        t.Fatal(err)
    }
}
```

## Snapshots

| Jest | Go |
|------|-----|
| `toMatchSnapshot()` | Golden files: compare against `testdata/*.golden` |
| Update snapshots | `go test -update` with custom flag |

## Coverage

| Jest | Go |
|------|-----|
| `jest --coverage` | `go test -cover ./...` |
| Coverage report | `go test -coverprofile=c.out && go tool cover -html=c.out` |

## Test helpers in gox

| Need | Package |
|------|---------|
| Ternary in tests | `gox/cond` |
| Slice assertions | `gox/slice` Contains, Every |
| Temp files | `gox/fs` with `t.TempDir()` |
| Env isolation | `gox/env` Set in test, restore after |

## Running tests

```bash
go test ./...                  # all packages
go test -race ./...            # race detector
go test -v ./http/...          # verbose single package
go test -run TestName ./...    # single test
```

CI is configured in `.github/workflows/ci.yml`.

## Further reading

- [Go testing package](https://pkg.go.dev/testing)
- [httptest package](https://pkg.go.dev/net/http/httptest)
