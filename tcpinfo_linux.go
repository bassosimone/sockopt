//go:build linux && !bassosimone_sockopt_simulate_nonlinux

//
// SPDX-License-Identifier: Apache-2.0
//
// Adapted from github.com/m-lab/tcp-info/tcp (Apache-2.0).
//

package sockopt

import (
	"net"
	"syscall"
	"unsafe"
)

// GetTCPInfo reads TCP_INFO from the socket.
//
// Fields not supported by the running kernel will be zero.
func GetTCPInfo(conn net.Conn) (*TCPInfo, error) {
	raw, err := syscallConnUnwrapTCPFn(conn)
	if err != nil {
		return nil, err
	}

	var (
		ti         TCPInfo
		tiLen      = uint32(unsafe.Sizeof(ti))
		syscallErr syscall.Errno
	)
	err = raw.Control(func(fd uintptr) {
		_, _, syscallErr = syscallSyscall6Fn(
			uintptr(syscall.SYS_GETSOCKOPT),
			fd,
			uintptr(syscall.SOL_TCP),
			uintptr(syscall.TCP_INFO),
			uintptr(unsafe.Pointer(&ti)),
			uintptr(unsafe.Pointer(&tiLen)),
			uintptr(0),
		)
	})
	if err != nil {
		return nil, err
	}
	if syscallErr != 0 {
		return nil, syscallErr
	}

	return &ti, nil
}
