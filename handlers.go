// Copyright © Loft Orbital Solutions Inc.
// Use of this source code is governed by a Apache-2.0-style
// license that can be found in the LICENSE file.

package slogx

import (
	"io"
	"log/slog"

	"github.com/lmittmann/tint"
)

// NewPrettyHandler returns a new [slog.Handler] for pretty printing.
//
// [slog.Handler]: https://pkg.go.dev/golang.org/x/exp/slog#Handler
func NewPrettyHandler(w io.Writer, opt *slog.HandlerOptions) slog.Handler {
	topt := &tint.Options{
		Level:     opt.Level,
		AddSource: opt.AddSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{}
			}
			return a
		},
	}

	return tint.NewHandler(w, topt)
}

// NewLogfmtHandler returns a new [slog.Handler] for logfmt printing.
//
// [slog.Handler]: https://pkg.go.dev/golang.org/x/exp/slog#Handler
func NewLogfmtHandler(w io.Writer, opt *slog.HandlerOptions) slog.Handler {
	return slog.NewTextHandler(w, opt)
}
