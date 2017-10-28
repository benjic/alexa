package directives

import (
	"github.com/benjic/alexa/request"
)

const (
	AudioPlayerPlaybackStartedType        = "AudioPlayer.PlaybackStarted"
	AudioPlayerPlaybackFinishedType       = "AudioPlayer.PlaybackFinished"
	AudioPlayerPlaybackStoppedType        = "AudioPlayer.PlaybackStopped"
	AudioPlayerPlaybackNearlyFinishedType = "AudioPlayer.PlaybackNearlyFinshed"
	AudioPlayerPlaybackFailedType         = "AudioPlayer.PlaybackFailed"
)

type AudioPlaybackStopperQueueClearerHandler func(AudioPlayerStopperQueueClearer, *AudioPlaybackRequest) error
type AudioPlaybackStoppedHandler func(*AudioPlaybackRequest) error
type AudioPlaybackDirectiveHandler func(AudioPlayerDirective, *AudioPlaybackRequest) error
type AudioPlaybackFailedHandler func(AudioPlayerDirective, *AudioPlaybackFailedRequest) error

type AudioPlaybackRequest struct {
	Version string                   `json:"version"`
	Context request.Context          `json:"context"`
	Request AudioPlaybackRequestData `json:"request"`
}

type AudioPlaybackFailedRequest struct {
	Version string                         `json:"version"`
	Context request.Context                `json:"context"`
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

type AudioPlayerDirective interface {
	AudioPlayer
	AudioStopper
	AudioQueueClearer
}

type AudioPlayerStopperQueueClearer interface {
	AudioStopper
	AudioQueueClearer
}
