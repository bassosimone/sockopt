//go:build linux && !bassosimone_sockopt_simulate_nonlinux

//
// SPDX-License-Identifier: Apache-2.0
//
// Adapted from github.com/m-lab/uuid/socookie (Apache-2.0).
//

package sockopt

import (
	"net"
	"syscall"
	"unsafe"
)

const soCookie = 57

// GetSOCookie reads the SO_COOKIE value from the socket.
func GetSOCookie(conn net.Conn) (uint64, error) {
	raw, err := syscallConnUnwrapTCPFn(conn)
	if err != nil {
		return 0, err
	}

	var (
		cookie     uint64
		cookieLen  = uint32(unsafe.Sizeof(cookie))
		syscallErr syscall.Errno
	)
	err = raw.Control(func(fd uintptr) {
		_, _, syscallErr = syscallSyscall6Fn(
			uintptr(syscall.SYS_GETSOCKOPT),
			fd,
			uintptr(syscall.SOL_SOCKET),
			uintptr(soCookie),
			uintptr(unsafe.Pointer(&cookie)),
			uintptr(unsafe.Pointer(&cookieLen)),
			uintptr(0),
		)
	})
	if err != nil {
		return 0, err
	}
	if syscallErr != 0 {
		return 0, syscallErr
	}
	return cookie, nil
}
