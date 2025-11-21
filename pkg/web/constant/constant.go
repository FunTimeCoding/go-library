package constant

import "net/http"

const (
	Success = "success"

	RequestStart = "request_start"
)

var OkayStatusCodes = []int{http.StatusOK, http.StatusCreated}
