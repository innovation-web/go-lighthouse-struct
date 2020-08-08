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
