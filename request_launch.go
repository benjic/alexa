package alexa

const (
	launchRequestType = "LaunchRequest"
)

type LaunchRequest struct {
	Version string         `json:"version"`
	Context RequestContext `json:"context"`
	Session Session        `json:"session"`
	Request LaunchData     `json:"request"`
}

type LaunchData struct {
	Timestamp string `json:"timestamp"`
	RequestID string `json:"requestId"`
	Locale    string `json:"locale"`
}
