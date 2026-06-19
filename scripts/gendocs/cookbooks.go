package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// packageCookbooks holds rich multi-recipe content per package (markdown body only).
var packageCookbooks = map[string]string{
	"cond": `# cond Cookbook

Ternary and nullish-coalescing patterns from JavaScript, expressed in typed Go.

## Pass / fail label

` + "```go" + `
import "github.com/sahilkhaire/gox/cond"

label := cond.If(score >= 60, "pass", "fail")
` + "```" + `

## Lazy branches (avoid expensive work)

` + "```go" + `
// Only the chosen branch runs — unlike cond.If, both sides are NOT evaluated upfront.
v := cond.IfLazy(useCache, loadFromCache, loadFromDB)
` + "```" + `

## Nullish coalescing chain

` + "```go" + `
name := cond.Coalesce(maybeName, fallbackName, "guest")
` + "```" + `

## Safe pointer fallback

` + "```go" + `
display := cond.CoalescePtr(user.Nickname, user.Name, "Anonymous")
` + "```",

	"slice": `# slice Cookbook

lodash / Array-style pipelines with Go generics.

## Filter → map pipeline

` + "```go" + `
import "github.com/sahilkhaire/gox/slice"

adults := slice.Filter(users, func(u User) bool { return u.Age >= 18 })
names := slice.Map(adults, func(u User) string { return u.Name })
` + "```" + `

## Reduce to sum (from tests)

` + "```go" + `
in := []int{1, 2, 3, 4}
total := slice.Reduce(
    slice.Filter(slice.Map(in, func(x int) int { return x * 2 }), func(x int) bool { return x > 4 }),
    0,
    func(a, b int) int { return a + b },
) // 14
` + "```" + `

## Group by role

` + "```go" + `
byRole := slice.GroupBy(users, func(u User) string { return u.Role })
` + "```" + `

## Unique IDs preserving order

` + "```go" + `
ids := slice.Uniq([]string{"a", "b", "a", "c"})
` + "```",

	"http": `# http Cookbook

Express-style apps with middleware, JSON, routers, and error handling.

## Minimal API server

` + "```go" + `
import (
    goxerr "github.com/sahilkhaire/gox/err"
    goxhttp "github.com/sahilkhaire/gox/http"
)

app := goxhttp.New()
app.Use(goxhttp.Logger(), goxhttp.Recover(), goxhttp.Security())

app.Get("/health", func(c *goxhttp.Ctx) error {
    return c.JSON(200, map[string]string{"status": "ok"})
})

app.Post("/users", func(c *goxhttp.Ctx) error {
    var u struct{ Name string ` + "`json:\"name\"`" + ` }
    if err := c.BindJSON(&u); err != nil {
        return goxerr.BadRequest("invalid json")
    }
    return c.JSON(201, u)
})

app.Listen(":3000")
` + "```" + `

## Mounted sub-router

` + "```go" + `
api := app.Router()
api.Get("/users/{id}", func(c *goxhttp.Ctx) error {
    return c.JSON(200, map[string]string{"id": c.Param("id")})
})
app.Mount("/api", api)
` + "```" + `

## CORS + typed errors

` + "```go" + `
app.Use(goxhttp.CORS(goxhttp.CORSOptions{AllowOrigins: []string{"https://app.example"}}))
app.Get("/missing", func(c *goxhttp.Ctx) error {
    return goxerr.NotFound("not found") // → 404 JSON automatically
})
` + "```" + `

See also: [HTTP guide](/guide/http) · [Express migration](/migration/express)

`,

	"env": `# env Cookbook

dotenv-style configuration with typed getters.

` + "```go" + `
import "github.com/sahilkhaire/gox/env"

_ = env.Load() // reads .env from cwd

port := env.Get("PORT", "8080")
debug := env.GetBool("DEBUG", false)
timeout := env.GetDuration("TIMEOUT", 30*time.Second)
` + "```",

	"db": `# db Cookbook

Knex-style chainable SQL on sqlx.

` + "```go" + `
import "github.com/sahilkhaire/gox/db"

database, err := db.Open(ctx, "sqlite", ":memory:")
defer database.Close()

var user struct {
    ID   int    ` + "`db:\"id\"`" + `
    Name string ` + "`db:\"name\"`" + `
}
err = database.From("users").WhereEq("id", 1).First(ctx, &user)

err = database.Transaction(ctx, func(tx *db.Tx) error {
    _, err := tx.Insert(ctx, "users", map[string]any{"name": "Ada"})
    return err
})
` + "```" + `

Escape hatch: **` + "database.SQL" + `** exposes the underlying *sqlx.DB*.`,

	"validate": "# validate Cookbook\n\nZod/Joi-style schemas plus struct tag validation.\n\n## Struct tags\n\n```go\ntype User struct {\n    Email string `json:\"email\" validate:\"required,email\"`\n    Age   int    `json:\"age\" validate:\"gte=18\"`\n}\nif err := validate.Validate(&user); err != nil { /* 400 */ }\n```\n\n## Fluent object schema\n\n```go\nsch := validate.Object(map[string]validate.Field{\n    \"email\": validate.String().Required().Email(),\n    \"age\":   validate.Int().Min(18),\n})\nif err := validate.ValidateSchema(sch, data); err != nil { /* handle */ }\n```",

	"jwt": `# jwt Cookbook

jsonwebtoken-style sign and verify.

` + "```go" + `
import "github.com/sahilkhaire/gox/jwt"

token, err := jwt.Sign(
    map[string]any{"sub": "user-123"},
    secret,
    jwt.SignOptions{ExpiresIn: time.Hour},
)
claims, err := jwt.Verify(token, secret)
` + "```",

	"async": `# async Cookbook

Promise.all / race / sleep with context cancellation.

` + "```go" + `
import "github.com/sahilkhaire/gox/async"

a, b, err := async.All(ctx, fetchUsers, fetchPosts)
winner, err := async.Race(ctx, fastPath, slowPath)
if err := async.Sleep(ctx, 2*time.Second); err != nil { return err }
` + "```",

	"client": `# client Cookbook

axios / fetch-style HTTP client.

` + "```go" + `
import "github.com/sahilkhaire/gox/client"

c := client.New()
res, err := c.Get(ctx, "https://api.example.com/users", client.WithQuery(url.Values{
    "page": {"1"},
}))
var users []User
if err := res.JSON(&users); err != nil { /* handle */ }
` + "```",

	"redis": `# redis Cookbook

ioredis-style commands with context.

` + "```go" + `
import "github.com/sahilkhaire/gox/redis"

rdb, err := redis.New("localhost:6379")
defer rdb.Close()

err = rdb.Set(ctx, "key", "value", time.Hour)
val, err := rdb.Get(ctx, "key")
` + "```",

	"mongo": `# mongo Cookbook

mongoose-style collections and filters.

` + "```go" + `
import "github.com/sahilkhaire/gox/mongo"

client, err := mongo.Connect(ctx, uri)
coll := client.DB("app").Collection("users")

var doc bson.M
err = coll.FindOne(ctx, mongo.Eq("email", email)).Decode(&doc)
` + "```",

	"maputil": `# maputil Cookbook

lodash object utilities.

` + "```go" + `
import "github.com/sahilkhaire/gox/maputil"

small := maputil.Pick(config, "host", "port")
merged := maputil.Merge(defaults, overrides)
city, _ := maputil.Get[string](nested, "address.city")
` + "```",

	"str": `# str Cookbook

String casing, slugging, and padding for APIs and UI.

` + "```go" + `
import "github.com/sahilkhaire/gox/str"

slug := str.Slug("Hello World!")       // hello-world
camel := str.Camel("foo_bar")        // fooBar
title := str.Capitalize("hello")     // Hello
short := str.Truncate(longText, 80)  // safe rune truncate
` + "```",

	"json": `# json Cookbook

Typed JSON.parse/stringify with file helpers.

` + "```go" + `
import "github.com/sahilkhaire/gox/json"

cfg, err := json.ParseFile[Config](ctx, "config.json")
out, err := json.Stringify(response)
pretty, err := json.Pretty(debugPayload)
` + "```",

	"err": `# err Cookbook

http-errors style responses in gox/http handlers.

` + "```go" + `
import goxerr "github.com/sahilkhaire/gox/err"

app.Get("/users/{id}", func(c *goxhttp.Ctx) error {
    user, ok := store.Find(c.Param("id"))
    if !ok {
        return goxerr.NotFound("user not found")
    }
    return c.JSON(200, user)
})
` + "```",

	"fs": `# fs Cookbook

Context-aware file I/O — read, write, exists, mkdir.

` + "```go" + `
import "github.com/sahilkhaire/gox/fs"

data, err := fs.ReadFile(ctx, "config.json")
if ok, _ := fs.Exists(ctx, "uploads"); !ok {
    err = fs.Mkdir(ctx, "uploads", 0755)
}
err = fs.WriteFile(ctx, "out.json", data)
` + "```",

	"crypto": `# crypto Cookbook

Hashing, HMAC, and bcrypt password helpers.

` + "```go" + `
import "github.com/sahilkhaire/gox/crypto"

hash, err := crypto.HashPassword(password)
ok := crypto.CheckPassword(password, hash)
digest := crypto.SHA256([]byte(payload))
` + "```",

	"log": `# log Cookbook

Structured logging with slog underneath.

` + "```go" + `
import "github.com/sahilkhaire/gox/log"

logger := log.New()
log.SetDefault(logger)
logger.Info("request", "method", "GET", "path", "/health")
` + "```",

	"ws": `# ws Cookbook

WebSocket server and client dial.

` + "```go" + `
import "github.com/sahilkhaire/gox/ws"

conn, err := ws.Dial(ctx, "wss://example.com/ws", nil)
srv := ws.NewServer()
srv.Handle("/chat", func(c *ws.Conn) error { /* ... */ })
` + "```",

	"cron": `# cron Cookbook

Schedule recurring jobs with cron expressions.

` + "```go" + `
import "github.com/sahilkhaire/gox/cron"

sched := cron.New()
sched.Schedule("0 * * * *", func() { cleanup() })
sched.Start()
` + "```",

	"cache": `# cache Cookbook

In-memory LRU cache with capacity bound.

` + "```go" + `
import "github.com/sahilkhaire/gox/cache"

c := cache.New(1000)
c.Set("user:1", user, time.Hour)
if v, ok := c.Get("user:1"); ok { _ = v }
` + "```",

	"auth": `# auth Cookbook

Passport-style Bearer, API key, and Basic middleware.

` + "```go" + `
import "github.com/sahilkhaire/gox/auth"

app.Use(auth.Bearer(jwtSecret))
app.Use(auth.APIKey(auth.APIKeyOpts{Header: "X-API-Key", Keys: keys}))
claims, _ := auth.GetClaims(ctx)
` + "```",
}

func writeCookbookPage(dir, pkgName string) error {
	body, ok := packageCookbooks[pkgName]
	if !ok {
		return nil
	}
	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString(fmt.Sprintf("title: %q\n", pkgName+" Cookbook"))
	b.WriteString(fmt.Sprintf("package: %q\n", pkgName))
	b.WriteString(fmt.Sprintf("gox-doc-version: %q\n", docVersion))
	b.WriteString("---\n\n")
	b.WriteString(writeAPIMetaHTML(pkgName, packageNodeAnalog[pkgName], modulePath+"/"+pkgName))
	b.WriteString("\n")
	b.WriteString(body)
	b.WriteString(fmt.Sprintf("\n\n← [Back to %s overview](/packages/%s/)\n", pkgName, pkgName))
	return os.WriteFile(filepath.Join(dir, pkgName, "cookbook.md"), []byte(b.String()), 0o644)
}

func hasCookbook(pkgName string) bool {
	_, ok := packageCookbooks[pkgName]
	return ok
}
