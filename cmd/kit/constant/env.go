package constant

const (

	// EnvKeyCliVerbose
	//	Provides the debug flag. This value is true when the command is open debug mode
	EnvKeyCliVerbose = "CLI_VERBOSE"

	EnvKeyDryRunDisable = "CLI_DRY_RUN_DISABLE"

	EnvKeyGitInfoScheme = "CLI_GIT_INFO_SCHEME"

	EnvKeySkipWorktreeCheck = "CLI_SKIP_WORKTREE_CHECK"

	// EnvKeyCliTimeoutSecond
	//	Provides the timeout second flag
	EnvKeyCliTimeoutSecond = "CLI_CONFIG_TIMEOUT_SECOND"

	// EnvLogLevel
	//	env ENV_WEB_LOG_LEVEL default ""
	EnvLogLevel string = "CLI_LOG_LEVEL"

	DefaultChangelogMarkdownFile    = "CHANGELOG.md"
	DefaultChangelogLastContentFile = "CHANGELOG.txt"

	VersionRcFileName = ".versionrc"
)
