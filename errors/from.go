package errors

func From(a any) error {
	switch cast := a.(type) {
	case error:
		return cast
	}

	return nil
}
