package lighthouse

import "time"

type Report struct {
	UserAgent         string      `json:"userAgent"`
	Environment       Environment `json:"environment"`
	LighthouseVersion string      `json:"lighthouseVersion"`
	FetchTime         time.Time   `json:"fetchTime"`
	FinalUrl          string      `json:"finalUrl"`
	RunWarnings       []string    `json:"runWarnings"`
	Audits            Audits      `json:"audits"`
	*ConfigSettings   `json:"configSettings"`
	Categories        interface{} `json:"categories"`
	CategoryGroups    interface{} `json:"categoryGroups"`
	*Timing           `json:"timing"`
	I18n              interface{}   `json:"i18n"`
	StackPacks        []interface{} `json:"stackPacks"`
}

type Environment struct {
	NetworkUserAgent string `json:"networkUserAgent"`
	HostUserAgent    string `json:"hostUserAgent"`
	BenchmarkIndex   int    `json:"benchmarkIndex"`
}

type Timing struct {
	Entries []struct {
		StartTime float64 `json:"startTime"`
		Name      string  `json:"name"`
		Duration  float64 `json:"duration"`
		EntryType string  `json:"entryType"`
	} `json:"entries"`
	Total float64 `json:"total"`
}

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
	AuditMode                            *bool   `json:"auditMode"`
	GatherMode                           *bool   `json:"gatherMode"`
	DisableStorageReset                  *bool   `json:"DisableStorageReset"`
	EmulatedFormFactor                   *string `json:"emulatedFormFactor"`
	InternalDisableDeviceScreenEmulation *bool   `json:"internalDisableDeviceScreenEmulation"`
	Channel                              *string `json:"channel"`
	// TODO "budgets": null,
	Locale *string `json:"locale"`
	// TODO "blockedUrlPatterns": null,
	// TODO "additionalTraceCategories": null,
	// TODO "extraHeaders": null,
	// TODO "precomputedLanternData": null,
	// TODO "onlyAudits": null,
	// TODO "onlyCategories": null,
	// TODO "skipAudits": null
}
