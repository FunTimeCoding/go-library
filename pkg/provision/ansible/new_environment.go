package ansible

import (
	"github.com/funtimecoding/go-library/pkg/provision/ansible/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.InventoryEnvironment))
}
