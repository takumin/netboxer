package netboxer

import (
	"errors"

	"github.com/digitalocean/go-netbox/netbox/client/dcim"
	"github.com/digitalocean/go-netbox/netbox/models"
	"github.com/gosimple/slug"
)

// SiteName Site Name
type SiteName string

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
func (n *NetBoxer) Sites(req string) error {
	if req == "" {
		return errors.New("Require Site Name")
	}

	n.getSites()

	if _, ok := n.sites[SiteName(req)]; ok {
		return nil
	}

	slugReq := slug.Make(req)
	res, err := n.client.Dcim.DcimSitesCreate(
		dcim.NewDcimSitesCreateParams().
			WithData(&models.WritableSite{
				Name: &req,
				Slug: &slugReq,
			}), nil)
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
