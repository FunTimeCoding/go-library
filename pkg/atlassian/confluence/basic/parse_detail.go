package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/web/detail_error"
)

func parseDetail(
	body []byte,
	status string,
) error {
	var e response.Error

	if json.Unmarshal(body, &e) == nil && e.Message != "" {
		return detail_error.New(e.Message, status)
	}

	return fmt.Errorf("%s", status)
}
