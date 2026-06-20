package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/argument/create_machine"
)

func convertCreateMachine(b *server.CreateMachineRequest) *create_machine.Machine {
	m := create_machine.New()
	m.Node = b.Node
	m.Name = b.Name

	if b.Identifier != nil {
		m.Identifier = int(*b.Identifier)
	}

	if b.Cores != nil {
		m.Cores = *b.Cores
	}

	if b.Memory != nil {
		m.Memory = *b.Memory
	}

	if b.DiskSize != nil {
		m.DiskSize = *b.DiskSize
	}

	if b.DiskStorage != nil {
		m.DiskStorage = *b.DiskStorage
	}

	if b.DiskImport != nil {
		m.DiskImport = *b.DiskImport
	}

	if b.Bridge != nil {
		m.Bridge = *b.Bridge
	}

	if b.OsType != nil {
		m.OSType = *b.OsType
	}

	if b.Start != nil {
		m.Start = *b.Start
	}

	if b.Tags != nil {
		m.Tags = *b.Tags
	}

	if b.CiUser != nil {
		m.CIUser = *b.CiUser
	}

	if b.CiPassword != nil {
		m.CIPassword = *b.CiPassword
	}

	if b.SshKeys != nil {
		m.SSHKeys = *b.SshKeys
	}

	if b.IpConfig != nil {
		m.IPConfiguration = *b.IpConfig
	}

	if b.Extras != nil {
		m.Extras = *b.Extras
	}

	return m
}
