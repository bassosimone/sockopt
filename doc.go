// SPDX-License-Identifier: Apache-2.0

// Package sockopt reads TCP socket options (TCP_INFO, TCP_CC_INFO)
// and sets the congestion control algorithm on a [net.Conn].
//
// On Linux, the functions use getsockopt/setsockopt via raw syscalls,
// while other platforms, they return [errors.ErrUnsupported].
//
// Struct fields use the kernel's native unsigned types. Callers
// with specific serialization needs should map to their own types.
package sockopt
