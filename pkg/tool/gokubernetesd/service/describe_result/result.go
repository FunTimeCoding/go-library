package describe_result

import "github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/service/response"

type Result struct {
	Resource map[string]any
	Events   []response.DescribeEvent
	Filtered []string
}
