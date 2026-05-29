package certificate_detail

func New(
	certificate map[string]interface{},
	latestRequest map[string]interface{},
	filtered []string,
) *Detail {
	return &Detail{
		Certificate:   certificate,
		LatestRequest: latestRequest,
		Filtered:      filtered,
	}
}
