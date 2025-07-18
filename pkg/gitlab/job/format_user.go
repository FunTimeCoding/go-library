package job

func (j *Job) formatUser() string {
	if j.Raw.User != nil {
		return j.Raw.User.Username
	}

	return NoUser
}
