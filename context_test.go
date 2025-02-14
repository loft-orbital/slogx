// Copyright © Loft Orbital Solutions Inc.
// Use of this source code is governed by a Apache-2.0-style
// license that can be found in the LICENSE file.

package slogx

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextWithLogger(t *testing.T) {
	t.Parallel()

	l := slog.Default()
	ctx := ContextWithLogger(context.Background(), l)
	al := ctx.Value(ctxKey{})
	assert.Equal(t, l, al)
}

func TestFromContext(t *testing.T) {
	t.Parallel()

	t.Run("logger found", func(t *testing.T) {
		t.Parallel()

		l := slog.Default()
		ctx := context.WithValue(context.Background(), ctxKey{}, l)
		al := FromContext(ctx)
		assert.Equal(t, l, al)
	})

	t.Run("logger not found", func(t *testing.T) {
		t.Parallel()

		al := FromContext(context.Background())
		assert.Equal(t, Default, al)
	})
}
