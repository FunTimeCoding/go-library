package advanced_option

type Alert struct {
	All             bool
	CriticalOnly    bool
	WarningOnly     bool
	InformationOnly bool
	Suppressed      bool
	Receiver        []string
}
