package status

func New(color bool) *Status {
	return &Status{color: color}
}
