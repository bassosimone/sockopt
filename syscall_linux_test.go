//go:build linux && !bassosimone_sockopt_simulate_nonlinux

// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyscallConnUnwrapTCP(t *testing.T) {
	t.Run("success with real TCP connection", func(t *testing.T) {
		ln, err := net.Listen("tcp", "127.0.0.1:0")
		assert.NoError(t, err)
		defer ln.Close()

		conn, err := net.Dial("tcp", ln.Addr().String())
		assert.NoError(t, err)
		defer conn.Close()

		raw, err := syscallConnUnwrapTCP(conn)
		assert.NoError(t, err)
		assert.NotNil(t, raw)
	})

	t.Run("error with non-TCP conn", func(t *testing.T) {
		conn := &net.UDPConn{}
		raw, err := syscallConnUnwrapTCP(conn)
		assert.ErrorIs(t, err, ErrNoTCPConn)
		assert.Nil(t, raw)
	})
}
