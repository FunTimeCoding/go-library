package zone

import "github.com/hetznercloud/hcloud-go/v2/hcloud"

func New(v *hcloud.Zone) *Zone {
	return &Zone{
		Identifier:  v.ID,
		Name:        v.Name,
		TTL:         v.TTL,
		RecordCount: v.RecordCount,
		Status:      string(v.Status),
		Raw:         v,
	}
}
