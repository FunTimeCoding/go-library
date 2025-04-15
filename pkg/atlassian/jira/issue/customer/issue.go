package customer

import "github.com/ctreminiom/go-atlassian/pkg/infra/models"

type Issue struct {
	Key    string
	Status string
	Title  string
	Body   string
	Link   string
	Raw    *models.CustomerRequestScheme
}
