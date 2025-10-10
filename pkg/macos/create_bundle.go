package macos

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"strings"
)

func CreateBundle(
	name string,
	path string,
	executablePath string,
	iconPath string,
	vendor string,
	version string,
) {
	lowerName := strings.ToLower(name)
	identifier := fmt.Sprintf("%s.%s", vendor, lowerName)
	fullName := fmt.Sprintf("%s.app", name)
	base := join.Absolute(system.WorkingDirectory(), path, fullName)
	contents := join.Absolute(base, "Contents")
	macos := join.Absolute(contents, "MacOS")
	resources := join.Absolute(contents, "Resources")
	executable := join.Absolute(macos, lowerName)
	system.MakeDirectory(macos)
	system.MakeDirectory(resources)
	system.CopyFile(executablePath, executable)
	system.Executable(executable)
	system.CopyFile(iconPath, join.Absolute(resources, "icon.icns"))
	info := fmt.Sprintf(
		`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" 
    "https://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleName</key>
    <string>%s</string>

    <key>CFBundleDisplayName</key>
    <string>%sDisplayName</string>

    <key>CFBundleIdentifier</key>
    <string>%s</string>

    <key>CFBundleVersion</key>
    <string>%s</string>

    <key>CFBundleShortVersionString</key>
    <string>%s</string>

    <key>CFBundleExecutable</key>
    <string>%s</string>

    <key>CFBundlePackageType</key>
    <string>APPL</string>

    <key>LSUIElement</key>
    <true/>

    <key>CFBundleIconFile</key>
    <string>icon</string>
</dict>
</plist>
`,
		name,
		name,
		identifier,
		version,
		version,
		lowerName,
	)
	system.SaveFile(join.Absolute(contents, "Info.plist"), info)
}
