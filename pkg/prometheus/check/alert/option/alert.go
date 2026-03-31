package option

type Alert struct {
	Notation    bool
	All         bool
	Critical    bool
	Warning     bool
	Information bool
	Extended    bool
	Suppressed  bool
	Rules       bool
	Firing      bool
	Fingerprint bool
	Copyable    bool
}
