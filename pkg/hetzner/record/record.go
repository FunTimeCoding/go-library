package record

import "github.com/hetznercloud/hcloud-go/v2/hcloud"

type Record struct {
	Name     string
	Type     string
	TTL      *int
	Values   []string
	ZoneName string
	Raw      *hcloud.ZoneRRSet
}
