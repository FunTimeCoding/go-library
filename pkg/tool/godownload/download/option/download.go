package option

type Download struct {
	Host           string
	Token          string
	Owner          string
	Repository     string
	PackageVersion string
	Output         string
	Verbose        bool

	Package string
}
