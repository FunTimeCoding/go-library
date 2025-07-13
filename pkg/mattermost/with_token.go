package mattermost

func WithToken(token string) OptionFunc {
	return func(o *Option) {
		o.Token = token
	}
}
