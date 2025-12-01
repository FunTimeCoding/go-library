package message

import "github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/response"

type Meta struct {
	Type      string
	Statistic response.Statistic
}
