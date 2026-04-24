package file_identity

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestClean(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "clean")
}

func TestMultipleIdentities(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "multi")
}

func TestMismatch(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "mismatch")
}

func TestMethods(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "methods")
}

func TestInterfaceMethod(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "iface")
}
