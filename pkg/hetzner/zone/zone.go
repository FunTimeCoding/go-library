package zone

import "github.com/hetznercloud/hcloud-go/v2/hcloud"

type Zone struct {
	Identifier  int64
	Name        string
	TTL         int
	RecordCount int
	Status      string
	Raw         *hcloud.Zone
}
