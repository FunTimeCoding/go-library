package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Printf(
			"Usage: %s EXECUTABLE VERSION MAINTAINER_NAME MAINTAINER_EMAIL"+
				"\n Version: 1.0"+
				"\n Maintainer name: AN Other"+
				"\n Maintainer email: another@example.org"+
				"\n",
			os.Args[0],
		)

		os.Exit(1)
	}

	executable := os.Args[1]
	version := os.Args[2]
	maintainerName := os.Args[3]
	maintainerMail := os.Args[4]

	work := system.WorkingDirectory()
	packageName := fmt.Sprintf("%s_%s-1_amd64", executable, version)
	root := system.Join(work, packageName)

	debian := system.Join(root, "DEBIAN")
	system.MakeDirectory(debian)
	control := system.Join(debian, "control")
	system.SaveFile(
		control,
		fmt.Sprintf(
			"Package: %s"+
				"\nVersion: %s"+
				"\nArchitecture: amd64"+
				"\nMaintainer: %s <%s>"+
				"\nDescription: Short stub description."+
				"\n Long stub description."+
				"\n",
			executable,
			version,
			maintainerName,
			maintainerMail,
		),
	)

	bin := system.Join(root, "usr", "local", "bin")
	system.Move(executable, system.Join(bin, executable))

	fmt.Println(
		system.Run(
			"dpkg-deb",
			"--build",
			"--root-owner-group",
			packageName,
		),
	)
}
