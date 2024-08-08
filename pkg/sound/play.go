package sound

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"runtime"
)

func Play(
	path string,
	volume float64,
) {
	switch p := runtime.GOOS; p {
	case constant.Linux:
		panic("not implemented")
	case constant.Darwin:
		system.Run(
			Afplay,
			VolumeArgument,
			fmt.Sprintf("%.2f", volume),
			path,
		)
	default:
		unexpected.String(p)
	}
}
