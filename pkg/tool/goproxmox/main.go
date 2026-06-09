package goproxmox

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmox/command_context"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmox/constant"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	var instance string
	c := command_context.New()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
		PersistentPreRun: func(
			_ *cobra.Command,
			_ []string,
		) {
			c.Initialize(instance)
		},
	}
	o.PersistentFlags().StringVar(
		&instance,
		"instance",
		"",
		"Proxmox instance name",
	)
	o.AddCommand(listInstances(c))
	o.AddCommand(listNodes(c))
	o.AddCommand(getNodeStatus(c))
	o.AddCommand(listMachines(c))
	o.AddCommand(getMachine(c))
	o.AddCommand(startMachine(c))
	o.AddCommand(stopMachine(c))
	o.AddCommand(shutdownMachine(c))
	o.AddCommand(resetMachine(c))
	o.AddCommand(listSnapshots(c))
	o.AddCommand(createSnapshot(c))
	o.AddCommand(rollbackSnapshot(c))
	o.AddCommand(deleteSnapshot(c))
	o.AddCommand(listContainers(c))
	o.AddCommand(getContainer(c))
	o.AddCommand(listNetworks(c))
	errors.PanicOnError(o.Execute())
}
