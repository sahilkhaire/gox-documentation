# GraphQL (Apollo) → Go ecosystem

gox does not ship a GraphQL server — use a dedicated Go library. This guide maps common Apollo/GraphQL-js patterns.

## Library choice

| Node.js | Go | Notes |
|---------|-----|-------|
| `apollo-server` | [gqlgen](https://github.com/99designs/gqlgen) | Schema-first, code generation |
| `graphql` (reference) | [graphql-go](https://github.com/graph-gophers/graphql-go) | Schema in Go structs |
| `@apollo/client` | [hasura/go-graphql-client](https://github.com/hasura/go-graphql-client) | Client only |

**Recommendation:** gqlgen for new APIs; graphql-go for minimal schemas.

## Server setup

### Apollo (Node)

```js
const { ApolloServer } = require('@apollo/server');
const { startStandaloneServer } = require('@apollo/server/standalone');

const server = new ApolloServer({ typeDefs, resolvers });
startStandaloneServer(server, { listen: { port: 4000 } });
```

### gqlgen (Go)

```go
import (
    "net/http"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
)

srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
http.Handle("/", playground.Handler("GraphQL", "/query"))
http.Handle("/query", srv)
http.ListenAndServe(":4000", nil)
```

## Resolvers

| Apollo | gqlgen |
|--------|--------|
| `Query: { user: (_, { id }) => getUser(id) }` | `func (r *queryResolver) User(ctx context.Context, id string) (*User, error)` |
| Context per request | `ctx context.Context` first parameter |
| DataLoader | Use [dataloadgen](https://github.com/vikstrous/dataloadgen) or hand-rolled batching |

## With gox/http

Mount gqlgen on a gox sub-router:

```go
import goxhttp "github.com/sahilkhaire/gox/http"

app := goxhttp.New()
gql := app.Router()
gql.Mount("/query", srv) // http.Handler from gqlgen
app.Mount("/graphql", gql)
app.Listen(":4000")
```

## Subscriptions

| Node.js | Go |
|---------|-----|
| `graphql-ws` / Apollo subscriptions | gqlgen subscriptions + WebSocket transport |
| `socket.io` rooms | Use `gox/ws` for raw WebSocket; gqlgen handles GraphQL subscription protocol |

See [ws.md](./ws.md) for WebSocket basics.

## Further reading

- [gqlgen getting started](https://gqlgen.com/getting-started/)
- [GraphQL Go comparison](https://graphql.org/code/)
