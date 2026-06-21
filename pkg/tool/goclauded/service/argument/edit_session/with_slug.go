package edit_session

func (o *Session) WithSlug(value string) *Session {
	o.Slug = new(value)

	return o
}
