package output

func New() *Settings {
	return &Settings{Format: Text, ShowDebug: false}
}
