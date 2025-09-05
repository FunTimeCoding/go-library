package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readDCIM(
	n *netbox.Client,
	f *option.Format,
) {
	// Data Center Infrastructure Management

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

		for _, i := range n.DeviceInterfaces(d.Identifier) {
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

	for _, t := range n.RackTypes() {
		fmt.Printf("RackType: %s\n", t.Format(f))
	}

	for _, r := range n.RackRoles() {
		fmt.Printf("RackRole: %s\n", r.Format(f))
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

	// TODO: How to create DeviceBay? "device not compatible"
	//  And where to create DeviceBayTemplate? API only?
	for _, t := range n.DeviceBayTemplates() {
		fmt.Printf("DeviceBayTemplate: %s\n", t.Format(f))
	}

	for _, p := range n.ModuleTypeProfiles() {
		fmt.Printf("ModuleTypeProfile: %s\n", p.Format(f))
	}
}
