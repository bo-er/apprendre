[Writing gRPC Interceptors in Go](https://shijuvar.medium.com/writing-grpc-interceptors-in-go-bf3e7671fe48)

In this post, we’ll take a look into how to write gRPC Interceptors in Go. When you write HTTP applications, you can wrap route specific application handlers with _HTTP middleware_ that let you execute some common logic before and after executing application handlers. We typically use _middleware_ to write cross-cutting components such as authorization, logging, caching, etc. The same kind of functionality can be implemented in gRPC by using a concept called _Interceptors_.

# gRPC Interceptors

By using Interceptors, you can intercept the execution of RPC methods on both the client and the server. On both the client and the server, there are two types of Interceptors:

- UnaryInterceptor
- StreamInterceptor

UnaryInterceptor intercepts the unary RPCs, while StreamInterceptor intercepts the streaming RPCs.

In Unary RPCs, the client sends a single request to the server and gets a single response back. In streaming RPCs, client or server, or both side (bi-directional streaming), get a stream to read a sequence of messages back, and then client or server reads from the returned stream until there are no more messages.

# Writing Interceptors in gRPC Client

In gRPC client applications, you write two types of Interceptors:

- **UnaryClientInterceptor:** UnaryClientInterceptor intercepts the execution of a unary RPC on the client.
- **StreamClientInterceptor:** StreamClientInterceptor intercepts the creation of ClientStream. It may return a custom ClientStream to intercept all I/O operations.

## UnaryClientInterceptor

In order to create a UnaryClientInterceptor, do call the _WithUnaryInterceptor_ function by providing a _UnaryClientInterceptor_ func value, which returns a _grpc.DialOption_ that specifies the interceptor for unary RPCs:

```
func WithUnaryInterceptor(f UnaryClientInterceptor) DialOption
```

The returned _grpc.DialOption_ value is then used as an argument to call _grpc.Dial_ function to apply the Interceptor for unary RPCs.

Here’s the definition of _UnaryClientInterceptor_ func type:

```
type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts …CallOption) error
```

The parameter _invoker_ is the handler to complete the RPC and it is the responsibility of the interceptor to call it. The _UnaryClientInterceptor_ func value provides interceptor logic. Here’s an example interceptor that implements the _UnaryClientInterceptor:_

```
func clientInterceptor(
   ctx context.Context,
   method string,
   req interface{},
   reply interface{},
   cc *grpc.ClientConn,
   invoker grpc.UnaryInvoker,
   opts ...grpc.CallOption,
) error {
   // Logic before invoking the invoker
   start := time.Now()
   // Calls the invoker to execute RPC
   err := invoker(ctx, method, req, reply, cc, opts...)
   // Logic after invoking the invoker
   grpcLog.Infof("Invoked RPC method=%s; Duration=%s; Error=%v", method,
      time.Since(start), err)
   return err
}
```

The function below returns an _grpc.DialOption_ value, which calls the _WithUnaryInterceptor_ function by providing the _UnaryClientInterceptor_ func value:

```
func withClientUnaryInterceptor() grpc.DialOption {
   return grpc.WithUnaryInterceptor(clientInterceptor)
}
```

The returned _grpc.DialOption_ value is used as an argument to call grpc.Dial function to apply the Interceptor:

```
conn, err := grpc.Dial(grpcUri, grpc.WithInsecure(), withClientUnaryInterceptor())
```

## StreamClientInterceptor

In order to create a StreamClientInterceptor, do call the _WithStreamInterceptor_ function by providing _StreamClientInterceptor_ func value, which returns a _grpc.DialOption_ that specifies the _Interceptor_ for streaming RPCs:

```
func WithStreamInterceptor(f StreamClientInterceptor) DialOption
```

The returned _grpc.DialOption_ value is then used as an argument to call _grpc.Dial_ function to apply the _Interceptor_ for streaming RPCs.

Here’s the definition of _StreamClientInterceptor_ func type:

```
type StreamClientInterceptor func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error)
```

In order to apply StreamClientInterceptor to streaming RPCs, just pass the returned _grpc.DialOption_ value of _WithStreamInterceptor_ function as an argument for calling *grpc.Dia*l function. You can pass both UnaryClientInterceptor and StreamClientInterceptor values to *grpc.Dia*l function.

```
conn, err := grpc.Dial(grpcUri, grpc.WithInsecure(), withClientUnaryInterceptor(), withClientStreamInterceptor)
```

# Writing Interceptors in gRPC Server

Like gRPC client applications, gRPC server applications provide two types of Interceptors:

- **UnaryServerInterceptor:** UnaryServerInterceptor provides a hook to intercept the execution of a unary RPC on the server.
- **StreamServerInterceptor:** StreamServerInterceptor provides a hook to intercept the execution of a streaming RPC on the server.

## UnaryServerInterceptor

In order to create a UnaryServerInterceptor, do call the _UnaryInterceptor_ function by providing the _UnaryServerInterceptor_ func value as an argument, which returns a _grpc.ServerOption_ value that sets the UnaryServerInterceptor for the server.

```
func UnaryInterceptor(i UnaryServerInterceptor) ServerOption
```

The returned _grpc.ServerOption_ value is then used to provide as an argument to _grpc.NewServer_ function to register as UnaryServerInterceptor.

Here’s the definition of _UnaryServerInterceptor_ func:

```
type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
```

The parameter _info_ contains all the information of this RPC, the interceptor can operate on. And _handler_ is the wrapper of the service method implementation. It is the responsibility of the interceptor to invoke _handler_ to complete the RPC.

## Example Server Unary Interceptor for Authorization

Here’s the example Interceptor for authorizing the RPC methods:

<iframe src="https://shijuvar.medium.com/media/5e24bb233d9e537cd13b51480205d8bd" allowfullscreen="" frameborder="0" height="1070" width="680" title="UnaryServerInterceptor for authorization." class="fd ej ef eo w" scrolling="auto" style="box-sizing: inherit; width: 680px; left: 0px; top: 0px; height: 1070px; position: absolute;"></iframe>

Note that the interceptor logic in the preceding code block, uses the packages _google.golang.org/grpc/codes_ and _google.golang.org/grpc/status._

**Sending metadata between client and server**

The _serverInterceptor_ function uses the _authorize_ function to authorize the RPC call, which receives the authorization token from _metadata_. gRPC supports sending metadata between client and server with Context value. The package **\*google.golang.org/grpc/metadata\*** provides the functionality for metadata.

Here’s the code block that sends JWT token from client to server:

```
ctx := context.Background()
md := metadata.Pairs("authorization", jwtToken)
ctx = metadata.NewOutgoingContext(ctx, md)
// Calls RPC method CreateEvent using the stub client
resp, err := client.CreateEvent(context.Background(), event)
```

Function _Pairs_ of _metadata_ package returns an _MD_ type (_type MD map[string][]string)_ formed by the mapping of key, value.

The code block below receives the metadata from gRPC server in our interceptor logic:

```
md, ok := metadata.FromIncomingContext(ctx)
if !ok {
   return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
}

authHeader, ok := md["authorization"]
if !ok {
   return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
}

token := authHeader[0]
// validateToken function validates the token
err := validateToken(token)
```

Function _FromIncomingContext_ of metadata package returns _MD_ type from which you can receive the metadata.

**Register Interceptor on Server**

Here’s the code block that register the Interceptor when creating the gRPC server:

<iframe src="https://shijuvar.medium.com/media/d172d1e00e910fd8b505c562f4a7673e" allowfullscreen="" frameborder="0" height="461" width="680" title="Register a server unary interceptor for authorization per RPC calls." class="fd ej ef eo w" scrolling="auto" style="box-sizing: inherit; width: 680px; left: 0px; top: 0px; height: 461px; position: absolute;"></iframe>

## StreamServerInterceptor

In order to create a StreamServerInterceptor, calls the _StreamInterceptor_ function by providing the _StreamServerInterceptor_ func value as an argument, which returns a _grpc.ServerOption_ value that sets the StreamServerInterceptor for the server.

```
func StreamInterceptor(i StreamServerInterceptor) ServerOption
```

The returned _grpc.ServerOption_ value is then used to provide as an argument to _grpc.NewServer_ function to register as UnaryServerInterceptor.

Here’s the definition of _StreamServerInterceptor_ func type:

```
type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
```

# Interceptor Chaining using Go gRPC Middleware

By default, gRPC doesn’t allow to apply more than one interceptor either on the client nor on the server side. By using the package [go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware), interceptor chaining that allows you to apply multiple interceptors.

Here’s an example for server chaining:

```
myServer := grpc.NewServer(
    grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(loggingStream, monitoringStream, authStream)),
    grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(loggingUnary, monitoringUnary, authUnary),
)
```

These interceptors will be executed from left to right: logging, monitoring and auth.

Here’s an example for client side chaining:

```
clientConn, err = grpc.Dial(
    address,
        grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(monitoringClientUnary, retryUnary)),
        grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(monitoringClientStream, retryStream)),
)
client = pb_testproto.NewTestServiceClient(clientConn)
resp, err := client.PingEmpty(s.ctx, &myservice.Request{Msg: "hello"})
```

These interceptors will be executed from left to right: monitoring and then retry logic.
