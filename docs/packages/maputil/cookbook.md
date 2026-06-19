---
title: "maputil Cookbook"
package: "maputil"
gox-doc-version: "10"
---

<div class="api-meta"><span class="api-badge node">Node: lodash object utils</span><span class="api-badge import">github.com/sahilkhaire/gox/maputil</span></div>
# maputil Cookbook

lodash object utilities.

```go
import "github.com/sahilkhaire/gox/maputil"

small := maputil.Pick(config, "host", "port")
merged := maputil.Merge(defaults, overrides)
city, _ := maputil.Get[string](nested, "address.city")
```

← [Back to maputil overview](/packages/maputil/)
