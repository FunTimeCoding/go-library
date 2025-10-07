package debian

import (
	"github.com/funtimecoding/go-library/pkg/debian/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func PrepareNetBoot(
	inputArchive string,
	outputArchive string,
	configurationPath string,
	temporaryDirectory string,
) {
	system.ExtractTarZip(inputArchive, temporaryDirectory)
	errors.PanicOnError(
		os.Rename(
			configurationPath,
			system.Join(temporaryDirectory, constant.PreseedConfiguration),
		),
	)
	system.CreateTarZip(temporaryDirectory, outputArchive)
}
