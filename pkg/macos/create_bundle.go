package macos

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
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
	base := system.Join(system.WorkingDirectory(), path, fullName)
	contents := system.Join(base, "Contents")
	macos := system.Join(contents, "MacOS")
	resources := system.Join(contents, "Resources")
	executable := system.Join(macos, lowerName)
	system.MakeDirectory(macos)
	system.MakeDirectory(resources)
	system.CopyFile(executablePath, executable)
	system.Executable(executable)
	system.CopyFile(iconPath, system.Join(resources, "icon.icns"))
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
	system.SaveFile(system.Join(contents, "Info.plist"), info)
}
