package astigandi

import (
	"fmt"
	"net/http"
)

// Rrset types
const (
	RrsetTypeA     = "A"
	RrsetTypeCNAME = "CNAME"
	RrsetTypeMX    = "MX"
	RrsetTypeTXT   = "TXT"
)

// Record represents a record
type Record struct {
	RrsetHref   string   `json:"rrset_href,omitempty"`
	RrsetName   string   `json:"rrset_name,omitempty"`
	RrsetTTL    int      `json:"rrset_ttl,omitempty"`
	RrsetType   string   `json:"rrset_type,omitempty"`
	RrsetValues []string `json:"rrset_values,omitempty"`
}

// DomainRecords list the domain's records
func (c *Client) DomainRecords(fqdn string) (rs []Record, err error) {
	// Send
	if err = c.send(http.MethodGet, "/domains/"+fqdn+"/records", nil, &rs); err != nil {
		err = fmt.Errorf("astigandi: sending failed: %w", err)
		return
	}
	return
}

// CreateDomainRecord creates a domain's record
func (c *Client) CreateDomainRecord(fqdn string, r Record) (err error) {
	// Send
	if err = c.send(http.MethodPost, "/domains/"+fqdn+"/records", r, nil); err != nil {
		err = fmt.Errorf("astigandi: sending failed: %w", err)
		return
	}
	return
}

// RemoveDomainRecords removes the domain's records
func (c *Client) RemoveDomainRecords(fqdn string, filter Record) (err error) {
	// Create url
	url := "/domains/" + fqdn + "/records"
	if filter.RrsetName != "" {
		url += "/" + filter.RrsetName
		if filter.RrsetType != "" {
			url += "/" + filter.RrsetType
		}
	}

	// Send
	if err = c.send(http.MethodDelete, url, nil, nil); err != nil {
		err = fmt.Errorf("astigandi: sending failed: %w", err)
		return
	}
	return
}
