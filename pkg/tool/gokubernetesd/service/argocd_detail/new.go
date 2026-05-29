package argocd_detail

func New(
	application map[string]interface{},
	filtered []string,
) *Detail {
	return &Detail{
		Application: application,
		Filtered:    filtered,
	}
}
