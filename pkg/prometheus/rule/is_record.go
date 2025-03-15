package rule

func (r *Rule) IsRecord() bool {
	return r.RawRecord != nil
}
