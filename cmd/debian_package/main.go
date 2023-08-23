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
				"\nVersion: 1.0.0"+
				"\nMaintainer name: AN Other"+
				"\nMaintainer email: another@example.org"+
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
	packageDirectory := system.Join(work, packageName)

	debian := system.Join(packageDirectory, "DEBIAN")
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

	unit := system.Join(packageDirectory, "lib", "systemd", "system")
	system.MakeDirectory(unit)
	system.SaveFile(
		system.Join(unit, fmt.Sprintf("%s.service", executable)),
		fmt.Sprintf(
			"[Unit]"+
				"\nDescription=%s stub description"+
				"\nAfter=network.target"+
				"\n"+
				"\n[Service]"+
				"\nType=simple"+
				"\nExecStart=/usr/local/bin/%s"+
				"\n"+
				"\n[Install]"+
				"\nWantedBy=multi-user.target"+
				"\n",
			executable,
			executable,
		),
	)

	bin := system.Join(packageDirectory, "usr", "local", "bin")
	system.MakeDirectory(bin)
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
