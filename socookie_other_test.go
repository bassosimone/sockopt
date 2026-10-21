//go:build !linux || bassosimone_sockopt_simulate_nonlinux

// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSOCookieUnsupported(t *testing.T) {
	cookie, err := GetSOCookie(&net.TCPConn{})
	assert.ErrorIs(t, err, errors.ErrUnsupported)
	assert.Equal(t, uint64(0), cookie)
}
