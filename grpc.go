// Copyright © Loft Orbital Solutions Inc.
// Use of this source code is governed by a Apache-2.0-style
// license that can be found in the LICENSE file.

// +build !nogrpc

package slogx

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
)

type wrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedServerStream) Context() context.Context {
	return w.ctx
}

// UnaryServerInterceptor returns a new gRPC unary server interceptor that will inject the logger into the context.
func UnaryServerInterceptor(l *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		return handler(ContextWithLogger(ctx, l), req)
	}
}

// StreamServerInterceptor returns a new gRPC stream server interceptor that will inject the logger into the context.
func StreamServerInterceptor(l *slog.Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ws := &wrappedServerStream{
			ServerStream: ss,
			ctx:          ContextWithLogger(ss.Context(), l),
		}

		return handler(srv, ws)
	}
}

// UnaryClientInterceptor returns a new gRPC unary client interceptor that will inject the logger into the context.
func UnaryClientInterceptor(l *slog.Logger) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		return invoker(ContextWithLogger(ctx, l), method, req, reply, cc, opts...)
	}
}

// StreamClientInterceptor returns a new gRPC stream client interceptor that will inject the logger into the context.
func StreamClientInterceptor(l *slog.Logger) grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		ctx = ContextWithLogger(ctx, l)
		return streamer(ctx, desc, cc, method, opts...)
	}
}
