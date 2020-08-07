package lighthouse

type LogItem struct {
	Method string `json:"method"`
	Params struct {
		FrameId   string  `json:"frameId"`
		LoaderId  string  `json:"loaderId"`
		Timestamp float64 `json:"timestamp"`

		Name        string          `json:"name"`
		Reason      string          `json:"reason"`
		URL         string          `json:"url"`
		Frame       *LogParamsFrame `json:"frame"`
		Delay       int             `json:"delay"`
		Disposition string          `json:"disposition"`

		RequestId                string              `json:"requestId"`
		Type                     string              `json:"type"`
		EncodedDataLength        int                 `json:"encodedDataLength"`
		DocumentURL              string              `json:"documentURL"`
		Request                  *LogParamsRequest   `json:"request"`
		RedirectResponse         *LogParamsResponse  `json:"redirectResponse"`
		WallTime                 float64             `json:"wallTime"`
		Initiator                *LogParamsInitiator `json:"initiator"`
		HasUserGesture           bool                `json:"hasUserGesture"`
		AssociatedCookies        []interface{}       `json:"associatedCookies"` // TODO
		Headers                  *LogParamsHeader    `json:"headers"`
		HeadersText              string              `json:"headersText"`
		BlockedCookies           []interface{}       `json:"blockedCookies"` // TODO
		Response                 *LogParamsResponse  `json:"response"`
		DataLength               int                 `json:"dataLength"`
		ShouldReportCorbBlocking bool                `json:"shouldReportCorbBlocking"`
	} `json:"params"`
}

type LogParamsFrame struct {
	Id             string `json:"id"`
	LoaderId       string `json:"loaderId"`
	URL            string `json:"url"`
	SecurityOrigin string `json:"securityOrigin"`
	MimeType       string `json:"mimeType"`
}

type LogParamsInitiator struct {
	Type       string    `json:"other"`
	URL        string    `json:"url"`
	LineNumber int       `json:"lineNumber"`
	Stack      *LogStack `json:"stack"`
}

type LogParamsHeader map[string]string

type LogParamsRequest struct {
	URL              string           `json:"url"`
	Method           string           `json:"method"`
	Headers          *LogParamsHeader `json:"headers"`
	MixedContentType string           `json:"mixedContentType"`
	InitialPriority  string           `json:"initialPriority"`
	ReferrerPolicy   string           `json:"referrerPolicy"`
	IsLinkPreload    bool             `json:"isLinkPreload"`
}

type LogParamsResponse struct {
	URL                string                    `json:"url"`
	Status             int                       `json:"status"`
	StatusText         string                    `json:"statusText"`
	Headers            *LogParamsHeader          `json:"headers"`
	MimeType           string                    `json:"mimeType"`
	RequestHeaders     *LogParamsHeader          `json:"requestHeaders"`
	ConnectionReused   bool                      `json:"connectionReused"`
	ConnectionId       int                       `json:"connectionId"`
	RemoteIPAddress    string                    `json:"remoteIPAddress"`
	RemotePort         int                       `json:"remotePort"`
	FromDiskCache      bool                      `json:"fromDiskCache"`
	FromServiceWorker  bool                      `json:"fromServiceWorker"`
	FromPrefetchCache  bool                      `json:"fromPrefetchCache"`
	EncodedDataLength  int                       `json:"encodedDataLength"`
	Timing             *LogParamsTiming          `json:"timing"`
	Protocol           string                    `json:"protocol"`
	SecurityState      string                    `json:"securityState"`
	SecurityDetails    *LogParamsSecurityDetails `json:"securityDetails"`
	HeadersText        string                    `json:"headersText"`
	RequestHeadersText string                    `json:"requestHeadersText"`
}

type LogParamsTiming struct {
	RequestTime       float64 `json:"requestTime"`
	ProxyStart        float64 `json:"proxyStart"`
	ProxyEnd          float64 `json:"proxyEnd"`
	DnsStart          float64 `json:"dnsStart"`
	DnsEnd            float64 `json:"dnsEnd"`
	ConnectStart      float64 `json:"connectStart"`
	ConnectEnd        float64 `json:"connectEnd"`
	SslStart          float64 `json:"sslStart"`
	SslEnd            float64 `json:"sslEnd"`
	WorkerStart       float64 `json:"workerStart"`
	WorkerReady       float64 `json:"workerReady"`
	SendStart         float64 `json:"sendStart"`
	SendEnd           float64 `json:"sendEnd"`
	PushStart         float64 `json:"pushStart"`
	PushEnd           float64 `json:"pushEnd"`
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"`
}

type LogParamsSecurityDetails struct {
	Protocol                          string        `json:"protocol"`
	KeyExchange                       string        `json:"keyExchange"`
	KeyExchangeGroup                  string        `json:"keyExchangeGroup"`
	Cipher                            string        `json:"cipher"`
	CertificateId                     int           `json:"certificateId"`
	SubjectName                       string        `json:"subjectName"`
	SanList                           []string      `json:"sanList"`
	Issuer                            string        `json:"issuer"`
	ValidFrom                         int           `json:"validFrom"`
	ValidTo                           int           `json:"validTo"`
	SignedCertificateTimestampList    []interface{} `json:"signedCertificateTimestampList"` // TODO
	CertificateTransparencyCompliance string        `json:"certificateTransparencyCompliance"`
}

type LogStack struct {
	Description string               `json:"description"`
	CallFrames  []LogStackCallFrames `json:"callFrames"`
	Parent      *LogStack            `json:"parent"`
}

type LogStackCallFrames struct {
	FunctionName string `json:"functionName"`
	ScriptId     string `json:"scriptId"`
	URL          string `json:"url"`
	LineNumber   int    `json:"lineNumber"`
	ColumnNumber int    `json:"columnNumber"`
}
