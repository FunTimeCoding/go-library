package advanced_option

type Alert struct {
	All          bool
	CriticalOnly bool
	WarningOnly  bool
	Old          bool
	Suppressed   bool
	Receiver     []string
}
