package post

func (p *Post) formatUser() string {
	if p.User == nil {
		return "unknown"
	}

	return p.User.Username
}
