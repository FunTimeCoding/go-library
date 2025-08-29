package internet_address

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func New(a *netbox.IPAddress) *Address {
	address, network, e := net.ParseCIDR(a.Display)
	errors.FatalOnError(e)
	var objectType string

	if a.AssignedObjectType.IsSet() {
		objectType = a.GetAssignedObjectType()

		if objectType != "" {
			validateObjectType(objectType)
		} else {
			fmt.Printf(
				"Warning: Empty object type for internet address %s\n",
				a.Display,
			)
		}
	} else {
		objectType = ""
	}

	var object int64

	if a.AssignedObjectId.IsSet() {
		object = a.GetAssignedObjectId()
	}

	return &Address{
		Identifier:       a.GetId(),
		Name:             a.GetDisplay(),
		Address:          address,
		Network:          network,
		ObjectType:       objectType,
		ObjectIdentifier: object,
		Raw:              a,
	}
}
