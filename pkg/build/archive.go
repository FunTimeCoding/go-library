package build

import (
	"archive/zip"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func Archive(name string, systemArchitecture string) {
	sourceFile := GuessBinaryPath(name, systemArchitecture)
	fmt.Printf("Source file: %s\n", sourceFile)

	archiveName := fmt.Sprintf("%s-%s.zip", name, systemArchitecture)
	fmt.Printf("Archive name: %s\n", archiveName)

	archive := system.Create(join.Relative(constant.Temporary, archiveName))
	defer errors.PanicClose(archive)
	w := zip.NewWriter(archive)
	defer errors.PanicClose(w)
	f := system.Open(sourceFile)
	defer errors.PanicClose(f)
	system.ZipAdd(w, f, name)
}
