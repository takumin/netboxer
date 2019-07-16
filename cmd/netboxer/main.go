package main

import (
	"github.com/takumin/netboxer"
)

const (
	defaultEndpointURL = "127.0.0.1:32768"
	defaultAPIToken    = "0123456789abcdef0123456789abcdef01234567"
)

var (
	name    = "netboxer"
	version = "unknown"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	netboxer.NewNetboxer(defaultEndpointURL, defaultAPIToken)
}
