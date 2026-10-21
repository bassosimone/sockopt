# TCP Socket Options

[![GoDoc](https://pkg.go.dev/badge/github.com/bassosimone/sockopt)](https://pkg.go.dev/github.com/bassosimone/sockopt) [![Build Status](https://github.com/bassosimone/sockopt/actions/workflows/go.yml/badge.svg)](https://github.com/bassosimone/sockopt/actions) [![codecov](https://codecov.io/gh/bassosimone/sockopt/branch/main/graph/badge.svg)](https://codecov.io/gh/bassosimone/sockopt)

The `sockopt` Go package reads TCP socket options (e.g., `TCP_INFO`,
`TCP_CC_INFO`, `SO_COOKIE`) and sets the congestion control algorithm
on a `net.Conn`. On non-Linux systems, the functions return
`errors.ErrUnsupported`.

```Go
import (
	"fmt"
	"log"
	"net"

	"github.com/bassosimone/sockopt"
)

conn, err := net.Dial("tcp", "example.com:80")
if err != nil {
	log.Fatal(err)
}
defer conn.Close()

info, err := sockopt.GetTCPInfo(conn)
if err != nil {
	log.Fatal(err)
}
fmt.Println("rtt (us):", info.RTT)
```

## Installation

To add this package as a dependency to your module:

```sh
go get github.com/bassosimone/sockopt
```

## Development

To run the tests:

```sh
go test -v .
```

To measure test coverage:

```sh
go test -v -cover .
```

To exercise the non-Linux code paths on Linux:

```sh
go test -v -tags bassosimone_sockopt_simulate_nonlinux .
```

## License

```
SPDX-License-Identifier: Apache-2.0
```

## History

Adapted from [m-lab/tcp-info](https://github.com/m-lab/tcp-info).
