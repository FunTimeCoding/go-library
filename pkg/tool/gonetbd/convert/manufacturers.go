package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func Manufacturers(v []*manufacturer.Manufacturer) []server.Manufacturer {
	result := make([]server.Manufacturer, 0, len(v))

	for _, m := range v {
		result = append(result, Manufacturer(m))
	}

	return result
}
