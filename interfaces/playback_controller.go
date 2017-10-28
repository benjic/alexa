package interfaces

import "github.com/benjic/alexa/request"

const (
	PlaybackControllerNextCommandIssuedType     = "PlaybackController.NextCommandIssued"
	PlaybackControllerPausedCommandIssuedType   = "PlaybackController.PausedCommandIssued"
	PlaybackControllerPlayCommandIssuedType     = "PlaybackController.PlayCommandIssued"
	PlaybackControllerPreviousCommandIssuedType = "PlaybackController.PreviousCommandIssued"
)

type PlaybackControllerRequestHandler func(AudioPlayerDirectives, *PlaybackControllerRequest) error

type PlaybackControllerRequest struct {
	Version string                        `json:"version"`
	Context request.Context               `json:"context"`
	Request PlaybackControllerRequestData `json:"request"`
}

type PlaybackControllerRequestData struct {
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Locale    string `json:"locale"`
}
