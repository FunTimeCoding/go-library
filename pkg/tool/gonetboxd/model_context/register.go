package model_context

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListDevices,
			mcp.WithDescription("List NetBox devices, optionally filtered by name"),
			mcp.WithString(
				parameter.Query,
				mcp.Description("Filter by name (case-insensitive contains)"),
			),
		),
		s.listDevices,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxGetDevice,
			mcp.WithDescription("Get a NetBox device by exact name"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.getDevice,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListInterfaces,
			mcp.WithDescription("List interfaces for a NetBox device"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.listInterfaces,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListAddresses,
			mcp.WithDescription("List IP addresses assigned to a NetBox device"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.listAddresses,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListSites,
			mcp.WithDescription("List all NetBox sites"),
		),
		s.listSites,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListTenants,
			mcp.WithDescription("List all NetBox tenants"),
		),
		s.listTenants,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListDeviceRoles,
			mcp.WithDescription("List all NetBox device roles"),
		),
		s.listDeviceRoles,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListManufacturers,
			mcp.WithDescription("List all NetBox manufacturers"),
		),
		s.listManufacturers,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListDeviceTypes,
			mcp.WithDescription("List all NetBox device types"),
		),
		s.listDeviceTypes,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateSite,
			mcp.WithDescription("Create a NetBox site"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Site name"),
			),
		),
		s.createSite,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateTenant,
			mcp.WithDescription("Create a NetBox tenant"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Tenant name"),
			),
		),
		s.createTenant,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateDeviceRole,
			mcp.WithDescription("Create a NetBox device role"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Role name"),
			),
		),
		s.createDeviceRole,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateManufacturer,
			mcp.WithDescription("Create a NetBox manufacturer"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Manufacturer name"),
			),
		),
		s.createManufacturer,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateDeviceType,
			mcp.WithDescription("Create a NetBox device type"),
			mcp.WithString(
				constant.Model,
				mcp.Required(),
				mcp.Description("Model name"),
			),
			mcp.WithString(
				constant.Manufacturer,
				mcp.Required(),
				mcp.Description("Manufacturer name (must exist)"),
			),
		),
		s.createDeviceType,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateDevice,
			mcp.WithDescription("Create a NetBox device. Requires site, role, and device type to exist."),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Role,
				mcp.Required(),
				mcp.Description("Device role name"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Required(),
				mcp.Description("Device type model name"),
			),
			mcp.WithString(
				constant.Site,
				mcp.Required(),
				mcp.Description("Site name"),
			),
			mcp.WithString(
				constant.Tenant,
				mcp.Description("Tenant name (optional)"),
			),
		),
		s.createDevice,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateInterface,
			mcp.WithDescription("Create a network interface on a NetBox device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Interface name (e.g. eno1, eth0)"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Required(),
				mcp.Description("Interface type (e.g. 1000base-t, 10gbase-t)"),
			),
		),
		s.createInterface,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateAddress,
			mcp.WithDescription("Assign an IP address to a device interface"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name (e.g. eno1)"),
			),
			mcp.WithString(
				constant.Address,
				mcp.Required(),
				mcp.Description("IP address in CIDR notation (e.g. 144.76.220.60/32)"),
			),
		),
		s.createAddress,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListClusterTypes,
			mcp.WithDescription("List all NetBox cluster types"),
		),
		s.listClusterTypes,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListClusters,
			mcp.WithDescription("List all NetBox clusters"),
		),
		s.listClusters,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListVirtualMachines,
			mcp.WithDescription("List all NetBox virtual machines"),
		),
		s.listVirtualMachines,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateClusterType,
			mcp.WithDescription("Create a NetBox cluster type"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Cluster type name"),
			),
		),
		s.createClusterType,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateCluster,
			mcp.WithDescription("Create a NetBox cluster (VM host grouping)"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Cluster name"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Required(),
				mcp.Description("Cluster type name"),
			),
			mcp.WithString(
				constant.Site,
				mcp.Required(),
				mcp.Description("Site name"),
			),
		),
		s.createCluster,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateVirtualMachine,
			mcp.WithDescription("Create a NetBox virtual machine assigned to a cluster"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Cluster,
				mcp.Required(),
				mcp.Description("Cluster name"),
			),
		),
		s.createVirtualMachine,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateVirtualInterface,
			mcp.WithDescription("Create an interface on a NetBox virtual machine"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Interface name (e.g. ens3)"),
			),
		),
		s.createVirtualInterface,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateVirtualAddress,
			mcp.WithDescription("Assign an IP address to a virtual machine interface"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name"),
			),
			mcp.WithString(
				constant.Address,
				mcp.Required(),
				mcp.Description("IP in CIDR notation"),
			),
		),
		s.createVirtualAddress,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxListTags,
			mcp.WithDescription("List all NetBox tags"),
		),
		s.listTags,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxCreateTag,
			mcp.WithDescription("Create a NetBox tag"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.createTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxAddDeviceTag,
			mcp.WithDescription("Add a tag to a device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.addDeviceTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxRemoveDeviceTag,
			mcp.WithDescription("Remove a tag from a device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.removeDeviceTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxAddVirtualTag,
			mcp.WithDescription("Add a tag to a virtual machine"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.addVirtualTag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NetboxRemoveVirtualTag,
			mcp.WithDescription("Remove a tag from a virtual machine"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		s.removeVirtualTag,
	)
}
