package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/web/subscription"
	"net/http"
	"strings"
)

func parseSubscribe(r *http.Request) subscription.Subscription {
	value := r.URL.Query().Get("subscribe")

	if value == "" {
		return subscription.Subscription{}
	}

	result := subscription.Subscription{}

	for _, s := range strings.Split(value, ",") {
		result[s] = true
	}

	return result
}
