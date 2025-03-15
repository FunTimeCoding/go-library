package rule

func (r *Rule) IsAlert() bool {
	return r.RawAlert != nil
}
