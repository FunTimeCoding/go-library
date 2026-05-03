package response

type Artifact struct {
	Uninstall           []Uninstall `json:"uninstall,omitempty"`
	App                 []any       `json:"app,omitempty"`
	Zap                 []Zap       `json:"zap,omitempty"`
	P                   []any       `json:"pkg,omitempty"`
	BashCompletion      []string    `json:"bash_completion,omitempty"`
	FishCompletion      []string    `json:"fish_completion,omitempty"`
	ZshCompletion       []string    `json:"zsh_completion,omitempty"`
	Binary              []any       `json:"binary,omitempty"`
	Postflight          any         `json:"postflight"`
	UninstallPostflight any         `json:"uninstall_postflight"`
	Preflight           any         `json:"preflight"`
	Font                []string    `json:"font,omitempty"`
	Manpage             []string    `json:"manpage,omitempty"`
	UninstallPreflight  any         `json:"uninstall_preflight"`
}
