package interfaces

import "github.com/benjic/alexa/request"

const (
	SystemExceptionEncounteredType = "System.ExceptionEncountered"
)

type SystemExceptionEncounteredHandler func(*SystemExceptionRequest) error

type SystemExceptionRequest struct {
	Version string                     `json:"version"`
	Context request.Context            `json:"context"`
	Request SystemExceptionRequestData `json:"request"`
}

type SystemExceptionRequestData struct {
	Type      string               `json:"type"`
	RequestID string               `json:"requestId"`
	Timestamp string               `json:"timestamp"`
	Locale    string               `json:"locale"`
	Error     SystemExceptionError `json:"error"`
	Cause     SystemExceptionCause `json:"cause"`
}

type SystemExceptionError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type SystemExceptionCause struct {
	RequestID string `json:"requestId"`
}
