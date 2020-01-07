package astigandi

import (
	"fmt"
	"net/http"
)

// Domain represents a domain
type Domain struct {
	DomainHref        string `json:"domain_href,omitempty"`
	DomainKeysHref    string `json:"domain_keys_href,omitempty"`
	DomainRecordsHref string `json:"domain_records_href,omitempty"`
	FQDN              string `json:"fqdn,omitempty"`
	ZoneHref          string `json:"zone_href,omitempty"`
	ZoneRecordsHref   string `json:"zone_records_href,omitempty"`
	ZoneUUID          string `json:"zone_uuid,omitempty"`
}

// Domains list the domains
func (c *Client) Domains() (ds []Domain, err error) {
	// Send
	if err = c.send(http.MethodGet, "/domains", nil, &ds); err != nil {
		err = fmt.Errorf("astigandi: sending failed: %w", err)
		return
	}
	return
}

// Domain gets a domain
func (c *Client) Domain(fqdn string) (d Domain, err error) {
	// Send
	if err = c.send(http.MethodGet, "/domains/"+fqdn, nil, &d); err != nil {
		err = fmt.Errorf("astigandi: sending failed: %w", err)
		return
	}
	return
}
