# axios / fetch → gox/client

## Client instance

| axios | gox/client |
|-------|------------|
| `axios.create({ baseURL, timeout, headers })` | `c := client.New(); c.BaseURL = base; c.HTTP.Timeout = timeout; c.DefaultHeaders = h` |
| `axios.get(url, { params })` | `c.Get(ctx, url, &client.RequestOpts{Query: q})` |
| `axios.post(url, data)` | `c.Post(ctx, url, client.JSONBody(data))` |

## Response

| axios | gox/client |
|-------|------------|
| `res.status` | `res.StatusCode()` |
| `res.data` (JSON) | `res.JSON(&v)` |
| `res.data` (text) | `string(bytes)` from `res.Bytes()` |
| `res.headers['x']` | `res.Header("X")` |

## fetch

| fetch | gox/client |
|-------|------------|
| `fetch(url)` | `client.Fetch(ctx, url, nil)` |
| `fetch(url, { method: 'POST', body })` | `client.Fetch(ctx, url, &client.RequestOpts{Method: "POST", Body: body})` |

## Retry

| axios `retry` plugin | gox/client |
|----------------------|------------|
| Retry on 5xx | `c.RetryOn5xx = true` (one extra attempt) |

## Context & cancellation

All requests take `context.Context` first (idiomatic Go). Pass `ctx` from HTTP handlers or `context.WithTimeout` for deadlines.

## Example

```go
c := client.New()
c.BaseURL = "https://api.example.com"

resp, err := c.Get(ctx, "/users", &client.RequestOpts{
    Query: url.Values{"page": {"1"}},
})
if err != nil { return err }
defer resp.Body.Close()

var users []User
if err := resp.JSON(&users); err != nil { return err }
```
