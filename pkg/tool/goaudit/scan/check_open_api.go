package scan

import "github.com/funtimecoding/go-library/pkg/system/virtual_file_system"

func (s *Service) checkOpenAPI(
	v *virtual_file_system.System,
	path string,
) {
	c := openAPIConcern(v, path, s.Name)

	if c != nil {
		s.Concerns = append(s.Concerns, c)
	}
}
