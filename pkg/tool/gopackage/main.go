package gopackage

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/build"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	system "github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"github.com/funtimecoding/go-library/pkg/tool/gopackage/constant"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	monitor.ParseBind(version, gitHash, buildDate)
	var runs int

	for _, name := range build.OutputDirectories() {
		fmt.Printf("Name: %s\n", name)
		outputDirectory := join.Relative(system.Temporary, name)
		fmt.Printf("Output directory: %s\n", outputDirectory)

		for _, systemArchitecture := range build.SystemArchitectures() {
			if build.GuessBinaryPath(name, systemArchitecture) != "" {
				build.Archive(name, systemArchitecture)
				runs++
			}
		}
	}

	if runs == 0 {
		fmt.Println("No archive created")
		os.Exit(1)
	}
}
