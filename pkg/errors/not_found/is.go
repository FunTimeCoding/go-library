package not_found

func (e *NotFoundError) Is(target error) bool {
	_, okay := target.(*NotFoundError)

	return okay
}
