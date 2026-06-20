package create_machine

import (
	"github.com/luthermonson/go-proxmox"
	"strings"
)

func parseExtras(s string) []proxmox.VirtualMachineOption {
	if s == "" {
		return nil
	}

	var result []proxmox.VirtualMachineOption

	for _, pair := range strings.Split(s, ",") {
		parts := strings.SplitN(pair, "=", 2)

		if len(parts) == 2 {
			result = append(
				result,
				proxmox.VirtualMachineOption{
					Name:  strings.TrimSpace(parts[0]),
					Value: strings.TrimSpace(parts[1]),
				},
			)
		}
	}

	return result
}
