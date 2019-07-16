package netboxer

import (
	"github.com/digitalocean/go-netbox/netbox"
	"github.com/digitalocean/go-netbox/netbox/client"
	"github.com/digitalocean/go-netbox/netbox/models"
	"github.com/takumin/netboxer"
)

type SiteID int64

type NetBoxer struct {
	client *client.NetBox

	sites map[SiteID]*models.Site
}

func NewNetboxer(endpointUrl, apiToken string) *NetBoxer {
	n := &NetBoxer{}
	n.client = netbox.NewNetboxWithAPIKey(endpointUrl, apiToken)
	n.sites = make(map[SiteID]*models.Site)
	n.getSites()
	return n
}

func (n *NetBoxer) getSites() error {
	res, err := n.client.Dcim.DcimSitesList(nil, nil)
	if err != nil {
		return err
	}

	for _, v := range res.Payload.Results {
		n.sites[SiteID(v.ID)] = v
	}

	return nil
}
