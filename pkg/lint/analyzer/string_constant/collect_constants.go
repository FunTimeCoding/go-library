package string_constant

func collectConstants(packageDirectory string) map[string]knownConstant {
	result := make(map[string]knownConstant)
	collectFromConstantFile(result, packageDirectory, "", 0)
	collectFromConstantDirectory(result, packageDirectory, "constant", 0)
	collectFromParents(result, packageDirectory)

	return result
}
