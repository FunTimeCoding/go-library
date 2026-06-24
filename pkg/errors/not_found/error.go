package not_found

func (e *NotFoundError) Error() string { return e.Message }
