package scan

import "github.com/funtimecoding/go-library/pkg/system/virtual_file_system"

func (s *Service) checkIdentity(
	v *virtual_file_system.System,
	path string,
) {
	if !s.ConstantDirectory {
		return
	}

	s.Concerns = append(s.Concerns, identityConcerns(v, path, s.Name)...)
}
