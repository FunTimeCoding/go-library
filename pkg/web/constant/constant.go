package constant

import "net/http"

const (
	Success = "success"
	Unauthorized = "unauthorized"

	RequestStart = "request_start"
)

var OkayStatusCodes = []int{http.StatusOK, http.StatusCreated}
