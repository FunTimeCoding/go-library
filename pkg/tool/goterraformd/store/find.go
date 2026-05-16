package store

func (s *Store) Find(identifier uint) (*TerraformRun, error) {
	var result TerraformRun

	return &result, s.mapper.First(&result, identifier).Error
}
