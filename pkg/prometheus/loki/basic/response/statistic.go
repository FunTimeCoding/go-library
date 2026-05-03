package response

type Statistic struct {
	Summary  Summary  `json:"summary"`
	Querier  Querier  `json:"querier"`
	Ingester Ingester `json:"ingester"`
}
