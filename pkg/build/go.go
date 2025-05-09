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
	p := *o

	if p.Output == "" {
		if p.Name == "" {
			panic("output empty and main not specified")
		} else {
			p.Output = system.Join(
				systemConstant.Temporary,
				p.Name,
				SystemArchitecture(p.OperatingSystem, p.Architecture),
				p.Name,
			)
		}
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Output: %s\n", p.Output)
	path := filepath.Dir(p.Output)
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

	if p.BuildTags != "" {
		s = append(s, constant.TagsArgument, p.BuildTags)
	}

	s = append(s, []string{constant.OutputArgument, p.Output, p.MainPath}...)

	r := run.New()
	r.Verbose = true
	r.Panic = false
	r.Environment(constant.NativeEnabled, constant.False)
	r.Environment(constant.System, p.OperatingSystem)
	r.Environment(constant.Architecture, p.Architecture)

	if t := r.Start(s...); t != "" {
		fmt.Printf("Output:\n%s", t)
	}

	errors.PanicOnError(r.Error)

	if p.CopyToBin &&
		runtime.GOOS == p.OperatingSystem &&
		runtime.GOARCH == p.Architecture {
		destination := system.Join(
			system.Home(),
			systemConstant.Binary,
			p.Name,
		)
		fmt.Printf("Source: %s\n", p.Output)
		fmt.Printf("Destination: %s\n", destination)
		system.CopyFile(p.Output, destination)
		system.Executable(destination)
	}

	fmt.Printf("Size: %dM\n", system.FileSize(p.Output)/1024/1024)
}
