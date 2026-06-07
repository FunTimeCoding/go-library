package scan

import (
	"github.com/funtimecoding/go-library/pkg/system/virtual_file_system"
	"path/filepath"
)

func (s *Service) checkTopLevelArgument(
	v *virtual_file_system.System,
	path string,
) {
	if v.DirectoryExists(filepath.Join(path, "argument")) {
		s.addConcern(TopLevelArgumentKey, TopLevelArgumentText, path)
	}
}
