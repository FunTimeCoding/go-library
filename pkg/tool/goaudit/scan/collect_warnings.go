package scan

import "github.com/funtimecoding/go-library/pkg/system/virtual_file_system"

func (s *Service) collectWarnings(
	v *virtual_file_system.System,
	path string,
) {
	s.checkStaleDirectories(v, path)
	s.checkModelContext(v, path)
	s.checkTopLevelArgument(v, path)
	s.checkTopLevelResponse(v, path)
	s.checkConstantFile(path)
	s.checkOption(path)
	s.checkRun(path)
	s.checkSuffix(path)
	s.checkIdentity(v, path)
	s.checkOpenAPI(v, path)
}
