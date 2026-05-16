package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Snapshots(items []*proxmox.Snapshot) []*server.Snapshot {
	result := make([]*server.Snapshot, len(items))

	for i, s := range items {
		result[i] = Snapshot(s)
	}

	return result
}
