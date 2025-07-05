package repository

func (r *Repository) HasConcerns() bool {
	return len(r.concern) > 0
}
