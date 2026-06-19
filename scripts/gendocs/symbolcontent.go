package main

import (
	"fmt"
	"strings"
)

func buildRichOverview(sym symbolDoc, e enrichment, hasEnrich bool) string {
	var parts []string
	if hasEnrich && e.Description != "" {
		parts = append(parts, collapseWS(e.Description))
	} else if sym.Doc != "" {
		parts = append(parts, escapeHTML(strings.TrimSpace(sym.Doc)))
	} else if sym.Summary != "" {
		parts = append(parts, escapeHTML(sym.Summary))
	}

	if sym.Node != "" {
		parts = append(parts, fmt.Sprintf("If you are coming from Node.js, the closest pattern is **`%s`**.", escapeHTML(sym.Node)))
	} else if analog := packageNodeAnalog[sym.Pkg]; analog != "" {
		parts = append(parts, fmt.Sprintf("Part of the **`%s`** package — Node.js analog: *%s*.", sym.Pkg, escapeHTML(analog)))
	}

	switch sym.Kind {
	case "type":
		parts = append(parts, fmt.Sprintf("`%s` is a type exported by gox. Methods on this type are documented separately.", sym.Name))
	case "method":
		parts = append(parts, fmt.Sprintf("Method on **`%s`** — call it on a value of that type after constructing or receiving one from a constructor.", sym.Receiver))
	case "const", "var":
		parts = append(parts, "Use this constant or variable directly — no function call required.")
	}

	return strings.Join(parts, "\n\n")
}

func buildExampleBlock(sym symbolDoc, pkgIdent string, e enrichment, hasEnrich bool) string {
	var body string
	if hasEnrich && e.Example != "" {
		body = e.Example
	} else {
		body = goxCodeBlock(sym, pkgIdent, e, hasEnrich)
	}
	if strings.TrimSpace(body) == "" {
		return ""
	}
	return fmt.Sprintf("## Example\n\n```go\nimport \"%s/%s\"\n\n%s\n```\n\n", modulePath, sym.Pkg, body)
}

func buildTipsBlock(sym symbolDoc, e enrichment, hasEnrich bool) string {
	if hasEnrich && e.Tips != "" {
		return "## Tips\n\n" + e.Tips + "\n\n"
	}
	if hasEnrich && e.When != "" {
		return "## Tips\n\n::: tip When to use\n" + e.When + "\n:::\n\n"
	}
	tips := autoTips(sym)
	if tips == "" {
		return ""
	}
	return "## Tips\n\n" + tips + "\n\n"
}

func autoTips(sym symbolDoc) string {
	switch sym.Pkg {
	case "cond":
		if sym.Name == "If" {
			return "Both branches are evaluated eagerly. Use `IfLazy` when one branch is expensive."
		}
		if strings.HasPrefix(sym.Name, "Coalesce") {
			return "Prefer explicit zero-value checks in performance-critical hot paths if the compiler cannot prove types."
		}
	case "slice":
		return "Chain `Filter`, `Map`, and `Reduce` for lodash-style pipelines. Results are new slices — inputs are never mutated."
	case "http":
		if sym.Receiver == "Ctx" {
			return "Return errors from handlers; `gox/err` types map to JSON error responses with the right status code."
		}
		return "Stack `Logger`, `Recover`, and `Security` middleware the way you would morgan + helmet in Express."
	case "env":
		return "Call `Load()` once at startup, then use typed getters instead of parsing strings manually."
	case "fs", "client", "db", "redis", "mongo", "exec":
		return "Pass `context.Context` as the first argument so cancellation and deadlines propagate correctly."
	case "async":
		return "All async helpers respect context cancellation — prefer them over raw goroutines when you need timeouts."
	}
	switch sym.Kind {
	case "func":
		return fmt.Sprintf("Import `%s/%s` and call `%s` directly. See the comparison below for the standard library equivalent.", modulePath, sym.Pkg, sym.Name)
	case "method":
		return fmt.Sprintf("Obtain a `%s` value first (see constructors on the package overview), then call `%s`.", sym.Receiver, sym.Name)
	case "type":
		return "Browse methods on this type in the sidebar for handler-style APIs and options structs."
	}
	return ""
}

func buildStdLibBlock(sym symbolDoc) string {
	note := standardGoNote(sym)
	if note == "" {
		return ""
	}
	return "## Standard library alternative\n\n" + note + "\n\n"
}

func nodeCompareCode(sym symbolDoc, e enrichment, hasEnrich bool) string {
	if hasEnrich && e.Node != "" {
		return e.Node
	}
	if sym.Node != "" {
		if strings.Contains(sym.Node, "`") || strings.Contains(sym.Node, "(") {
			return sym.Node
		}
		return fmt.Sprintf("// %s", sym.Node)
	}
	analog := packageNodeAnalog[sym.Pkg]
	if analog != "" {
		return fmt.Sprintf("// Typical %s pattern in Node.js", analog)
	}
	return fmt.Sprintf("// Node.js equivalent for %s.%s", sym.Pkg, sym.Name)
}
