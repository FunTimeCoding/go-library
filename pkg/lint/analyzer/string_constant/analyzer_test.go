package string_constant

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestFlagged(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "flagged")
}

func TestClean(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "clean")
}

func TestNoConstant(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "no_constant")
}
