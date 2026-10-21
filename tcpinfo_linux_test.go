//go:build linux && !bassosimone_sockopt_simulate_nonlinux

// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"errors"
	"net"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTCPInfo(t *testing.T) {
	t.Run("unwrap error", func(t *testing.T) {
		expectedErr := errors.New("unwrap failed")
		syscallConnUnwrapTCPFn = func(conn net.Conn) (syscall.RawConn, error) {
			return nil, expectedErr
		}
		defer func() { syscallConnUnwrapTCPFn = syscallConnUnwrapTCP }()

		info, err := GetTCPInfo(&net.TCPConn{})
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

		info, err := GetTCPInfo(&net.TCPConn{})
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

		info, err := GetTCPInfo(&net.TCPConn{})
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

		info, err := GetTCPInfo(&net.TCPConn{})
		assert.NoError(t, err)
		assert.NotNil(t, info)
	})
}
