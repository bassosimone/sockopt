// SPDX-License-Identifier: Apache-2.0

package sockopt

import "net"

// SetLinger unwraps a [net.Conn] to the underlying [*net.TCPConn]
// and sets SO_LINGER to the given value. Works through TLS and other
// wrapping layers that implement NetConn().
func SetLinger(conn net.Conn, sec int) error {
	tc, err := UnwrapTCPConn(conn)
	if err != nil {
		return err
	}
	return tc.SetLinger(sec)
}
