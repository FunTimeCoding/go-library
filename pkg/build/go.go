package build

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"path/filepath"
	"runtime"
)

func Go(
	name string,
	mainPath string,
	output string,
	buildTags string,
	copyToBin bool,
	operatingSystem string,
	architecture string,
) {
	if output == "" {
		if name == "" {
			panic("output empty and main not specified")
		} else {
			output = system.Join(
				systemConstant.Temporary,
				name,
				SystemArchitecture(operatingSystem, architecture),
				name,
			)
		}
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Output: %s\n", output)
	path := filepath.Dir(output)
	fmt.Printf("Path: %s\n", path)
	system.EnsurePathExists(path)

	s := []string{
		constant.Go,
		constant.Build,
		constant.LinkerFlagsArgument,
		join.Space(
			constant.LinkerSetVariable,
			fmt.Sprintf("main.Version=%s", GitTag()),
			constant.LinkerSetVariable,
			fmt.Sprintf("main.GitHash=%s", GitHash()),
			constant.LinkerSetVariable,
			fmt.Sprintf("main.BuildDate=%s", Date()),
		),
	}

	if buildTags != "" {
		s = append(s, constant.TagsArgument, buildTags)
	}

	s = append(s, []string{constant.OutputArgument, output, mainPath}...)

	r := run.New()
	r.Verbose = true
	r.Panic = false
	r.Environment(constant.NativeEnabled, constant.False)
	r.Environment(constant.System, operatingSystem)
	r.Environment(constant.Architecture, architecture)

	if t := r.Start(s...); t != "" {
		fmt.Printf("Output:\n%s", t)
	}

	errors.PanicOnError(r.Error)

	if copyToBin &&
		runtime.GOOS == operatingSystem &&
		runtime.GOARCH == architecture {
		destination := system.Join(system.Home(), systemConstant.Binary, name)
		fmt.Printf("Source: %s\n", output)
		fmt.Printf("Destination: %s\n", destination)
		system.CopyFile(output, destination)
		system.Executable(destination)
	}

	fmt.Printf("Size: %dM\n", system.FileSize(output)/1024/1024)
}
