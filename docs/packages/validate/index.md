---
title: "validate"
package: "validate"
gox-doc-version: "7"
---

<div class="package-hero"><div class="api-meta"><span class="api-badge node">Node: zod, joi</span></div><h1>validate</h1>
<p class="subtitle">Package validate provides Zod/Joi-like validation on top of go-playground/validator. Node equivalent: zod, joi, yup</p><div class="import-line">import "github.com/sahilkhaire/gox/validate"</div></div>
## Quick start

A minimal example to get going:

```go
import "github.com/sahilkhaire/gox/validate"

if err := validate.Validate(&user); err != nil { /* return 400 */ }
```

::: tip Full cookbook
See the [**validate cookbook**](/packages/validate/cookbook) for multi-step recipes and real-world patterns.
:::

## What's in this package

If you have used **zod, joi** in Node.js, this package is your starting point. Browse **14 documented symbols** below — each page includes Node.js, standard Go, and gox side-by-side examples.

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/validate/array">Array</a></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/bool">Bool</a></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/float">Float</a></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/int">Int</a></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/must-validate">MustValidate</a></td><td><span class="kind-pill">func</span></td><td>MustValidate panics if validation fails.</td></tr>
<tr><td><a href="/packages/validate/object-field">ObjectField</a></td><td><span class="kind-pill">func</span></td><td>ObjectField nests an object schema as a field.</td></tr>
<tr><td><a href="/packages/validate/parse-json">ParseJSON</a></td><td><span class="kind-pill">func</span></td><td>ParseJSON decodes JSON from r into v and validates struct tags on v.</td></tr>
<tr><td><a href="/packages/validate/string">String</a></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/validate">Validate</a></td><td><span class="kind-pill">func</span></td><td>Validate runs struct-tag validation on structs.</td></tr>
<tr><td><a href="/packages/validate/validate-schema">ValidateSchema</a></td><td><span class="kind-pill">func</span></td><td>ValidateSchema validates v against sch (map or struct).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/validate/field">Field</a></td><td><span class="kind-pill">type</span></td><td>Field validates a single value in a programmatic schema.</td></tr>
<tr><td><a href="/packages/validate/schema">Schema</a></td><td><span class="kind-pill">type</span></td><td>Schema is a map-based object schema (Joi.object).</td></tr>
<tr><td><a href="/packages/validate/validation-error">ValidationError</a></td><td><span class="kind-pill">type</span></td><td>ValidationError holds human-readable validation messages.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/validate/schema-object">Schema.Object</a></td><td><span class="kind-pill">method</span></td><td>Object builds a schema from field definitions.</td></tr>
</tbody></table>

