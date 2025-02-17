# slogx

[![Godoc](https://godoc.org/github.com/loft-orbital/slogx?status.svg)](https://pkg.go.dev/github.com/loft-orbital/slogx)
[![Release](https://img.shields.io/github/release/loft-orbital/slogx.svg)](https://github.com/loft-orbital/slogx/releases/latest)

Go library to enhances the [slog](https://pkg.go.dev/golang.org/x/exp/slog) package with additional handlers and context helpers.

## Installation

To install use the following command:

```shell
go get github.com/loft-orbital/slogx
```

## Features

Refer to the [godoc](https://pkg.go.dev/github.com/loft-orbital/slogx) for the full usage documentation.

### Simple logger creation

```go
package main

import "github.com/loft-orbital/slogx"

func main() {
  log := slogx.New(os.Stdout, slogx.JSON, false)
  log.Info("Hello world!")
}
```

### Context Injection

```go
package main

import (
  "context"
  "os"

  "github.com/loft-orbital/slogx"
)

func foo(ctx context.Context) {
  log := slogx.FromContext(ctx)
  log.Info("Hello world!")
}

func main() {
  log := slogx.New(os.Stdout, slogx.JSON, false)
  
  ctx := slogx.ContextWithLogger(context.Background(), log)
  foo(ctx)
}
```

### gRPC interceptors

```go
package main

import (
  "google.golang.org/grpc"
  "github.com/loft-orbital/slogx"
)

func main() {
  dopts := []grpc.DialOptions {
    grpc.WithStreamInterceptor(slogx.StreamClientInterceptor(slogx.Default)),
    // ... others gRPC dial options
  }

  client, err := grpc.NewClient("grpchost:443", dopts...)
  if err!=nil {
    panic(err)
  }

  // gRPC service created using this client will have the default slogx logger
  // available in each stream's context.
}
```

## Contributing

Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file for information on how to contribute to this project.
