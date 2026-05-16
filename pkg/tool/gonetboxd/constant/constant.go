package constant

import "github.com/funtimecoding/go-library/pkg/identity"

var Identity = identity.New(
	"gonetboxd",
	"NetBox infrastructure inventory bridge",
	"gonetboxd",
).WithInstructions(
	"NetBox infrastructure inventory - devices, interfaces, IP addresses, sites, clusters, virtual machines. Both read and create operations.",
)

const (
	NetboxListDevices            = "netbox_list_devices"
	NetboxGetDevice              = "netbox_get_device"
	NetboxListInterfaces         = "netbox_list_interfaces"
	NetboxListAddresses          = "netbox_list_addresses"
	NetboxListSites              = "netbox_list_sites"
	NetboxListTenants            = "netbox_list_tenants"
	NetboxListDeviceRoles        = "netbox_list_device_roles"
	NetboxListManufacturers      = "netbox_list_manufacturers"
	NetboxListDeviceTypes        = "netbox_list_device_types"
	NetboxCreateSite             = "netbox_create_site"
	NetboxCreateTenant           = "netbox_create_tenant"
	NetboxCreateDeviceRole       = "netbox_create_device_role"
	NetboxCreateManufacturer     = "netbox_create_manufacturer"
	NetboxCreateDeviceType       = "netbox_create_device_type"
	NetboxCreateDevice           = "netbox_create_device"
	NetboxCreateInterface        = "netbox_create_interface"
	NetboxCreateAddress          = "netbox_create_address"
	NetboxListClusterTypes       = "netbox_list_cluster_types"
	NetboxListClusters           = "netbox_list_clusters"
	NetboxListVirtualMachines    = "netbox_list_virtual_machines"
	NetboxCreateClusterType      = "netbox_create_cluster_type"
	NetboxCreateCluster          = "netbox_create_cluster"
	NetboxCreateVirtualMachine   = "netbox_create_virtual_machine"
	NetboxCreateVirtualInterface = "netbox_create_virtual_interface"
	NetboxCreateVirtualAddress   = "netbox_create_virtual_address"
	NetboxListTags               = "netbox_list_tags"
	NetboxCreateTag              = "netbox_create_tag"
	NetboxAddDeviceTag           = "netbox_add_device_tag"
	NetboxRemoveDeviceTag        = "netbox_remove_device_tag"
	NetboxAddVirtualTag                    = "netbox_add_virtual_tag"
	NetboxRemoveVirtualTag                 = "netbox_remove_virtual_tag"
	NetboxListTunnelGroups                 = "netbox_list_tunnel_groups"
	NetboxListTunnels                      = "netbox_list_tunnels"
	NetboxListTunnelTerminations           = "netbox_list_tunnel_terminations"
	NetboxCreateTunnelGroup                = "netbox_create_tunnel_group"
	NetboxCreateTunnel                     = "netbox_create_tunnel"
	NetboxCreateDeviceTunnelTermination    = "netbox_create_device_tunnel_termination"
	NetboxCreateVirtualTunnelTermination   = "netbox_create_virtual_tunnel_termination"

	Address        = "address"
	Cluster        = "cluster"
	ClusterType    = "cluster_type"
	Device         = "device"
	Encapsulation  = "encapsulation"
	Group          = "group"
	Interface      = "interface"
	Manufacturer   = "manufacturer"
	Model          = "model"
	Role           = "role"
	Site           = "site"
	Tag            = "tag"
	Tenant         = "tenant"
	Tunnel         = "tunnel"
	Type           = "type"
	VirtualMachine = "virtual_machine"
)
