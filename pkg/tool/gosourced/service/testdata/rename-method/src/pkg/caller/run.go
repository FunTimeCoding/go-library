package caller

import "example/pkg/target"

func Run(s *target.Store) string {
	return s.FindByName("test")
}
