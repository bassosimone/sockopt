//go:build linux && !bassosimone_sockopt_simulate_nonlinux

// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"net"
	"syscall"
	"unsafe"
)

// tcpCCInfo is the TCP_CC_INFO socket option.
const tcpCCInfo = 26

// bbrInfoKernel mirrors the kernel's struct tcp_bbr_info layout.
type bbrInfoKernel struct {
	BwLo       uint32
	BwHi       uint32
	MinRTT     uint32
	PacingGain uint32
	CwndGain   uint32
}

// SetCCA attempts to set the congestion control algorithm.
func SetCCA(conn net.Conn, cca string) error {
	raw, err := syscallConnUnwrapTCPFn(conn)
	if err != nil {
		return err
	}

	var setsockoptErr error
	err = raw.Control(func(fd uintptr) {
		setsockoptErr = syscallSetsockoptStringFn(
			int(fd), syscall.IPPROTO_TCP, syscall.TCP_CONGESTION, cca,
		)
	})
	if err != nil {
		return err
	}

	return setsockoptErr
}

// GetBBRInfo reads BBR congestion control information from the socket.
func GetBBRInfo(conn net.Conn) (*BBRInfo, error) {
	raw, err := syscallConnUnwrapTCPFn(conn)
	if err != nil {
		return nil, err
	}

	var (
		ki         bbrInfoKernel
		kiLen      = uint32(unsafe.Sizeof(ki))
		syscallErr syscall.Errno
	)
	err = raw.Control(func(fd uintptr) {
		_, _, syscallErr = syscallSyscall6Fn(
			uintptr(syscall.SYS_GETSOCKOPT),
			fd,
			uintptr(syscall.IPPROTO_TCP),
			uintptr(tcpCCInfo),
			uintptr(unsafe.Pointer(&ki)),
			uintptr(unsafe.Pointer(&kiLen)),
			uintptr(0),
		)
	})
	if err != nil {
		return nil, err
	}
	if syscallErr != 0 {
		return nil, syscallErr
	}

	return parseBBRInfo(&ki, kiLen)
}

func parseBBRInfo(ki *bbrInfoKernel, kiLen uint32) (*BBRInfo, error) {
	if kiLen != uint32(unsafe.Sizeof(*ki)) {
		return nil, syscall.EINVAL
	}
	info := &BBRInfo{
		BW:         uint64(ki.BwHi)<<32 | uint64(ki.BwLo),
		MinRTT:     ki.MinRTT,
		PacingGain: ki.PacingGain,
		CwndGain:   ki.CwndGain,
	}
	return info, nil
}
