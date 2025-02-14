// Copyright © Loft Orbital Solutions Inc.
// Use of this source code is governed by a Apache-2.0-style
// license that can be found in the LICENSE file.

// Package slogx enhances the [slog] package with additional handlers and context helpers.
//
// [slog]: https://pkg.go.dev/golang.org/x/exp/slog
package slogx

import (
	"io"
	"log/slog"
	"os"
)

// Format is a log handler format.
type Format uint8

const (
	// Pretty print format.
	Pretty Format = iota
	// JSON format.
	JSON
	// Logfmt (key = value) format.
	Logfmt
)

// UnmarshalText implements [encoding.TextUnmarshaler] for Format.
// Can be used for flag parsing.
func (f *Format) UnmarshalText(text []byte) error {
	switch string(text) {
	case "pretty":
		*f = Pretty
	case "json":
		*f = JSON
	case "logfmt":
		*f = Logfmt
	default:
		return ErrInvalidFormat
	}
	return nil
}

// Default is the default logger.
// It logs to [os.Stdout] using the [Pretty] format.
var Default = New(os.Stdout, Pretty, false)

// New returns a new [slog.Logger] with the given writer and format.
// If verbose is true, the logger will log debug messages.
// The default format is [Pretty].
//
// [slog.Logger]: https://pkg.go.dev/golang.org/x/exp/slog#Logger
func New(w io.Writer, format Format, verbose bool) *slog.Logger {
	opt := &slog.HandlerOptions{Level: slog.LevelInfo}
	if verbose {
		opt.Level = slog.LevelDebug
		opt.AddSource = true
	}

	switch format {
	case Logfmt:
		return slog.New(NewLogfmtHandler(w, opt))
	case JSON:
		return slog.New(slog.NewJSONHandler(w, opt))
	case Pretty:
		return slog.New(NewPrettyHandler(w, opt))
	default:
		return slog.New(NewPrettyHandler(w, opt))
	}
}
