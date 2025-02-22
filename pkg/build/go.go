package build

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Go(
	main string,
	output string,
	buildTags string,
	copyToBin bool,
) {
	var name string

	if main == "" {
		main = "."
	} else {
		name = filepath.Base(filepath.Dir(main))
		fmt.Printf("Name: %s\n", name)
	}

	if output == "" {
		if name == "" {
			panic("output empty and main not specified")
		} else {
			output = system.Join(systemConstant.Temporary, name)
		}
	}

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

	s = append(s, []string{constant.OutputArgument, output, main}...)

	fmt.Printf("Command: %s\n", join.Space(s...))
	c := exec.Command(s[0], s[1:]...)
	c.Env = os.Environ()
	c.Env = append(
		c.Env,
		key_value.Equals(constant.NativeEnabled, constant.False),
	)
	c.Env = append(
		c.Env,
		key_value.Equals(constant.System, runtime.GOOS),
	)
	combined, e := c.CombinedOutput()

	if t := string(combined); t != "" {
		fmt.Printf("Output:\n%s", t)
	}

	errors.PanicOnError(e)

	if copyToBin {
		destination := system.Join(system.Home(), systemConstant.Binary, name)
		fmt.Printf("Source: %s\n", output)
		fmt.Printf("Destination: %s\n", destination)
		system.CopyFile(output, destination)
		system.Executable(destination)
	}

	fmt.Printf("Size: %dM\n", system.FileSize(output)/1024/1024)
}
