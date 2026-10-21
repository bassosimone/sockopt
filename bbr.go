//
// SPDX-License-Identifier: Apache-2.0
//
// Adapted from github.com/m-lab/ndt-server/bbr (Apache-2.0).
//

package sockopt

// BBRInfo contains BBR congestion control information.
type BBRInfo struct {
	BW         uint64
	MinRTT     uint32
	PacingGain uint32
	CwndGain   uint32
}
