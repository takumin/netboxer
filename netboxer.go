package netboxer

import (
	"github.com/digitalocean/go-netbox/netbox"
	"github.com/digitalocean/go-netbox/netbox/client"
)

type NetBoxer struct {
	client *client.NetBox
}

func NewNetboxer(endpointUrl, apiToken string) (client *NetBoxer) {
	client.client = netbox.NewNetboxWithAPIKey(endpointUrl, apiToken)
	return client
}
