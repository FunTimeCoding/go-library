package defer_close

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "defer_close",
	Doc:  "flags defer x.Close() where x implements io.Closer; use defer errors.PanicClose(x)",
	Run:  run,
}
