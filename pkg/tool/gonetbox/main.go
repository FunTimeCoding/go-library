package gonetbox

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbox/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/client"
	"github.com/spf13/cobra"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	c := client.NewEnvironment()
	root := &cobra.Command{
		Use:     constant.Name,
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	root.AddCommand(listDevices(c))
	root.AddCommand(getDevice(c))
	root.AddCommand(listInterfaces(c))
	root.AddCommand(listAddresses(c))
	root.AddCommand(listSites(c))
	root.AddCommand(listTenants(c))
	root.AddCommand(listDeviceRoles(c))
	root.AddCommand(listManufacturers(c))
	root.AddCommand(listDeviceTypes(c))
	root.AddCommand(createSite(c))
	root.AddCommand(createTenant(c))
	root.AddCommand(createDeviceRole(c))
	root.AddCommand(createManufacturer(c))
	root.AddCommand(createDeviceType(c))
	root.AddCommand(createDevice(c))
	root.AddCommand(createInterface(c))
	root.AddCommand(createAddress(c))
	root.AddCommand(listClusterTypes(c))
	root.AddCommand(listClusters(c))
	root.AddCommand(listVirtualMachines(c))
	root.AddCommand(createClusterType(c))
	root.AddCommand(createCluster(c))
	root.AddCommand(createVirtualMachine(c))
	root.AddCommand(createVirtualInterface(c))
	root.AddCommand(createVirtualAddress(c))
	root.AddCommand(listTags(c))
	root.AddCommand(createTag(c))
	root.AddCommand(addDeviceTag(c))
	root.AddCommand(removeDeviceTag(c))
	root.AddCommand(addVirtualTag(c))
	root.AddCommand(removeVirtualTag(c))

	if f := root.Execute(); f != nil {
		errors.Printf("%v", f)
		os.Exit(1)
	}
}
