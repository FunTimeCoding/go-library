package gobuild

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/build"
	"github.com/funtimecoding/go-library/pkg/build/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	pflag.String(
		argument.Main,
		"",
		"Path to main.go, defaults to cmd/$NAME/main.go",
	)
	pflag.String(
		argument.Output,
		"",
		"Output path, defaults to tmp/$NAME/$OS-$ARCH/$NAME",
	)
	pflag.String(argument.BuildTags, "", "Build tags")
	pflag.Bool(build.CopyToBinFlag, false, "Copy to $HOME/bin")
	pflag.Bool(build.CgoFlag, false, "Enable CGO")
	pflag.Bool(constant.LinuxAMD64, false, "Linux AMD64")
	pflag.Bool(constant.DarwinARM64, false, "Darwin ARM64")
	pflag.Bool(constant.DarwinAMD64, false, "Darwin AMD64")
	monitor.ParseBind(version, gitHash, buildDate)
	linuxAMD64 := viper.GetBool(constant.LinuxAMD64)
	darwinARM64 := viper.GetBool(constant.DarwinARM64)
	darwinAMD64 := viper.GetBool(constant.DarwinAMD64)

	if !linuxAMD64 && !darwinARM64 && !darwinAMD64 {
		linuxAMD64 = true
		darwinARM64 = true
		darwinAMD64 = true
	}

	name := argument.Positional(0)

	if name == "" {
		for _, n := range system.Directories(
			system.Join(system.WorkingDirectory(), constant.CommandPath),
		) {
			if n == build.ExamplePath {
				continue
			}

			fmt.Printf("Build %s\n", n)
			mainPath := build.GuessMainPath(n)

			if mainPath == "" {
				log.Panicf("could not find main.go for %s", n)
			}

			o := option.New()
			o.Name = n
			o.MainPath = mainPath
			o.Output = viper.GetString(argument.Output)
			o.BuildTags = viper.GetString(argument.BuildTags)
			o.CopyToBin = viper.GetBool(build.CopyToBinFlag)
			o.Cgo = viper.GetBool(build.CgoFlag)
			o.LinuxAMD64 = linuxAMD64
			o.DarwinARM64 = darwinARM64
			o.DarwinAMD64 = darwinAMD64
			build.Architectures(o)
		}

		return
	}

	mainPath := viper.GetString(argument.Main)

	if mainPath == "" {
		if mainPath = build.GuessMainPath(name); mainPath == "" {
			log.Panicf("could not find main.go for %s", name)
		}
	}

	o := option.New()
	o.Name = name
	o.MainPath = mainPath
	o.Output = viper.GetString(argument.Output)
	o.BuildTags = viper.GetString(argument.BuildTags)
	o.CopyToBin = viper.GetBool(build.CopyToBinFlag)
	o.Cgo = viper.GetBool(build.CgoFlag)
	o.LinuxAMD64 = linuxAMD64
	o.DarwinARM64 = darwinARM64
	o.DarwinAMD64 = darwinAMD64
	build.Architectures(o)
}
