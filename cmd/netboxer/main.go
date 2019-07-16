package main

import (
	"flag"

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
	var (
		site string
	)
	flag.StringVar(&site, "Site", "", "NetBox Site")
	flag.Parse()
	netbox := netboxer.NewNetboxer(defaultEndpointURL, defaultAPIToken)
	netbox.Sites(site)
}
