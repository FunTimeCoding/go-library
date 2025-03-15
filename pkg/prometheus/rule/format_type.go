package rule

func (r *Rule) formatType() string {
	if r.RawAlert != nil {
		return AlertType
	}

	if r.RawRecord != nil {
		return RecordType
	}

	return UnknownType
}
