package request

const (
	LaunchRequestType = "LaunchRequest"
)

type Launch struct {
	Version string     `json:"version"`
	Context Context    `json:"context"`
	Session Session    `json:"session"`
	Request LaunchData `json:"request"`
}

type LaunchData struct {
	Timestamp string `json:"timestamp"`
	RequestID string `json:"requestId"`
	Locale    string `json:"locale"`
}
