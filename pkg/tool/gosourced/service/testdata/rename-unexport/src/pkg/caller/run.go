package caller

import "example/pkg/target"

func Run() bool {
	return target.Validate("test")
}
