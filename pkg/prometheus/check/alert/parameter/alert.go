package parameter

type Alert struct {
	All      bool
	Notation bool

	Critical bool
	Warning  bool

	Extended   bool
	Old        bool
	Suppressed bool

	Rules  bool
	Firing bool
}
