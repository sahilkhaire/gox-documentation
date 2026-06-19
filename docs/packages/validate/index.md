---
title: "validate"
package: "validate"
gox-doc-version: "11"
---

<PackageOverview
  name="validate"
  node="zod, joi"
  import-path="github.com/sahilkhaire/gox/validate"
  subtitle="Package validate provides Zod/Joi-like validation on top of go-playground/validator. Node equivalent: zod, joi, yup"
  :symbol-count=14
/>
## API reference

Select a symbol below — each page explains what it does, shows Node.js vs Go comparisons, and includes a runnable example.

<SymbolFilter placeholder="Filter symbols…" />

### Functions

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/validate/array">Array</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/bool">Bool</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/float">Float</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/int">Int</a></td><td><code class="node-cell">z.number().min(18)</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/must-validate">MustValidate</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>MustValidate panics if validation fails.</td></tr>
<tr><td><a href="/packages/validate/object-field">ObjectField</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ObjectField nests an object schema as a field.</td></tr>
<tr><td><a href="/packages/validate/parse-json">ParseJSON</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>ParseJSON decodes JSON from r into v and validates struct tags on v.</td></tr>
<tr><td><a href="/packages/validate/string">String</a></td><td><code class="node-cell">z.string().email()</code></td><td><span class="kind-pill">func</span></td><td></td></tr>
<tr><td><a href="/packages/validate/validate">Validate</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">func</span></td><td>Validate runs struct-tag validation on structs.</td></tr>
<tr><td><a href="/packages/validate/validate-schema">ValidateSchema</a></td><td><code class="node-cell">schema.parse(data)</code></td><td><span class="kind-pill">func</span></td><td>ValidateSchema validates v against sch (map or struct).</td></tr>
</tbody></table>

### Types

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/validate/field">Field</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Field validates a single value in a programmatic schema.</td></tr>
<tr><td><a href="/packages/validate/schema">Schema</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>Schema is a map-based object schema (Joi.object).</td></tr>
<tr><td><a href="/packages/validate/validation-error">ValidationError</a></td><td><code class="node-cell">—</code></td><td><span class="kind-pill">type</span></td><td>ValidationError holds human-readable validation messages.</td></tr>
</tbody></table>

### Methods

<table class="symbol-table"><thead><tr><th>Symbol</th><th>Node.js</th><th>Kind</th><th>Summary</th></tr></thead><tbody>
<tr><td><a href="/packages/validate/schema-object">Schema.Object</a></td><td><code class="node-cell">z.object({...})</code></td><td><span class="kind-pill">method</span></td><td>Object builds a schema from field definitions.</td></tr>
</tbody></table>

