# slogx

[![Godoc](https://godoc.org/github.com/loft-orbital/slogx?status.svg)](https://pkg.go.dev/github.com/loft-orbital/slogx)
[![Release](https://img.shields.io/github/release/loft-orbital/slogx.svg)](https://github.com/loft-orbital/slogx/releases/latest)

A Go library that enhances the standard [slog](https://pkg.go.dev/log/slog) package with additional handlers, context helpers, and optional gRPC integration.

## Key Features

- **Multiple output formats**: Pretty (colorized), JSON, and Logfmt
- **Context-aware logging**: Pass loggers through context instead of function parameters
- **gRPC interceptors**: Automatic logger injection for gRPC services
- **Simple API**: Easy-to-use wrapper around the standard slog package
- **Optional gRPC support**: Use build tags to exclude gRPC dependencies for smaller binaries

## Installation

```shell
go get github.com/loft-orbital/slogx
```

### Reducing Binary Size (nogrpc build tag)

If you don't need gRPC functionality, you can significantly reduce your binary size by using the `nogrpc` build tag:

```shell
go build -tags nogrpc ./...
```

This excludes the gRPC dependencies and interceptors from your build, resulting in a smaller binary footprint.

## Usage Examples

### Available Formats

slogx supports three output formats:

- **`slogx.Pretty`**: Colorized, human-readable output (ideal for development)
- **`slogx.JSON`**: Structured JSON output (ideal for production and log aggregation)
- **`slogx.Logfmt`**: Key-value pair format (compatible with traditional logging systems)

You can also use the default logger which outputs to stdout with Pretty format:

```go
slogx.Default.Info("Using the default logger")
```

### Simple Logger Creation

```go
package main

import (
	"os"

	"github.com/loft-orbital/slogx"
)

func main() {
	// Create a JSON logger
	jsonLog := slogx.New(os.Stdout, slogx.JSON, false)
	jsonLog.Info("Hello world!", "user", "alice", "action", "login")
	
	// Create a pretty logger with colors (great for development)
	prettyLog := slogx.New(os.Stdout, slogx.Pretty, false)
	prettyLog.Info("Hello world!", "user", "alice", "action", "login")
	
	// Create a logfmt logger
	logfmtLog := slogx.New(os.Stdout, slogx.Logfmt, false)
	logfmtLog.Info("Hello world!", "user", "alice", "action", "login")
	
	// Enable verbose mode to include source location
	verboseLog := slogx.New(os.Stdout, slogx.Pretty, true)
	verboseLog.Debug("Debugging info", "variable", "value")
}
```

### Context Injection

Pass loggers through your application using context:

```go
package main

import (
	"context"
	"os"

	"github.com/loft-orbital/slogx"
)

func processRequest(ctx context.Context, userID string) {
	// Retrieve logger from context - falls back to default logger if not found
	log := slogx.FromContext(ctx)
	log.Info("Processing request", "user_id", userID)
	
	// Add request-specific attributes
	requestLog := log.With("request_id", "12345")
	ctx = slogx.ContextWithLogger(ctx, requestLog)
	
	// Pass context to other functions
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	log := slogx.FromContext(ctx)
	log.Info("Doing something") // Will include request_id attribute
}

func main() {
	// Create a logger with custom configuration
	log := slogx.New(os.Stdout, slogx.JSON, false)
	
	// Inject logger into context
	ctx := slogx.ContextWithLogger(context.Background(), log)
	
	processRequest(ctx, "user123")
}
```

### gRPC Interceptors

Automatically inject loggers into gRPC service contexts:

```go
package main

import (
	"log/slog"
	"os"
	
	"google.golang.org/grpc"
	"github.com/loft-orbital/slogx"
)

// Server-side interceptors
func setupGRPCServer() *grpc.Server {
	// Create a custom logger
	logger := slogx.New(os.Stdout, slogx.JSON, false)
	
	// Add server interceptors
	server := grpc.NewServer(
		grpc.UnaryInterceptor(slogx.UnaryServerInterceptor(logger)),
		grpc.StreamInterceptor(slogx.StreamServerInterceptor(logger)),
	)
	
	// Your gRPC handlers can now use slogx.FromContext(ctx) to get the logger
	return server
}

// Client-side interceptors
func setupGRPCClient() (*grpc.ClientConn, error) {
	// Use the default logger or create a custom one
	logger := slogx.Default
	
	// Add client interceptors
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithUnaryInterceptor(slogx.UnaryClientInterceptor(logger)),
		grpc.WithStreamInterceptor(slogx.StreamClientInterceptor(logger)),
		// ... other dial options
	)
	
	return conn, err
}
```

**Note**: If you're not using gRPC in your application, you can exclude these interceptors by building with the `nogrpc` tag to reduce binary size.

## Build Constraints

### Excluding gRPC Support

The gRPC functionality is optional and can be excluded using build tags. This is useful when:
- You want to minimize binary size
- You don't need gRPC functionality
- You want to avoid gRPC dependencies

To build without gRPC support:

```shell
# Build your application
go build -tags nogrpc ./...

# Run tests without gRPC
go test -tags nogrpc ./...

# Install a tool without gRPC
go install -tags nogrpc ./cmd/mytool
```

When built with the `nogrpc` tag, the following functions will not be available:
- `UnaryServerInterceptor`
- `StreamServerInterceptor`
- `UnaryClientInterceptor`
- `StreamClientInterceptor`

## API Reference

For complete API documentation, see the [godoc](https://pkg.go.dev/github.com/loft-orbital/slogx).

### Core Functions

- `New(w io.Writer, format Format, verbose bool) *slog.Logger` - Create a new logger
- `FromContext(ctx context.Context) *slog.Logger` - Get logger from context
- `ContextWithLogger(ctx context.Context, l *slog.Logger) context.Context` - Add logger to context

### Variables

- `Default` - The default logger (Pretty format to stdout)

### Types

- `Format` - Logger format type (Pretty, JSON, or Logfmt)

## Contributing

Please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file for information on how to contribute to this project.
