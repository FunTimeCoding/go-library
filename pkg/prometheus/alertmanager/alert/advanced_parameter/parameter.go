package advanced_parameter

type Parameter struct {
	All          bool
	CriticalOnly bool
	WarningOnly  bool
	Old          bool
	Suppressed   bool
	Receiver     []string
}
