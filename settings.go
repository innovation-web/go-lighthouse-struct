package lighthouse

const (
	ThrottlingMethodDevtools = "devtools"
	ThrottlingMethodProvided = "provided"
	ThrottlingMethodSimulate = "simulate"

	OutputJson = "json"
	OutputHtml = "html"
	OutputCsv  = "csv"

	EmulatedFormFactorMobile  = "mobile"
	EmulatedFormFactorDesktop = "desktop"
	EmulatedFormFactorNone    = "none"

	CategoryAccessibility = "accessibility"
	CategoryBestPractices = "best-practices"
	CategoryPerformance   = "performance"
	CategoryPwa           = "pwa"
	CategorySeo           = "seo"
)

type ConfigSettings struct {
	/* The type(s) of report output to be produced. */
	Output []string `json:"output"`
	/* The locale to use for the output. */
	Locale *string `json:"locale"`
	/* The maximum amount of time to wait for a page content render, in ms. If no content is rendered within this limit, the run is aborted with an error. */
	MaxWaitForFcp *int `json:"maxWaitForFcp"`
	/* The maximum amount of time to wait for a page to load, in ms. */
	MaxWaitForLoad *int `json:"maxWaitForLoad"`
	/* List of URL patterns to block. */
	BlockedUrlPatterns []string `json:"blockedUrlPatterns"`
	/* Comma-delimited list of trace categories to include. */
	AdditionalTraceCategories []string `json:"additionalTraceCategories"`
	/* Flag indicating the run should only audit. */
	AuditMode *bool `json:"auditMode"`
	/* Flag indicating the run should only gather. */
	GatherMode *bool `json:"gatherMode"`
	/** Flag indicating that the browser storage should not be reset for the audit. */
	DisableStorageReset *bool `json:"disableStorageReset"`

	/* How Lighthouse should interpret this run in regards to scoring performance metrics and skipping mobile-only tests in desktop. Must be set even if throttling/emulation is being applied outside of Lighthouse. */
	EmulatedFormFactor *string `json:"emulatedFormFactor"`

	/* The method used to throttle the network. */
	ThrottlingMethod *string `json:"throttlingMethod"`
	/* The throttling config settings. */
	Throttling struct {
		RttMs                  *int     `json:"rttMs"`
		ThroughputKbps         *float64 `json:"throughputKbps"`
		RequestLatencyMs       *float64 `json:"requestLatencyMs"`
		DownloadThroughputKbps *float64 `json:"downloadThroughputKbps"`
		UploadThroughputKbps   *float64 `json:"uploadThroughputKbps"`
		CpuSlowdownMultiplier  *int     `json:"cpuSlowdownMultiplier"`
	} `json:"throttling"`
	/* If present, the run should only conduct this list of audits. */
	OnlyAudits []string `json:"onlyAudits"`
	/* If present, the run should only conduct this list of categories. */
	OnlyCategories []string `json:"onlyCategories"`
	/* If present, the run should skip this list of audits. */
	SkipAudits []string `json:"skipAudits"`
	/* List of extra HTTP Headers to include. */
	ExtraHeaders *string `json:"extraHeaders"`
	/* How Lighthouse was run, e.g. from the Chrome extension or from the npm module */
	Channel *string `json:"channel"`

	InternalDisableDeviceScreenEmulation *bool `json:"internalDisableDeviceScreenEmulation"`

	/* The budget.json object for LightWallet. */
	// TODO "budgets": null,
	/* Precomputed lantern estimates to use instead of observed analysis. */
	// TODO "precomputedLanternData": null,
}
