package option

type Build struct {
	Name            string
	MainPath        string
	Output          string
	BuildTags       string
	CopyToBin       bool
	Cgo             bool
	OperatingSystem string
	Architecture    string

	LinuxAMD64  bool
	DarwinARM64 bool
	DarwinAMD64 bool
}
