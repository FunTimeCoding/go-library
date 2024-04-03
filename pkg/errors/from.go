package errors

import "fmt"

func From(a any) error {
	switch cast := a.(type) {
	case nil:
		return nil
	case error:
		return cast
	default:
		v := fmt.Sprintf("%v", a)

		if v != "" {
			return fmt.Errorf("non-error: %s", v)
		}
	}

	return nil
}
