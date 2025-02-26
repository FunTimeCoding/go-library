package command

import "github.com/funtimecoding/go-library/pkg/linux/systemd/constant"

func Failed() []string {
	return []string{
		constant.Command,
		constant.ListUnits,
		constant.Output,
		constant.Notation,
		constant.State,
		constant.FailedState,
	}
}
