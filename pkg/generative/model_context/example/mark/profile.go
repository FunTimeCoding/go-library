package mark

import "fmt"

func profile(identifier string) (string, error) {
	if identifier == "123" {
		return `{"id": "123", "name": "John Doe", "email": "john@gmail.com"}`, nil
	}

	return "", fmt.Errorf("user not found")
}
