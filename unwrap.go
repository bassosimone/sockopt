// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"errors"
	"net"
)

// ErrNoTCPConn indicates that the connection could not be unwrapped to a *net.TCPConn.
var ErrNoTCPConn = errors.New("sockopt: cannot unwrap to *net.TCPConn")

// UnwrapTCPConn recursively unwraps a net.Conn through layers that implement NetConn(),
// such [*tls.Conn], until it finds the underlying [*net.TCPConn].
func UnwrapTCPConn(conn net.Conn) (*net.TCPConn, error) {
	for {
		if tc, ok := conn.(*net.TCPConn); ok {
			return tc, nil
		}
		if w, ok := conn.(interface{ NetConn() net.Conn }); ok {
			conn = w.NetConn()
			continue
		}
		return nil, ErrNoTCPConn
	}
}
