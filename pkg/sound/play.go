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
	verbose bool,
) {
	switch p := runtime.GOOS; p {
	case constant.Linux:
		panic("not implemented")
	case constant.Darwin:
		result, e := system.RunError(
			Afplay,
			VolumeArgument,
			fmt.Sprintf("%.2f", volume),
			path,
		)

		if e != nil && verbose {
			fmt.Printf("Sound error: %s\n", e)
			fmt.Printf("Output: %s\n", result)
		}
	default:
		unexpected.String(p)
	}
}
