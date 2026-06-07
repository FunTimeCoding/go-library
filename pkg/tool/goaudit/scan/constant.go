package scan

const (
	StaleRouteKey  = "stale_route"
	StaleRouteText = "route/ exists (should be server/)"

	StaleToolKey  = "stale_tool"
	StaleToolText = "tool/ exists (should be model_context/)"

	StaleToolsetKey  = "stale_toolset"
	StaleToolsetText = "toolset/ exists (should be model_context/)"

	StalePollerKey  = "stale_poller"
	StalePollerText = "poller/ exists (should be worker/)"

	MissingMountKey  = "missing_mount"
	MissingMountText = "model_context/ missing mount.go"

	MissingCaptureFailKey  = "missing_capture_fail"
	MissingCaptureFailText = "model_context/ missing capture_fail.go"

	StaleNestedKey  = "stale_nested"
	StaleNestedText = "model_context/ contains nested.go (should be mount.go)"

	TopLevelArgumentKey  = "top_level_argument"
	TopLevelArgumentText = "argument/ is top-level (should be model_context/argument/)"

	TopLevelResponseOrphanKey  = "top_level_response_orphan"
	TopLevelResponseOrphanText = "response/ is top-level with no server/ or model_context/"

	TopLevelResponseMCPKey  = "top_level_response_mcp"
	TopLevelResponseMCPText = "response/ is top-level (should be convert/ or model_context/response/)"

	ConstantFileKey  = "constant_file"
	ConstantFileText = "constant.go should be constant/ directory"

	MissingOptionKey  = "missing_option"
	MissingOptionText = "missing option/"

	MissingRunKey  = "missing_run"
	MissingRunText = "missing run.go"

	MissingSuffixKey  = "missing_suffix"
	MissingSuffixText = "no standard service suffix (expected *d)"

	MissingSentryKey  = "missing_sentry"
	MissingSentryText = "missing sentry reporter"

	IdentityMissingFileKey  = "identity_missing_file"
	IdentityMissingFileText = "constant/constant.go not found"

	IdentityNotParseableKey  = "identity_not_parseable"
	IdentityNotParseableText = "constant/constant.go not parseable"

	IdentityMissingVariableKey  = "identity_missing_variable"
	IdentityMissingVariableText = "missing Identity in constant/"

	IdentityNoInitializerKey  = "identity_no_initializer"
	IdentityNoInitializerText = "Identity has no initializer"

	IdentityNotNewKey  = "identity_not_new"
	IdentityNotNewText = "Identity not initialized with identity.New()"

	IdentityNoArgumentKey  = "identity_no_argument"
	IdentityNoArgumentText = "identity.New() has no arguments"

	IdentityNotStringKey  = "identity_not_string"
	IdentityNotStringText = "identity.New() first argument is not a string literal"

	IdentityMismatchKey = "identity_mismatch"
	OpenAPIMismatchKey  = "openapi_mismatch"
)
