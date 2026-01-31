package build

import (
	"github.com/funtimecoding/go-library/pkg/provision/packer/build/builder"
	"github.com/funtimecoding/go-library/pkg/provision/packer/build/provisioner"
)

type Build struct {
	architecture        string
	CharacterDevicePort int `json:"-"`
	username            string
	password            string

	Builders     []builder.Builder         `json:"builders"`
	Provisioners []provisioner.Provisioner `json:"provisioners"`
}
