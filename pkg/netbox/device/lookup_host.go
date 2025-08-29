package device

import (
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/system"
)

func (d *Device) LookupHost() string {
	if d.PrimaryAddress == constant.NoPrimaryAddress {
		return ""
	}

	return system.LookupFirst(d.PrimaryAddress)
}
