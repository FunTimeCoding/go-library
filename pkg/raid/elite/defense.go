package elite

type Defense struct {
	DamageTaken       int `json:"damageTaken"`
	DownCount         int `json:"downCount"`
	DeadCount         int `json:"deadCount"`
	ConditionCleanses int `json:"conditionCleanses"`
	DodgeCount        int `json:"dodgeCount"`
	EvadedCount       int `json:"evadedCount"`
	BlockedCount      int `json:"blockedCount"`
	InvulnedCount     int `json:"invulnedCount"`
	InterruptedCount  int `json:"interruptedCount"`
}
