package model_context

import "github.com/mark3labs/mcp-go/mcp"

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_devices",
			mcp.WithDescription("List NetBox devices, optionally filtered by name"),
			mcp.WithString(
				"query",
				mcp.Description("Filter by name (case-insensitive contains)"),
			),
		),
		s.listDevices,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_get_device",
			mcp.WithDescription("Get a NetBox device by exact name"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.getDevice,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_interfaces",
			mcp.WithDescription("List interfaces for a NetBox device"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.listInterfaces,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_addresses",
			mcp.WithDescription("List IP addresses assigned to a NetBox device"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.listAddresses,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_sites",
			mcp.WithDescription("List all NetBox sites"),
		),
		s.listSites,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_tenants",
			mcp.WithDescription("List all NetBox tenants"),
		),
		s.listTenants,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_device_roles",
			mcp.WithDescription("List all NetBox device roles"),
		),
		s.listDeviceRoles,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_manufacturers",
			mcp.WithDescription("List all NetBox manufacturers"),
		),
		s.listManufacturers,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_device_types",
			mcp.WithDescription("List all NetBox device types"),
		),
		s.listDeviceTypes,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_site",
			mcp.WithDescription("Create a NetBox site"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Site name"),
			),
		),
		s.createSite,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_tenant",
			mcp.WithDescription("Create a NetBox tenant"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Tenant name"),
			),
		),
		s.createTenant,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_device_role",
			mcp.WithDescription("Create a NetBox device role"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Role name"),
			),
		),
		s.createDeviceRole,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_manufacturer",
			mcp.WithDescription("Create a NetBox manufacturer"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Manufacturer name"),
			),
		),
		s.createManufacturer,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_device_type",
			mcp.WithDescription("Create a NetBox device type"),
			mcp.WithString(
				"model",
				mcp.Required(),
				mcp.Description("Model name"),
			),
			mcp.WithString(
				"manufacturer",
				mcp.Required(),
				mcp.Description("Manufacturer name (must exist)"),
			),
		),
		s.createDeviceType,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_device",
			mcp.WithDescription("Create a NetBox device. Requires site, role, and device type to exist."),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				"role",
				mcp.Required(),
				mcp.Description("Device role name"),
			),
			mcp.WithString(
				"type",
				mcp.Required(),
				mcp.Description("Device type model name"),
			),
			mcp.WithString(
				"site",
				mcp.Required(),
				mcp.Description("Site name"),
			),
			mcp.WithString(
				"tenant",
				mcp.Description("Tenant name (optional)"),
			),
		),
		s.createDevice,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_interface",
			mcp.WithDescription("Create a network interface on a NetBox device"),
			mcp.WithString(
				"device",
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Interface name (e.g. eno1, eth0)"),
			),
			mcp.WithString(
				"type",
				mcp.Required(),
				mcp.Description("Interface type (e.g. 1000base-t, 10gbase-t)"),
			),
		),
		s.createInterface,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_address",
			mcp.WithDescription("Assign an IP address to a device interface"),
			mcp.WithString(
				"device",
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				"interface",
				mcp.Required(),
				mcp.Description("Interface name (e.g. eno1)"),
			),
			mcp.WithString(
				"address",
				mcp.Required(),
				mcp.Description("IP address in CIDR notation (e.g. 144.76.220.60/32)"),
			),
		),
		s.createAddress,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_cluster_types",
			mcp.WithDescription("List all NetBox cluster types"),
		),
		s.listClusterTypes,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_clusters",
			mcp.WithDescription("List all NetBox clusters"),
		),
		s.listClusters,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_virtual_machines",
			mcp.WithDescription("List all NetBox virtual machines"),
		),
		s.listVirtualMachines,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_cluster_type",
			mcp.WithDescription("Create a NetBox cluster type"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Cluster type name"),
			),
		),
		s.createClusterType,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_cluster",
			mcp.WithDescription("Create a NetBox cluster (VM host grouping)"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Cluster name"),
			),
			mcp.WithString(
				"type",
				mcp.Required(),
				mcp.Description("Cluster type name"),
			),
			mcp.WithString(
				"site",
				mcp.Required(),
				mcp.Description("Site name"),
			),
		),
		s.createCluster,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_virtual_machine",
			mcp.WithDescription("Create a NetBox virtual machine assigned to a cluster"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				"cluster",
				mcp.Required(),
				mcp.Description("Cluster name"),
			),
		),
		s.createVirtualMachine,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_virtual_interface",
			mcp.WithDescription("Create an interface on a NetBox virtual machine"),
			mcp.WithString(
				"virtual_machine",
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Interface name (e.g. ens3)"),
			),
		),
		s.createVirtualInterface,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_virtual_address",
			mcp.WithDescription("Assign an IP address to a virtual machine interface"),
			mcp.WithString(
				"virtual_machine",
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				"interface",
				mcp.Required(),
				mcp.Description("Interface name"),
			),
			mcp.WithString(
				"address",
				mcp.Required(),
				mcp.Description("IP in CIDR notation"),
			),
		),
		s.createVirtualAddress,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_list_tags",
			mcp.WithDescription("List all NetBox tags"),
		),
		s.listTags,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_create_tag",
			mcp.WithDescription("Create a NetBox tag"),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.createTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_add_device_tag",
			mcp.WithDescription("Add a tag to a device"),
			mcp.WithString(
				"device",
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				"tag",
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.addDeviceTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_remove_device_tag",
			mcp.WithDescription("Remove a tag from a device"),
			mcp.WithString(
				"device",
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				"tag",
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.removeDeviceTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_add_virtual_tag",
			mcp.WithDescription("Add a tag to a virtual machine"),
			mcp.WithString(
				"virtual_machine",
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				"tag",
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.addVirtualTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			"netbox_remove_virtual_tag",
			mcp.WithDescription("Remove a tag from a virtual machine"),
			mcp.WithString(
				"virtual_machine",
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				"tag",
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.removeVirtualTag,
	)
}
