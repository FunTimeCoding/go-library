package response

type OS struct {
	Cellar string `json:"cellar"`
	Url    string `json:"url"`
	Sha256 string `json:"sha256"`
}

type Installed struct {
	Formulae []struct {
		Name              string   `json:"name"`
		FullName          string   `json:"full_name"`
		Tap               string   `json:"tap"`
		OldNames          []string `json:"oldnames"`
		Aliases           []string `json:"aliases"`
		VersionedFormulae []string `json:"versioned_formulae"`
		Desc              string   `json:"desc"`
		License           *string  `json:"license"`
		Homepage          string   `json:"homepage"`
		Versions          struct {
			Stable string  `json:"stable"`
			Head   *string `json:"head"`
			Bottle bool    `json:"bottle"`
		} `json:"versions"`
		Urls struct {
			Stable struct {
				Url      string  `json:"url"`
				Tag      *string `json:"tag"`
				Revision *string `json:"revision"`
				Using    *string `json:"using"`
				Checksum *string `json:"checksum"`
			} `json:"stable"`
			Head struct {
				Url    string  `json:"url"`
				Branch *string `json:"branch"`
				Using  *string `json:"using"`
			} `json:"head,omitempty"`
		} `json:"urls"`
		Revision          int         `json:"revision"`
		VersionScheme     int         `json:"version_scheme"`
		Autobump          bool        `json:"autobump"`
		NoAutobumpMessage interface{} `json:"no_autobump_message"`
		SkipLivecheck     bool        `json:"skip_livecheck"`
		Bottle            struct {
			Stable struct {
				Rebuild int    `json:"rebuild"`
				RootUrl string `json:"root_url"`
				Files   struct {
					Arm64Sequoia struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"arm64_sequoia,omitempty"`
					Arm64Sonoma struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"arm64_sonoma,omitempty"`
					Arm64Ventura struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"arm64_ventura,omitempty"`
					Sonoma struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"sonoma,omitempty"`
					Ventura struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"ventura,omitempty"`
					Arm64Linux struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"arm64_linux,omitempty"`
					X8664Linux struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"x86_64_linux,omitempty"`
					All struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"all,omitempty"`
					Arm64Monterey struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"arm64_monterey,omitempty"`
					Arm64BigSur struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"arm64_big_sur,omitempty"`
					Monterey struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"monterey,omitempty"`
					BigSur struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"big_sur,omitempty"`
					Catalina struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"catalina,omitempty"`
					Arm64Tahoe struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"arm64_tahoe,omitempty"`
					Sequoia struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"sequoia,omitempty"`
					Mojave struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"mojave,omitempty"`
					HighSierra struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"high_sierra,omitempty"`
					Sierra struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"sierra,omitempty"`
					ElCapitan struct {
						Cellar string `json:"cellar"`
						Url    string `json:"url"`
						Sha256 string `json:"sha256"`
					} `json:"el_capitan,omitempty"`
				} `json:"files"`
			} `json:"stable,omitempty"`
		} `json:"bottle"`
		PourBottleOnlyIf *string `json:"pour_bottle_only_if"`
		KegOnly          bool    `json:"keg_only"`
		KegOnlyReason    *struct {
			Reason      string `json:"reason"`
			Explanation string `json:"explanation"`
		} `json:"keg_only_reason"`
		Options                 []interface{} `json:"options"`
		BuildDependencies       []string      `json:"build_dependencies"`
		Dependencies            []string      `json:"dependencies"`
		TestDependencies        []string      `json:"test_dependencies"`
		RecommendedDependencies []interface{} `json:"recommended_dependencies"`
		OptionalDependencies    []interface{} `json:"optional_dependencies"`
		UsesFromMacos           []interface{} `json:"uses_from_macos"`
		UsesFromMacosBounds     []struct {
			Since string `json:"since,omitempty"`
		} `json:"uses_from_macos_bounds"`
		Requirements []struct {
			Name     string   `json:"name"`
			Cask     any      `json:"cask"`
			Download any      `json:"download"`
			Version  *string  `json:"version"`
			Contexts []string `json:"contexts"`
			Specs    []string `json:"specs"`
		} `json:"requirements"`
		ConflictsWith        []string `json:"conflicts_with"`
		ConflictsWithReasons []string `json:"conflicts_with_reasons"`
		LinkOverwrite        []string `json:"link_overwrite"`
		Caveats              *string  `json:"caveats"`
		Installed            []struct {
			Version             string `json:"version"`
			UsedOptions         []any  `json:"used_options"`
			BuiltAsBottle       bool   `json:"built_as_bottle"`
			PouredFromBottle    bool   `json:"poured_from_bottle"`
			Time                int    `json:"time"`
			RuntimeDependencies []struct {
				FullName         string `json:"full_name"`
				Version          string `json:"version"`
				Revision         int    `json:"revision,omitempty"`
				BottleRebuild    int    `json:"bottle_rebuild,omitempty"`
				PkgVersion       string `json:"pkg_version,omitempty"`
				DeclaredDirectly bool   `json:"declared_directly"`
			} `json:"runtime_dependencies"`
			InstalledAsDependency bool `json:"installed_as_dependency"`
			InstalledOnRequest    bool `json:"installed_on_request"`
		} `json:"installed"`
		LinkedKeg                     *string `json:"linked_keg"`
		Pinned                        bool    `json:"pinned"`
		Outdated                      bool    `json:"outdated"`
		Deprecated                    bool    `json:"deprecated"`
		DeprecationDate               *string `json:"deprecation_date"`
		DeprecationReason             *string `json:"deprecation_reason"`
		DeprecationReplacementFormula any     `json:"deprecation_replacement_formula"`
		DeprecationReplacementCask    any     `json:"deprecation_replacement_cask"`
		Disabled                      bool    `json:"disabled"`
		DisableDate                   *string `json:"disable_date"`
		DisableReason                 *string `json:"disable_reason"`
		DisableReplacementFormula     any     `json:"disable_replacement_formula"`
		DisableReplacementCask        any     `json:"disable_replacement_cask"`
		PostInstallDefined            bool    `json:"post_install_defined"`
		Service                       *struct {
			Name struct {
				Macos string `json:"macos"`
			} `json:"name,omitempty"`
			Run       any    `json:"run,omitempty"`
			RunType   string `json:"run_type,omitempty"`
			KeepAlive struct {
				Always bool `json:"always"`
			} `json:"keep_alive,omitempty"`
			EnvironmentVariables struct {
				ETCDUNSUPPORTEDARCH  string `json:"ETCD_UNSUPPORTED_ARCH,omitempty"`
				PATH                 string `json:"PATH,omitempty"`
				OLLAMAFLASHATTENTION string `json:"OLLAMA_FLASH_ATTENTION,omitempty"`
				OLLAMAKVCACHETYPE    string `json:"OLLAMA_KV_CACHE_TYPE,omitempty"`
			} `json:"environment_variables,omitempty"`
			WorkingDir   string `json:"working_dir,omitempty"`
			LogPath      string `json:"log_path,omitempty"`
			ErrorLogPath string `json:"error_log_path,omitempty"`
			RequireRoot  bool   `json:"require_root,omitempty"`
		} `json:"service"`
		TapGitHead         string `json:"tap_git_head"`
		RubySourcePath     string `json:"ruby_source_path"`
		RubySourceChecksum struct {
			Sha256 string `json:"sha256"`
		} `json:"ruby_source_checksum"`
		HeadDependencies struct {
			BuildDependencies       []string `json:"build_dependencies"`
			Dependencies            []string `json:"dependencies"`
			TestDependencies        []string `json:"test_dependencies"`
			RecommendedDependencies []any    `json:"recommended_dependencies"`
			OptionalDependencies    []any    `json:"optional_dependencies"`
			UsesFromMacos           []any    `json:"uses_from_macos"`
			UsesFromMacosBounds     []struct {
				Since string `json:"since,omitempty"`
			} `json:"uses_from_macos_bounds"`
		} `json:"head_dependencies,omitempty"`
	} `json:"formulae"`
	Casks []struct {
		Token     string   `json:"token"`
		FullToken string   `json:"full_token"`
		OldTokens []string `json:"old_tokens"`
		Tap       string   `json:"tap"`
		Name      []string `json:"name"`
		Desc      *string  `json:"desc"`
		Homepage  string   `json:"homepage"`
		Url       string   `json:"url"`
		UrlSpecs  struct {
			Verified string `json:"verified,omitempty"`
			Branch   string `json:"branch,omitempty"`
			OnlyPath string `json:"only_path,omitempty"`
		} `json:"url_specs"`
		Version            string  `json:"version"`
		Autobump           bool    `json:"autobump"`
		NoAutobumpMessage  any     `json:"no_autobump_message"`
		SkipLivecheck      bool    `json:"skip_livecheck"`
		Installed          string  `json:"installed"`
		InstalledTime      int     `json:"installed_time"`
		BundleVersion      *string `json:"bundle_version"`
		BundleShortVersion *string `json:"bundle_short_version"`
		Outdated           bool    `json:"outdated"`
		Sha256             string  `json:"sha256"`
		Artifacts          []struct {
			Uninstall []struct {
				Quit      string `json:"quit,omitempty"`
				LoginItem string `json:"login_item,omitempty"`
				Pkgutil   any    `json:"pkgutil,omitempty"`
				Launchctl any    `json:"launchctl,omitempty"`
				Script    struct {
					Executable string   `json:"executable"`
					Sudo       bool     `json:"sudo"`
					Input      []string `json:"input,omitempty"`
					Args       []string `json:"args,omitempty"`
				} `json:"script,omitempty"`
				Delete any      `json:"delete,omitempty"`
				Rmdir  string   `json:"rmdir,omitempty"`
				Signal []string `json:"signal,omitempty"`
				Trash  string   `json:"trash,omitempty"`
			} `json:"uninstall,omitempty"`
			App []any `json:"app,omitempty"`
			Zap []struct {
				Trash     any      `json:"trash"`
				Rmdir     any      `json:"rmdir,omitempty"`
				Launchctl []string `json:"launchctl,omitempty"`
			} `json:"zap,omitempty"`
			Pkg                 []any    `json:"pkg,omitempty"`
			BashCompletion      []string `json:"bash_completion,omitempty"`
			FishCompletion      []string `json:"fish_completion,omitempty"`
			ZshCompletion       []string `json:"zsh_completion,omitempty"`
			Binary              []any    `json:"binary,omitempty"`
			Postflight          any      `json:"postflight"`
			UninstallPostflight any      `json:"uninstall_postflight"`
			Preflight           any      `json:"preflight"`
			Font                []string `json:"font,omitempty"`
			Manpage             []string `json:"manpage,omitempty"`
			UninstallPreflight  any      `json:"uninstall_preflight"`
		} `json:"artifacts"`
		Caveats   *string `json:"caveats"`
		DependsOn struct {
			Macos struct {
				Field1 []string `json:">="`
			} `json:"macos"`
			Arch []struct {
				Type string `json:"type"`
				Bits int    `json:"bits"`
			} `json:"arch,omitempty"`
			Cask    []string `json:"cask,omitempty"`
			Formula []string `json:"formula,omitempty"`
		} `json:"depends_on"`
		ConflictsWith *struct {
			Cask []string `json:"cask"`
		} `json:"conflicts_with"`
		Container                     any      `json:"container"`
		Rename                        []any    `json:"rename"`
		AutoUpdates                   *bool    `json:"auto_updates"`
		Deprecated                    bool     `json:"deprecated"`
		DeprecationDate               any      `json:"deprecation_date"`
		DeprecationReason             any      `json:"deprecation_reason"`
		DeprecationReplacementFormula any      `json:"deprecation_replacement_formula"`
		DeprecationReplacementCask    any      `json:"deprecation_replacement_cask"`
		Disabled                      bool     `json:"disabled"`
		DisableDate                   any      `json:"disable_date"`
		DisableReason                 any      `json:"disable_reason"`
		DisableReplacementFormula     any      `json:"disable_replacement_formula"`
		DisableReplacementCask        any      `json:"disable_replacement_cask"`
		TapGitHead                    string   `json:"tap_git_head"`
		Languages                     []string `json:"languages"`
		RubySourcePath                string   `json:"ruby_source_path"`
		RubySourceChecksum            struct {
			Sha256 string `json:"sha256"`
		} `json:"ruby_source_checksum"`
	} `json:"casks"`
}
