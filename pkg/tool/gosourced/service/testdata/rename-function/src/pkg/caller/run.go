package caller

import "example/pkg/target"

func Run(content string) bool {
	return target.IsGeneratedHeader(content)
}
