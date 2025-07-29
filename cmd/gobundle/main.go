package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/macos"
)

func main() {
	argument.ParseBind()
	name := argument.RequiredPositional(0, "NAME")
	path := argument.RequiredPositional(1, "PATH")
	executable := argument.RequiredPositional(2, "EXECUTABLE")
	icon := argument.RequiredPositional(3, "ICON")
	vendor := argument.RequiredPositional(4, "VENDOR")
	version := argument.RequiredPositional(5, "VERSION")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Executable: %s\n", executable)
	fmt.Printf("Icon: %s\n", icon)
	fmt.Printf("Vendor: %s\n", vendor)
	fmt.Printf("Version: %s\n", version)
	macos.CreateBundle(name, path, executable, icon, vendor, version)
}
