package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
)

func Read() {
	n := internal.NetBox()
	f := constant.Format

	for _, g := range n.TenantGroups() {
		fmt.Printf("TenantGroup: %s\n", g.Format(f))
	}

	for _, t := range n.Tenants() {
		fmt.Printf("Tenant: %s\n", t.Format(f))
	}

	for _, l := range n.Locations() {
		fmt.Printf("Location: %s\n", l.Format(f))
	}

	for _, s := range n.Sites() {
		fmt.Printf("Site: %s\n", s.Format(f))
	}

	for _, r := range n.DeviceRoles() {
		fmt.Printf("DeviceRole: %s\n", r.Format(f))
	}

	for _, p := range n.Platforms() {
		fmt.Printf("Platform: %s\n", p.Format(f))
	}

	for _, m := range n.Manufacturers() {
		fmt.Printf("Manufacturer: %s\n", m.Format(f))
	}

	for _, t := range n.DeviceTypes() {
		fmt.Printf("DeviceType: %s\n", t.Format(f))
	}

	for _, d := range n.Devices() {
		fmt.Printf("Device: %s\n", d.Format(f))

		for _, i := range n.DeviceInterfaces(d.Name) {
			fmt.Printf("  Interface: %s\n", i.Format(f))
		}

		for _, i := range n.DeviceModuleBays(d.Name) {
			fmt.Printf("  ModuleBay: %s\n", i.Format(f))
		}
	}

	for _, r := range n.Racks() {
		fmt.Printf("Rack: %s\n", r.Format(f))
	}

	for _, p := range n.Cables() {
		fmt.Printf("Cable: %s\n", p.Format(f))
	}

	for _, c := range n.VirtualChassis() {
		fmt.Printf("VirtualChassis: %s\n", c.Format(f))
	}

	for _, r := range n.RackRoles() {
		fmt.Printf("RackRole: %s\n", r.Format(f))
	}

	for _, t := range n.RackTypes() {
		fmt.Printf("RackType: %s\n", t.Format(f))
	}

	for _, g := range n.ClusterGroups() {
		fmt.Printf("ClusterGroup: %s\n", g.Format(f))
	}

	for _, t := range n.ClusterTypes() {
		fmt.Printf("ClusterType: %s\n", t.Format(f))
	}

	for _, c := range n.Clusters() {
		fmt.Printf("Cluster: %s\n", c.Format(f))
	}

	for _, m := range n.VirtualMachines() {
		fmt.Printf("VirtualMachine: %s\n", m.Format(f))
	}

	for _, t := range n.ServiceTemplates() {
		fmt.Printf("ServiceTemplate: %s\n", t.Format(f))
	}

	for _, s := range n.Services() {
		fmt.Printf("Services: %s\n", s.Format(f))
	}

	for _, g := range n.ContactGroups() {
		fmt.Printf("ContactGroup: %s\n", g.Format(f))
	}

	for _, r := range n.ContactRoles() {
		fmt.Printf("ContactRole: %s\n", r.Format(f))
	}

	for _, c := range n.Contacts() {
		fmt.Printf("Contact: %s\n", c.Format(f))
	}

	for _, g := range n.UserGroups() {
		fmt.Printf("UserGroup: %s\n", g.Format(f))
	}

	for _, g := range n.Users() {
		fmt.Printf("User: %s\n", g.Format(f))
	}

	for _, g := range n.NotificationGroups() {
		fmt.Printf("NotificationGroup: %s\n", g.Format(f))
	}

	for _, t := range n.ModuleTypes() {
		fmt.Printf("ModuleType: %s\n", t.Format(f))
	}

	for _, r := range n.InventoryItemRoles() {
		fmt.Printf("InventoryItemRole: %s\n", r.Format(f))
	}

	for _, i := range n.InventoryItems() {
		fmt.Printf("InventoryItem: %s\n", i.Format(f))
	}

	for _, i := range n.ModuleBays() {
		fmt.Printf("ModuleBay: %s\n", i.Format(f))
	}

	for _, e := range n.PowerFeeds() {
		fmt.Printf("PowerFeed: %s\n", e.Format(f))
	}

	for _, p := range n.PowerPanels() {
		fmt.Printf("PowerPanel: %s\n", p.Format(f))
	}

	for _, p := range n.ConsoleServerPorts() {
		fmt.Printf("ConsoleServerPort: %s\n", p.Format(f))
	}

	for _, p := range n.ConsolePorts() {
		fmt.Printf("ConsolePort: %s\n", p.Format(f))
	}

	for _, o := range n.PowerOutlets() {
		fmt.Printf("PowerOutlet: %s\n", o.Format(f))
	}

	for _, p := range n.PowerPorts() {
		fmt.Printf("PowerPort: %s\n", p.Format(f))
	}

	for _, p := range n.FrontPorts() {
		fmt.Printf("FrontPort: %s\n", p.Format(f))
	}

	for _, p := range n.RearPorts() {
		fmt.Printf("FrontPort: %s\n", p.Format(f))
	}

	for _, c := range n.VirtualDeviceContexts() {
		fmt.Printf("VirtualDeviceContext: %s\n", c.Format(f))
	}

	// TODO: How to create DeviceBay? "device not compatible"

	// TODO: WirelessLan, WirelessLanGroup

	// TODO: WirelessLink

	// TODO: VPN Tunnel, TunnelGroup

	// TODO: Circuits
}
