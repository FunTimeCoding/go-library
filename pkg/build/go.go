package build

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/build/option"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"path/filepath"
	"runtime"
)

func Go(o *option.Build) {
	if o.Output == "" {
		if o.Name == "" {
			panic("output empty and main not specified")
		} else {
			o.Output = system.Join(
				systemConstant.Temporary,
				o.Name,
				SystemArchitecture(o.OperatingSystem, o.Architecture),
				o.Name,
			)
		}
	}

	fmt.Printf("Name: %s\n", o.Name)
	fmt.Printf("Output: %s\n", o.Output)
	path := filepath.Dir(o.Output)
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

	if o.BuildTags != "" {
		s = append(s, constant.TagsArgument, o.BuildTags)
	}

	s = append(s, []string{constant.OutputArgument, o.Output, o.MainPath}...)

	r := run.New()
	r.Verbose = true
	r.Panic = false
	r.Environment(constant.NativeEnabled, constant.False)
	r.Environment(constant.System, o.OperatingSystem)
	r.Environment(constant.Architecture, o.Architecture)

	if t := r.Start(s...); t != "" {
		fmt.Printf("Output:\n%s", t)
	}

	errors.PanicOnError(r.Error)

	if o.CopyToBin &&
		runtime.GOOS == o.OperatingSystem &&
		runtime.GOARCH == o.Architecture {
		destination := system.Join(
			system.Home(),
			systemConstant.Binary,
			o.Name,
		)
		fmt.Printf("Source: %s\n", o.Output)
		fmt.Printf("Destination: %s\n", destination)
		system.CopyFile(o.Output, destination)
		system.Executable(destination)
	}

	fmt.Printf("Size: %dM\n", system.FileSize(o.Output)/1024/1024)
}
