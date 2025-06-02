package gobuild

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/build"
	"github.com/funtimecoding/go-library/pkg/build/option"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

func Main() {
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
	pflag.String(
		argument.BuildTags,
		"",
		"Build tags",
	)
	pflag.Bool(
		build.CopyToBinFlag,
		false,
		"Copy to $HOME/bin",
	)
	pflag.Bool(argument.All, false, "Build all")
	pflag.Bool(constant.LinuxAMD64, false, "Linux AMD64")
	pflag.Bool(constant.DarwinARM64, false, "Darwin ARM64")
	pflag.Bool(constant.DarwinAMD64, false, "Darwin AMD64")
	argument.ParseBind()
	linuxAMD64 := viper.GetBool(constant.LinuxAMD64)
	darwinARM64 := viper.GetBool(constant.DarwinARM64)
	darwinAMD64 := viper.GetBool(constant.DarwinAMD64)

	if !linuxAMD64 && !darwinARM64 && !darwinAMD64 {
		linuxAMD64 = true
		darwinARM64 = true
		darwinAMD64 = true
	}

	if viper.GetBool(argument.All) {
		for _, name := range system.Directories(
			system.Join(system.WorkingDirectory(), constant.CommandPath),
		) {
			if name == build.ExamplePath {
				continue
			}

			fmt.Printf("Build %s\n", name)
			mainPath := build.GuessMainPath(name)

			if mainPath == "" {
				log.Panicf("could not find main.go for %s", name)
			}

			o := option.New()
			o.Name = name
			o.MainPath = mainPath
			o.Output = viper.GetString(argument.Output)
			o.BuildTags = viper.GetString(argument.BuildTags)
			o.CopyToBin = viper.GetBool(build.CopyToBinFlag)
			o.LinuxAMD64 = linuxAMD64
			o.DarwinARM64 = darwinARM64
			o.DarwinAMD64 = darwinAMD64
			build.Architectures(o)
		}

		return
	}

	name := argument.RequiredPositional(0, "NAME", 1)
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
	o.LinuxAMD64 = linuxAMD64
	o.DarwinARM64 = darwinARM64
	o.DarwinAMD64 = darwinAMD64
	build.Architectures(o)
}
