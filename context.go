// Copyright © Loft Orbital Solutions Inc.
// Use of this source code is governed by a Apache-2.0-style
// license that can be found in the LICENSE file.

package slogx

import (
	"context"
	"log/slog"
)

type ctxKey struct{}

// ContextWithLogger returns a new [context] containing the given logger.
func ContextWithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, l)
}

// FromContext returns the context's [slog.Logger].
// If no logger is found, it returns the [Default] logger.
//
// [slog.Logger]: https://pkg.go.dev/golang.org/x/exp/slog#Logger
func FromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*slog.Logger); ok {
		return l
	}
	return Default
}
