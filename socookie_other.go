//go:build !linux || bassosimone_sockopt_simulate_nonlinux

// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"errors"
	"net"
)

// GetSOCookie reads the SO_COOKIE value from the socket.
func GetSOCookie(conn net.Conn) (uint64, error) {
	return 0, errors.ErrUnsupported
}
