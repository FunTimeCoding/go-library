package detail_error

func New(
	detail string,
	status string,
) *Detail {
	return &Detail{Detail: detail, Status: status}
}
