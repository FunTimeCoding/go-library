package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/spf13/cobra"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gonetb", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	c := netb.NewEnvironment()
	root := &cobra.Command{
		Use:     "gonetb",
		Version: fmt.Sprintf("%s (%s %s)", version, gitHash, buildDate),
	}
	root.AddCommand(listDevicesCommand(c))
	root.AddCommand(getDeviceCommand(c))
	root.AddCommand(listInterfacesCommand(c))
	root.AddCommand(listAddressesCommand(c))
	root.AddCommand(listSitesCommand(c))
	root.AddCommand(listTenantsCommand(c))
	root.AddCommand(listDeviceRolesCommand(c))
	root.AddCommand(listManufacturersCommand(c))
	root.AddCommand(listDeviceTypesCommand(c))
	root.AddCommand(createSiteCommand(c))
	root.AddCommand(createTenantCommand(c))
	root.AddCommand(createDeviceRoleCommand(c))
	root.AddCommand(createManufacturerCommand(c))
	root.AddCommand(createDeviceTypeCommand(c))
	root.AddCommand(createDeviceCommand(c))
	root.AddCommand(createInterfaceCommand(c))
	root.AddCommand(createAddressCommand(c))
	root.AddCommand(listClusterTypesCommand(c))
	root.AddCommand(listClustersCommand(c))
	root.AddCommand(listVirtualMachinesCommand(c))
	root.AddCommand(createClusterTypeCommand(c))
	root.AddCommand(createClusterCommand(c))
	root.AddCommand(createVirtualMachineCommand(c))
	root.AddCommand(createVirtualInterfaceCommand(c))
	root.AddCommand(createVirtualAddressCommand(c))
	root.AddCommand(listTagsCommand(c))
	root.AddCommand(createTagCommand(c))
	root.AddCommand(addDeviceTagCommand(c))
	root.AddCommand(removeDeviceTagCommand(c))
	root.AddCommand(addVirtualTagCommand(c))
	root.AddCommand(removeVirtualTagCommand(c))

	if f := root.Execute(); f != nil {
		errors.Printf("%v", f)
		os.Exit(1)
	}
}
