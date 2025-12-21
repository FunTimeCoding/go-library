package physical_address

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/netbox-community/go-netbox/v4"
)

func New(a *netbox.MACAddress) *Address {
	var objectType string

	if a.AssignedObjectType.IsSet() {
		objectType = a.GetAssignedObjectType()

		if objectType != "" {
			validateObjectType(objectType)
		} else {
			errors.Warning(
				"empty object type for physical address %s\n",
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

	var d *netbox.BriefInterface

	if r := a.AssignedObjectType.Get(); r != nil {
		if *r == constant.InterfaceAddress {
			notation.DecodeStrict(
				notation.Encode(a.AssignedObject, false),
				&d,
				false,
			)
		}
	}

	return &Address{
		Identifier:       a.GetId(),
		Name:             a.GetDisplay(),
		Address:          network.PhysicalAddress(a.GetDisplay()),
		ObjectType:       objectType,
		ObjectIdentifier: object,
		Interface:        d,
		Raw:              a,
	}
}
