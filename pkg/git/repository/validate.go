package repository

func (r *Repository) Validate() {
	if !r.IsClean {
		r.concern = append(r.concern, NotClean)
	}
}
