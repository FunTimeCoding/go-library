package constant

import "github.com/funtimecoding/go-library/pkg/console/status/option"

const (
	HostEnvironment  = "NET_BOX_HOST"
	TokenEnvironment = "NET_BOX_TOKEN"

	NoName           = "no name"
	NoGroup          = "no group"
	NoTenant         = "no tenant"
	NoPrimaryAddress = "no primary address"
	NoComment        = "no comment"

	PageLimit int32 = 1000

	// InterfaceAddress when assigning an IP or MAC address to a network device
	InterfaceAddress = "dcim.interface"

	Interface = "/api"

	SignatureHeader = "X-Hook-Signature"
)

var Format = option.Color.Copy()
