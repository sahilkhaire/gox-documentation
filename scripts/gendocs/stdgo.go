package main

import (
	"fmt"
	"strings"
)

func standardGoCode(sym symbolDoc, e enrichment, hasEnrich bool) string {
	if hasEnrich && e.StdGo != "" {
		return e.StdGo
	}
	if code := heuristicStdGo(sym); code != "" {
		return code
	}
	return defaultStdGoNote(sym.Pkg)
}

func heuristicStdGo(sym symbolDoc) string {
	name := sym.Name
	pkg := sym.Pkg
	key := pkg + "." + name
	if sym.Receiver != "" {
		key = pkg + "." + sym.Receiver + "." + name
	}

	switch pkg {
	case "cond":
		switch name {
		case "If", "IfLazy":
			return "var result T\nif cond {\n    result = a\n} else {\n    result = b\n}"
		case "Coalesce", "CoalescePtr", "CoalesceFn":
			return "result := fallback\nif v != zero {\n    result = v\n}"
		}
	case "slice":
		return "// Manual loop or Go 1.21+ slices package\nfor _, v := range items {\n    // transform or filter v\n}"
	case "maputil":
		switch name {
		case "Pick", "Omit":
			return "out := make(map[string]any)\nfor k, v := range obj { /* copy keys */ }"
		case "Merge":
			return "out := maps.Clone(base)\nmaps.Copy(out, overlay)"
		case "Get":
			return "// Walk nested maps manually or use a struct"
		case "Keys", "Values":
			return "keys := make([]string, 0, len(m))\nfor k := range m { keys = append(keys, k) }"
		}
	case "str":
		switch name {
		case "Slug", "Camel", "Snake", "Pascal", "Capitalize":
			return "// Manual string transformation with strings/unicode"
		case "Truncate":
			return "if len([]rune(s)) > max {\n    s = string([]rune(s)[:max]) + suffix\n}"
		case "PadStart", "PadEnd":
			return "s = fmt.Sprintf(\"%*s\", width, s) // or custom padding"
		case "IsBlank":
			return "blank := strings.TrimSpace(s) == \"\""
		}
	case "set":
		switch name {
		case "New":
			return "s := make(map[T]struct{})\nfor _, v := range items { s[v] = struct{}{} }"
		case "Union", "Intersection", "Difference":
			return "// Iterate maps/sets manually"
		}
	case "buffer":
		switch name {
		case "From", "FromString", "Concat":
			return "b := append([]byte(nil), parts...)"
		case "Equals", "Compare":
			return "ok := bytes.Equal(a, b)"
		}
	case "async":
		switch name {
		case "All":
			return "errg, ctx := errgroup.WithContext(ctx)\n// errg.Go(func() error { ... })"
		case "Race":
			return "select {\ncase <-doneA:\ncase <-doneB:\ncase <-ctx.Done():\n}"
		case "Sleep":
			return "select {\ncase <-time.After(d):\ncase <-ctx.Done():\n    return ctx.Err()\n}"
		case "Retry", "Timeout":
			return "// Loop with backoff or context.WithTimeout"
		}
	case "path":
		switch name {
		case "Join", "Resolve", "Dirname", "Basename", "Extname", "IsAbs":
			return fmt.Sprintf("// filepath.%s(...)", exportPathName(name))
		}
	case "url":
		switch name {
		case "Parse":
			return "u, err := url.Parse(rawURL)"
		case "ParseQuery", "EncodeQuery":
			return "vals, err := url.ParseQuery(q)\n// or url.Values.Encode()"
		case "Format", "Resolve":
			return "u := &url.URL{Scheme: \"https\", Host: host, Path: path}"
		}
	case "env":
		switch name {
		case "Load":
			return "// godotenv.Load() or manual os.Setenv from .env file"
		case "Get", "MustGet", "GetInt", "GetBool", "GetDuration", "Set":
			return "v := os.Getenv(\"KEY\")\nif v == \"\" {\n    v = \"default\"\n}"
		}
	case "fs":
		switch name {
		case "ReadFile":
			return "data, err := os.ReadFile(path)"
		case "WriteFile":
			return "err := os.WriteFile(path, data, 0644)"
		case "Exists", "Stat", "Mkdir", "Remove", "Copy", "ReadDir":
			return fmt.Sprintf("// os / io helpers — e.g. os.%s", strings.ToLower(name))
		}
	case "err":
		switch name {
		case "NotFound", "BadRequest", "Unauthorized", "Forbidden", "Internal":
			return fmt.Sprintf("return fmt.Errorf(\"%%w: message\", http.Err%s)", strings.ToUpper(name[:1])+name[1:])
		case "Is", "As", "Wrap":
			return "errors.Is(err, target)\nerrors.As(err, &target)\nfmt.Errorf(\"context: %w\", err)"
		}
	case "exec":
		return "cmd := exec.CommandContext(ctx, name, args...)\nout, err := cmd.CombinedOutput()"
	case "osutil":
		switch name {
		case "Hostname", "Homedir", "TempDir", "Platform", "Arch":
			return fmt.Sprintf("val, err := os.%s()", exportOsutilName(name))
		}
	case "stream":
		return "io.Copy(dst, src)\n// or io.ReadAll(r)"
	case "compress":
		switch name {
		case "Gzip", "Gunzip":
			return "// compress/gzip NewWriter / NewReader"
		}
	case "csv":
		return "r := csv.NewReader(f)\nrecords, err := r.ReadAll()"
	case "archive":
		return "// archive/zip or archive/tar from stdlib"
	case "http":
		if sym.Receiver == "Ctx" && name == "JSON" {
			return "w.Header().Set(\"Content-Type\", \"application/json\")\njson.NewEncoder(w).Encode(v)"
		}
		return "func handler(w http.ResponseWriter, r *http.Request) {\n    // chi or net/http\n}"
	case "client":
		return "resp, err := http.NewRequestWithContext(ctx, method, url, body)\nclient.Do(req)"
	case "log":
		return "logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))\nlogger.Info(\"msg\", \"key\", val)"
	case "ws":
		return "// gorilla/websocket Upgrader or Dialer"
	case "json":
		switch name {
		case "Parse", "MustParse", "ParseFile":
			return "err := json.Unmarshal([]byte(raw), &v)"
		case "Stringify", "MustStringify", "Pretty", "WriteFile":
			return "b, err := json.Marshal(v)"
		}
	case "validate":
		return "if err := validator.Struct(v); err != nil { /* handle */ }"
	case "db":
		return "db, err := sqlx.Connect(\"postgres\", dsn)\ndb.GetContext(ctx, &row, query, args...)"
	case "redis":
		return "rdb := redis.NewClient(&redis.Options{Addr: addr})\nval, err := rdb.Get(ctx, key).Result()"
	case "mongo":
		return "client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))"
	case "crypto":
		switch name {
		case "SHA256", "HMACSHA256":
			return "h := sha256.Sum256(data)\n// or hmac.New(sha256.New, key)"
		case "HashPassword", "CheckPassword":
			return "hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)"
		case "RandomBytes", "RandomString":
			return "b := make([]byte, n)\n_, err := rand.Read(b)"
		}
	case "jwt":
		return "// github.com/golang-jwt/jwt/v5 token.Sign / Parse"
	case "id":
		switch name {
		case "NewUUID", "ParseUUID", "MustParseUUID":
			return "id := uuid.New()\nu, err := uuid.Parse(s)"
		case "NewNanoID":
			return "// generate random URL-safe string"
		}
	case "auth":
		return "// net/http middleware checking Authorization header"
	case "cache":
		return "// sync.Map or lru cache library"
	case "event":
		return "// sync.Mutex + map[string][]func() or channels"
	case "cron":
		return "c := cron.New()\nc.AddFunc(spec, fn)\nc.Start()"
	case "queue":
		return "// hibiken/asynq client + worker"
	case "mail":
		return "smtp.SendMail(addr, auth, from, to, msg)"
	case "time":
		switch name {
		case "Now", "Parse", "Format", "Add", "Diff":
			return "t := time.Now()\nt.Format(time.RFC3339)"
		}
	case "semver":
		return "// github.com/Masterminds/semver/v3"
	}
	_ = key
	return ""
}

func exportPathName(name string) string {
	switch name {
	case "Join":
		return "Join"
	case "Resolve":
		return "Join" // closest
	case "Dirname":
		return "Dir"
	case "Basename":
		return "Base"
	case "Extname":
		return "Ext"
	case "IsAbs":
		return "IsAbs"
	default:
		return name
	}
}

func exportOsutilName(name string) string {
	switch name {
	case "Homedir":
		return "UserHomeDir"
	case "TempDir":
		return "TempDir"
	default:
		return name
	}
}

func defaultStdGoNote(pkg string) string {
	switch pkg {
	case "slice":
		return "// Manual for loop or Go 1.21+ slices package"
	case "cond":
		return "if condition {\n    result = a\n} else {\n    result = b\n}"
	case "http":
		return "func handler(w http.ResponseWriter, r *http.Request) {\n    json.NewEncoder(w).Encode(data)\n}"
	case "env":
		return "v := os.Getenv(\"KEY\")\nif v == \"\" {\n    v = \"default\"\n}"
	case "fs":
		return "data, err := os.ReadFile(path)"
	default:
		return "// Use the underlying stdlib or driver directly.\n// See package overview for escape hatches."
	}
}

func isGenericWhenToUse(s string) bool {
	return strings.Contains(s, "Familiar API if you are migrating from Node.js")
}

func buildOverviewDescription(sym symbolDoc, e enrichment, hasEnrich bool) string {
	if hasEnrich && e.Description != "" {
		return collapseWS(e.Description)
	}
	if sym.Summary != "" {
		return escapeHTML(sym.Summary)
	}
	return ""
}
