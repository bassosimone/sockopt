// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

type wrappedConn struct {
	net.Conn
	inner net.Conn
}

func (w *wrappedConn) NetConn() net.Conn {
	return w.inner
}

func TestUnwrapTCPConn(t *testing.T) {
	t.Run("direct *net.TCPConn", func(t *testing.T) {
		tc := &net.TCPConn{}
		got, err := UnwrapTCPConn(tc)
		assert.NoError(t, err)
		assert.Same(t, tc, got)
	})

	t.Run("wrapped conn with NetConn()", func(t *testing.T) {
		tc := &net.TCPConn{}
		wrapped := &wrappedConn{inner: tc}
		got, err := UnwrapTCPConn(wrapped)
		assert.NoError(t, err)
		assert.Same(t, tc, got)
	})

	t.Run("non-TCP conn returns ErrNoTCPConn", func(t *testing.T) {
		conn := &net.UDPConn{}
		got, err := UnwrapTCPConn(conn)
		assert.ErrorIs(t, err, ErrNoTCPConn)
		assert.Nil(t, got)
	})
}
