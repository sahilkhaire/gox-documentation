---
title: "validate"
package: "validate"
gox-doc-version: "10"
---

<PackageOverview
  name="validate"
  node="zod, joi"
  import-path="github.com/sahilkhaire/gox/validate"
  subtitle="Package validate provides Zod/Joi-like validation on top of go-playground/validator. Node equivalent: zod, joi, yup"
  :symbol-count=14
  :has-cookbook=true
  migration-link=""
  narrative="Zod/Joi-style validation via struct tags and fluent object schemas."
/>
## Start here

<div class="featured-grid">
<a class="featured-card" href="/packages/validate/validate"><div class="featured-name">Validate</div><div class="featured-summary">schema.validate</div></a>
<a class="featured-card" href="/packages/validate/schema-object"><div class="featured-name">Object</div><div class="featured-summary">z.object</div></a>
<a class="featured-card" href="/packages/validate/string"><div class="featured-name">String</div><div class="featured-summary">z.string</div></a>
</div>

## Common use cases

- Validate request bodies
- Build reusable field schemas
- Return 400 with field errors

## npm → gox

Quick mapping from Node.js patterns to gox APIs:

<table class="npm-map-table"><thead><tr><th>Node.js</th><th>gox</th></tr></thead><tbody>
<tr><td><code>z.object({...})</code></td><td><a href="/packages/validate/object"><code>validate.Object(map[string]validate.Field{...})</code></a></td></tr>
<tr><td><code>z.string().email()</code></td><td><a href="/packages/validate/string"><code>validate.String().Email()</code></a></td></tr>
<tr><td><code>z.number().min(18)</code></td><td><a href="/packages/validate/int"><code>validate.Int().Min(18)</code></a></td></tr>
<tr><td><code>schema.parse(data)</code></td><td><a href="/packages/validate/validate-schema"><code>validate.ValidateSchema(sch, data)</code></a></td></tr>
</tbody></table>

## Quick start

Copy-paste a minimal example:

```go
import "github.com/sahilkhaire/gox/validate"

if err := validate.Validate(&user); err != nil { /* return 400 */ }
```

::: tip Full cookbook
See the [**validate cookbook**](/packages/validate/cookbook) for multi-step recipes and real-world patterns.
:::

## API reference

Browse **14 documented symbols** — each page includes Node.js, standard Go, and gox side-by-side examples.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/validate/array">Array</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/bool">Bool</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/float">Float</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/int">Int</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/must-validate">MustValidate</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>MustValidate panics if validation fails.</td></tr>
<tr><td><a href="/packages/validate/object-field">ObjectField</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ObjectField nests an object schema as a field.</td></tr>
<tr><td><a href="/packages/validate/parse-json">ParseJSON</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ParseJSON decodes JSON from r into v and validates struct tags on v.</td></tr>
<tr><td><a href="/packages/validate/string">String</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/validate">Validate</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Validate runs struct-tag validation on structs.</td></tr>
<tr><td><a href="/packages/validate/validate-schema">ValidateSchema</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ValidateSchema validates v against sch (map or struct).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/validate/field">Field</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Field validates a single value in a programmatic schema.</td></tr>
<tr><td><a href="/packages/validate/schema">Schema</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Schema is a map-based object schema (Joi.object).</td></tr>
<tr><td><a href="/packages/validate/validation-error">ValidationError</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>ValidationError holds human-readable validation messages.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/validate/schema-object">Schema.Object</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">method</span></td><td>Object builds a schema from field definitions.</td></tr>
</tbody></table>

