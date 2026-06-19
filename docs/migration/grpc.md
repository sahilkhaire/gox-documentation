# gRPC (@grpc/grpc-js) → Go

gox does not wrap gRPC — use the official Go implementation. This guide maps Node grpc-js patterns.

## Library

| Node.js | Go |
|---------|-----|
| `@grpc/grpc-js` | `google.golang.org/grpc` |
| `@grpc/proto-loader` | `protoc` + `protoc-gen-go` + `protoc-gen-go-grpc` |
| `grpc-tools` | `buf` or `protoc` in CI |

```bash
go get google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Server

### Node.js

```js
const grpc = require('@grpc/grpc-js');
const server = new grpc.Server();
server.addService(UserService, {
  getUser: (call, callback) => {
    callback(null, { id: call.request.id, name: 'Alice' });
  },
});
server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
  server.start();
});
```

### Go

```go
import (
    "net"
    "google.golang.org/grpc"
    pb "myapp/proto"
)

func main() {
    lis, _ := net.Listen("tcp", ":50051")
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &userServer{})
    s.Serve(lis)
}

func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
    return &pb.User{Id: req.Id, Name: "Alice"}, nil
}
```

## Client

| Node.js | Go |
|---------|-----|
| `new UserServiceClient(address, credentials)` | `pb.NewUserServiceClient(conn)` |
| `client.getUser({ id }, callback)` | `client.GetUser(ctx, &pb.GetUserRequest{Id: id})` |
| `grpc.credentials.createInsecure()` | `grpc.WithTransportCredentials(insecure.NewCredentials())` |
| Metadata | `metadata.NewOutgoingContext(ctx, md)` |

```go
conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
defer conn.Close()
client := pb.NewUserServiceClient(conn)
user, err := client.GetUser(ctx, &pb.GetUserRequest{Id: "1"})
```

## Interceptors (middleware)

| Node.js | Go |
|---------|-----|
| Server interceptors | `grpc.UnaryInterceptor(fn)` |
| Client interceptors | `grpc.WithUnaryInterceptor(fn)` |

Use interceptors for logging (`gox/log`), auth (`gox/auth` patterns), and tracing (OpenTelemetry).

## REST + gRPC together

Run Express-style REST with `gox/http` on `:3000` and gRPC on `:50051` in the same process using goroutines:

```go
go func() { app.Listen(":3000") }()
lis, _ := net.Listen("tcp", ":50051")
grpcServer.Serve(lis)
```

## Further reading

- [gRPC Go quick start](https://grpc.io/docs/languages/go/quickstart/)
- [Buf](https://buf.build/) for proto management
