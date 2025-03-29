package gobuild

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/build"
	"github.com/funtimecoding/go-library/pkg/build/option"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
	pflag.Bool(constant.LinuxAMD64, false, "Linux AMD64")
	pflag.Bool(constant.DarwinARM64, false, "Darwin ARM64")
	pflag.Bool(constant.DarwinAMD64, false, "Darwin AMD64")
	argument.ParseAndBind()
	linuxAMD64 := viper.GetBool(constant.LinuxAMD64)
	darwinARM64 := viper.GetBool(constant.DarwinARM64)
	darwinAMD64 := viper.GetBool(constant.DarwinAMD64)

	if !linuxAMD64 && !darwinARM64 && !darwinAMD64 {
		linuxAMD64 = true
		darwinARM64 = true
		darwinAMD64 = true
	}

	name := argument.RequiredPositional(
		0,
		"NAME",
		1,
	)
	mainPath := viper.GetString(argument.Main)

	if mainPath == "" {
		if mainPath = build.GuessMainPath(name); mainPath == "" {
			argument.RequiredStringFlag(argument.Main, 1)
		}
	}

	o := option.New()
	o.Name = name
	o.MainPath = mainPath
	o.Output = viper.GetString(argument.Output)
	o.BuildTags = viper.GetString(argument.BuildTags)
	o.CopyToBin = viper.GetBool(build.CopyToBinFlag)

	if linuxAMD64 {
		o.OperatingSystem = constant.Linux
		o.Architecture = constant.AMD64
		build.Go(o)
	}

	if darwinARM64 {
		o.OperatingSystem = constant.Darwin
		o.Architecture = constant.ARM64
		build.Go(o)
	}

	if darwinAMD64 {
		o.OperatingSystem = constant.Darwin
		o.Architecture = constant.AMD64
		build.Go(o)
	}
}
