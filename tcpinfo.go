//
// SPDX-License-Identifier: Apache-2.0
//
// Adapted from github.com/m-lab/tcp-info/tcp (Apache-2.0).
//

package sockopt

// TCPInfo mirrors the kernel's `struct tcp_info`.
//
// See `include/uapi/linux/tcp.h`.
type TCPInfo struct {
	State       uint8
	CAState     uint8
	Retransmits uint8
	Probes      uint8
	Backoff     uint8
	Options     uint8
	WScale      uint8
	AppLimited  uint8

	RTO    uint32
	ATO    uint32
	SndMSS uint32
	RcvMSS uint32

	Unacked uint32
	Sacked  uint32
	Lost    uint32
	Retrans uint32
	Fackets uint32

	LastDataSent uint32
	LastAckSent  uint32
	LastDataRecv uint32
	LastAckRecv  uint32

	PMTU        uint32
	RcvSsThresh uint32
	RTT         uint32
	RTTVar      uint32
	SndSsThresh uint32
	SndCwnd     uint32
	AdvMSS      uint32
	Reordering  uint32

	RcvRTT   uint32
	RcvSpace uint32

	TotalRetrans uint32

	PacingRate    uint64
	MaxPacingRate uint64

	BytesAcked    uint64
	BytesReceived uint64
	SegsOut       uint32
	SegsIn        uint32

	NotsentBytes uint32
	MinRTT       uint32
	DataSegsIn   uint32
	DataSegsOut  uint32

	DeliveryRate uint64

	BusyTime      uint64
	RWndLimited   uint64
	SndBufLimited uint64

	Delivered   uint32
	DeliveredCE uint32

	BytesSent    uint64
	BytesRetrans uint64

	DSackDups uint32
	ReordSeen uint32

	RcvOooPack uint32

	SndWnd uint32
	RcvWnd uint32

	Rehash uint32

	TotalRTO           uint16
	TotalRTORecoveries uint16
	TotalRTOTime       uint32

	ReceivedCE       uint32
	DeliveredE1Bytes uint32
	DeliveredE0Bytes uint32
	DeliveredCEBytes uint32
	ReceivedE1Bytes  uint32
	ReceivedE0Bytes  uint32
	ReceivedCEBytes  uint32

	ECNFlags uint32
}
