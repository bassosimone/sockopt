//go:build linux && !bassosimone_sockopt_simulate_nonlinux

// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"net"
	"syscall"
)

// syscallConnUnwrapTCP is the function we use to unwrap a [*net.TCPConn] from
// the [net.Conn] and get a [syscall.RawConn] from the [*net.TCPConn].
//
// Callers should use [syscallConnUnwrapTCPFn] so that tests could swap the
// real function away and operate on a mocked connection instead.
func syscallConnUnwrapTCP(conn net.Conn) (syscall.RawConn, error) {
	tc, err := UnwrapTCPConn(conn)
	if err != nil {
		return nil, err
	}
	return tc.SyscallConn()
}

// syscallConnUnwrapTCPFn allows to call [syscallConnUnwrapTCP] indirectly
// so that tests could swap the real function with a mocked function.
var syscallConnUnwrapTCPFn = syscallConnUnwrapTCP

// syscallSetsockoptStringFn allows to call [syscall.SetsockoptString] indirectly
// so that tests could swap the real function with a mocked function.
var syscallSetsockoptStringFn = syscall.SetsockoptString

// syscallSyscall6Fn allows to call [syscall.Syscall6] indirectly
// so that tests could swap the real function with a mocked function.
var syscallSyscall6Fn = syscall.Syscall6
