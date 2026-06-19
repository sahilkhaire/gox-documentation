package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Curated enrichment for API docs: Node.js snippets, stdlib Go alternatives, gox examples.

type enrichment struct {
	Node        string
	StdGo       string
	Gox         string
	When        string
	Description string
	Tips        string
	Pitfall     string
	Example     string // from tests or curated full example
}

var testExamples map[string]string

var enrichments = map[string]enrichment{
	// cond
	"cond.If": {
		Description: "Returns `a` when `cond` is true, otherwise `b`. The closest equivalent to JavaScript's ternary operator — both branches must be the same type in Go.",
		Node:        "const label = age >= 18 ? 'adult' : 'minor';",
		StdGo:       "label := \"minor\"\nif age >= 18 {\n    label = \"adult\"\n}",
		Gox:         "label := cond.If(age >= 18, \"adult\", \"minor\")",
		Example:     "adult := cond.If(age >= 18, \"adult\", \"minor\")",
		Tips:        "Both branches are evaluated eagerly. Use `cond.IfLazy` when one branch is expensive and should be skipped.",
	},
	"cond.IfLazy": {
		Node:  "const v = cond ? expensive() : fallback();",
		StdGo: "var v T\nif cond {\n    v = expensive()\n} else {\n    v = fallback()\n}",
		Gox:   "v := cond.IfLazy(cond, expensive, fallback)",
	},
	"cond.Coalesce": {
		Node:  "const name = maybeName ?? 'guest';",
		StdGo: "name := maybeName\nif name == \"\" {\n    name = \"guest\"\n}",
		Gox:   "name := cond.Coalesce(maybeName, \"guest\")",
	},
	"cond.CoalescePtr": {
		Node:  "const name = user?.name ?? 'guest';",
		StdGo: "name := \"guest\"\nif user != nil {\n    name = user.Name\n}",
		Gox:   "name := cond.CoalescePtr(&user.Name, \"guest\")",
	},

	// slice
	"slice.Map": {
		Description: "Transforms each element of a slice, returning a new slice — identical mental model to `Array.prototype.map`. Uses Go generics so input and output types can differ safely.",
		Node:        "const names = users.map(u => u.name);",
		StdGo:       "names := make([]string, len(users))\nfor i, u := range users {\n    names[i] = u.Name\n}",
		Gox:         "names := slice.Map(users, func(u User) string { return u.Name })",
		Tips:        "Chain with `slice.Filter` and `slice.Reduce` for lodash-style pipelines. Pre-allocates output slice for performance.",
	},
	"slice.Filter": {
		Description: "Keeps elements matching a predicate — same as `Array.filter`. Returns a new slice containing only matching items.",
		Node:        "const adults = users.filter(u => u.age >= 18);",
		StdGo:       "var adults []User\nfor _, u := range users {\n    if u.Age >= 18 {\n        adults = append(adults, u)\n    }\n}",
		Gox:         "adults := slice.Filter(users, func(u User) bool { return u.Age >= 18 })",
		Tips:        "Combine with `slice.Map` for filter-then-map pipelines without nested loops.",
	},
	"slice.Reduce": {
		Node:  "const total = nums.reduce((acc, n) => acc + n, 0);",
		StdGo: "total := 0\nfor _, n := range nums {\n    total += n\n}",
		Gox:   "total := slice.Reduce(nums, 0, func(acc, n int) int { return acc + n })",
	},
	"slice.Find": {
		Node:  "const user = users.find(u => u.id === id);",
		StdGo: "var found User\nvar ok bool\nfor _, u := range users {\n    if u.ID == id {\n        found, ok = u, true\n        break\n    }\n}",
		Gox:   "user, ok := slice.Find(users, func(u User) bool { return u.ID == id })",
	},
	"slice.GroupBy": {
		Node:  "const byRole = _.groupBy(users, u => u.role);",
		StdGo: "byRole := make(map[string][]User)\nfor _, u := range users {\n    byRole[u.Role] = append(byRole[u.Role], u)\n}",
		Gox:   "byRole := slice.GroupBy(users, func(u User) string { return u.Role })",
	},
	"slice.Uniq": {
		Node:  "const unique = _.uniq(items);",
		StdGo: "seen := make(map[T]struct{})\nvar unique []T\nfor _, v := range items { /* dedupe */ }",
		Gox:   "unique := slice.Uniq(items)",
	},

	// maputil
	"maputil.Pick": {
		Node:  "const subset = _.pick(obj, ['a', 'b']);",
		StdGo: "subset := map[string]any{}\nfor _, k := range []string{\"a\", \"b\"} {\n    if v, ok := obj[k]; ok {\n        subset[k] = v\n    }\n}",
		Gox:   "subset := maputil.Pick(obj, \"a\", \"b\")",
	},
	"maputil.Get": {
		Node:  "const city = _.get(obj, 'address.city', 'unknown');",
		StdGo: "city := \"unknown\"\nif addr, ok := obj[\"address\"].(map[string]any); ok {\n    if c, ok := addr[\"city\"].(string); ok {\n        city = c\n    }\n}",
		Gox:   "city, _ := maputil.Get[string](obj, \"address.city\")",
	},

	// env
	"env.Load": {
		Node:  "require('dotenv').config();",
		StdGo: "// use os.Getenv after loading .env manually or via a config lib",
		Gox:   "if err := env.Load(); err != nil {\n    log.Fatal(err)\n}",
	},
	"env.Get": {
		Node:  "const port = process.env.PORT || '8080';",
		StdGo: "port := os.Getenv(\"PORT\")\nif port == \"\" {\n    port = \"8080\"\n}",
		Gox:   "port := env.Get(\"PORT\", \"8080\")",
	},

	// fs
	"fs.ReadFile": {
		Node:  "const data = await fs.promises.readFile('file.txt');",
		StdGo: "data, err := os.ReadFile(\"file.txt\")",
		Gox:   "data, err := fs.ReadFile(ctx, \"file.txt\")",
	},

	// async
	"async.All": {
		Node:  "const [a, b] = await Promise.all([fetchA(), fetchB()]);",
		StdGo: "errg, ctx := errgroup.WithContext(ctx)\n// run goroutines...",
		Gox:   "a, b, err := async.All(ctx, fetchA, fetchB)",
	},
	"async.Sleep": {
		Node:  "await new Promise(r => setTimeout(r, 1000));",
		StdGo: "select {\ncase <-time.After(time.Second):\ncase <-ctx.Done():\n    return ctx.Err()\n}",
		Gox:   "if err := async.Sleep(ctx, time.Second); err != nil { return err }",
	},

	// http
	"http.New": {
		Description: "Creates a new Express-style application with chi router underneath. Register routes with `Get`/`Post`/… and global middleware with `Use`.",
		Node:        "const app = express();",
		StdGo:       "mux := http.NewServeMux()\nhttp.ListenAndServe(\":8080\", mux)",
		Gox:         "app := goxhttp.New()",
		Tips:        "Stack `Logger`, `Recover`, and `Security` middleware like you would morgan + helmet in Express. See the [HTTP guide](/guide/http).",
	},
	"http.Ctx.JSON": {
		Description: "Writes a JSON response with the given status code — combines `res.status()` and `res.json()` from Express into one call that returns an error for the framework to handle.",
		Node:        "res.status(200).json({ ok: true });",
		StdGo:       "w.Header().Set(\"Content-Type\", \"application/json\")\nw.WriteHeader(200)\njson.NewEncoder(w).Encode(map[string]bool{\"ok\": true})",
		Gox:         "return c.JSON(200, map[string]bool{\"ok\": true})",
		Tips:        "Return the error from handlers — `*err.AppError` sets HTTP status automatically.",
	},

	// client
	"client.Fetch": {
		Node:  "const res = await fetch(url);",
		StdGo: "resp, err := http.Get(url)",
		Gox:   "res, err := client.Fetch(ctx, url)",
	},

	// db
	"db.Open": {
		Node:  "const db = knex({ client: 'pg', connection: process.env.DATABASE_URL });",
		StdGo: "sqlx.Connect(\"postgres\", os.Getenv(\"DATABASE_URL\"))",
		Gox:   "database, err := db.Open(ctx, \"postgres\", os.Getenv(\"DATABASE_URL\"))",
	},

	// json
	"json.Parse": {
		Node:  "const obj = JSON.parse(str);",
		StdGo: "json.Unmarshal([]byte(str), &obj)",
		Gox:   "obj, err := json.Parse[MyType](str)",
	},
	"json.Stringify": {
		Node:  "const str = JSON.stringify(obj);",
		StdGo: "b, err := json.Marshal(obj)",
		Gox:   "str, err := json.Stringify(obj)",
	},

	// jwt
	"jwt.Sign": {
		Node:  "const token = jwt.sign(payload, secret);",
		StdGo: "// use golang-jwt directly",
		Gox:   "token, err := jwt.Sign(claims, secret, jwt.SignOptions{})",
	},
	"jwt.Verify": {
		Node:  "const payload = jwt.verify(token, secret);",
		StdGo: "// use golang-jwt directly",
		Gox:   "claims, err := jwt.Verify(token, secret)",
	},

	// err
	// err
	"err.NotFound": {
		Node:  "next(createError(404, 'not found'));",
		StdGo: "return fmt.Errorf(\"not found: %w\", ErrNotFound)",
		Gox:   "return err.NotFound(\"not found\")",
	},
	"err.Forbidden": {
		Node:  "next(createError(403, 'forbidden'));",
		StdGo: "return fmt.Errorf(\"forbidden\")",
		Gox:   "return err.Forbidden(\"forbidden\")",
	},
	"err.BadRequest": {
		Node:  "next(createError(400, 'bad request'));",
		Gox:   "return err.BadRequest(\"bad request\")",
	},
	"err.Unauthorized": {
		Node:  "next(createError(401, 'unauthorized'));",
		Gox:   "return err.Unauthorized(\"unauthorized\")",
	},
	"err.Internal": {
		Node:  "next(createError(500, 'internal'));",
		Gox:   "return err.Internal(\"internal error\")",
	},

	// str
	"str.Slug": {
		Node:  "slugify('Hello World');",
		Gox:   "slug := str.Slug(\"Hello World\")",
	},
	"str.Camel": {
		Node:  "_.camelCase('foo_bar');",
		Gox:   "s := str.Camel(\"foo_bar\")",
	},
	"str.Truncate": {
		Node:  "_.truncate(str, { length: 10 });",
		Gox:   "short := str.Truncate(long, 10)",
	},

	// set
	"set.New": {
		Node:  "const s = new Set([1, 2, 3]);",
		Gox:   "s := set.New(1, 2, 3)",
	},
	"set.Union": {
		Node:  "new Set([...a, ...b]);",
		Gox:   "u := set.Union(a, b)",
	},

	// path
	"path.Join": {
		Node:  "path.join('a', 'b', 'c');",
		StdGo: "filepath.Join(\"a\", \"b\", \"c\")",
		Gox:   "p := path.Join(\"a\", \"b\", \"c\")",
	},
	"path.Basename": {
		Node:  "path.basename('/foo/bar.txt');",
		Gox:   "name := path.Basename(\"/foo/bar.txt\")",
	},

	// url
	"url.Parse": {
		Node:  "new URL('https://example.com/path');",
		Gox:   "u, err := url.Parse(\"https://example.com/path\")",
	},
	"url.ParseQuery": {
		Node:  "querystring.parse('a=1&b=2');",
		Gox:   "vals, err := url.ParseQuery(\"a=1&b=2\")",
	},

	// buffer
	"buffer.From": {
		Node:  "Buffer.from([1, 2, 3]);",
		Gox:   "b := buffer.From([]byte{1, 2, 3})",
	},
	"buffer.Concat": {
		Node:  "Buffer.concat([buf1, buf2]);",
		Gox:   "combined := buffer.Concat(buf1, buf2)",
	},

	// validate
	"validate.Validate": {
		Node:  "schema.validate(data);",
		Gox:   "if err := validate.Validate(&v); err != nil { /* handle */ }",
	},
	"validate.Object": {
		Node:  "z.object({ name: z.string() });",
		Gox:   "sch := validate.Object(map[string]validate.Field{\"name\": validate.String().Required()})",
	},

	// redis
	"redis.New": {
		Node:  "const redis = new Redis();",
		Gox:   "rdb, err := redis.New(\"localhost:6379\")",
	},
	"redis.Get": {
		Node:  "await redis.get('key');",
		Gox:   "val, err := rdb.Get(ctx, \"key\")",
	},

	// mongo
	"mongo.Connect": {
		Node:  "await mongoose.connect(uri);",
		Gox:   "client, err := mongo.Connect(ctx, uri)",
	},

	// queue
	"queue.New": {
		Node:  "const queue = new Queue('tasks');",
		Gox:   "client, err := queue.New(\"localhost:6379\")",
	},

	// event
	"event.New": {
		Node:  "const emitter = new EventEmitter();",
		Gox:   "em := event.New()",
	},

	// cron
	"cron.New": {
		Node:  "const job = cron.schedule('* * * * *', fn);",
		Gox:   "sched := cron.New()\nsched.Schedule(\"* * * * *\", fn)",
	},

	// cache
	"cache.New": {
		Node:  "const cache = new LRUCache({ max: 100 });",
		Gox:   "c := cache.New(100)",
	},

	// ws
	"ws.NewServer": {
		Node:  "const wss = new WebSocket.Server({ server });",
		Gox:   "srv := ws.NewServer()",
	},
	"ws.Dial": {
		Node:  "const ws = new WebSocket('wss://example.com');",
		Gox:   "conn, err := ws.Dial(ctx, \"wss://example.com\", nil)",
	},

	// crypto
	"crypto.SHA256": {
		Node:  "crypto.createHash('sha256').update(data).digest('hex');",
		Gox:   "hash := crypto.SHA256(data)",
	},
	"crypto.HashPassword": {
		Node:  "await bcrypt.hash(password, 10);",
		Gox:   "hash, err := crypto.HashPassword(password)",
	},

	// id
	"id.NewUUID": {
		Node:  "uuid.v4();",
		Gox:   "id := id.NewUUID()",
	},
	"id.NewNanoID": {
		Node:  "nanoid();",
		Gox:   "id := id.NewNanoID(21)",
	},

	// log
	"log.New": {
		Node:  "winston.createLogger();",
		Gox:   "logger := log.New()",
	},

	// compress
	"compress.Gzip": {
		Node:  "zlib.gzipSync(data);",
		Gox:   "compressed, err := compress.Gzip(data)",
	},

	// semver
	"semver.Parse": {
		Node:  "semver.parse('1.2.3');",
		Gox:   "v, err := semver.Parse(\"1.2.3\")",
	},

	// exec
	"exec.Command": {
		Node:  "spawn('ls', ['-la']);",
		Gox:   "cmd := exec.Command(ctx, \"ls\", \"-la\")",
	},

	// osutil
	"osutil.Hostname": {
		Node:  "os.hostname();",
		Gox:   "name, err := osutil.Hostname()",
	},
}

func lookupEnrichment(pkg, name, receiver string) (enrichment, string, bool) {
	key := pkg + "." + name
	if receiver != "" {
		key = pkg + "." + receiver + "." + name
	}
	if e, ok := enrichments[key]; ok {
		if ex, ok2 := testExamples[key]; ok2 && e.Example == "" {
			e.Example = ex
		}
		return e, key, true
	}
	if receiver != "" {
		key = pkg + "." + name
	}
	if e, ok := enrichments[key]; ok {
		if ex, ok2 := testExamples[key]; ok2 && e.Example == "" {
			e.Example = ex
		}
		return e, key, true
	}
	if e, ok := autoEnrichments[key]; ok {
		if ex, ok := testExamples[key]; ok && e.Example == "" {
			e.Example = ex
		}
		return e, key, true
	}
	if receiver != "" {
		key = pkg + "." + receiver + "." + name
		if e, ok := autoEnrichments[key]; ok {
			if ex, ok := testExamples[key]; ok && e.Example == "" {
				e.Example = ex
			}
			return e, key, true
		}
	}
	// test-only enrichment
	if ex, ok := testExamples[pkg+"."+name]; ok {
		return enrichment{Example: ex}, pkg + "." + name, true
	}
	if receiver != "" {
		if ex, ok := testExamples[pkg+"."+receiver+"."+name]; ok {
			return enrichment{Example: ex}, pkg + "." + receiver + "." + name, true
		}
	}
	return enrichment{}, "", false
}

// autoEnrichments is populated from gox/docs/node-mapping.md at build time.
var autoEnrichments = map[string]enrichment{}

func loadAutoEnrichments(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	rowRe := regexp.MustCompile(`^\|\s*` + "`" + `([^` + "`" + `]+)` + "`" + `\s*\|\s*` + "`" + `([^` + "`" + `]+)` + "`" + `\s*\|`)
	var currentPkg string
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "## ") {
			h := strings.TrimPrefix(line, "## ")
			if idx := strings.Index(h, " — "); idx > 0 {
				currentPkg = strings.Fields(h[:idx])[0]
			} else {
				currentPkg = strings.Fields(h)[0]
			}
		}
		matches := rowRe.FindStringSubmatch(line)
		if len(matches) != 3 || currentPkg == "" {
			continue
		}
		nodeExpr, goxExpr := strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2])
		sym := extractSymbol(goxExpr)
		if sym == "" {
			continue
		}
		key := currentPkg + "." + sym
		if _, exists := enrichments[key]; exists {
			continue
		}
		pkgIdent := currentPkg
		if currentPkg == "time" {
			pkgIdent = "timex"
		}
		gox := goxExpr
		if !strings.Contains(gox, "(") && !strings.HasPrefix(gox, pkgIdent) {
			gox = pkgIdent + "." + gox
		}
		autoEnrichments[key] = enrichment{
			Node:        nodeExpr,
			Gox:         gox,
			Description: fmt.Sprintf("Maps the Node.js pattern `%s` to gox `%s`. %s", nodeExpr, goxExpr, packageBlurb(currentPkg)),
		}
	}
}

func packageBlurb(pkg string) string {
	if analog, ok := packageNodeAnalogForEnrich[pkg]; ok {
		return "Part of the " + pkg + " package — Node.js analog: " + analog + "."
	}
	return ""
}

// mirror of main.go analogs for auto descriptions
var packageNodeAnalogForEnrich = map[string]string{
	"cond": "ternary ? :, nullish ??", "slice": "lodash / Array.*", "http": "express",
	"client": "axios", "env": "dotenv", "db": "knex", "validate": "zod/joi",
}
