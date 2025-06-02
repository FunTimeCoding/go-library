package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/macos"
)

func main() {
	argument.ParseBind()
	name := argument.RequiredPositional(0, "NAME", 1)
	path := argument.RequiredPositional(1, "PATH", 1)
	executable := argument.RequiredPositional(2, "EXECUTABLE", 1)
	icon := argument.RequiredPositional(3, "ICON", 1)
	vendor := argument.RequiredPositional(4, "VENDOR", 1)
	version := argument.RequiredPositional(5, "VERSION", 1)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Executable: %s\n", executable)
	fmt.Printf("Icon: %s\n", icon)
	fmt.Printf("Vendor: %s\n", vendor)
	fmt.Printf("Version: %s\n", version)
	macos.CreateBundle(name, path, executable, icon, vendor, version)
}
