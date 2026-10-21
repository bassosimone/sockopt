//go:build !linux || bassosimone_sockopt_simulate_nonlinux

// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"errors"
	"net"
)

// SetCCA attempts to set the congestion control algorithm.
func SetCCA(conn net.Conn, cca string) error {
	return errors.ErrUnsupported
}

// GetBBRInfo reads BBR congestion control information from the socket.
func GetBBRInfo(conn net.Conn) (*BBRInfo, error) {
	return nil, errors.ErrUnsupported
}
