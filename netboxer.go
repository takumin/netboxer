package netboxer

import (
	"github.com/digitalocean/go-netbox/netbox"
	"github.com/digitalocean/go-netbox/netbox/client"
	"github.com/digitalocean/go-netbox/netbox/client/dcim"
	"github.com/digitalocean/go-netbox/netbox/models"
)

// SiteName Site Name
type SiteName string

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

func (n *NetBoxer) getSites() error {
	res, err := n.client.Dcim.DcimSitesList(nil, nil)
	if err != nil {
		return err
	}

	for _, v := range res.Payload.Results {
		n.sites[SiteName(*v.Name)] = v
	}

	return nil
}

// Sites Declare
func (n *NetBoxer) Sites(req *models.Site) error {
	n.getSites()

	if _, ok := n.sites[SiteName(*req.Name)]; ok {
		return nil
	}

	res, err := n.client.Dcim.DcimSitesCreate(nil, nil)
	if err != nil {
		return err
	}

	param := &dcim.DcimSitesReadParams{
		ID: res.Payload.ID,
	}

	site, err := n.client.Dcim.DcimSitesRead(param, nil)
	if err != nil {
		return err
	}

	n.sites[SiteName(*site.Payload.Name)] = site.Payload

	return nil
}
