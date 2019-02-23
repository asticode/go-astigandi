package main

import (
	"flag"
	"github.com/asticode/go-astilog"
	"github.com/asticode/go-astitools/flag"
	"github.com/asticode/rename-1"
	"github.com/pkg/errors"
)

// Flags
var (
	f = flag.String("fqdn", "","the FQDN")
	n = flag.String("name", "", "the name")
	t = flag.String("type", "", "the type")
	v = flag.String("value", "", "the value")
)

func main() {
	// Parse flags
	s := astiflag.Subcommand()
	flag.Parse()
	astilog.FlagInit()

	// Create client
	c := astigandi.New(astigandi.FlagConfig())

	// Switch on subcommand
	switch s {
	case "create-domain-record":
		err := c.CreateDomainRecord(*f, astigandi.Record{
			RrsetName: *n,
			RrsetTTL: 10800,
			RrsetType: *t,
			RrsetValues: []string{*v},
		})
		if err != nil {
			astilog.Fatal(errors.Wrap(err, "main: creating domain record failed"))
		}
		astilog.Info("record created")
	case "get-domain":
		d, err := c.Domain(*f)
		if err != nil {
			astilog.Fatal(errors.Wrap(err, "main: getting domain failed"))
		}
		astilog.Infof("domain is %+v", d)
	case "list-domains":
		ds, err := c.Domains()
		if err != nil {
			astilog.Fatal(errors.Wrap(err, "main: listing domains failed"))
		}
		astilog.Infof("domains are %+v", ds)
	case "list-domain-records":
		rs, err := c.DomainRecords(*f)
		if err != nil {
			astilog.Fatal(errors.Wrap(err, "main: listing domain records failed"))
		}
		astilog.Infof("domain records are %+v", rs)
	case "remove-domain-records":
		err := c.RemoveDomainRecords(*f, astigandi.Record{
			RrsetName: *n,
			RrsetType: *t,
		})
		if err != nil {
			astilog.Fatal(errors.Wrap(err, "main: removing domain records failed"))
		}
		astilog.Info("records removed")
	default:
		astilog.Errorf("unknown subcommand %s", s)
	}
}
