---
title: "str Cookbook"
package: "str"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: string helpers</span><span class="api-badge import">github.com/sahilkhaire/gox/str</span></div>
# str Cookbook

String casing, slugging, and padding for APIs and UI.

```go
import "github.com/sahilkhaire/gox/str"

slug := str.Slug("Hello World!")       // hello-world
camel := str.Camel("foo_bar")        // fooBar
title := str.Capitalize("hello")     // Hello
short := str.Truncate(longText, 80)  // safe rune truncate
```

← [Back to str overview](/packages/str/)
