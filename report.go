package lighthouse

import "time"

type Report struct {
	/* The URL that was supplied to Lighthouse and initially navigated to. */
	RequestedUrl string `json:"requestedUrl"`
	/* The post-redirects URL that Lighthouse loaded. */
	FinalUrl string `json:"finalUrl"`
	/* The ISO-8601 timestamp of when the results were generated. */
	FetchTime time.Time `json:"fetchTime"`
	/* The version of Lighthouse with which these results were generated. */
	LighthouseVersion string `json:"lighthouseVersion"`
	/* An object containing the results of the audits, keyed by the audits' `id` identifier. */
	Audits Audits `json:"audits"`
	/* The top-level categories, their overall scores, and member audits. */
	Categories interface{} `json:"categories"`
	/* Descriptions of the groups referenced by CategoryMembers. */
	CategoryGroups interface{} `json:"categoryGroups"`

	/* The config settings used for these results. */
	*ConfigSettings `json:"configSettings"`
	/* List of top-level warnings for this Lighthouse run. */
	RunWarnings []string `json:"runWarnings"`
	/* A top-level error message that, if present, indicates a serious enough problem that this Lighthouse result may need to be discarded. */
	RuntimeError *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"runtimeError"`
	/* The User-Agent string of the browser used run Lighthouse for these results. */
	UserAgent string `json:"userAgent"`
	/* Information about the environment in which Lighthouse was run. */
	Environment Environment `json:"environment"`
	/* Execution timings for the Lighthouse run */
	*Timing `json:"timing"`
	/* The record of all formatted string locations in the LHR and their corresponding source values. */
	I18n interface{} `json:"i18n"`
	/* An array containing the result of all stack packs. */
	StackPacks []StackPacksItem `json:"stackPacks"`
}

type Environment struct {
	NetworkUserAgent string  `json:"networkUserAgent"`
	HostUserAgent    string  `json:"hostUserAgent"`
	BenchmarkIndex   float64 `json:"benchmarkIndex"`
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

type StackPacksItem struct {
	/* The unique string ID for this stack pack. */
	Id string `json:"id"`
	/* The title of the stack pack, to be displayed in the report. */
	Title string `json:"title"`
	/* A base64 data url to be used as the stack pack's icon. */
	IconDataURL string `json:"iconDataURL"`
	/* A set of descriptions for some of Lighthouse's audits, keyed by audit `id`. */
	Descriptions map[string]string `json:"descriptions"`
}
