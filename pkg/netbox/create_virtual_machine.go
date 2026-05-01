package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/cluster"
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
	upstream "github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateVirtualMachine(
	name string,
	cl *cluster.Cluster,
) (*virtual_machine.Machine, error) {
	q := upstream.NewWritableVirtualMachineWithConfigContextRequest(name)
	q.SetCluster(
		upstream.BriefClusterRequestAsDeviceWithConfigContextRequestCluster(
			upstream.NewBriefClusterRequest(cl.Name),
		),
	)
	q.SetStatus(upstream.PATCHEDWRITABLEVIRTUALMACHINEWITHCONFIGCONTEXTREQUESTSTATUS_ACTIVE)
	result, _, e := c.client.VirtualizationAPI.VirtualizationVirtualMachinesCreate(
		c.context,
	).WritableVirtualMachineWithConfigContextRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return virtual_machine.New(result), nil
}
