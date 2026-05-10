package salt

import "github.com/funtimecoding/go-library/pkg/provision/salt/basic"

type Client struct {
	basic          *basic.Client
	authentication string
}
