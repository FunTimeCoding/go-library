package store

func (s *Store) Find(id uint) (*TerraformRun, error) {
	var result TerraformRun

	return &result, s.mapper.First(&result, id).Error
}
