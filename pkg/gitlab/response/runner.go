package response

type Runner struct {
	Data struct {
		Runner struct {
			ID          string `json:"id"`
			Description string `json:"description"`
			Status      string `json:"status"`
			RunnerType  string `json:"runnerType"`
			Managers    struct {
				Nodes []struct {
					SystemID  string `json:"systemId"`
					IPAddress string `json:"ipAddress"`
					Version   string `json:"version"`
					Revision  string `json:"revision"`
				} `json:"nodes"`
			} `json:"managers"`
		} `json:"runner"`
	}
}
