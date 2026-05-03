package response

type InstalledCask struct {
	Token                         string                   `json:"token"`
	FullToken                     string                   `json:"full_token"`
	OldTokens                     []string                 `json:"old_tokens"`
	Tap                           string                   `json:"tap"`
	Name                          []string                 `json:"name"`
	Desc                          *string                  `json:"desc"`
	Homepage                      string                   `json:"homepage"`
	Locator                       string                   `json:"url"`
	LocatorSpecification          CaskLocatorSpecification `json:"url_specs"`
	Version                       string                   `json:"version"`
	Autobump                      bool                     `json:"autobump"`
	NoAutobumpMessage             any                      `json:"no_autobump_message"`
	SkipLivecheck                 bool                     `json:"skip_livecheck"`
	Installed                     string                   `json:"installed"`
	InstalledTime                 int                      `json:"installed_time"`
	BundleVersion                 *string                  `json:"bundle_version"`
	BundleShortVersion            *string                  `json:"bundle_short_version"`
	Outdated                      bool                     `json:"outdated"`
	Sha256                        string                   `json:"sha256"`
	Artifacts                     []Artifact               `json:"artifacts"`
	Caveats                       *string                  `json:"caveats"`
	DependsOn                     CaskDependency           `json:"depends_on"`
	ConflictsWith                 *CaskConflicts           `json:"conflicts_with"`
	Container                     any                      `json:"container"`
	Rename                        []any                    `json:"rename"`
	AutoUpdates                   *bool                    `json:"auto_updates"`
	Deprecated                    bool                     `json:"deprecated"`
	DeprecationDate               any                      `json:"deprecation_date"`
	DeprecationReason             any                      `json:"deprecation_reason"`
	DeprecationReplacementFormula any                      `json:"deprecation_replacement_formula"`
	DeprecationReplacementCask    any                      `json:"deprecation_replacement_cask"`
	Disabled                      bool                     `json:"disabled"`
	DisableDate                   any                      `json:"disable_date"`
	DisableReason                 any                      `json:"disable_reason"`
	DisableReplacementFormula     any                      `json:"disable_replacement_formula"`
	DisableReplacementCask        any                      `json:"disable_replacement_cask"`
	TapGitHead                    string                   `json:"tap_git_head"`
	Languages                     []string                 `json:"languages"`
	RubySourcePath                string                   `json:"ruby_source_path"`
	RubySourceChecksum            RubySourceChecksum       `json:"ruby_source_checksum"`
}
