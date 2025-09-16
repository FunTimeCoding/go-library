package main

import "github.com/funtimecoding/go-library/pkg/defect_dojo"

func main() {
	// https://github.com/truemilk/go-defectdojo
	//  Fork: https://github.com/blackaichi/go-defectdojo
	// https://github.com/egongu90/defectdojo-go
	// Make own openapi-generated client or build basic client?
	//  https://docs.defectdojo.com/en/api/api-v2-docs/
	defect_dojo.NewEnvironment()
}
