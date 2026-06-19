# Passport → gox/auth

Authentication middleware for `gox/http` apps.

## JWT bearer

| Passport | gox/auth |
|----------|----------|
| `passport.use(JwtStrategy, ...)` | `app.Use(auth.Bearer(secret, nil))` |
| `req.user` / JWT payload | `auth.GetClaims(c)` |

```go
import (
    goxhttp "github.com/sahilkhaire/gox/http"
    "github.com/sahilkhaire/gox/auth"
    goxjwt "github.com/sahilkhaire/gox/jwt"
)

secret := []byte(os.Getenv("JWT_SECRET"))
app := goxhttp.New()
app.Use(auth.Bearer(secret, nil))
app.Get("/profile", func(c *goxhttp.Ctx) error {
    claims := auth.GetClaims(c)
    return c.JSON(200, claims)
})
```

Sign tokens with `gox/jwt.Sign` and `SignOptions{ExpiresIn: time.Hour}`.

## API key

| Passport / custom | gox/auth |
|-------------------|----------|
| Header check | `auth.APIKey("X-API-Key", auth.APIKeyOpts{Keys: map[string]bool{"k": true}})` |
| Validated key | `auth.GetAPIKey(c)` |

## HTTP Basic

| Passport | gox/auth |
|----------|----------|
| `passport-http` BasicStrategy | `auth.Basic(func(user, pass string) bool { ... })` |
| `req.user` | `auth.GetBasicUser(c)` |

Failed auth returns `err.Unauthorized` (401 JSON via `gox/http` error handler).
