package alexa

const (
	sessionEndedRequestType = "SessionEndedRequest"
)

type SessionEndedRequestHandler func(*SessionEndedRequest) error

type SessionEndedRequest struct {
	Version string           `json:"version"`
	Context RequestContext   `json:"context"`
	Session Session          `json:"session"`
	Request SessionEndedData `json:"request"`
}

type SessionEndedData struct {
	Timestamp string            `json:"timestamp"`
	RequestID string            `json:"requestId"`
	Reason    string            `json:"reason"`
	Error     SessionEndedError `json:"error"`
	Locale    string            `json:"locale"`
}

type SessionEndedError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
