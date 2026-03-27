package token_check

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func rawSearch(
	host string,
	user string,
	token string,
	project string,
) {
	rawGet(
		locator.New(host).Path("rest/api/2/search/jql").Set(
			"jql",
			fmt.Sprintf("project = %s ORDER BY updated DESC", project),
		).Set("maxResults", "1").String(),
		user,
		token,
	)
}
