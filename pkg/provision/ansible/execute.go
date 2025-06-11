package ansible

import (
	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Execute(p *playbook.AnsiblePlaybookCmd) {
	errors.PanicOnError(
		execute.NewDefaultExecute(
			execute.WithCmd(p),
			execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
		).Execute(c.context),
	)
}
