package lighthouse

type Audits struct {
	IsOnHttps                     *AuditDetails               `json:"is-on-https"`
	RedirectsHttp                 *Audit                      `json:"redirects-http"`
	ServiceWorker                 *Audit                      `json:"service-worker"`
	WorksOffline                  *Audit                      `json:"works-offline"`
	Viewport                      *Audit                      `json:"viewport"`
	WithoutJavascript             *Audit                      `json:"without-javascript"`
	FirstContentfulPaint          *Audit                      `json:"first-contentful-paint"`
	LargestContentfulPaint        *Audit                      `json:"largest-contentful-paint"`
	FirstMeaningfulPaint          *Audit                      `json:"first-meaningful-paint"`
	LoadFastEnoughForPwa          *Audit                      `json:"load-fast-enough-for-pwa"`
	SpeedIndex                    *Audit                      `json:"speed-index"`
	ScreenshotThumbnails          *AuditScreenshotThumbnails  `json:"screenshot-thumbnails"`
	FinalScreenshot               *AuditFinalScreenshot       `json:"final-screenshot"`
	EstimatedInputLatency         *Audit                      `json:"estimated-input-latency"`
	TotalBlockingTime             *Audit                      `json:"total-blocking-time"`
	MaxPotentialFid               *Audit                      `json:"max-potential-fid"`
	CumulativeLayoutShift         *AuditDetails               `json:"cumulative-layout-shift"`
	ErrorsInConsole               *AuditDetails               `json:"errors-in-console"`
	ServerResponseTime            *AuditDetails               `json:"server-response-time"`
	FirstCpuIdle                  *Audit                      `json:"first-cpu-idle"`
	Interactive                   *Audit                      `json:"interactive"`
	UserTimings                   *AuditDetails               `json:"user-timings"`
	CriticalRequestChains         *AuditCriticalRequestChains `json:"critical-request-chains"`
	Redirects                     *AuditDetails               `json:"redirects"`
	InstallableManifest           *AuditDetails               `json:"installable-manifest"`
	AppleTouchIcon                *Audit                      `json:"apple-touch-icon"`
	SplashScreen                  *AuditDetails               `json:"splash-screen"`
	ThemedOmnibox                 *AuditDetails               `json:"themed-omnibox"`
	MaskableIcon                  *Audit                      `json:"maskable-icon"`
	ContentWidth                  *Audit                      `json:"content-width"`
	ImageAspectRatio              *AuditDetails               `json:"image-aspect-ratio"`
	ImageSizeResponsive           *AuditDetails               `json:"image-size-responsive"`
	Deprecations                  *AuditDetails               `json:"deprecations"`
	MainthreadWorkBreakdown       *AuditDetails               `json:"mainthread-work-breakdown"`
	BootupTime                    *AuditDetails               `json:"bootup-time"`
	UsesRelPreload                *AuditDetails               `json:"uses-rel-preload"`
	UsesRelPreconnect             *AuditDetails               `json:"uses-rel-preconnect"`
	FontDisplay                   *AuditDetails               `json:"font-display"`
	Diagnostics                   *AuditDiagnostics           `json:"diagnostics"`
	NetworkRequests               *AuditDetails               `json:"network-requests"`
	NetworkRtt                    *AuditDetails               `json:"network-rtt"`
	NetworkServerLatency          *AuditDetails               `json:"network-server-latency"`
	MainThreadTasks               *AuditDetails               `json:"main-thread-tasks"`
	Metrics                       *AuditMetrics               `json:"metrics"`
	OfflineStartUrl               *Audit                      `json:"offline-start-url"`
	PerformanceBudget             *Audit                      `json:"performance-budget"`
	TimingBudget                  *Audit                      `json:"timing-budget"`
	ResourceSummary               *AuditDetails               `json:"resource-summary"`
	ThirdPartySummary             *AuditDetails               `json:"third-party-summary"`
	LargestContentfulPaintElement *AuditDetails               `json:"largest-contentful-paint-element"`
	LayoutShiftElements           *AuditDetails               `json:"layout-shift-elements"`
	LongTasks                     *AuditDetails               `json:"long-tasks"`
	PwaCrossBrowser               *Audit                      `json:"pwa-cross-browser"`
	PwaPageTransitions            *Audit                      `json:"pwa-page-transitions"`
	PwaEachPageHasUrl             *Audit                      `json:"pwa-each-page-has-url"`
	Accesskeys                    *Audit                      `json:"accesskeys"`
	AriaAllowedAttr               *Audit                      `json:"aria-allowed-attr"`
	AriaHiddenBody                *AuditDetails               `json:"aria-hidden-body"`
	AriaHiddenFocus               *Audit                      `json:"aria-hidden-focus"`
	AriaInputFieldName            *Audit                      `json:"aria-input-field-name"`
	AriaRequiredAttr              *Audit                      `json:"aria-required-attr"`
	AriaRequiredChildren          *Audit                      `json:"aria-required-children"`
	AriaRequiredParent            *Audit                      `json:"aria-required-parent"`
	AriaRoles                     *Audit                      `json:"aria-roles"`
	AriaToggleFieldName           *Audit                      `json:"aria-toggle-field-name"`
	AriaValidAttrValue            *Audit                      `json:"aria-valid-attr-value"`
	AriaValidAttr                 *Audit                      `json:"aria-valid-attr"`
	ButtonName                    *AuditDetails               `json:"button-name"`
	Bypass                        *AuditDetails               `json:"bypass"`
	ColorContrast                 *AuditDetails               `json:"color-contrast"`
	DefinitionList                *Audit                      `json:"definition-list"`
	Dlitem                        *Audit                      `json:"dlitem"`
	DocumentTitle                 *AuditDetails               `json:"document-title"`
	DuplicateIdActive             *AuditDetails               `json:"duplicate-id-active"`
	DuplicateIdAria               *AuditDetails               `json:"duplicate-id-aria"`
	FormFieldMultipleLabels       *Audit                      `json:"form-field-multiple-labels"`
	FrameTitle                    *Audit                      `json:"frame-title"`
	HeadingOrder                  *Audit                      `json:"heading-order"`
	HtmlHasLang                   *AuditDetails               `json:"html-has-lang"`
	HtmlLangValid                 *Audit                      `json:"html-lang-valid"`
	ImageAlt                      *AuditDetails               `json:"image-alt"`
	InputImageAlt                 *Audit                      `json:"input-image-alt"`
	Label                         *AuditDetails               `json:"label"`
	LayoutTable                   *Audit                      `json:"layout-table"`
	LinkName                      *AuditDetails               `json:"link-name"`
	List                          *Audit                      `json:"list"`
	Listitem                      *Audit                      `json:"listitem"`
	MetaRefresh                   *Audit                      `json:"meta-refresh"`
	MetaViewport                  *AuditDetails               `json:"meta-viewport"`
	ObjectAlt                     *Audit                      `json:"object-alt"`
	Tabindex                      *Audit                      `json:"tabindex"`
	TdHeadersAttr                 *Audit                      `json:"td-headers-attr"`
	ThHasDataCells                *Audit                      `json:"th-has-data-cells"`
	ValidLang                     *Audit                      `json:"valid-lang"`
	VideoCaption                  *Audit                      `json:"video-caption"`
	VideoDescription              *Audit                      `json:"video-description"`
	CustomControlsLabels          *Audit                      `json:"custom-controls-labels"`
	CustomControlsRoles           *Audit                      `json:"custom-controls-roles"`
	FocusTraps                    *Audit                      `json:"focus-traps"`
	FocusableControls             *Audit                      `json:"focusable-controls"`
	InteractiveElementAffordance  *Audit                      `json:"interactive-element-affordance"`
	LogicalTabOrder               *Audit                      `json:"logical-tab-order"`
	ManagedFocus                  *Audit                      `json:"managed-focus"`
	OffscreenContentHidden        *Audit                      `json:"offscreen-content-hidden"`
	UseLandmarks                  *Audit                      `json:"use-landmarks"`
	VisualOrderFollowsDom         *Audit                      `json:"visual-order-follows-dom"`
	UsesLongCacheTtl              *AuditDetails               `json:"uses-long-cache-ttl"`
	TotalByteWeight               *AuditDetails               `json:"total-byte-weight"`
	OffscreenImages               *AuditDetails               `json:"offscreen-images"`
	RenderBlockingResources       *AuditDetails               `json:"render-blocking-resources"`
	UnminifiedCss                 *AuditDetails               `json:"unminified-css"`
	UnminifiedJavascript          *AuditDetails               `json:"unminified-javascript"`
	UnusedCssRules                *AuditDetails               `json:"unused-css-rules"`
	UnusedJavascript              *AuditDetails               `json:"unused-javascript"`
	UsesWebpImages                *AuditDetails               `json:"uses-webp-images"`
	UsesOptimizedImages           *AuditDetails               `json:"uses-optimized-images"`
	UsesTextCompression           *AuditDetails               `json:"uses-text-compression"`
	UsesResponsiveImages          *AuditDetails               `json:"uses-responsive-images"`
	EfficientAnimatedContent      *AuditDetails               `json:"efficient-animated-content"`
	AppcacheManifest              *AuditDetails               `json:"appcache-manifest"`
	Doctype                       *Audit                      `json:"doctype"`
	Charset                       *Audit                      `json:"charset"`
	DomSize                       *AuditDetails               `json:"dom-size"`
	ExternalAnchorsUseRelNoopener *AuditDetails               `json:"external-anchors-use-rel-noopener"`
	GeolocationOnStart            *AuditDetails               `json:"geolocation-on-start"`
	NoDocumentWrite               *AuditDetails               `json:"no-document-write"`
	NoVulnerableLibraries         *AuditDetails               `json:"no-vulnerable-libraries"`
	JsLibraries                   *AuditDetails               `json:"js-libraries"`
	NotificationOnStart           *AuditDetails               `json:"notification-on-start"`
	PasswordInputsCanBePastedInto *AuditDetails               `json:"password-inputs-can-be-pasted-into"`
	UsesHttp2                     *AuditDetails               `json:"uses-http2"`
	UsesPassiveEventListeners     *AuditDetails               `json:"uses-passive-event-listeners"`
	MetaDescription               *Audit                      `json:"meta-description"`
	HttpStatusCode                *Audit                      `json:"http-status-code"`
	FontSize                      *AuditDetails               `json:"font-size"`
	LinkText                      *AuditDetails               `json:"link-text"`
	CrawlableAnchors              *AuditDetails               `json:"crawlable-anchors"`
	IsCrawlable                   *AuditDetails               `json:"is-crawlable"`
	RobotsTxt                     *AuditDetails               `json:"robots-txt"`
	TapTargets                    *AuditDetails               `json:"tap-targets"`
	Hreflang                      *AuditDetails               `json:"hreflang"`
	Plugins                       *AuditDetails               `json:"plugins"`
	Canonical                     *Audit                      `json:"canonical"`
	StructuredData                *Audit                      `json:"structured-data"`
	CustomAudits
}

const (
	ScoreDisplayModeNotApplicable = "notApplicable"
	ScoreDisplayModeBinary        = "binary"
	ScoreDisplayModeNumeric       = "numeric"
	ScoreDisplayModeInformative   = "informative"

	AuditDetailsTypeTable                = "table"
	AuditDetailsTypeFilmstrip            = "filmstrip"
	AuditDetailsTypeScreenshot           = "screenshot"
	AuditDetailsTypeDebugData            = "debugdata"
	AuditDetailsTypeOpportunity          = "opportunity"
	AuditDetailsTypeCriticalRequestChain = "criticalrequestchain"

	NumericUnitByte        = "byte"
	NumericUnitMillisecond = "millisecond"
	NumericUnitElement     = "element"
	NumericUnitUnitless    = "unitless"
)

type Audit struct {
	/* The string identifier of the audit, in kebab case. */
	Id string `json:"id"`
	/* Short, user-visible title for the audit. The text can change depending on if the audit passed or failed. */
	Title string `json:"title"`
	/* A more detailed description that describes why the audit is important and links to Lighthouse documentation on the audit; markdown links supported. */
	Description string `json:"description"`
	/* The scored value of the audit, provided in the range `0-1`, or null if `scoreDisplayMode` indicates not scored. */
	Score *float32 `json:"score"`
	/* A string identifying how the score should be interpreted for display. */
	ScoreDisplayMode string        `json:"scoreDisplayMode"`
	Warnings         []interface{} `json:"warnings"`
	/* An explanation of why the audit failed on the test page. */
	Explanation interface{} `json:"explanation"`
	/* A numeric value that has a meaning specific to the audit, e.g. the number of nodes in the DOM or the timestamp of a specific load event. More information can be found in the audit details, if present. */
	NumericValue float64 `json:"numericValue"`
	/** The unit of `numericValue`, used when the consumer wishes to convert numericValue to a display string. A superset of https://tc39.es/proposal-unified-intl-numberformat/section6/locales-currencies-tz_proposed_out.html#sec-issanctionedsimpleunitidentifier */
	NumericUnit string `json:"numericUnit"`
	/* The i18n'd string value that the audit wishes to display for its results. This value is not necessarily the string version of the `numericValue`. */
	DisplayValue interface{} `json:"displayValue"`
	/* Error message from any exception thrown while running this audit. */
	ErrorMessage interface{} `json:"errorMessage"`
	/* Overrides scoreDisplayMode with notApplicable if set to true */
	NotApplicable *bool `json:"notApplicable"`
	/* If an audit encounters unusual execution circumstances, strings can be put in this optional array to add top-level warnings to the LHR. */
	RunWarnings []interface{} `json:"runWarnings"`
}

type AuditDetails struct {
	Audit
	Details `json:"details"`
}

/* Extra information about the page provided by some types of audits, in one of several possible forms that can be rendered in the HTML report. */
type Details struct {
	Type     string                   `json:"type"`
	Scale    *int                     `json:"scale"`
	Headings []AuditDetailsHeading    `json:"headings"`
	Items    []map[string]interface{} `json:"items"`
	/* The relative time from navigationStart to this frame, in milliseconds. */
	Timing *float64 `json:"timing"`
	/* The raw timestamp of this frame, in microseconds. */
	Timestamp *float64 `json:"timestamp"`
	/* The data URL encoding of this frame. */
	Data                *string  `json:"data"`
	OverallSavingsMs    *float64 `json:"overallSavingsMs"`
	OverallSavingsBytes *int     `json:"overallSavingsBytes"`
	Summary             *struct {
		WastedBytes float64 `json:"wastedBytes"`
	} `json:"summary"`
}

type AuditDetailsHeading struct {
	/*
		The name of the property within items being described.

		If null, subItemsHeading must be defined, and the first table row in this column for
		every item will be empty.

		See legacy-javascript for an example.
	*/
	Key string `json:"key"`
	/*
		The data format of the column of values being described. Usually
		those values will be primitives rendered as this type, but the values
		could also be objects with their own type to override this field.
	*/
	ItemType string `json:"itemType"`
	/* Readable text label of the field. */
	Text interface{} `json:"text"`
}

type AuditChain struct {
	Request struct {
		URL                  string  `json:"url"`
		StartTime            float64 `json:"startTime"`
		EndTime              float64 `json:"endTime"`
		ResponseReceivedTime float64 `json:"responseReceivedTime"`
		TransferSize         int     `json:"transferSize"`
	} `json:"request"`
	Children map[string]AuditChain
}

type AuditScreenshotThumbnails struct {
	Audit
	Details struct {
		Type  string `json:"type"`
		Scale int    `json:"scale"`
		Items []struct {
			Timing    float64 `json:"timing"`
			Timestamp float64 `json:"timestamp"`
			Data      string  `json:"data"`
		} `json:"items"`
	} `json:"details"`
}

type AuditFinalScreenshot struct {
	Audit
	Details struct {
		Type      string  `json:"type"`
		Timing    float64 `json:"timing"`
		Timestamp float64 `json:"timestamp"`
		Data      string  `json:"data"`
	} `json:"details"`
}

type AuditCriticalRequestChains struct {
	Audit
	Details struct {
		Type   string `json:"type"`
		Chains *map[string]struct {
			Request struct {
				URL                  string  `json:"url"`
				StartTime            float64 `json:"startTime"`
				EndTime              float64 `json:"endTime"`
				ResponseReceivedTime float64 `json:"responseReceivedTime"`
				TransferSize         int     `json:"transferSize"`
			} `json:"request"`
			Children map[string]AuditChain `json:"children"`
		} `json:"chains"`
		LongestChain *struct {
			Duration     float64 `json:"duration"`
			Length       int     `json:"length"`
			TransferSize int     `json:"transferSize"`
		} `json:"longestChain"`
	} `json:"details"`
}

type AuditDiagnostics struct {
	Audit
	Details struct {
		Type  string `json:"type"`
		Items []struct {
			NumRequests              int     `json:"numRequests"`
			NumScripts               int     `json:"numScripts"`
			NumStylesheets           int     `json:"numStylesheets"`
			NumFonts                 int     `json:"numFonts"`
			NumTasks                 int     `json:"numTasks"`
			NumTasksOver10ms         int     `json:"numTasksOver10ms"`
			NumTasksOver25ms         int     `json:"numTasksOver25ms"`
			NumTasksOver50ms         int     `json:"numTasksOver50ms"`
			NumTasksOver100ms        int     `json:"numTasksOver100ms"`
			NumTasksOver500ms        int     `json:"numTasksOver500ms"`
			Rtt                      float64 `json:"rtt"`
			Throughput               float64 `json:"throughput"`
			MaxRtt                   float64 `json:"maxRtt"`
			MaxServerLatency         float64 `json:"maxServerLatency"`
			TotalByteWeight          int     `json:"totalByteWeight"`
			TotalTaskTime            float64 `json:"totalTaskTime"`
			MainDocumentTransferSize int     `json:"mainDocumentTransferSize"`
		} `json:"items"`
	} `json:"details"`
}

type AuditMetrics struct {
	Audit
	Details struct {
		Type  string `json:"type"`
		Items []struct {
			FirstContentfulPaint             int     `json:"firstContentfulPaint"`
			FirstMeaningfulPaint             int     `json:"firstMeaningfulPaint"`
			LargestContentfulPaint           int     `json:"largestContentfulPaint"`
			FirstCPUIdle                     int     `json:"firstCPUIdle"`
			Interactive                      int     `json:"interactive"`
			SpeedIndex                       int     `json:"speedIndex"`
			EstimatedInputLatency            int     `json:"estimatedInputLatency"`
			TotalBlockingTime                int     `json:"totalBlockingTime"`
			MaxPotentialFID                  int     `json:"maxPotentialFID"`
			CumulativeLayoutShift            float64 `json:"cumulativeLayoutShift"`
			ObservedNavigationStart          int     `json:"observedNavigationStart"`
			ObservedNavigationStartTs        int     `json:"observedNavigationStartTs"`
			ObservedFirstPaint               int     `json:"observedFirstPaint"`
			ObservedFirstPaintTs             int     `json:"observedFirstPaintTs"`
			ObservedFirstContentfulPaint     int     `json:"observedFirstContentfulPaint"`
			ObservedFirstContentfulPaintTs   int     `json:"observedFirstContentfulPaintTs"`
			ObservedFirstMeaningfulPaint     int     `json:"observedFirstMeaningfulPaint"`
			ObservedFirstMeaningfulPaintTs   int     `json:"observedFirstMeaningfulPaintTs"`
			ObservedLargestContentfulPaint   int     `json:"observedLargestContentfulPaint"`
			ObservedLargestContentfulPaintTs int     `json:"observedLargestContentfulPaintTs"`
			ObservedTraceEnd                 int     `json:"observedTraceEnd"`
			ObservedTraceEndTs               int     `json:"observedTraceEndTs"`
			ObservedLoad                     int     `json:"observedLoad"`
			ObservedLoadTs                   int     `json:"observedLoadTs"`
			ObservedDomContentLoadedTs       int     `json:"observedDomContentLoadedTs"`
			ObservedCumulativeLayoutShift    float64 `json:"observedCumulativeLayoutShift"`
			ObservedFirstVisualChange        int     `json:"observedFirstVisualChange"`
			ObservedFirstVisualChangeTs      int     `json:"observedFirstVisualChangeTs"`
			ObservedLastVisualChange         int     `json:"observedLastVisualChange"`
			ObservedLastVisualChangeTs       int     `json:"observedLastVisualChangeTs"`
			ObservedSpeedIndex               int     `json:"observedSpeedIndex"`
			ObservedSpeedIndexTs             int     `json:"observedSpeedIndexTs"`
			LcpInvalidated                   *bool   `json:"lcpInvalidated"`
		}
	} `json:"details"`
}

type CustomAudits struct {
	FullSizeScreenshot *AuditFinalScreenshot `json:"fullsize-screenshot"`
	VideoScreenCast    *AuditFinalScreenshot `json:"video-screencast"`
}
