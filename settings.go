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
	Output           []string `json:"output"`
	MaxWaitForFcp    *int     `json:"maxWaitForFcp"`
	MaxWaitForLoad   *int     `json:"maxWaitForLoad"`
	ThrottlingMethod *string  `json:"throttlingMethod"`
	Throttling       *struct {
		RttMs                  *int     `json:"rttMs"`
		ThroughputKbps         *float64 `json:"throughputKbps"`
		RequestLatencyMs       *float64 `json:"requestLatencyMs"`
		DownloadThroughputKbps *float64 `json:"downloadThroughputKbps"`
		UploadThroughputKbps   *float64 `json:"uploadThroughputKbps"`
		CpuSlowdownMultiplier  *int     `json:"cpuSlowdownMultiplier"`
	} `json:"throttling"`
	AuditMode                            *bool    `json:"auditMode"`
	GatherMode                           *bool    `json:"gatherMode"`
	DisableStorageReset                  *bool    `json:"DisableStorageReset"`
	EmulatedFormFactor                   *string  `json:"emulatedFormFactor"`
	InternalDisableDeviceScreenEmulation *bool    `json:"internalDisableDeviceScreenEmulation"`
	Channel                              *string  `json:"channel"`
	Locale                               *string  `json:"locale"`
	BlockedUrlPatterns                   []string `json:"blockedUrlPatterns"`
	AdditionalTraceCategories            []string `json:"additionalTraceCategories"`
	ExtraHeaders                         *string  `json:"extraHeaders"`
	OnlyAudits                           []string `json:"onlyAudits"`
	OnlyCategories                       []string `json:"onlyCategories"`
	SkipAudits                           []string `json:"skipAudits"`
	// TODO "budgets": null,
	// TODO "precomputedLanternData": null,
}
