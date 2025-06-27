package option

type Alert struct {
	Notation bool
	All      bool

	Critical bool
	Warning  bool

	Extended   bool
	Old        bool
	Suppressed bool

	Rules  bool
	Firing bool

	Fingerprint bool
}
