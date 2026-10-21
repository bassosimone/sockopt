//go:build !linux || bassosimone_sockopt_simulate_nonlinux

// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"errors"
	"net"
)

// GetTCPInfo reads TCP_INFO from the socket.
//
// Fields not supported by the running kernel will be zero.
func GetTCPInfo(conn net.Conn) (*TCPInfo, error) {
	return nil, errors.ErrUnsupported
}
