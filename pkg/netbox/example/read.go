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

		if false {
			// TODO: on load: panic: no value given for required property device
			//  Even if name is set
			//  But this worked yesterday..?
			for _, i := range n.DeviceModuleBays(d.Name) {
				fmt.Printf("  ModuleBay: %s\n", i.Format(f))
			}
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

	for _, m := range n.Modules() {
		fmt.Printf("Module: %s\n", m.Format(f))
	}

	for _, r := range n.InventoryItemRoles() {
		fmt.Printf("InventoryItemRole: %s\n", r.Format(f))
	}

	for _, i := range n.InventoryItems() {
		fmt.Printf("InventoryItem: %s\n", i.Format(f))
	}

	if false {
		// TODO: on load: panic: no value given for required property device
		//  This also worked - is something corrupt in the devices?
		for _, i := range n.ModuleBays() {
			fmt.Printf("ModuleBay: %s\n", i.Format(f))
		}
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

	for _, t := range n.Tags() {
		fmt.Printf("Tag: %s\n", t.Format(f))
	}

	// TODO: How to create DeviceBay? "device not compatible"
	//  And where to create DeviceBayTemplate? API only?

	for _, t := range n.DeviceBayTemplates() {
		fmt.Printf("DeviceBayTemplate: %s\n", t.Format(f))
	}

	for _, p := range n.ModuleTypeProfiles() {
		fmt.Printf("ModuleTypeProfile: %s\n", p.Format(f))
	}

	if false {
		// TODO: on load: panic: no value given for required property data_path
		for _, c := range n.ConfigContexts() {
			fmt.Printf("ConfigContext: %s\n", c.Format(f))
		}
	}

	// TODO: Must specify either local content or a data file
	//  How, what is this for?
	for _, t := range n.ConfigTemplates() {
		fmt.Printf("ConfigTemplate: %s\n", t.Format(f))
	}

	if false {
		// TODO: on load: panic: json: cannot unmarshal number into Go struct field _PaginatedVLANGroupList.results.vid_ranges of type []int32
		for _, g := range n.VirtualNetworkGroups() {
			fmt.Printf("VirtualNetworkGroup: %s\n", g.Format(f))
		}
	}

	for _, e := range n.VirtualNetworks() {
		fmt.Printf("VirtualNetwork: %s\n", e.Format(f))
	}

	for _, u := range n.SystemNumbers() {
		fmt.Printf("SystemNumber: %s\n", u.Format(f))
	}

	for _, g := range n.WirelessNetworkGroups() {
		fmt.Printf("WirelessNetworkGroup: %s\n", g.Format(f))
	}

	for _, e := range n.WirelessNetworks() {
		fmt.Printf("WirelessNetwork: %s\n", e.Format(f))
	}

	if false {
		// TODO: What must devices have to show up in the picker?
		for _, l := range n.WirelessLinks() {
			fmt.Printf("WirelessLink: %s\n", l.Format(f))
		}
	}

	for _, g := range n.TunnelGroups() {
		fmt.Printf("TunnelGroup: %s\n", g.Format(f))
	}

	for _, t := range n.Tunnels() {
		fmt.Printf("Tunnel: %s\n", t.Format(f))
	}

	// TODO: DataSource: Requires local, git or S3 source
	//  Is this used to import entities?
	for _, s := range n.DataSources() {
		fmt.Printf("DataSource: %s\n", s.Format(f))
	}
	// TODO: ExportTemplate, this optionally depends on DataSource
	//  Whats a good use case?
	for _, t := range n.ExportTemplates() {
		fmt.Printf("ExportTemplate: %s\n", t.Format(f))
	}

	// TODO: CustomField, CustomFieldChoice, CustomLink
	//  Whats a good use case?
	for _, i := range n.CustomFields() {
		fmt.Printf("CustomField: %s\n", i.Format(f))
	}
	for _, c := range n.CustomFieldChoices() {
		fmt.Printf("CustomFieldChoice: %s\n", c.Format(f))
	}
	for _, l := range n.CustomLinks() {
		fmt.Printf("CustomLink: %s\n", l.Format(f))
	}

	// TODO: Circuits
}
