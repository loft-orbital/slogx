// Copyright © Loft Orbital Solutions Inc.
// Use of this source code is governed by a Apache-2.0-style
// license that can be found in the LICENSE file.

package slogx

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	grpcmock "github.com/loft-orbital/slogx/internal/mocks/google.golang.org/grpc"
)

func TestUnaryServerInterceptor(t *testing.T) {
	l := New(io.Discard, Pretty, false)

	h := func(ctx context.Context, req any) (any, error) {
		assert.NotNil(t, ctx)
		log := FromContext(ctx)
		assert.NotNil(t, log)
		assert.Equal(t, l, log)
		return nil, nil
	}

	i := UnaryServerInterceptor(l)
	assert.NotNil(t, i)
	i(context.TODO(), nil, nil, h)
}

func TestStreamServerInterceptor(t *testing.T) {
	l := New(io.Discard, Pretty, false)

	ssm := &grpcmock.MockServerStream{}
	ssm.EXPECT().Context().Return(context.TODO())

	h := func(srv interface{}, ss grpc.ServerStream) error {
		assert.NotNil(t, ss)
		log := FromContext(ss.Context())
		assert.NotNil(t, log)
		assert.Equal(t, l, log)
		return nil
	}

	i := StreamServerInterceptor(l)
	assert.NotNil(t, i)
	i(nil, ssm, nil, h)

	ssm.AssertExpectations(t)
}

func TestUnaryClientInterceptor(t *testing.T) {
	l := New(io.Discard, Pretty, false)

	invoker := func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		opts ...grpc.CallOption,
	) error {
		assert.NotNil(t, ctx)
		log := FromContext(ctx)
		assert.NotNil(t, log)
		assert.Equal(t, l, log)
		return nil
	}

	i := UnaryClientInterceptor(l)
	assert.NotNil(t, i)
	i(context.TODO(), "", nil, nil, nil, invoker)
}

func TestStreamClientInterceptor(t *testing.T) {
	l := New(io.Discard, Pretty, false)

	streamer := func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		assert.NotNil(t, ctx)
		log := FromContext(ctx)
		assert.NotNil(t, log)
		assert.Equal(t, l, log)
		return nil, nil
	}

	i := StreamClientInterceptor(l)
	assert.NotNil(t, i)
	i(context.TODO(), nil, nil, "", streamer)
}
