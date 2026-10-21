//go:build !linux || bassosimone_sockopt_simulate_nonlinux

// SPDX-License-Identifier: Apache-2.0

package sockopt

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTCPInfoUnsupported(t *testing.T) {
	info, err := GetTCPInfo(&net.TCPConn{})
	assert.ErrorIs(t, err, errors.ErrUnsupported)
	assert.Nil(t, info)
}
