package main

import (
	"github.com/takumin/netboxer"
)

const (
	defaultEndpointURL = "127.0.0.1:32768"
	defaultAPIToken    = "0123456789abcdef0123456789abcdef01234567"
)

func main() {
	netboxer.NewNetboxer(defaultEndpointURL, defaultAPIToken)
}
