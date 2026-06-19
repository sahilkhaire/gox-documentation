package main

import (
	"fmt"
	"regexp"
	"strings"
)

// resolveGoxExample returns a concrete, runnable gox usage snippet (no placeholders).
func resolveGoxExample(sym symbolDoc, pkgIdent string, e enrichment, hasEnrich bool) string {
	if hasEnrich {
		if ex := strings.TrimSpace(e.Example); ex != "" {
			return ex
		}
		if ex := strings.TrimSpace(e.Gox); ex != "" {
			return ex
		}
	}
	key := symbolKey(sym)
	if ex, ok := curatedExamples[key]; ok {
		return ex
	}
	if ex := generateExampleFromSignature(sym, pkgIdent); ex != "" {
		return ex
	}
	return ""
}

func symbolKey(sym symbolDoc) string {
	if sym.Receiver != "" {
		return sym.Pkg + "." + sym.Receiver + "." + sym.Name
	}
	return sym.Pkg + "." + sym.Name
}

func generateExampleFromSignature(sym symbolDoc, pkgIdent string) string {
	switch sym.Kind {
	case "const", "var":
		return constVarExample(sym, pkgIdent)
	case "type":
		return typeUsageExample(sym, pkgIdent)
	case "method":
		return methodUsageExample(sym, pkgIdent)
	case "func":
		return funcUsageExample(sym, pkgIdent)
	}
	return ""
}

func funcUsageExample(sym symbolDoc, pkgIdent string) string {
	_, params, _ := parseFuncSignature(sym.Signature)
	args := sampleArgs(params, pkgIdent, sym.Pkg, sym.Name)
	call := fmt.Sprintf("%s.%s(%s)", pkgIdent, sym.Name, args)
	if strings.Contains(sym.Signature, " error") && !strings.Contains(sym.Signature, "(error)") {
		return fmt.Sprintf("result, err := %s\nif err != nil {\n    return err\n}", call)
	}
	if strings.Contains(sym.Signature, ") (") {
		parts := strings.Split(call, ".")
		if len(parts) >= 2 {
			return fmt.Sprintf("result, err := %s", call)
		}
	}
	return call
}

func methodUsageExample(sym symbolDoc, pkgIdent string) string {
	key := symbolKey(sym)
	if ex, ok := curatedExamples[key]; ok {
		return ex
	}
	setup := methodReceiverSetup(sym, pkgIdent)
	_, params, receiver := parseFuncSignature(sym.Signature)
	_ = receiver
	recv := receiverVar(sym.Receiver)
	args := sampleArgs(params, pkgIdent, sym.Pkg, sym.Name)
	call := fmt.Sprintf("%s.%s(%s)", recv, sym.Name, args)
	if setup != "" {
		if strings.Contains(sym.Signature, " error") {
			return setup + "\n" + fmt.Sprintf("if err := %s; err != nil {\n    return err\n}", call)
		}
		return setup + "\n" + call
	}
	if strings.Contains(sym.Signature, " error") {
		return fmt.Sprintf("if err := %s; err != nil {\n    return err\n}", call)
	}
	return call
}

func receiverVar(typeName string) string {
	typeName = strings.TrimPrefix(typeName, "*")
	if typeName == "" {
		return "v"
	}
	r := []rune(typeName)
	r[0] = rune(strings.ToLower(string(r[0]))[0])
	return string(r)
}

func methodReceiverSetup(sym symbolDoc, pkgIdent string) string {
	key := sym.Pkg + "." + sym.Receiver
	if setup, ok := receiverSetup[key]; ok {
		return setup
	}
	recv := receiverVar(sym.Receiver)
	switch sym.Pkg {
	case "set":
		return fmt.Sprintf("%s := %s.New(\"alpha\", \"beta\")", recv, pkgIdent)
	case "buffer":
		return fmt.Sprintf("%s := %s.FromString(\"hello\")", recv, pkgIdent)
	case "event":
		return fmt.Sprintf("%s := %s.NewEmitter()", recv, pkgIdent)
	case "cache":
		return fmt.Sprintf("%s := %s.New(100)", recv, pkgIdent)
	case "cron":
		return fmt.Sprintf("%s := %s.NewScheduler()", recv, pkgIdent)
	case "queue":
		if sym.Receiver == "Worker" {
			return fmt.Sprintf("%s := %s.NewWorker(\"localhost:6379\", mux)", pkgIdent)
		}
		if sym.Receiver == "ServeMux" {
			return fmt.Sprintf("mux := %s.NewServeMux()", pkgIdent)
		}
		if sym.Receiver == "Client" {
			return fmt.Sprintf("%s, err := %s.New(\"localhost:6379\")\nif err != nil {\n    return err\n}", recv, pkgIdent)
		}
	case "redis":
		return fmt.Sprintf("%s, err := %s.New(\"localhost:6379\")\nif err != nil {\n    return err\n}", recv, pkgIdent)
	case "db":
		return fmt.Sprintf("%s, err := %s.Open(ctx, \"postgres\", dsn)\nif err != nil {\n    return err\n}", recv, pkgIdent)
	case "mongo":
		if sym.Receiver == "Client" {
			return fmt.Sprintf("%s, err := %s.Connect(ctx, uri)\nif err != nil {\n    return err\n}", recv, pkgIdent)
		}
		if sym.Receiver == "Database" {
			return fmt.Sprintf("%s := client.DB(\"app\")", recv)
		}
		if sym.Receiver == "Collection" {
			return fmt.Sprintf("%s := db.Collection(\"users\")", recv)
		}
	case "http":
		switch sym.Receiver {
		case "App":
			return fmt.Sprintf("%s := %s.New()", recv, pkgIdent)
		case "Ctx":
			return "func handler(c *http.Ctx) error {"
		case "EventStream":
			return "stream := http.NewSSE(w, r)"
		case "MemoryStore":
			return fmt.Sprintf("store := %s.NewMemoryStore()", pkgIdent)
		case "MultipartForm":
			return "form, err := http.ParseMultipart(c)\nif err != nil {\n    return err\n}"
		}
	case "client":
		return fmt.Sprintf("%s := %s.New()", recv, pkgIdent)
	case "log":
		return fmt.Sprintf("%s := %s.New()", recv, pkgIdent)
	case "ws":
		if sym.Receiver == "Conn" {
			return "// conn from ws upgrade or Dial"
		}
		if sym.Receiver == "Server" {
			return fmt.Sprintf("%s := %s.NewServer()", recv, pkgIdent)
		}
	case "validate":
		if sym.Receiver == "Schema" {
			return "sch := validate.Object(map[string]validate.Field{\n    \"email\": validate.String().Email(),\n})"
		}
		if sym.Receiver == "Field" {
			return "field := validate.String().Min(1).Max(100)"
		}
	case "exec":
		return fmt.Sprintf("%s := %s.Command(ctx, \"echo\", \"hello\")", recv, pkgIdent)
	case "err":
		if sym.Receiver == "AppError" {
			return fmt.Sprintf("%s := %s.NotFound(\"resource missing\")", recv, pkgIdent)
		}
	case "compress":
		if sym.Receiver == "GzipWriter" {
			return "var buf bytes.Buffer\nw := compress.NewGzipWriter(&buf)"
		}
		if sym.Receiver == "GzipReader" {
			return "r, err := compress.NewGzipReader(bytes.NewReader(data))"
		}
	}
	return ""
}

func constVarExample(sym symbolDoc, pkgIdent string) string {
	key := symbolKey(sym)
	if ex, ok := curatedExamples[key]; ok {
		return ex
	}
	switch sym.Pkg {
	case "time":
		return fmt.Sprintf("formatted := %s.Format(t, %s.%s)", pkgIdent, pkgIdent, sym.Name)
	case "ws":
		if sym.Name == "DefaultUpgrader" {
			return "conn, err := ws.DefaultUpgrader.Upgrade(w, r, nil)"
		}
	}
	return fmt.Sprintf("%s.%s", pkgIdent, sym.Name)
}

func typeUsageExample(sym symbolDoc, pkgIdent string) string {
	key := symbolKey(sym)
	if ex, ok := curatedExamples[key]; ok {
		return ex
	}
	switch sym.Pkg + "." + sym.Name {
	case "http.App":
		return fmt.Sprintf("app := %s.New()\napp.Get(\"/\", func(c *http.Ctx) error {\n    return c.JSON(200, map[string]string{\"ok\": \"true\"})\n})", pkgIdent)
	case "http.Ctx":
		return "func handler(c *http.Ctx) error {\n    return c.JSON(200, map[string]string{\"hello\": \"world\"})\n}"
	case "http.Handler":
		return fmt.Sprintf("var _ %s.Handler = func(c *http.Ctx) error { return nil }", pkgIdent)
	case "http.Middleware":
		return fmt.Sprintf("app.Use(%s.Logger(), %s.Recover())", pkgIdent, pkgIdent)
	case "client.Client":
		return fmt.Sprintf("c := %s.New()\nresp, err := c.Get(ctx, \"https://api.example.com/users\", nil)", pkgIdent)
	case "redis.Client":
		return fmt.Sprintf("rdb, err := %s.New(\"localhost:6379\")\nif err != nil {\n    return err\n}\nval, err := rdb.Get(ctx, \"session:1\")", pkgIdent)
	case "db.DB":
		return fmt.Sprintf("db, err := %s.Open(ctx, \"postgres\", dsn)\nif err != nil {\n    return err\n}", pkgIdent)
	case "mongo.Client":
		return fmt.Sprintf("client, err := %s.Connect(ctx, \"mongodb://localhost:27017\")\nif err != nil {\n    return err\n}", pkgIdent)
	case "event.Emitter":
		return fmt.Sprintf("emitter := %s.NewEmitter()\nemitter.On(\"ready\", func(args ...any) {\n    fmt.Println(\"ready\", args)\n})", pkgIdent)
	case "cache.Cache":
		return fmt.Sprintf("c := %s.New(500)\nc.Set(\"user:1\", user, time.Minute)", pkgIdent)
	case "cron.Scheduler":
		return fmt.Sprintf("sched := %s.NewScheduler()\nsched.Every(time.Hour, func() { fmt.Println(\"tick\") })", pkgIdent)
	case "queue.Client":
		return fmt.Sprintf("q, err := %s.New(\"localhost:6379\")\nif err != nil {\n    return err\n}", pkgIdent)
	case "mail.Message":
		return fmt.Sprintf("msg := %s.Message{\n    From: \"noreply@example.com\",\n    To: []string{\"user@example.com\"},\n    Subject: \"Welcome\",\n    Body: \"Hello!\",\n}", pkgIdent)
	case "mail.SMTPConfig":
		return fmt.Sprintf("cfg := %s.SMTPConfig{\n    Host: \"smtp.example.com\",\n    Port: 587,\n    Username: \"user\",\n    Password: \"pass\",\n}", pkgIdent)
	case "semver.Version":
		return fmt.Sprintf("v, err := %s.Parse(\"1.2.3\")\nif err != nil {\n    return err\n}", pkgIdent)
	case "set.Set":
		return fmt.Sprintf("s := %s.New(\"go\", \"node\", \"rust\")\n_ = s.Has(\"go\")", pkgIdent)
	case "buffer.Buffer":
		return fmt.Sprintf("b := %s.FromString(\"hello world\")\n_ = b.ToString()", pkgIdent)
	case "validate.Schema":
		return "sch := validate.Object(map[string]validate.Field{\n    \"name\": validate.String().Required(),\n})"
	case "validate.Field":
		return "field := validate.String().Email().Required()"
	case "validate.ValidationError":
		return "if err := validate.Validate(v); err != nil {\n    var verr validate.ValidationError\n    errors.As(err, &verr)\n}"
	case "jwt.SignOptions":
		return fmt.Sprintf("opts := &%s.SignOptions{ExpiresIn: time.Hour}", pkgIdent)
	case "auth.BearerOptions":
		return fmt.Sprintf("opts := &%s.BearerOptions{}", pkgIdent)
	case "auth.APIKeyOpts":
		return fmt.Sprintf("opts := %s.APIKeyOpts{Keys: map[string]bool{\"secret-key\": true}}", pkgIdent)
	case "url.URL":
		return fmt.Sprintf("u, err := %s.Parse(\"https://example.com/path?q=1\")", pkgIdent)
	case "async.RetryConfig":
		return fmt.Sprintf("cfg := %s.RetryConfig{MaxAttempts: 3, Delay: time.Second}", pkgIdent)
	default:
		if ex := defaultTypeExample(sym, pkgIdent); ex != "" {
			return ex
		}
		return fmt.Sprintf("// See methods on %s in the sidebar.", sym.Name)
	}
}

func defaultTypeExample(sym symbolDoc, pkgIdent string) string {
	switch sym.Pkg + "." + sym.Name {
	case "queue.Worker":
		return "mux := queue.NewServeMux()\nmux.HandleFunc(\"email\", handleEmail)\nworker := queue.NewWorker(\"localhost:6379\", mux)\nworker.Run()"
	case "queue.ServeMux":
		return "mux := queue.NewServeMux()\nmux.HandleFunc(\"email\", func(ctx context.Context, payload []byte) error { return nil })"
	case "queue.Handler":
		return "var _ queue.Handler = func(ctx context.Context, payload []byte) error { return nil }"
	case "queue.EnqueueOpts":
		return "opts := &queue.EnqueueOpts{MaxRetry: 3, Queue: \"default\"}"
	case "cron.JobID":
		return "var id cron.JobID = \"nightly-backup\""
	case "mongo.Database":
		return "db := client.DB(\"app\")"
	case "mongo.Collection":
		return "coll := db.Collection(\"users\")"
	case "redis.Message":
		return "msg := <-subChan // redis.Message with Channel and Payload fields"
	case "db.Tx":
		return "err := db.Transaction(ctx, func(tx *db.Tx) error {\n    return tx.Insert(ctx, \"users\", row)\n})"
	case "db.Query":
		return "q := db.From(\"users\").Select(\"id\", \"email\").WhereEq(\"active\", true)"
	case "ws.WSHandler":
		return "var _ ws.WSHandler = func(conn *ws.Conn) error { return conn.WriteJSON(map[string]string{\"ok\": \"1\"}) }"
	case "ws.Upgrader":
		return "upgrader := ws.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}"
	case "ws.Server":
		return "srv := ws.NewServer()\nsrv.Handle(\"/ws\", handler)"
	case "ws.Conn":
		return "if err := conn.WriteJSON(map[string]string{\"event\": \"ping\"}); err != nil {\n    return err\n}"
	case "log.Logger":
		return "logger := log.New()\nlogger.Info(\"server started\", \"port\", 8080)"
	case "client.Response":
		return "if resp.StatusCode() != 200 {\n    return fmt.Errorf(\"unexpected status %d\", resp.StatusCode())\n}\nvar body User\nif err := resp.JSON(&body); err != nil {\n    return err\n}"
	case "client.RequestOpts":
		return "opts := &client.RequestOpts{\n    Method: \"GET\",\n    Query: map[string]string{\"page\": \"1\"},\n    Headers: map[string]string{\"Accept\": \"application/json\"},\n}"
	case "http.WSHandler":
		return "app.HandleWS(\"/ws\", func(conn *ws.Conn) error { return nil })"
	case "http.SessionStore":
		return "var store http.SessionStore = http.NewMemoryStore()"
	case "http.SessionOptions":
		return "opts := http.SessionOptions{MaxAge: time.Hour, Secure: true, HttpOnly: true}"
	case "http.RateLimitOptions":
		return "opts := http.RateLimitOptions{Max: 100, Window: time.Minute}"
	case "http.MultipartForm":
		return "form, err := http.ParseMultipart(c)\nif err != nil {\n    return err\n}\nfile := form.File(\"avatar\")"
	case "http.MemoryStore":
		return "store := http.NewMemoryStore()\nsession := store.Get(sessionID)"
	case "http.EventStream":
		return "stream, err := http.SSE(w, r)\nif err != nil {\n    return err\n}\nstream.Send(\"message\", \"hello\")"
	case "archive.FileEntry":
		return "entry := archive.FileEntry{Name: \"readme.txt\", Mode: 0644, Body: []byte(\"hello\")}"
	case "osutil.User":
		return "info, err := osutil.UserInfo()\nif err != nil {\n    return err\n}\n_ = info.Username"
	case "exec.Cmd":
		return "cmd := exec.Command(ctx, \"git\", \"status\")\nout, err := cmd.Output()"
	}
	return ""
}

func parseFuncSignature(sig string) (name, params, receiver string) {
	sig = strings.TrimSpace(sig)
	if !strings.HasPrefix(sig, "func ") {
		return "", "", ""
	}
	rest := sig[5:]
	rest = strings.TrimSpace(skipTypeParams(rest))

	if strings.HasPrefix(rest, "(") {
		end := matchingParen(rest, 0)
		if end < 0 {
			return "", "", ""
		}
		receiver = strings.TrimSpace(rest[1:end])
		rest = strings.TrimSpace(rest[end+1:])
	}

	paren := strings.Index(rest, "(")
	if paren < 0 {
		return strings.TrimSpace(rest), "", receiver
	}
	name = strings.TrimSpace(rest[:paren])
	end := matchingParen(rest, paren)
	if end < 0 {
		return name, "", receiver
	}
	params = rest[paren+1 : end]
	return name, params, receiver
}

func skipTypeParams(s string) string {
	if !strings.HasPrefix(s, "[") {
		return s
	}
	end := strings.Index(s, "]")
	if end < 0 {
		return s
	}
	return strings.TrimSpace(s[end+1:])
}

func matchingParen(s string, open int) int {
	if open < 0 || open >= len(s) || s[open] != '(' {
		return -1
	}
	depth := 0
	for i := open; i < len(s); i++ {
		switch s[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}

func splitParams(params string) []string {
	params = strings.TrimSpace(params)
	if params == "" {
		return nil
	}
	var parts []string
	depth := 0
	start := 0
	for i := 0; i < len(params); i++ {
		switch params[i] {
		case '(', '[', '{':
			depth++
		case ')', ']', '}':
			depth--
		case ',':
			if depth == 0 {
				parts = append(parts, strings.TrimSpace(params[start:i]))
				start = i + 1
			}
		}
	}
	parts = append(parts, strings.TrimSpace(params[start:]))
	return parts
}

func parseParam(p string) (name, typ string) {
	p = strings.TrimSpace(p)
	if p == "" {
		return "", ""
	}
	if strings.HasPrefix(p, "...") {
		return "", strings.TrimPrefix(p, "...")
	}
	// "fn func(T) U" or "ctx context.Context"
	fields := strings.Fields(p)
	if len(fields) == 1 {
		return "", fields[0]
	}
	// last token(s) form type; handle "map[string]any"
	lastIdent := len(fields) - 1
	for lastIdent > 0 {
		candidate := strings.Join(fields[lastIdent:], " ")
		if strings.ContainsAny(candidate, "[]*(") || isTypeKeyword(fields[lastIdent]) {
			return strings.Join(fields[:lastIdent], " "), candidate
		}
		lastIdent--
	}
	return fields[0], strings.Join(fields[1:], " ")
}

func isTypeKeyword(w string) bool {
	switch w {
	case "string", "int", "bool", "byte", "error", "any", "float64", "time.Duration", "time.Time":
		return true
	}
	return strings.Contains(w, ".") || strings.HasPrefix(w, "[]") || strings.HasPrefix(w, "map[")
}

func sampleArgs(params, pkgIdent, pkg, funcName string) string {
	parts := splitParams(params)
	if len(parts) == 0 {
		return ""
	}
	var args []string
	for _, p := range parts {
		name, typ := parseParam(p)
		args = append(args, sampleValue(name, typ, pkgIdent, pkg, funcName))
	}
	return strings.Join(args, ", ")
}

func sampleValue(name, typ, pkgIdent, pkg, funcName string) string {
	typ = strings.TrimSpace(typ)
	name = strings.TrimSpace(name)

	if strings.HasPrefix(typ, "...") {
		typ = strings.TrimPrefix(typ, "...")
		inner := sampleValue(name, typ, pkgIdent, pkg, funcName)
		return inner
	}
	if strings.Contains(typ, "context.Context") {
		return "ctx"
	}
	if strings.Contains(typ, "http.ResponseWriter") {
		return "w"
	}
	if strings.Contains(typ, "*http.Request") || strings.Contains(typ, "http.Request") {
		return "r"
	}
	if strings.Contains(typ, "io.Writer") {
		return "w"
	}
	if strings.Contains(typ, "io.Reader") {
		return "strings.NewReader(\"data\")"
	}
	if strings.Contains(typ, "stdtime.Time") || typ == "time.Time" {
		if name == "a" || name == "b" {
			return name
		}
		return "time.Now()"
	}
	if strings.Contains(typ, "stdtime.Duration") || typ == "time.Duration" {
		return "time.Hour"
	}
	if strings.Contains(typ, "[]byte") {
		return "[]byte(\"hello\")"
	}
	if strings.Contains(typ, "string") {
		return sampleString(name, pkg, funcName, pkgIdent)
	}
	if strings.Contains(typ, "int") {
		return sampleInt(name, pkg, funcName)
	}
	if strings.Contains(typ, "bool") {
		return "true"
	}
	if strings.Contains(typ, "float") {
		return "3.14"
	}
	if strings.Contains(typ, "map[") {
		return "map[string]any{\"key\": \"value\"}"
	}
	if strings.Contains(typ, "[]") {
		return sampleSlice(pkg, funcName)
	}
	if strings.Contains(typ, "func(") {
		return sampleFunc(typ, pkg)
	}
	if strings.Contains(typ, "*") || strings.Contains(typ, "Opts") || strings.Contains(typ, "Options") || strings.Contains(typ, "Config") {
		return sampleStruct(pkg, funcName, typ, pkgIdent)
	}
	if name != "" {
		return name
	}
	return "value"
}

func sampleString(name, pkg, funcName, pkgIdent string) string {
	n := strings.ToLower(name)
	switch {
	case n == "layout" && pkg == "time":
		return pkgIdent + ".LayoutISO"
	case n == "value" && pkg == "time":
		return `"2024-06-15T12:00:00Z"`
	case strings.Contains(n, "path"), strings.Contains(n, "file"):
		return `"config.json"`
	case strings.Contains(n, "url"):
		return `"https://api.example.com/users"`
	case strings.Contains(n, "key"):
		return `"session:1"`
	case strings.Contains(n, "name"):
		return `"example"`
	case strings.Contains(n, "table"):
		return `"users"`
	case strings.Contains(n, "query"):
		return `"SELECT * FROM users WHERE id = $1"`
	case strings.Contains(n, "cmd"), strings.Contains(n, "command"):
		return `"echo hello"`
	case strings.Contains(n, "spec"), strings.Contains(n, "expr"):
		return `"0 * * * *"`
	case strings.Contains(n, "token"):
		return `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	case strings.Contains(n, "secret"):
		return `"change-me"`
	case strings.Contains(n, "header"):
		return `"Authorization"`
	case strings.Contains(n, "field"):
		return `"email"`
	case strings.Contains(n, "event"):
		return `"user.created"`
	case strings.Contains(n, "host"):
		return `"localhost:6379"`
	case strings.Contains(n, "uri"), strings.Contains(n, "dsn"):
		return `"postgres://user:pass@localhost/db"`
	case strings.Contains(n, "version"), strings.Contains(n, "constraint"), strings.Contains(n, "range"):
		return `"1.2.3"`
	case strings.Contains(n, "part"):
		return `"minor"`
	case strings.Contains(n, "suffix"):
		return `"..."`
	case strings.Contains(n, "prefix"):
		return `"/api"`
	case strings.Contains(n, "method"):
		return `"GET"`
	case strings.Contains(n, "id"):
		return `"550e8400-e29b-41d4-a716-446655440000"`
	case strings.Contains(n, "password"):
		return `"correct-horse-battery-staple"`
	case strings.Contains(n, "hash"):
		return `"$2a$10$..."`
	case strings.Contains(n, "subject"):
		return `"Welcome"`
	case strings.Contains(n, "body"), strings.Contains(n, "html"):
		return `"Hello from gox"`
	case strings.Contains(n, "from"):
		return `"noreply@example.com"`
	case strings.Contains(n, "s"):
		return `"hello world"`
	case n == "t":
		return `"2024-06-15"`
	}
	return `"example"`
}

func sampleInt(name, pkg, funcName string) string {
	n := strings.ToLower(name)
	switch {
	case strings.Contains(n, "port"):
		return "8080"
	case strings.Contains(n, "code"), strings.Contains(n, "status"):
		return "200"
	case strings.Contains(n, "max"), strings.Contains(n, "limit"), strings.Contains(n, "size"):
		return "100"
	case strings.Contains(n, "n"), strings.Contains(n, "count"):
		return "16"
	case strings.Contains(n, "width"):
		return "10"
	}
	return "42"
}

func sampleSlice(pkg, funcName string) string {
	switch pkg {
	case "slice":
		return "[]int{1, 2, 3}"
	case "set":
		return "[]string{\"a\", \"b\"}"
	case "csv":
		return "[][]string{{\"name\", \"age\"}, {\"Ada\", \"36\"}}"
	default:
		return "[]string{\"a\", \"b\", \"c\"}"
	}
}

func sampleFunc(typ, pkg string) string {
	switch pkg {
	case "async":
		if strings.Contains(typ, "error") && !strings.Contains(typ, "any") {
			return "func(ctx context.Context) error { return nil }"
		}
		return "func(ctx context.Context) (string, error) { return \"ok\", nil }"
	case "auth":
		return "func(user, pass string) bool { return user == \"admin\" && pass == \"secret\" }"
	case "cron", "queue", "event":
		return "func() { fmt.Println(\"done\") }"
	case "slice", "maputil":
		return "func(v string) bool { return v != \"\" }"
	case "validate":
		return "func(v any) error { return nil }"
	case "stream":
		return "func(p []byte) ([]byte, error) { return p, nil }"
	case "db", "mongo":
		return "func(ctx context.Context) error { return nil }"
	default:
		return "func() error { return nil }"
	}
}

func sampleStruct(pkg, funcName, typ, pkgIdent string) string {
	switch {
	case strings.Contains(typ, "RetryConfig"):
		return pkgIdent + ".RetryConfig{MaxAttempts: 3, Delay: time.Second}"
	case strings.Contains(typ, "SignOptions"):
		return "&" + pkgIdent + ".SignOptions{ExpiresIn: time.Hour}"
	case strings.Contains(typ, "BearerOptions"):
		return "&" + pkgIdent + ".BearerOptions{}"
	case strings.Contains(typ, "APIKeyOpts"):
		return pkgIdent + ".APIKeyOpts{Keys: map[string]bool{\"dev-key\": true}}"
	case strings.Contains(typ, "EnqueueOpts"):
		return "&" + pkgIdent + ".EnqueueOpts{MaxRetry: 3}"
	case strings.Contains(typ, "RequestOpts"):
		return "&" + pkgIdent + ".RequestOpts{Query: map[string]string{\"page\": \"1\"}}"
	case strings.Contains(typ, "CORSOptions"):
		return pkgIdent + ".CORSOptions{AllowOrigins: []string{\"*\"}}"
	case strings.Contains(typ, "SessionOptions"):
		return pkgIdent + ".SessionOptions{MaxAge: time.Hour}"
	case strings.Contains(typ, "RateLimitOptions"):
		return pkgIdent + ".RateLimitOptions{Max: 100, Window: time.Minute}"
	case strings.Contains(typ, "SMTPConfig"):
		return pkgIdent + ".SMTPConfig{Host: \"smtp.example.com\", Port: 587}"
	case strings.Contains(typ, "Message"):
		return "&" + pkgIdent + ".Message{To: []string{\"user@example.com\"}, Subject: \"Hi\", Body: \"Hello\"}"
	case strings.Contains(typ, "Upgrader"):
		return pkgIdent + ".DefaultUpgrader"
	case strings.Contains(typ, "FileEntry"):
		return pkgIdent + ".FileEntry{Name: \"readme.txt\", Body: []byte(\"hello\")}"
	default:
		return "nil"
	}
}

var receiverSetup = map[string]string{
	"queue.Worker":  "mux := queue.NewServeMux()",
	"queue.ServeMux": "mux := queue.NewServeMux()",
}

// curatedExamples overrides generated snippets where context matters.
var curatedExamples = map[string]string{
	"time.Now":           "now := timex.Now()",
	"time.Parse":         "t, err := timex.Parse(timex.LayoutISO, \"2024-06-15T12:00:00Z\")",
	"time.Add":           "later := timex.Add(time.Now(), time.Hour)",
	"time.Diff":          "gap := timex.Diff(deadline, time.Now())",
	"time.StartOfDay":    "start := timex.StartOfDay(time.Now())",
	"time.EndOfDay":      "end := timex.EndOfDay(time.Now())",
	"time.StartOfMonth":  "start := timex.StartOfMonth(time.Now())",
	"time.EndOfMonth":    "end := timex.EndOfMonth(time.Now())",
	"time.ParseDuration": "d, err := timex.ParseDuration(\"2d3h\")",
	"time.Format":        "s := timex.Format(time.Now(), timex.LayoutDate)",

	"env.Load":         "if err := env.Load(\".env\"); err != nil {\n    return err\n}",
	"env.Get":          "port := env.Get(\"PORT\")",
	"env.MustGet":      "secret := env.MustGet(\"JWT_SECRET\")",
	"env.GetInt":       "port, err := env.GetInt(\"PORT\", 8080)",
	"env.GetBool":      "debug, err := env.GetBool(\"DEBUG\", false)",
	"env.GetDuration":  "timeout, err := env.GetDuration(\"TIMEOUT\", time.Second)",
	"env.Set":          "env.Set(\"APP_ENV\", \"development\")",

	"set.Intersection": "both := set.Intersection(set.New(\"a\", \"b\"), set.New(\"b\", \"c\"))",
	"set.Difference":   "onlyA := set.Difference(set.New(\"a\", \"b\"), set.New(\"b\", \"c\"))",

	"async.Retry":   "err := async.Retry(ctx, async.RetryConfig{MaxAttempts: 3, Delay: time.Second}, func(ctx context.Context) error {\n    return fetch(ctx)\n})",
	"async.Timeout": "err := async.Timeout(ctx, 5*time.Second, func(ctx context.Context) error {\n    return work(ctx)\n})",

	"json.MustParse":     "var user User\njson.MustParse([]byte(raw), &user)",
	"json.MustStringify": "raw := json.MustStringify(user)",

	"id.MustParseUUID": "id := id.MustParseUUID(\"550e8400-e29b-41d4-a716-446655440000\")",

	"crypto.RandomString": "token, err := crypto.RandomString(32)",

	"auth.Bearer":       "app.Use(auth.Bearer([]byte(jwtSecret), nil))",
	"auth.APIKey":       "app.Use(auth.APIKey(\"X-API-Key\", auth.APIKeyOpts{Keys: map[string]bool{\"dev-key\": true}}))",
	"auth.Basic":        "app.Use(auth.Basic(func(user, pass string) bool { return user == \"admin\" }))",
	"auth.GetClaims":    "claims, ok := auth.GetClaims(c)",
	"auth.GetAPIKey":    "key, ok := auth.GetAPIKey(c)",
	"auth.GetBasicUser": "user, pass, ok := auth.GetBasicUser(c)",

	"validate.MustValidate": "validate.MustValidate(signup{Email: \"a@b.com\", Age: 21})",
	"validate.ParseJSON":    "var payload signup\nif err := validate.ParseJSON(r, &payload); err != nil {\n\treturn err\n}",
	"validate.ValidateSchema": "sch := validate.Object(map[string]validate.Field{\n\t\"name\": validate.String().MinLen(2),\n})\nif err := validate.ValidateSchema(sch, data); err != nil {\n\treturn err\n}",
	"validate.Validate": "type signup struct {\n\tEmail string `validate:\"required,email\"`\n\tAge   int    `validate:\"gte=18\"`\n}\nif err := validate.Validate(signup{Email: \"a@b.com\", Age: 20}); err != nil {\n\treturn err\n}",
	"validate.String":     "field := validate.String().Required().Email().MinLen(3)",
	"validate.Int":          "field := validate.Int().Min(18).Max(120)",
	"validate.Float":        "field := validate.Float().Min(0)",
	"validate.Bool":         "field := validate.Bool().Required()",
	"validate.Array":        "field := validate.Array(validate.String().MinLen(1)).MinLen(1)",
	"validate.ObjectField":  "field := validate.ObjectField(map[string]validate.Field{\"city\": validate.String()})",
	"validate.Schema":       "sch := validate.Object(map[string]validate.Field{\"email\": validate.String().Email()})",
	"validate.Field":        "var field validate.Field = validate.String().Email()",
	"validate.ValidationError": "var verr validate.ValidationError\nif errors.As(err, &verr) {\n\tfmt.Println(verr.Error())\n}",
	"validate.Schema.Validate": "if err := sch.Validate(data); err != nil {\n\treturn err\n}",
	"validate.stringField.Email":    "field := validate.String().Email()",
	"validate.stringField.Required": "field := validate.String().Required()",
	"validate.stringField.MinLen":   "field := validate.String().MinLen(2)",
	"validate.intField.Min":         "field := validate.Int().Min(18)",
	"validate.arrayField.MinLen":    "field := validate.Array(validate.String()).MinLen(1)",

	"mail.BuildMIME": "raw := mail.BuildMIME(msg)",

	"cron.Parse": "spec, err := cron.Parse(\"0 * * * *\")",

	"queue.Worker.NewWorker":     "worker := queue.NewWorker(\"localhost:6379\", mux)",
	"queue.ServeMux.NewServeMux": "mux := queue.NewServeMux()",

	"redis.Client.NewFromClient": "wrapper := redis.NewFromClient(rdb)",

	"db.DB.MustOpen": "db := db.MustOpen(ctx, \"postgres\", dsn)",

	"mongo.Eq": "filter := mongo.Eq(\"status\", \"active\")",
	"mongo.Gt": "filter := mongo.Gt(\"age\", 18)",
	"mongo.In": "filter := mongo.In(\"role\", \"admin\", \"editor\")",
	"mongo.Set": "update := mongo.Set(\"name\", \"Ada\")",

	"client.SetDefaultClient": "client.SetDefaultClient(c)",

	"client.RequestOpts.WithQuery":   "opts := client.RequestOpts{}.WithQuery(map[string]string{\"page\": \"1\"})",
	"client.RequestOpts.WithMethod":  "opts := client.RequestOpts{}.WithMethod(\"POST\")",
	"client.RequestOpts.WithHeaders": "opts := client.RequestOpts{}.WithHeaders(map[string]string{\"Authorization\": \"Bearer token\"})",
	"client.RequestOpts.JSONBody":    "opts := client.RequestOpts{}.JSONBody(map[string]string{\"name\": \"Ada\"})",

	"log.SetDefault":          "log.SetDefault(log.New())",
	"log.Logger.NewWithLevel": "logger := log.NewWithLevel(log.LevelDebug)",
	"log.Logger.Default":      "logger := log.Default()",

	"ws.Conn.Upgrade": "conn, err := upgrader.Upgrade(w, r, nil)",

	"http.Middleware.SessionMiddleware": "app.Use(http.SessionMiddleware(http.SessionOptions{}, store))",
	"http.Handler.SSEHandler":           "http.HandleSSE(app, \"/events\", handler)",
	"http.EventStream.SSE":              "stream, err := http.SSE(w, r)",
	"http.Middleware.Recover":           "app.Use(http.Recover())",
	"http.MultipartForm.ParseMultipart": "form, err := http.ParseMultipart(c)",
	"http.MemoryStore.NewMemoryStore":   "store := http.NewMemoryStore()",
	"http.SaveUploadedFile":             "path, err := http.SaveUploadedFile(ctx, file, \"./uploads\")",

	"archive.ZipCreate":        "err := archive.ZipCreate(ctx, \"out.zip\", entries)",
	"archive.ZipCreateEntries": "err := archive.ZipCreateEntries(ctx, w, entries)",
	"archive.ZipExtract":       "err := archive.ZipExtract(ctx, \"archive.zip\", \"./out\")",
	"archive.TarCreate":        "err := archive.TarCreate(ctx, w, entries)",
	"archive.TarExtract":       "err := archive.TarExtract(ctx, r, \"./out\")",

	"csv.Read":        "rows, err := csv.Read(ctx, r, csv.ReadOptions{Header: true})",
	"csv.ReadAll":     "rows, err := csv.ReadAll(ctx, \"data.csv\")",
	"csv.ParseString": "rows, err := csv.ParseString(raw, csv.ReadOptions{Header: true})",
	"csv.Write":       "err := csv.Write(ctx, w, rows, []string{\"name\", \"age\"})",

	"compress.GzipWriter.NewGzipWriter": "w := compress.NewGzipWriter(&buf)",
	"compress.GzipReader.NewGzipReader": "r, err := compress.NewGzipReader(bytes.NewReader(data))",

	"stream.ReadAll":   "data, err := stream.ReadAll(ctx, r)",
	"stream.TeeReader": "reader := stream.TeeReader(r, auditWriter)",
	"stream.Transform": "out := stream.Transform(r, func(p []byte) ([]byte, error) { return p, nil })",

	"osutil.UserInfo.Numeric": "uid, gid, err := osutil.UserInfoNumeric()",

	"err.Is":            "if err.Is(err, target) { /* ... */ }",
	"err.As":            "var appErr err.AppError\nif err.As(err, &appErr) { /* ... */ }",
	"err.AppError.New":  "e := err.AppError.New(404, \"not found\")",
	"err.AppError.Wrap": "return err.AppError.Wrap(cause, 500, \"internal error\")",
}

var placeholderRE = regexp.MustCompile(`/\*\s*(args|handle)\s*\*/|//\s*timex\s*$`)

func exampleLooksPlaceholder(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return true
	}
	if placeholderRE.MatchString(s) {
		return true
	}
	if strings.HasPrefix(s, "_ = "+strings.Fields(s)[1]) && len(s) < 40 {
		// "_ = pkg.Symbol" with no context
		if strings.Count(s, "\n") == 0 && !strings.Contains(s, "// package-level") {
			return strings.HasPrefix(s, "_ = ")
		}
	}
	return false
}
