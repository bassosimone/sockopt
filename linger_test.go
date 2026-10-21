// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLinger(t *testing.T) {
	t.Run("non-TCP conn returns ErrNoTCPConn", func(t *testing.T) {
		err := SetLinger(&net.UDPConn{}, 0)
		assert.ErrorIs(t, err, ErrNoTCPConn)
	})

	t.Run("wrapped non-TCP conn returns ErrNoTCPConn", func(t *testing.T) {
		err := SetLinger(&wrappedConn{inner: &net.UDPConn{}}, 0)
		assert.ErrorIs(t, err, ErrNoTCPConn)
	})

	t.Run("real TCP connection", func(t *testing.T) {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		conn, err := net.Dial("tcp", ln.Addr().String())
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		err = SetLinger(conn, 0)
		assert.NoError(t, err)
	})

	t.Run("wrapped real TCP connection", func(t *testing.T) {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		conn, err := net.Dial("tcp", ln.Addr().String())
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		err = SetLinger(&wrappedConn{inner: conn}, 0)
		assert.NoError(t, err)
	})
}
