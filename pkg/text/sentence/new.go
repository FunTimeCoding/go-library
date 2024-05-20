package sentence

func New(action string) *Sentence {
	return &Sentence{action: action}
}
