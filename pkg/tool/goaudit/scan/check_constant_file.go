package scan

func (s *Service) checkConstantFile(path string) {
	if s.ConstantFile {
		s.addConcern(ConstantFileKey, ConstantFileText, path)
	}
}
