package constant

import "net/http"

const (
	Success      = "success"
	Unauthorized = "unauthorized"

	RequestStart = "request_start"

	OriginAll = "*"
)

var OkayStatusCodes = []int{http.StatusOK, http.StatusCreated}
