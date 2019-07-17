package netboxer

import (
	"github.com/digitalocean/go-netbox/netbox"
	"github.com/digitalocean/go-netbox/netbox/client"
	"github.com/digitalocean/go-netbox/netbox/models"
)

// NetBoxer Client
type NetBoxer struct {
	client *client.NetBox

	sites map[SiteName]*models.Site
}

// NewNetboxer New Client
func NewNetboxer(endpointURL, apiToken string) *NetBoxer {
	n := &NetBoxer{}
	n.client = netbox.NewNetboxWithAPIKey(endpointURL, apiToken)
	n.sites = make(map[SiteName]*models.Site)
	return n
}
