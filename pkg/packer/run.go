package packer

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/go-library/pkg/debian/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/packer/build"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"log"
	"runtime"
)

func (c *Client) Run(
	architecture string,
	imageName string,
	checksum string,
	v *semver.Version,
	preseed string,
	username string,
	password string,
	headless bool,
) {
	if headless {
		log.Panicf("headless not supported")
	}

	system.Move(
		preseed,
		system.Join(c.packerWebDirectory, constant.PreseedConfiguration),
	)
	// packer plugins install github.com/hashicorp/qemu
	b := build.New(architecture, 4444, username, password)
	b.SetBuilder(
		c.packerWebDirectory,
		c.packerOutputDirectory,
		headless,
		system.Join(c.workDirectory, imageName),
		checksum,
		v,
	)
	template := system.Join(c.workDirectory, "debian.json")
	system.SaveFile(
		template,
		fmt.Sprintf("%s\n", notation.Format(b)),
	)
	arguments := []string{
		"packer",
		"build",
		"-machine-readable",
		"-on-error",
		"cleanup",
		template,
	}
	fmt.Printf("Command: %s\n", join.Space(arguments...))

	if headless {
		var connectCommand string

		switch runtime.GOOS {
		case systemConstant.Linux:
			connectCommand = fmt.Sprintf(
				"nc -v 127.0.0.1 %d",
				b.CharacterDevicePort,
			)
		case systemConstant.Darwin:
			connectCommand = fmt.Sprintf(
				"socat -,raw,echo=0,escape=0x11 tcp:127.0.0.1:%d",
				b.CharacterDevicePort,
			)
		}

		// TODO: Typing happens via VNC, the character device does not reflect that
		//  Check VNC if typing actually happens
		//  Then proceed to actual unattended install

		if connectCommand != "" {
			fmt.Printf("Character device: %s\n", connectCommand)
		}
	}

	output, e := system.RunError(arguments...)

	if e != nil {
		fmt.Println("Error:")
	}

	fmt.Println(output)
}
