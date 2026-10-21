//go:build linux && !bassosimone_sockopt_simulate_nonlinux

// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"errors"
	"net"
	"syscall"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// mockRawConn is a fake [syscall.RawConn] that calls the provided
// controlFn when Control is invoked.
type mockRawConn struct {
	controlFn func(func(fd uintptr)) error
}

func (m *mockRawConn) Control(f func(fd uintptr)) error {
	return m.controlFn(f)
}

func (m *mockRawConn) Read(func(fd uintptr) (done bool)) error {
	return errors.New("not implemented")
}

func (m *mockRawConn) Write(func(fd uintptr) (done bool)) error {
	return errors.New("not implemented")
}

func TestSetCCA(t *testing.T) {
	t.Run("unwrap error", func(t *testing.T) {
		expectedErr := errors.New("unwrap failed")
		syscallConnUnwrapTCPFn = func(conn net.Conn) (syscall.RawConn, error) {
			return nil, expectedErr
		}
		defer func() { syscallConnUnwrapTCPFn = syscallConnUnwrapTCP }()

		err := SetCCA(&net.TCPConn{}, "bbr")
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("Control error", func(t *testing.T) {
		expectedErr := errors.New("control failed")
		syscallConnUnwrapTCPFn = func(conn net.Conn) (syscall.RawConn, error) {
			return &mockRawConn{controlFn: func(f func(fd uintptr)) error {
				return expectedErr
			}}, nil
		}
		defer func() { syscallConnUnwrapTCPFn = syscallConnUnwrapTCP }()

		err := SetCCA(&net.TCPConn{}, "bbr")
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("setsockopt error", func(t *testing.T) {
		expectedErr := errors.New("setsockopt failed")
		syscallConnUnwrapTCPFn = func(conn net.Conn) (syscall.RawConn, error) {
			return &mockRawConn{controlFn: func(f func(fd uintptr)) error {
				f(0)
				return nil
			}}, nil
		}
		syscallSetsockoptStringFn = func(fd, level, opt int, s string) error {
			return expectedErr
		}
		defer func() {
			syscallConnUnwrapTCPFn = syscallConnUnwrapTCP
			syscallSetsockoptStringFn = syscall.SetsockoptString
		}()

		err := SetCCA(&net.TCPConn{}, "bbr")
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("success", func(t *testing.T) {
		syscallConnUnwrapTCPFn = func(conn net.Conn) (syscall.RawConn, error) {
			return &mockRawConn{controlFn: func(f func(fd uintptr)) error {
				f(0)
				return nil
			}}, nil
		}
		syscallSetsockoptStringFn = func(fd, level, opt int, s string) error {
			return nil
		}
		defer func() {
			syscallConnUnwrapTCPFn = syscallConnUnwrapTCP
			syscallSetsockoptStringFn = syscall.SetsockoptString
		}()

		err := SetCCA(&net.TCPConn{}, "bbr")
		assert.NoError(t, err)
	})
}

func TestGetBBRInfo(t *testing.T) {
	t.Run("unwrap error", func(t *testing.T) {
		expectedErr := errors.New("unwrap failed")
		syscallConnUnwrapTCPFn = func(conn net.Conn) (syscall.RawConn, error) {
			return nil, expectedErr
		}
		defer func() { syscallConnUnwrapTCPFn = syscallConnUnwrapTCP }()

		info, err := GetBBRInfo(&net.TCPConn{})
		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, info)
	})

	t.Run("Control error", func(t *testing.T) {
		expectedErr := errors.New("control failed")
		syscallConnUnwrapTCPFn = func(conn net.Conn) (syscall.RawConn, error) {
			return &mockRawConn{controlFn: func(f func(fd uintptr)) error {
				return expectedErr
			}}, nil
		}
		defer func() { syscallConnUnwrapTCPFn = syscallConnUnwrapTCP }()

		info, err := GetBBRInfo(&net.TCPConn{})
		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, info)
	})

	t.Run("syscall error", func(t *testing.T) {
		syscallConnUnwrapTCPFn = func(conn net.Conn) (syscall.RawConn, error) {
			return &mockRawConn{controlFn: func(f func(fd uintptr)) error {
				f(0)
				return nil
			}}, nil
		}
		syscallSyscall6Fn = func(trap, a1, a2, a3, a4, a5, a6 uintptr) (uintptr, uintptr, syscall.Errno) {
			return 0, 0, syscall.ENOENT
		}
		defer func() {
			syscallConnUnwrapTCPFn = syscallConnUnwrapTCP
			syscallSyscall6Fn = syscall.Syscall6
		}()

		info, err := GetBBRInfo(&net.TCPConn{})
		assert.ErrorIs(t, err, syscall.ENOENT)
		assert.Nil(t, info)
	})

	t.Run("success", func(t *testing.T) {
		syscallConnUnwrapTCPFn = func(conn net.Conn) (syscall.RawConn, error) {
			return &mockRawConn{controlFn: func(f func(fd uintptr)) error {
				f(0)
				return nil
			}}, nil
		}
		syscallSyscall6Fn = func(trap, a1, a2, a3, a4, a5, a6 uintptr) (uintptr, uintptr, syscall.Errno) {
			return 0, 0, 0
		}
		defer func() {
			syscallConnUnwrapTCPFn = syscallConnUnwrapTCP
			syscallSyscall6Fn = syscall.Syscall6
		}()

		info, err := GetBBRInfo(&net.TCPConn{})
		assert.NoError(t, err)
		assert.NotNil(t, info)
	})
}

func TestParseBBRInfo(t *testing.T) {
	t.Run("size mismatch", func(t *testing.T) {
		ki := &bbrInfoKernel{}
		info, err := parseBBRInfo(ki, 0)
		assert.ErrorIs(t, err, syscall.EINVAL)
		assert.Nil(t, info)
	})

	t.Run("success", func(t *testing.T) {
		ki := &bbrInfoKernel{
			BwLo:       0xDEADBEEF,
			BwHi:       0x00000001,
			MinRTT:     1000,
			PacingGain: 256,
			CwndGain:   512,
		}
		info, err := parseBBRInfo(ki, uint32(unsafe.Sizeof(*ki)))
		assert.NoError(t, err)
		assert.Equal(t, uint64(0x1DEADBEEF), info.BW)
		assert.Equal(t, uint32(1000), info.MinRTT)
		assert.Equal(t, uint32(256), info.PacingGain)
		assert.Equal(t, uint32(512), info.CwndGain)
	})
}
