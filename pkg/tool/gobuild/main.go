package gobuild

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/build"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main() {
	pflag.String(argument.Main, "", "Path to main.go")
	pflag.StringP(
		argument.Output,
		"o",
		"",
		"Output path",
	)
	pflag.StringP(
		argument.BuildTags,
		"t",
		"",
		"Build tags",
	)
	pflag.BoolP(
		build.CopyToBinFlag,
		"b",
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

	name := argument.RequiredPositional(0, "NAME", 1)
	mainPath := viper.GetString(argument.Main)

	if mainPath == "" {
		if mainPath = build.GuessMainPath(name); mainPath == "" {
			argument.RequiredStringFlag(argument.Main, 1)
		}
	}

	output := viper.GetString(argument.Output)
	buildTags := viper.GetString(argument.BuildTags)
	copyToBin := viper.GetBool(build.CopyToBinFlag)

	if linuxAMD64 {
		build.Go(
			name,
			mainPath,
			output,
			buildTags,
			copyToBin,
			constant.Linux,
			constant.AMD64,
		)
	}

	if darwinARM64 {
		build.Go(
			name,
			mainPath,
			output,
			buildTags,
			copyToBin,
			constant.Darwin,
			constant.ARM64,
		)
	}

	if darwinAMD64 {
		build.Go(
			name,
			mainPath,
			output,
			buildTags,
			copyToBin,
			constant.Darwin,
			constant.AMD64,
		)
	}
}
