package response

type Data struct {
	ResultType string    `json:"resultType"`
	Result     []Result  `json:"result"`
	Stats      Statistic `json:"stats"`
}
