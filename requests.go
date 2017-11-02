package alexa

const (
	audioPlayerPlaybackStartedType              = "AudioPlayer.PlaybackStarted"
	audioPlayerPlaybackFinishedType             = "AudioPlayer.PlaybackFinished"
	audioPlayerPlaybackStoppedType              = "AudioPlayer.PlaybackStopped"
	audioPlayerPlaybackNearlyFinishedType       = "AudioPlayer.PlaybackNearlyFinshed"
	audioPlayerPlaybackFailedType               = "AudioPlayer.PlaybackFailed"
	intentRequestType                           = "IntentRequest"
	launchRequestType                           = "LaunchRequest"
	playbackControllerNextCommandIssuedType     = "PlaybackController.NextCommandIssued"
	playbackControllerPausedCommandIssuedType   = "PlaybackController.PausedCommandIssued"
	playbackControllerPlayCommandIssuedType     = "PlaybackController.PlayCommandIssued"
	playbackControllerPreviousCommandIssuedType = "PlaybackController.PreviousCommandIssued"
	sessionEndedRequestType                     = "SessionEndedRequest"
	systemExceptionEncounteredType              = "System.ExceptionEncountered"
)

type AudioPlaybackStopperQueueClearerHandler func(AudioPlayerStopperQueueClearer, *AudioPlaybackRequest) error

type AudioPlaybackStoppedHandler func(*AudioPlaybackRequest) error

type AudioPlaybackDirectiveHandler func(AudioPlayerDirectives, *AudioPlaybackRequest) error

type AudioPlaybackFailedHandler func(AudioPlayerDirectives, *AudioPlaybackFailedRequest) error

// A PlaybackControllerRequestHandler is a function that will receive a request
// payload when the controller state updates.
type PlaybackControllerRequestHandler func(AudioPlayerDirectives, *PlaybackControllerRequest) error

// A SessionEndedRequestHandler is a function that will receive a request
// payload when a session is ended.
type SessionEndedRequestHandler func(*SessionEndedRequest) error

// A SystemExceptionEncounteredHandler is a function that can receive a request
// payload when a hardware exception is encountered.
type SystemExceptionEncounteredHandler func(*SystemExceptionEncounteredRequest) error

// An AudioPlaybackFailedRequest represents the payload provided by Amazon when
// audio playback enters a failed state.
type AudioPlaybackFailedRequest struct {
	Version string `json:"version"`
	Context struct {
		System struct {
			APIEndpoint string `json:"apiEndpoint"`
			Application struct {
				ID string `json:"applicationId"`
			} `json:"application"`
			User struct {
				AccessToken string `json:"accessToken"`
				ID          string `json:"userId"`
				Permissions struct {
					ConsentToken string `json:"consentToken"`
				} `json:"permissions"`
			} `json:"user"`
		} `json:"system"`
	} `json:"context"`
	Request struct {
		Type                 string `json:"type"`
		RequestID            string `json:"requestId"`
		Timestamp            string `json:"timestamp"`
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
		Locale               string `json:"locale"`
		Error                struct {
			Type          string `json:"type"`
			Message       string `json:"message"`
			PlaybackState struct {
				Token                string `json:"token"`
				OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
				PlayerActivity       string `json:"playerActivity"`
			} `json:"currentPlaybackState"`
		} `json:"error"`
	} `json:"request"`
}

// An AudioPlaybackRequest represents the payload provided by amazon when the
// audio playback changes state.
type AudioPlaybackRequest struct {
	Version string `json:"version"`
	Context struct {
		System struct {
			APIEndpoint string `json:"apiEndpoint"`
			Application struct {
				ID string `json:"applicationId"`
			} `json:"application"`
			User struct {
				AccessToken string `json:"accessToken"`
				ID          string `json:"userId"`
				Permissions struct {
					ConsentToken string `json:"consentToken"`
				} `json:"permissions"`
			} `json:"user"`
		} `json:"system"`
	} `json:"context"`
	Request struct {
		Type                 string `json:"type"`
		RequestID            string `json:"requestId"`
		Timestamp            string `json:"timestamp"`
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
		Locale               string `json:"locale"`
	} `json:"request"`
}

// IntentRequest represents they payload provided by Amazon when an Alexa Intent
// request is made.
type IntentRequest struct {
	Version string `json:"version"`
	Context struct {
		System struct {
			APIEndpoint string `json:"apiEndpoint"`
			Application struct {
				ID string `json:"applicationId"`
			} `json:"application"`
			User struct {
				AccessToken string `json:"accessToken"`
				ID          string `json:"userId"`
				Permissions struct {
					ConsentToken string `json:"consentToken"`
				} `json:"permissions"`
			} `json:"user"`
		} `json:"system"`
	} `json:"context"`
	Session struct {
		Application struct {
			ID string `json:"applicationId"`
		} `json:"application"`
		Attributes map[string]interface{} `json:"attributes"`
		ID         string                 `json:"sessionId"`
		New        bool                   `json:"new"`
		User       struct {
			AccessToken string `json:"accessToken"`
			ID          string `json:"userId"`
			Permissions struct {
				ConsentToken string `json:"consentToken"`
			} `json:"permissions"`
		} `json:"user"`
	} `json:"session"`
	Request struct {
		RequestID   string `json:"requestId"`
		Timestamp   string `json:"timestamp"`
		DialogState string `json:"dialogState"`
		Locale      string `json:"locale"`
		Intent      struct {
			Name               string `json:"name"`
			ConfirmationStatus string `json:"confirmationStatus"`
			Slots              map[string]struct {
				Name               string `json:"name"`
				Value              string `json:"value"`
				ConfirmationStatus string `json:"confirmationStatus"`
				Resolutions        struct {
					ResolutionsPerAuthority []struct {
						Authority string `json:"authority"`
						Status    struct {
							Code string `json:"code"`
						} `json:"status"`
						Values []struct {
							Value struct {
								Name string `json:"name"`
								ID   string `json:"id"`
							} `json:"value"`
						} `json:"values"`
					} `json:"resolutionsPerAuthority"`
				} `json:"resolutions"`
			} `json:"slots"`
		} `json:"intent"`
	} `json:"request"`
}

// A LaunchRequest represents the payload provided by Amazon when a launch
// request is made.
type LaunchRequest struct {
	Version string `json:"version"`
	Context struct {
		System struct {
			APIEndpoint string `json:"apiEndpoint"`
			Application struct {
				ID string `json:"applicationId"`
			} `json:"application"`
			User struct {
				AccessToken string `json:"accessToken"`
				ID          string `json:"userId"`
				Permissions struct {
					ConsentToken string `json:"consentToken"`
				} `json:"permissions"`
			} `json:"user"`
		} `json:"system"`
	} `json:"context"`
	Session struct {
		Application struct {
			ID string `json:"applicationId"`
		} `json:"application"`
		Attributes map[string]interface{} `json:"attributes"`
		ID         string                 `json:"sessionId"`
		New        bool                   `json:"new"`
		User       struct {
			AccessToken string `json:"accessToken"`
			ID          string `json:"userId"`
			Permissions struct {
				ConsentToken string `json:"consentToken"`
			} `json:"permissions"`
		} `json:"user"`
	} `json:"session"`
	Request struct {
		Timestamp string `json:"timestamp"`
		RequestID string `json:"requestId"`
		Locale    string `json:"locale"`
	} `json:"request"`
}

// A PlaybackControllerRequest represents the payload provided by Amazon when
// a controller state updates.
type PlaybackControllerRequest struct {
	Version string `json:"version"`
	Context struct {
		System struct {
			APIEndpoint string `json:"apiEndpoint"`
			Application struct {
				ID string `json:"applicationId"`
			} `json:"application"`
			User struct {
				AccessToken string `json:"accessToken"`
				ID          string `json:"userId"`
				Permissions struct {
					ConsentToken string `json:"consentToken"`
				} `json:"permissions"`
			} `json:"user"`
		} `json:"system"`
	} `json:"context"`
	Request struct {
		RequestID string `json:"requestId"`
		Timestamp string `json:"timestamp"`
		Locale    string `json:"locale"`
	} `json:"request"`
}

// A SessionEndedRequest represents the payload provided by Amazon when a
// session ended request is made.
type SessionEndedRequest struct {
	Version string `json:"version"`
	Context struct {
		System struct {
			APIEndpoint string `json:"apiEndpoint"`
			Application struct {
				ID string `json:"applicationId"`
			} `json:"application"`
			User struct {
				AccessToken string `json:"accessToken"`
				ID          string `json:"userId"`
				Permissions struct {
					ConsentToken string `json:"consentToken"`
				} `json:"permissions"`
			} `json:"user"`
		} `json:"system"`
	} `json:"context"`
	Session struct {
		Application struct {
			ID string `json:"applicationId"`
		} `json:"application"`
		Attributes map[string]interface{} `json:"attributes"`
		ID         string                 `json:"sessionId"`
		New        bool                   `json:"new"`
		User       struct {
			AccessToken string `json:"accessToken"`
			ID          string `json:"userId"`
			Permissions struct {
				ConsentToken string `json:"consentToken"`
			} `json:"permissions"`
		} `json:"user"`
	} `json:"session"`
	Request struct {
		Timestamp string `json:"timestamp"`
		RequestID string `json:"requestId"`
		Reason    string `json:"reason"`
		Error     struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		} `json:"error"`
		Locale string `json:"locale"`
	} `json:"request"`
}

// A SystemExceptionEncounteredRequest represents the payload provided by Amazon
// when a system exception request is made.
type SystemExceptionEncounteredRequest struct {
	Version string `json:"version"`
	Context struct {
		System struct {
			APIEndpoint string `json:"apiEndpoint"`
			Application struct {
				ID string `json:"applicationId"`
			} `json:"application"`
			User struct {
				AccessToken string `json:"accessToken"`
				ID          string `json:"userId"`
				Permissions struct {
					ConsentToken string `json:"consentToken"`
				} `json:"permissions"`
			} `json:"user"`
		} `json:"system"`
	} `json:"context"`
	Request struct {
		Type      string `json:"type"`
		RequestID string `json:"requestId"`
		Timestamp string `json:"timestamp"`
		Locale    string `json:"locale"`
		Error     struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		} `json:"error"`
		Cause struct {
			RequestID string `json:"requestId"`
		} `json:"cause"`
	} `json:"request"`
}
