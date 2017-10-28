package alexa

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/benjic/alexa/interfaces"
	"github.com/benjic/alexa/request"
)

// Handler allows for custom behavior to be attributed to specific request
// types.
type Handler struct {
	// Standard Request Handlers

	IntentRequest       func(Response, *request.Intent) error
	LaunchRequest       func(Response, *request.Launch) error
	SessionEndedRequest request.SessionEndedRequestHandler

	// Audio Request Handlers

	AudioPlaybackStartedRequest        interfaces.AudioPlaybackStopperQueueClearerHandler
	AudioPlaybackFinishedRequest       interfaces.AudioPlaybackStopperQueueClearerHandler
	AudioPlaybackStoppedRequest        interfaces.AudioPlaybackStoppedHandler
	AudioPlaybackNearlyFinishedRequest interfaces.AudioPlaybackDirectiveHandler
	AudioPlaybackFailedRequest         interfaces.AudioPlaybackFailedHandler

	// Playback Controller Handlers

	PlaybackControllerNextCommandRequest     interfaces.PlaybackControllerRequestHandler
	PlaybackControllerPausedCommandRequest   interfaces.PlaybackControllerRequestHandler
	PlaybackControllerPlayCommandRequest     interfaces.PlaybackControllerRequestHandler
	PlaybackControllerPreviousCommandRequest interfaces.PlaybackControllerRequestHandler

	SystemExceptionRequest interfaces.SystemExceptionEncounteredHandler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := validateRequest(r); err != nil {
		// Invalid requests are dropped immediately.
		return
	}

	var buf bytes.Buffer
	// TODO(benjic): Handle ReadFrom error
	buf.ReadFrom(r.Body)

	resp, err := h.routeRequest(buf.Bytes())

	if err != nil {
		// Any error should respond with a 500.
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if resp != nil {
		// Any non nil response should be written.
		w.Header().Add("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		// TODO(benjic): Handle json error
		json.NewEncoder(w).Encode(resp)
	}
}

func (h *Handler) routeRequest(bs []byte) (Response, error) {
	var p struct {
		Request struct {
			Type string `json:"type"`
		} `json:"request"`
	}

	if err := json.Unmarshal(bs, &p); err != nil {
		return nil, err
	}

	resp := &responseBuilder{Version: version, Response: &response{}}

	switch p.Request.Type {
	case request.LaunchRequestType:
		if h.LaunchRequest != nil {
			req := &request.Launch{}
			if err := json.Unmarshal(bs, req); err != nil {
				return resp, err
			}
			return resp, h.LaunchRequest(resp, req)
		}
	case request.IntentRequestType:
		if h.IntentRequest != nil {
			req := &request.Intent{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return resp, h.IntentRequest(resp, req)
		}
	case request.SessionEndedRequestType:
		if h.SessionEndedRequest != nil {
			req := &request.SessionEnded{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return nil, h.SessionEndedRequest(req)
		}
	case interfaces.AudioPlayerPlaybackFailedType:
		if h.AudioPlaybackFailedRequest != nil {
			req := &interfaces.AudioPlaybackFailedRequest{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return nil, h.AudioPlaybackFailedRequest(resp, req)
		}
	case interfaces.AudioPlayerPlaybackStartedType:
		if h.AudioPlaybackStartedRequest != nil {
			req := &interfaces.AudioPlaybackRequest{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return resp, h.AudioPlaybackStartedRequest(resp, req)
		}
	case interfaces.AudioPlayerPlaybackStoppedType:
		if h.AudioPlaybackStoppedRequest != nil {
			req := &interfaces.AudioPlaybackRequest{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return resp, h.AudioPlaybackStoppedRequest(req)
		}
	case interfaces.AudioPlayerPlaybackFinishedType:
		if h.AudioPlaybackFinishedRequest != nil {
			req := &interfaces.AudioPlaybackRequest{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return resp, h.AudioPlaybackFinishedRequest(resp, req)
		}
	case interfaces.AudioPlayerPlaybackNearlyFinishedType:
		if h.AudioPlaybackNearlyFinishedRequest != nil {
			req := &interfaces.AudioPlaybackRequest{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return resp, h.AudioPlaybackNearlyFinishedRequest(resp, req)
		}
	case interfaces.PlaybackControllerNextCommandIssuedType:
		if h.PlaybackControllerNextCommandRequest != nil {
			req := &interfaces.PlaybackControllerRequest{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return nil, h.PlaybackControllerNextCommandRequest(resp, req)
		}
	case interfaces.PlaybackControllerPlayCommandIssuedType:
		if h.PlaybackControllerPlayCommandRequest != nil {
			req := &interfaces.PlaybackControllerRequest{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return nil, h.PlaybackControllerPlayCommandRequest(resp, req)
		}
	case interfaces.PlaybackControllerPausedCommandIssuedType:
		if h.PlaybackControllerPausedCommandRequest != nil {
			req := &interfaces.PlaybackControllerRequest{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return nil, h.PlaybackControllerPausedCommandRequest(resp, req)
		}
	case interfaces.PlaybackControllerPreviousCommandIssuedType:
		if h.PlaybackControllerPreviousCommandRequest != nil {
			req := &interfaces.PlaybackControllerRequest{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return nil, h.PlaybackControllerPreviousCommandRequest(resp, req)
		}
	case interfaces.SystemExceptionEncounteredType:
		if h.SystemExceptionRequest != nil {
			req := &interfaces.SystemExceptionRequest{}
			if err := json.Unmarshal(bs, req); err != nil {
				return nil, err
			}
			return nil, h.SystemExceptionRequest(req)
		}
	}
	return nil, nil
}

func validateRequest(r *http.Request) error {
	// TODO(benjic): validate request headers
	return nil
}
