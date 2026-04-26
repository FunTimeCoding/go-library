package constant

import "net/http"

const (
	Success      = "success"
	Unauthorized = "unauthorized"

	RequestStart = "request_start"

	OriginAll = "*"

	FormMethod = "method"
)

var OkayStatusCodes = []int{http.StatusOK, http.StatusCreated}
