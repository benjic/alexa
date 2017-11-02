package alexa

const (
	audioPlayerPlaybackStartedType        = "AudioPlayer.PlaybackStarted"
	audioPlayerPlaybackFinishedType       = "AudioPlayer.PlaybackFinished"
	audioPlayerPlaybackStoppedType        = "AudioPlayer.PlaybackStopped"
	audioPlayerPlaybackNearlyFinishedType = "AudioPlayer.PlaybackNearlyFinshed"
	audioPlayerPlaybackFailedType         = "AudioPlayer.PlaybackFailed"
)

type AudioPlaybackStopperQueueClearerHandler func(AudioPlayerStopperQueueClearer, *AudioPlaybackRequest) error
type AudioPlaybackStoppedHandler func(*AudioPlaybackRequest) error
type AudioPlaybackDirectiveHandler func(AudioPlayerDirectives, *AudioPlaybackRequest) error
type AudioPlaybackFailedHandler func(AudioPlayerDirectives, *AudioPlaybackFailedRequest) error

type AudioPlaybackRequest struct {
	Version string                   `json:"version"`
	Context RequestContext           `json:"context"`
	Request AudioPlaybackRequestData `json:"request"`
}

type AudioPlaybackFailedRequest struct {
	Version string                         `json:"version"`
	Context RequestContext                 `json:"context"`
	Request AudioPlaybackFailedRequestData `json:"request"`
}

type AudioPlaybackRequestData struct {
	Type                 string `json:"type"`
	RequestID            string `json:"requestId"`
	Timestamp            string `json:"timestamp"`
	Token                string `json:"token"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	Locale               string `json:"locale"`
}

type AudioPlaybackFailedRequestData struct {
	AudioPlaybackRequestData
	Error AudioPlaybackFailedError `json:"error"`
}

type AudioPlaybackFailedError struct {
	Type          string                           `json:"type"`
	Message       string                           `json:"message"`
	PlaybackState AudioPlaybackFailedPlaybackState `json:"currentPlaybackState"`
}

type AudioPlaybackFailedPlaybackState struct {
	Token                string `json:"token"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	PlayerActivity       string `json:"playerActivity"`
}

type AudioPlaybackStopOrClearResponse interface {
	StopAudio()
	ClearEnqueuedAudio()
	ClearAllAudio()
}

type AudioPlayer interface {
	ReplaceAllAudio(token, url string, offsetInMilliseconds int)
	EnqueueAudio(token, url, expectedPreviousToken string, offsetInMilliseconds int)
	ReplacedEnqueuedAudio(token, url string, offsetInMilliseconds int)
}

type AudioStopper interface {
	StopAudio()
}

type AudioQueueClearer interface {
	ClearEnqueuedAudio()
	ClearAllAudio()
}

type AudioPlayerDirectives interface {
	AudioPlayer
	AudioStopper
	AudioQueueClearer
}

type AudioPlayerStopperQueueClearer interface {
	AudioStopper
	AudioQueueClearer
}
