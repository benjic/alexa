package alexa

const (
	playbackControllerNextCommandIssuedType     = "PlaybackController.NextCommandIssued"
	playbackControllerPausedCommandIssuedType   = "PlaybackController.PausedCommandIssued"
	playbackControllerPlayCommandIssuedType     = "PlaybackController.PlayCommandIssued"
	playbackControllerPreviousCommandIssuedType = "PlaybackController.PreviousCommandIssued"
)

type PlaybackControllerRequestHandler func(AudioPlayerDirectives, *PlaybackControllerRequest) error

type PlaybackControllerRequest struct {
	Version string                        `json:"version"`
	Context RequestContext                `json:"context"`
	Request PlaybackControllerRequestData `json:"request"`
}

type PlaybackControllerRequestData struct {
	RequestID string `json:"requestId"`
	Timestamp string `json:"timestamp"`
	Locale    string `json:"locale"`
}
