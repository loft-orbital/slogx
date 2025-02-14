// Copyright © Loft Orbital Solutions Inc.
// Use of this source code is governed by a Apache-2.0-style
// license that can be found in the LICENSE file.

package slogx

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("format pretty", func(t *testing.T) {
		t.Parallel()

		buf := new(bytes.Buffer)
		l := New(buf, Pretty, false)
		if assert.NotNil(t, l) {
			l.Info("test")
			assert.Equal(t, "\x1b[92mINF\x1b[0m test\n", buf.String())
		}
	})

	t.Run("format JSON", func(t *testing.T) {
		t.Parallel()

		buf := new(bytes.Buffer)
		l := New(buf, JSON, false)
		if assert.NotNil(t, l) {
			l.Info("test")

			var actual map[string]interface{}
			if assert.NoError(t, json.Unmarshal(buf.Bytes(), &actual)) {
				assert.Equal(t, "INFO", actual["level"])
				assert.Equal(t, "test", actual["msg"])
			}
		}
	})

	t.Run("format logfmt", func(t *testing.T) {
		t.Parallel()

		buf := new(bytes.Buffer)
		l := New(buf, Logfmt, false)
		if assert.NotNil(t, l) {
			l.Info("test")
			assert.Contains(t, buf.String(), "level=INFO msg=test\n")
		}
	})

	t.Run("format default", func(t *testing.T) {
		t.Parallel()

		buf := new(bytes.Buffer)
		l := New(buf, Format(255), false)
		if assert.NotNil(t, l) {
			l.Info("test")
			assert.Equal(t, "\x1b[92mINF\x1b[0m test\n", buf.String())
		}
	})

	t.Run("format verbose", func(t *testing.T) {
		t.Parallel()

		buf := new(bytes.Buffer)
		l := New(buf, Pretty, true)
		if assert.NotNil(t, l) {
			l.Debug("test")
			assert.NotEmpty(t, buf.String())
		}
	})

	t.Run("not verbose", func(t *testing.T) {
		t.Parallel()

		buf := new(bytes.Buffer)
		l := New(buf, Pretty, false)
		if assert.NotNil(t, l) {
			l.Debug("test")
			assert.Empty(t, buf.String())
			l.Info("test")
			assert.NotEmpty(t, buf.String())
		}
	})
}

func TestFormat_UnmarshalText(t *testing.T) {
	t.Parallel()

	t.Run("json", func(t *testing.T) {
		t.Parallel()

		var f Format
		err := f.UnmarshalText([]byte("json"))
		if assert.NoError(t, err) {
			assert.Equal(t, JSON, f)
		}
	})

	t.Run("pretty", func(t *testing.T) {
		t.Parallel()

		var f Format
		err := f.UnmarshalText([]byte("pretty"))
		if assert.NoError(t, err) {
			assert.Equal(t, Pretty, f)
		}
	})

	t.Run("logfmt", func(t *testing.T) {
		t.Parallel()

		var f Format
		err := f.UnmarshalText([]byte("logfmt"))
		if assert.NoError(t, err) {
			assert.Equal(t, Logfmt, f)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		var f Format
		err := f.UnmarshalText([]byte("invalid"))
		if assert.Error(t, err) {
			assert.Equal(t, ErrInvalidFormat, err)
		}
	})
}
