package ansible

import "github.com/apenella/go-ansible/v2/pkg/playbook"

func (c *Client) Playbook(name string) *playbook.AnsiblePlaybookCmd {
	return playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(name),
		playbook.WithPlaybookOptions(
			&playbook.AnsiblePlaybookOptions{
				Become:     true,
				Connection: "local",
				Inventory:  c.inventory,
			},
		),
	)
}
