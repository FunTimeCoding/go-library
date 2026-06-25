package technitium

import "github.com/funtimecoding/go-library/pkg/technitium/zone"

type zonesResponse struct {
	Zones []*zone.Zone `json:"zones"`
}
