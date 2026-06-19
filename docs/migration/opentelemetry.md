# OpenTelemetry (Node) → Go

gox does not wrap OpenTelemetry — use the official Go SDK. Pair with `gox/log` for application logs and OTel for traces/metrics.

## Libraries

| Node.js | Go |
|---------|-----|
| `@opentelemetry/sdk-node` | `go.opentelemetry.io/otel/sdk` |
| `@opentelemetry/api` | `go.opentelemetry.io/otel` |
| `@opentelemetry/instrumentation-express` | `go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp` |
| `@opentelemetry/exporter-trace-otlp-http` | `go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp` |
| `prom-client` | `github.com/prometheus/client_golang` |

```bash
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/sdk
go get go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
```

## Basic setup

### Node.js

```js
const { NodeSDK } = require('@opentelemetry/sdk-node');
const { getNodeAutoInstrumentations } = require('@opentelemetry/auto-instrumentations-node');

const sdk = new NodeSDK({ instrumentations: [getNodeAutoInstrumentations()] });
sdk.start();
```

### Go

```go
import (
    "context"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func initTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
    exp, err := otlptracehttp.New(ctx)
    if err != nil {
        return nil, err
    }
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exp),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceName("my-service"),
        )),
    )
    otel.SetTracerProvider(tp)
    return tp, nil
}
```

## Instrument gox/http

Wrap the handler or use otelhttp middleware:

```go
import (
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    goxhttp "github.com/sahilkhaire/gox/http"
)

app := goxhttp.New()
// otelhttp works with any net/http handler
instrumented := otelhttp.NewHandler(app, "api")
http.ListenAndServe(":3000", instrumented)
```

Or trace inside handlers:

```go
tracer := otel.Tracer("my-service")
app.Get("/users", func(c *goxhttp.Ctx) error {
    ctx, span := tracer.Start(c.Request.Context(), "getUsers")
    defer span.End()
    c.Request = c.Request.WithContext(ctx)
    // ...
    return c.JSON(200, users)
})
```

## Logging vs tracing

| Concern | Use |
|---------|-----|
| Application logs | `gox/log` (slog) |
| Request traces | OpenTelemetry |
| HTTP access logs | `gox/http` `Logger()` middleware |
| Metrics | Prometheus or OTel metrics SDK |

Correlate logs with traces by adding `trace_id` to slog output from the active span context.

## Environment

| Node.js | Go |
|---------|-----|
| `OTEL_EXPORTER_OTLP_ENDPOINT` | Same env var supported by OTLP exporters |
| `OTEL_SERVICE_NAME` | Set via `resource.WithAttributes(semconv.ServiceName(...))` |

Use `gox/env` to load `.env` with OTel variables:

```go
env.Load()
endpoint := env.Get("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4318")
```

## Further reading

- [OpenTelemetry Go getting started](https://opentelemetry.io/docs/languages/go/getting-started/)
- [otelhttp instrumentation](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp)
