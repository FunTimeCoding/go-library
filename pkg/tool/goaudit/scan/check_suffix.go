package scan

import "strings"

func (s *Service) checkSuffix(path string) {
	if !strings.HasSuffix(s.Name, "d") && s.hasCapability() {
		s.addConcern(MissingSuffixKey, MissingSuffixText, path)
	}
}
