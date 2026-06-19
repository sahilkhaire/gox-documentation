package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type packageNarrative struct {
	Intro         string
	UseCases      []string
	MigrationLink string
}

type featuredSymbol struct {
	Name    string
	Slug    string
	Summary string
}

var packageNarratives = map[string]packageNarrative{
	"cond":      {Intro: "Replace JavaScript ternary and nullish coalescing with typed, generic helpers. Ideal for config defaults, labels, and safe fallbacks without nested if/else noise.", UseCases: []string{"Pass/fail labels from boolean conditions", "Default values when strings or pointers are empty", "Lazy evaluation when one branch is expensive"}, MigrationLink: ""},
	"slice":     {Intro: "lodash and Array methods as generic Go functions — map, filter, reduce, groupBy, and more with compile-time type safety.", UseCases: []string{"Transform API response arrays", "Filter collections before mapping", "Group records by key for dashboards"}, MigrationLink: ""},
	"maputil":   {Intro: "Object utilities from lodash — pick, omit, merge, deep get/set — for config maps and dynamic JSON.", UseCases: []string{"Pick allowed keys from request bodies", "Merge defaults with overrides", "Read nested paths like address.city"}, MigrationLink: ""},
	"str":       {Intro: "String casing, slugging, padding, and truncation helpers matching lodash and String prototype patterns.", UseCases: []string{"Generate URL slugs from titles", "Normalize casing for APIs", "Truncate UI labels safely by rune count"}, MigrationLink: ""},
	"set":       {Intro: "Set operations with union, intersection, and difference — a typed alternative to JavaScript Set utilities.", UseCases: []string{"Dedupe tag lists", "Compare permission sets", "Merge unique IDs from multiple sources"}, MigrationLink: ""},
	"buffer":    {Intro: "Buffer-style byte helpers for concatenation, comparison, and construction from strings or slices.", UseCases: []string{"Concatenate binary payloads", "Compare constant-time byte slices", "Build buffers from UTF-8 strings"}, MigrationLink: ""},
	"async":     {Intro: "Promise.all, race, sleep, retry, and timeout patterns with context cancellation built in.", UseCases: []string{"Fetch multiple resources in parallel", "Implement deadlines on outbound calls", "Retry flaky operations with backoff"}, MigrationLink: ""},
	"path":      {Intro: "Node path helpers on top of filepath — join, basename, extname, and absolute checks.", UseCases: []string{"Build cross-platform file paths", "Extract extensions for uploads", "Resolve relative config paths"}, MigrationLink: ""},
	"url":       {Intro: "Parse, format, and query-string utilities matching Node url and querystring modules.", UseCases: []string{"Parse redirect URLs", "Build query strings for APIs", "Resolve relative links"}, MigrationLink: ""},
	"env":       {Intro: "dotenv-style loading plus typed getters for strings, ints, booleans, and durations.", UseCases: []string{"Load .env in local dev", "Read PORT with defaults", "Parse feature flags as booleans"}, MigrationLink: ""},
	"fs":        {Intro: "fs.promises-style file helpers with context-first I/O — read, write, copy, mkdir, and stat.", UseCases: []string{"Read config files at startup", "Write uploaded content to disk", "Check existence before processing"}, MigrationLink: ""},
	"err":       {Intro: "http-errors style typed HTTP errors with status codes — return from handlers for automatic JSON responses.", UseCases: []string{"Return 404 from REST handlers", "Wrap validation failures as 400", "Map domain errors to status codes"}, MigrationLink: ""},
	"exec":      {Intro: "child_process-style command execution with context timeouts and captured stdout/stderr.", UseCases: []string{"Run CLI tools from Go services", "Execute migration scripts", "Capture subprocess output"}, MigrationLink: ""},
	"osutil":    {Intro: "os module helpers — hostname, homedir, temp dirs, platform, and architecture.", UseCases: []string{"Detect runtime environment", "Resolve home directory paths", "Create temp workspaces"}, MigrationLink: ""},
	"stream":    {Intro: "Node stream patterns — pipe, tee, transform, and read-all for io.Reader chains.", UseCases: []string{"Pipe HTTP bodies to storage", "Duplicate readers for checksum + upload", "Drain streams to bytes"}, MigrationLink: ""},
	"compress":  {Intro: "zlib-style gzip and gunzip helpers for compressing HTTP payloads and archives.", UseCases: []string{"Compress API responses", "Decompress uploaded gzip files", "Shrink log batches"}, MigrationLink: ""},
	"csv":       {Intro: "papaparse-style CSV read, write, and parse for tabular data import/export.", UseCases: []string{"Import user CSV uploads", "Export reports", "Parse inline CSV strings"}, MigrationLink: ""},
	"archive":   {Intro: "archiver-style zip and tar create/extract for bundling files.", UseCases: []string{"Bundle download artifacts", "Extract uploaded archives", "Stream tar entries"}, MigrationLink: ""},
	"http":      {Intro: "Express-style HTTP apps on chi — routes, middleware, JSON, sessions, uploads, SSE, rate limits, and WebSockets.", UseCases: []string{"Build REST APIs with middleware stacks", "Serve JSON with typed errors", "Add CORS, security headers, and logging"}, MigrationLink: "/migration/express"},
	"client":    {Intro: "axios/fetch-style HTTP client with query params, headers, JSON bodies, and response helpers.", UseCases: []string{"Call third-party APIs", "POST JSON with typed responses", "Set default client for services"}, MigrationLink: "/migration/axios"},
	"log":       {Intro: "winston/pino-style structured logging with slog underneath.", UseCases: []string{"Structured request logging", "Set default logger for packages", "Level-based log output"}, MigrationLink: ""},
	"ws":        {Intro: "WebSocket server, client dial, and upgrade helpers built on gorilla/websocket.", UseCases: []string{"Real-time chat backends", "Upgrade HTTP handlers to WS", "Connect to external WS APIs"}, MigrationLink: "/migration/ws"},
	"json":      {Intro: "JSON.parse/stringify with generics, pretty printing, and file helpers.", UseCases: []string{"Parse API payloads into structs", "Pretty-print debug output", "Read/write JSON config files"}, MigrationLink: ""},
	"validate":  {Intro: "Zod/Joi-style validation via struct tags and fluent object schemas.", UseCases: []string{"Validate request bodies", "Build reusable field schemas", "Return 400 with field errors"}, MigrationLink: ""},
	"db":        {Intro: "Knex-style chainable SQL on sqlx — queries, inserts, transactions, and escape hatches to raw SQL.", UseCases: []string{"CRUD with WhereEq chains", "Run migrations in transactions", "Access underlying *sqlx.DB* when needed"}, MigrationLink: "/migration/knex"},
	"redis":     {Intro: "ioredis-style Redis client with get/set, pub/sub, and context-aware commands.", UseCases: []string{"Cache hot keys", "Pub/sub event fan-out", "Session or rate-limit storage"}, MigrationLink: ""},
	"mongo":     {Intro: "mongoose-style MongoDB client, collections, and filter builders.", UseCases: []string{"Connect with URI", "Find documents by field", "Use Eq/In/Gt filter helpers"}, MigrationLink: ""},
	"crypto":    {Intro: "Node crypto and bcrypt patterns — hashing, HMAC, random bytes, and password helpers.", UseCases: []string{"Hash passwords for storage", "Sign webhook payloads", "Generate secure tokens"}, MigrationLink: ""},
	"jwt":       {Intro: "jsonwebtoken-style sign, verify, and decode with golang-jwt underneath.", UseCases: []string{"Issue session tokens", "Verify Bearer tokens in middleware", "Inspect claims without verifying (decode)"}, MigrationLink: ""},
	"id":        {Intro: "uuid and nanoid generation and parsing for identifiers.", UseCases: []string{"Generate request IDs", "Parse UUID path params", "Short public IDs with nanoid"}, MigrationLink: ""},
	"auth":      {Intro: "Passport-style Bearer, API key, and Basic auth middleware with context claims.", UseCases: []string{"Protect routes with JWT Bearer", "Accept API keys from headers", "Read authenticated user from context"}, MigrationLink: "/migration/passport"},
	"cache":     {Intro: "lru-cache-style in-memory LRU with TTL-style eviction by capacity.", UseCases: []string{"Memoize expensive lookups", "Cache parsed config", "Bound memory for hot keys"}, MigrationLink: ""},
	"event":     {Intro: "EventEmitter-style pub/sub for in-process events.", UseCases: []string{"Decouple modules with events", "Notify listeners on state changes", "Build plugin hooks"}, MigrationLink: ""},
	"cron":      {Intro: "node-cron-style schedulers for recurring jobs.", UseCases: []string{"Run cleanup every hour", "Schedule report generation", "Parse cron expressions"}, MigrationLink: ""},
	"queue":     {Intro: "Bull-like Redis task queues with asynq — enqueue, workers, and handlers.", UseCases: []string{"Background email sending", "Retry failed jobs", "Separate worker processes"}, MigrationLink: ""},
	"mail":      {Intro: "nodemailer-style SMTP sending and MIME building.", UseCases: []string{"Send transactional email", "Build multipart messages", "Configure SMTP transport"}, MigrationLink: ""},
	"time":      {Intro: "moment/dayjs-style date parsing, formatting, and calendar boundaries. Import as timex to avoid clashing with stdlib time.", UseCases: []string{"Format ISO timestamps", "Start/end of day calculations", "Parse duration strings"}, MigrationLink: ""},
	"semver":    {Intro: "semver module helpers — parse, compare, increment, and satisfy ranges.", UseCases: []string{"Gate features by app version", "Compare dependency versions", "Bump version segments"}, MigrationLink: ""},
}

var featuredSymbols = map[string][]featuredSymbol{
	"cond":      {{Name: "If", Slug: "if", Summary: "Ternary ? : operator"}, {Name: "Coalesce", Slug: "coalesce", Summary: "Nullish ?? chain"}, {Name: "IfLazy", Slug: "if-lazy", Summary: "Lazy ternary branches"}},
	"slice":     {{Name: "Map", Slug: "map", Summary: "Array.map"}, {Name: "Filter", Slug: "filter", Summary: "Array.filter"}, {Name: "GroupBy", Slug: "group-by", Summary: "lodash groupBy"}},
	"maputil":   {{Name: "Pick", Slug: "pick", Summary: "_.pick"}, {Name: "Get", Slug: "get", Summary: "_.get path"}, {Name: "Merge", Slug: "merge", Summary: "_.merge"}},
	"str":       {{Name: "Slug", Slug: "slug", Summary: "URL-friendly slug"}, {Name: "Camel", Slug: "camel", Summary: "camelCase"}, {Name: "Truncate", Slug: "truncate", Summary: "_.truncate"}},
	"set":       {{Name: "New", Slug: "set-new", Summary: "new Set()"}, {Name: "Union", Slug: "set-union", Summary: "Set union"}, {Name: "Intersection", Slug: "set-intersection", Summary: "Set intersection"}},
	"buffer":    {{Name: "From", Slug: "buffer-from", Summary: "Buffer.from"}, {Name: "Concat", Slug: "buffer-concat", Summary: "Buffer.concat"}, {Name: "Equals", Slug: "equals", Summary: "Compare bytes"}},
	"async":     {{Name: "All", Slug: "all", Summary: "Promise.all"}, {Name: "Sleep", Slug: "sleep", Summary: "setTimeout"}, {Name: "Retry", Slug: "retry", Summary: "Retry with backoff"}},
	"path":      {{Name: "Join", Slug: "join", Summary: "path.join"}, {Name: "Basename", Slug: "basename", Summary: "path.basename"}, {Name: "Extname", Slug: "extname", Summary: "path.extname"}},
	"url":       {{Name: "Parse", Slug: "url-parse", Summary: "new URL()"}, {Name: "ParseQuery", Slug: "parse-query", Summary: "querystring.parse"}, {Name: "EncodeQuery", Slug: "encode-query", Summary: "querystring.stringify"}},
	"env":       {{Name: "Load", Slug: "load", Summary: "dotenv.config"}, {Name: "Get", Slug: "get", Summary: "process.env"}, {Name: "GetInt", Slug: "get-int", Summary: "Typed int env"}},
	"fs":        {{Name: "ReadFile", Slug: "read-file", Summary: "fs.readFile"}, {Name: "WriteFile", Slug: "write-file", Summary: "fs.writeFile"}, {Name: "Exists", Slug: "exists", Summary: "fs.access"}},
	"err":       {{Name: "NotFound", Slug: "app-error-not-found", Summary: "404 error"}, {Name: "BadRequest", Slug: "app-error-bad-request", Summary: "400 error"}, {Name: "Unauthorized", Slug: "app-error-unauthorized", Summary: "401 error"}},
	"exec":      {{Name: "Command", Slug: "cmd-command", Summary: "spawn"}, {Name: "Exec", Slug: "exec", Summary: "execFile"}},
	"osutil":    {{Name: "Hostname", Slug: "hostname", Summary: "os.hostname"}, {Name: "Homedir", Slug: "homedir", Summary: "os.homedir"}, {Name: "TempDir", Slug: "temp-dir", Summary: "os.tmpdir"}},
	"stream":    {{Name: "Pipe", Slug: "pipe", Summary: "stream.pipe"}, {Name: "ReadAll", Slug: "read-all", Summary: "stream consumers"}, {Name: "Transform", Slug: "transform", Summary: "Transform stream"}},
	"compress":  {{Name: "Gzip", Slug: "gzip", Summary: "zlib.gzip"}, {Name: "Gunzip", Slug: "gunzip", Summary: "zlib.gunzip"}},
	"csv":       {{Name: "Read", Slug: "read", Summary: "Parse CSV file"}, {Name: "Write", Slug: "write", Summary: "Write CSV"}, {Name: "ParseString", Slug: "parse-string", Summary: "Parse inline CSV"}},
	"archive":   {{Name: "ZipCreate", Slug: "zip-create", Summary: "Create zip"}, {Name: "ZipExtract", Slug: "zip-extract", Summary: "Extract zip"}, {Name: "TarCreate", Slug: "tar-create", Summary: "Create tar"}},
	"http":      {{Name: "New", Slug: "app-new", Summary: "express()"}, {Name: "Ctx.JSON", Slug: "ctx", Summary: "res.json()"}, {Name: "CORS", Slug: "middleware-cors", Summary: "cors middleware"}},
	"client":    {{Name: "Get", Slug: "client", Summary: "axios.get"}, {Name: "Fetch", Slug: "response-fetch", Summary: "fetch()"}, {Name: "New", Slug: "client-new", Summary: "axios.create"}},
	"log":       {{Name: "New", Slug: "logger-new", Summary: "winston.createLogger"}, {Name: "Default", Slug: "logger-default", Summary: "Default logger"}},
	"ws":        {{Name: "NewServer", Slug: "server-new-server", Summary: "WebSocket.Server"}, {Name: "Dial", Slug: "conn-dial", Summary: "new WebSocket()"}, {Name: "Upgrade", Slug: "conn-upgrade", Summary: "Upgrade handler"}},
	"json":      {{Name: "Parse", Slug: "parse", Summary: "JSON.parse"}, {Name: "Stringify", Slug: "stringify", Summary: "JSON.stringify"}, {Name: "Pretty", Slug: "pretty", Summary: "Pretty print"}},
	"validate":  {{Name: "Validate", Slug: "validate", Summary: "schema.validate"}, {Name: "Object", Slug: "schema-object", Summary: "z.object"}, {Name: "String", Slug: "string", Summary: "z.string"}},
	"db":        {{Name: "Open", Slug: "db-open", Summary: "knex()"}, {Name: "From", Slug: "query", Summary: "knex('table')"}, {Name: "Transaction", Slug: "tx", Summary: "knex.transaction"}},
	"redis":     {{Name: "New", Slug: "client-new", Summary: "new Redis()"}, {Name: "Get", Slug: "client", Summary: "redis.get"}, {Name: "Set", Slug: "client", Summary: "redis.set"}},
	"mongo":     {{Name: "Connect", Slug: "client-connect", Summary: "mongoose.connect"}, {Name: "Collection", Slug: "collection", Summary: "Model collection"}, {Name: "Eq", Slug: "eq", Summary: "Filter helper"}},
	"crypto":    {{Name: "SHA256", Slug: "sha256", Summary: "createHash"}, {Name: "HashPassword", Slug: "hash-password", Summary: "bcrypt.hash"}, {Name: "RandomString", Slug: "random-string", Summary: "crypto.randomBytes"}},
	"jwt":       {{Name: "Sign", Slug: "sign", Summary: "jwt.sign"}, {Name: "Verify", Slug: "verify", Summary: "jwt.verify"}, {Name: "Decode", Slug: "decode", Summary: "jwt.decode"}},
	"id":        {{Name: "NewUUID", Slug: "new-uuid", Summary: "uuid.v4"}, {Name: "NewNanoID", Slug: "new-nano-id", Summary: "nanoid"}, {Name: "ParseUUID", Slug: "parse-uuid", Summary: "Parse UUID string"}},
	"auth":      {{Name: "Bearer", Slug: "bearer", Summary: "JWT Bearer auth"}, {Name: "APIKey", Slug: "api-key", Summary: "API key header"}, {Name: "Basic", Slug: "basic", Summary: "Basic auth"}},
	"cache":     {{Name: "New", Slug: "cache-new", Summary: "new LRUCache"}, {Name: "Get", Slug: "cache", Summary: "cache.get"}, {Name: "Set", Slug: "cache", Summary: "cache.set"}},
	"event":     {{Name: "New", Slug: "emitter-new", Summary: "EventEmitter"}, {Name: "On", Slug: "emitter", Summary: "emitter.on"}, {Name: "Emit", Slug: "emitter", Summary: "emitter.emit"}},
	"cron":      {{Name: "New", Slug: "scheduler-new", Summary: "node-cron"}, {Name: "Schedule", Slug: "scheduler", Summary: "cron.schedule"}, {Name: "Parse", Slug: "parse", Summary: "Parse expression"}},
	"queue":     {{Name: "New", Slug: "client-new", Summary: "Bull queue"}, {Name: "Enqueue", Slug: "client", Summary: "queue.add"}, {Name: "Worker", Slug: "worker-new-worker", Summary: "Worker process"}},
	"mail":      {{Name: "Send", Slug: "send", Summary: "nodemailer.send"}, {Name: "BuildMIME", Slug: "build-mime", Summary: "Build MIME"}, {Name: "Message", Slug: "message", Summary: "Email message"}},
	"time":      {{Name: "Now", Slug: "now", Summary: "moment()"}, {Name: "Format", Slug: "format", Summary: "moment.format"}, {Name: "Parse", Slug: "parse", Summary: "moment.parse"}},
	"semver":    {{Name: "Parse", Slug: "version-parse", Summary: "semver.parse"}, {Name: "Compare", Slug: "compare", Summary: "semver.compare"}, {Name: "Satisfies", Slug: "satisfies", Summary: "semver.satisfies"}},
}

var packageQuickStart = map[string]string{
	"cond":     "label := cond.If(score >= 60, \"pass\", \"fail\")",
	"slice":    "evens := slice.Filter(nums, func(n int) bool { return n%2 == 0 })",
	"maputil":  "picked := maputil.Pick(cfg, \"host\", \"port\")",
	"str":      "slug := str.Slug(\"Hello World!\")",
	"set":      "u := set.Union(tagsA, tagsB)",
	"buffer":   "b := buffer.FromString(\"hello\")",
	"async":    "a, b, err := async.All(ctx, fetchA, fetchB)",
	"path":     "p := path.Join(\"var\", \"log\", \"app.log\")",
	"url":      "u, err := url.Parse(\"https://example.com/api?page=1\")",
	"env":      "port := env.Get(\"PORT\", \"8080\")",
	"fs":       "data, err := fs.ReadFile(ctx, \"config.json\")",
	"err":      "return err.NotFound(\"user not found\")",
	"exec":     "out, err := exec.Command(ctx, \"git\", \"status\").Output()",
	"osutil":   "home, err := osutil.Homedir()",
	"stream":   "n, err := stream.ReadAll(reader)",
	"compress": "gz, err := compress.Gzip(data)",
	"csv":      "rows, err := csv.ReadAll(ctx, \"data.csv\")",
	"archive":  "err := archive.ZipCreate(\"out.zip\", entries)",
	"http":     "app := goxhttp.New()\napp.Get(\"/health\", func(c *goxhttp.Ctx) error {\n    return c.JSON(200, map[string]string{\"status\": \"ok\"})\n})",
	"client":   "res, err := client.Get(ctx, \"https://api.example.com/users\")",
	"log":      "logger := log.New()\nlogger.Info(\"server started\", \"port\", 8080)",
	"ws":       "conn, err := ws.Dial(ctx, \"wss://example.com/ws\", nil)",
	"json":     "obj, err := json.Parse[Config](jsonStr)",
	"validate": "if err := validate.Validate(&user); err != nil { /* return 400 */ }",
	"db":       "row, err := db.From(\"users\").WhereEq(\"id\", id).First(ctx, &user)",
	"redis":    "rdb, err := redis.New(\"localhost:6379\")",
	"mongo":    "client, err := mongo.Connect(ctx, uri)",
	"crypto":   "hash, err := crypto.HashPassword(password)",
	"jwt":      "token, err := jwt.Sign(claims, secret, jwt.SignOptions{ExpiresIn: time.Hour})",
	"id":       "uid := id.NewUUID()",
	"auth":     "app.Use(auth.Bearer(secret))",
	"cache":    "c := cache.New(1000)",
	"event":    "em := event.New()\nem.On(\"ready\", func() { /* ... */ })",
	"cron":     "sched := cron.New()\nsched.Schedule(\"0 * * * *\", cleanup)",
	"queue":    "client, err := queue.New(\"localhost:6379\")",
	"mail":     "err := mail.Send(ctx, cfg, msg)",
	"time":     "now := timex.Now()\nformatted := timex.Format(now, timex.LayoutISO)",
	"semver":   "v, err := semver.Parse(\"1.2.3\")",
}

type npmGoxRow struct {
	Node string
	Gox  string
	Sym  string
}

func loadNpmGoxRows(path string) map[string][]npmGoxRow {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	rowRe := regexp.MustCompile(`^\|\s*` + "`" + `([^` + "`" + `]+)` + "`" + `\s*\|\s*` + "`" + `([^` + "`" + `]+)` + "`" + `\s*\|`)
	var currentPkg string
	out := make(map[string][]npmGoxRow)
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
		out[currentPkg] = append(out[currentPkg], npmGoxRow{Node: nodeExpr, Gox: goxExpr, Sym: extractSymbol(goxExpr)})
	}
	return out
}

func writePackageOverviewComponent(pkgName, analog, importPath, subtitle string, symCount int) string {
	return fmt.Sprintf(`<PackageOverview
  name="%s"
  node="%s"
  import-path="%s"
  subtitle="%s"
  :symbol-count=%d
/>
`, escapeAttr(pkgName), escapeAttr(analog), escapeAttr(importPath), escapeAttr(subtitle), symCount)
}

func writeFeaturedSection(pkgName string) string {
	feat := featuredSymbols[pkgName]
	if len(feat) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("## Start here\n\n")
	b.WriteString(`<div class="featured-grid">` + "\n")
	for _, f := range feat {
		b.WriteString(fmt.Sprintf(`<a class="featured-card" href="/packages/%s/%s"><div class="featured-name">%s</div><div class="featured-summary">%s</div></a>`+"\n",
			pkgName, f.Slug, escapeHTML(f.Name), escapeHTML(f.Summary)))
	}
	b.WriteString("</div>\n\n")
	return b.String()
}

func writeNpmGoxSection(pkgName string, rows []npmGoxRow) string {
	if len(rows) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("## npm → gox\n\n")
	b.WriteString("Quick mapping from Node.js patterns to gox APIs:\n\n")
	b.WriteString("<table class=\"npm-map-table\"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>\n")
	limit := len(rows)
	if limit > 12 {
		limit = 12
	}
	for _, r := range rows[:limit] {
		link := ""
		if r.Sym != "" {
			slug := slugify(r.Sym)
			link = fmt.Sprintf("<tr><td><code>%s</code></td><td><a href=\"/packages/%s/%s\"><code>%s</code></a></td></tr>\n",
				escapeHTML(r.Node), pkgName, slug, escapeHTML(r.Gox))
		} else {
			link = fmt.Sprintf("<tr><td><code>%s</code></td><td><code>%s</code></td></tr>\n",
				escapeHTML(r.Node), escapeHTML(r.Gox))
		}
		b.WriteString(link)
	}
	b.WriteString("</tbody></table>\n\n")
	if len(rows) > limit {
		b.WriteString(fmt.Sprintf("::: info\nShowing %d of %d mappings. Browse the symbol table below for the full API.\n:::\n\n", limit, len(rows)))
	}
	return b.String()
}

func writeUseCasesSection(pkgName string) string {
	narr, ok := packageNarratives[pkgName]
	if !ok || len(narr.UseCases) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("## Common use cases\n\n")
	for _, u := range narr.UseCases {
		b.WriteString("- " + u + "\n")
	}
	b.WriteString("\n")
	return b.String()
}
