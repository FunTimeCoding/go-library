package response

type WheelResult struct {
	Success  bool          `json:"success"`
	Affected WheelAffected `json:"return"`
}
