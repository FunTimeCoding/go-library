package create_machine

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/luthermonson/go-proxmox"
	"testing"
)

func findOption(
	options []proxmox.VirtualMachineOption,
	name string,
) (any, bool) {
	for _, o := range options {
		if o.Name == name {
			return o.Value, true
		}
	}

	return nil, false
}

func requireOption(
	t *testing.T,
	options []proxmox.VirtualMachineOption,
	name string,
) any {
	t.Helper()
	v, okay := findOption(options, name)

	if !okay {
		t.Fatalf("option %q not found", name)
	}

	return v
}

func TestBuildOptionsDefaults(t *testing.T) {
	m := &Machine{Name: "test-vm"}
	options := m.BuildOptions()
	assert.String(t, "test-vm", requireOption(t, options, "name").(string))
	assert.Integer(t, 2, requireOption(t, options, "cores").(int))
	assert.Integer(t, 1, requireOption(t, options, "sockets").(int))
	assert.Integer(t, 2048, requireOption(t, options, "memory").(int))
	assert.String(
		t,
		"virtio-scsi-single",
		requireOption(t, options, "scsihw").(string),
	)
	assert.String(
		t,
		"local-lvm:32",
		requireOption(t, options, "scsi0").(string),
	)
	assert.String(
		t,
		"order=scsi0",
		requireOption(t, options, "boot").(string),
	)
	assert.String(
		t,
		"virtio,bridge=vnet0",
		requireOption(t, options, "net0").(string),
	)
	_, hasIDE := findOption(options, "ide2")
	assert.Boolean(t, false, hasIDE)
	_, hasCI := findOption(options, "ipconfig0")
	assert.Boolean(t, false, hasCI)
}

func TestBuildOptionsCustomValues(t *testing.T) {
	m := &Machine{
		Name:    "custom",
		Cores:   8,
		Memory:  16384,
		Bridge:  "vmbr0",
		OSType:  "l26",
		Tags:    "prod;web",
	}
	options := m.BuildOptions()
	assert.Integer(t, 8, requireOption(t, options, "cores").(int))
	assert.Integer(t, 16384, requireOption(t, options, "memory").(int))
	assert.String(
		t,
		"virtio,bridge=vmbr0",
		requireOption(t, options, "net0").(string),
	)
	assert.String(t, "l26", requireOption(t, options, "ostype").(string))
	assert.String(t, "prod;web", requireOption(t, options, "tags").(string))
}

func TestBuildOptionsDiskImport(t *testing.T) {
	m := &Machine{
		Name:       "import-vm",
		DiskImport: "local:import/debian-13-generic-amd64.qcow2",
	}
	options := m.BuildOptions()
	assert.String(
		t,
		"local-lvm:0,import-from=local:import/debian-13-generic-amd64.qcow2",
		requireOption(t, options, "scsi0").(string),
	)
}

func TestBuildOptionsDiskImportCustomStorage(t *testing.T) {
	m := &Machine{
		Name:        "import-vm",
		DiskImport:  "local:import/debian-13.qcow2",
		DiskStorage: "ceph-pool",
	}
	options := m.BuildOptions()
	assert.String(
		t,
		"ceph-pool:0,import-from=local:import/debian-13.qcow2",
		requireOption(t, options, "scsi0").(string),
	)
}

func TestBuildOptionsCDROM(t *testing.T) {
	m := &Machine{
		Name:  "iso-vm",
		CDROM: "local:iso/debian-13.iso",
	}
	options := m.BuildOptions()
	assert.String(
		t,
		"local:iso/debian-13.iso,media=cdrom",
		requireOption(t, options, "ide2").(string),
	)
	assert.String(
		t,
		"order=ide2;scsi0",
		requireOption(t, options, "boot").(string),
	)
}

func TestBuildOptionsCDROMWithDiskImport(t *testing.T) {
	m := &Machine{
		Name:       "combo-vm",
		CDROM:      "local:iso/drivers.iso",
		DiskImport: "local:import/base.qcow2",
	}
	options := m.BuildOptions()
	assert.String(
		t,
		"local-lvm:0,import-from=local:import/base.qcow2",
		requireOption(t, options, "scsi0").(string),
	)
	assert.String(
		t,
		"local:iso/drivers.iso,media=cdrom",
		requireOption(t, options, "ide2").(string),
	)
	assert.String(
		t,
		"order=ide2;scsi0",
		requireOption(t, options, "boot").(string),
	)
}

func TestBuildOptionsCloudInit(t *testing.T) {
	m := &Machine{
		Name:   "ci-vm",
		CIUser: "admin",
	}
	options := m.BuildOptions()
	assert.String(t, "admin", requireOption(t, options, "ciuser").(string))
	assert.String(
		t,
		"ip=dhcp",
		requireOption(t, options, "ipconfig0").(string),
	)
}

func TestBuildOptionsCloudInitFull(t *testing.T) {
	m := &Machine{
		Name:            "ci-vm",
		CIUser:          "deploy",
		CIPassword:      "secret",
		SSHKeys:         "ssh-ed25519 AAAA key1\nssh-ed25519 BBBB key2",
		IPConfiguration: "ip=10.0.0.5/24,gw=10.0.0.1",
	}
	options := m.BuildOptions()
	assert.String(t, "deploy", requireOption(t, options, "ciuser").(string))
	assert.String(
		t,
		"secret",
		requireOption(t, options, "cipassword").(string),
	)
	assert.String(
		t,
		"ip=10.0.0.5/24,gw=10.0.0.1",
		requireOption(t, options, "ipconfig0").(string),
	)
	_, hasKeys := findOption(options, "sshkeys")
	assert.Boolean(t, true, hasKeys)
}

func TestBuildOptionsCloudInitSSHKeysOnly(t *testing.T) {
	m := &Machine{
		Name:    "ssh-vm",
		SSHKeys: "ssh-ed25519 AAAA key1",
	}
	options := m.BuildOptions()
	assert.String(
		t,
		"ip=dhcp",
		requireOption(t, options, "ipconfig0").(string),
	)
}

func TestBuildOptionsExtras(t *testing.T) {
	m := &Machine{
		Name:   "extras-vm",
		Extras: "serial0=socket,vga=serial0",
	}
	options := m.BuildOptions()
	assert.String(
		t,
		"socket",
		requireOption(t, options, "serial0").(string),
	)
	assert.String(
		t,
		"serial0",
		requireOption(t, options, "vga").(string),
	)
}

func TestBuildOptionsCustomDiskSize(t *testing.T) {
	m := &Machine{
		Name:     "big-vm",
		DiskSize: 100,
	}
	options := m.BuildOptions()
	assert.String(
		t,
		"local-lvm:100",
		requireOption(t, options, "scsi0").(string),
	)
}
