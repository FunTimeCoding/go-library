package build

import "github.com/funtimecoding/go-library/pkg/provision/packer/build/provisioner"

func (b *Build) SetProvisioner(commands []string) {
	b.Provisioners = []provisioner.Provisioner{
		{Type: "shell", Inline: commands},
	}
}
