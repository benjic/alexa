// Package alexa provides a way to write typed handlers for requests made from
// the Alexa Skill service.
//
//
package alexa

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type body struct {
	Context struct {
		System struct {
			Application struct {
				ApplicationID string `json:"applicationId"`
			} `json:"application"`
		} `json:"system"`
	} `json:"context"`
	Request struct {
		Type      string `json:"type"`
		Timestamp string `json:"timestamp"`
	}
	bs []byte
}

// Handler allows for custom behavior to be attributed to specific request
// types.
type Handler struct {
	// Standard Request Handlers

	IntentRequest       func(Response, *IntentRequest) error
	LaunchRequest       func(Response, *LaunchRequest) error
	SessionEndedRequest SessionEndedRequestHandler

	// Audio Request Handlers

	AudioPlaybackFailedRequest         AudioPlaybackFailedHandler
	AudioPlaybackFinishedRequest       AudioStopperQueueClearerHandler
	AudioPlaybackNearlyFinishedRequest AudioPlayerStopperQueueClearerHandler
	AudioPlaybackStartedRequest        AudioStopperQueueClearerHandler
	AudioPlaybackStoppedRequest        AudioPlaybackStoppedHandler

	// Playback Controller Handlers

	PlaybackControllerNextCommandRequest     PlaybackControllerRequestHandler
	PlaybackControllerPausedCommandRequest   PlaybackControllerRequestHandler
	PlaybackControllerPlayCommandRequest     PlaybackControllerRequestHandler
	PlaybackControllerPreviousCommandRequest PlaybackControllerRequestHandler

	SystemExceptionRequest SystemExceptionEncounteredHandler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := parseRequestBody(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := verifyRequest(r, body); err != nil {
		// Invalid requests are dropped immediately.
		return
	}

	resp, err := h.routeRequest(body)
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

func (h *Handler) routeRequest(b *body) (Response, error) {
	resp := &responseBuilder{Version: version, Response: &response{}}

	switch b.Request.Type {
	case launchRequestType:
		if h.LaunchRequest != nil {
			req := &LaunchRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return resp, err
			}
			return resp, h.LaunchRequest(resp, req)
		}
	case intentRequestType:
		if h.IntentRequest != nil {
			req := &IntentRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return resp, h.IntentRequest(resp, req)
		}
	case sessionEndedRequestType:
		if h.SessionEndedRequest != nil {
			req := &SessionEndedRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return nil, h.SessionEndedRequest(req)
		}
	case audioPlayerPlaybackFailedType:
		if h.AudioPlaybackFailedRequest != nil {
			req := &AudioPlaybackFailedRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return nil, h.AudioPlaybackFailedRequest(resp, req)
		}
	case audioPlayerPlaybackStartedType:
		if h.AudioPlaybackStartedRequest != nil {
			req := &AudioPlaybackRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return resp, h.AudioPlaybackStartedRequest(resp, req)
		}
	case audioPlayerPlaybackStoppedType:
		if h.AudioPlaybackStoppedRequest != nil {
			req := &AudioPlaybackRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return resp, h.AudioPlaybackStoppedRequest(req)
		}
	case audioPlayerPlaybackFinishedType:
		if h.AudioPlaybackFinishedRequest != nil {
			req := &AudioPlaybackRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return resp, h.AudioPlaybackFinishedRequest(resp, req)
		}
	case audioPlayerPlaybackNearlyFinishedType:
		if h.AudioPlaybackNearlyFinishedRequest != nil {
			req := &AudioPlaybackRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return resp, h.AudioPlaybackNearlyFinishedRequest(resp, req)
		}
	case playbackControllerNextCommandIssuedType:
		if h.PlaybackControllerNextCommandRequest != nil {
			req := &PlaybackControllerRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return nil, h.PlaybackControllerNextCommandRequest(resp, req)
		}
	case playbackControllerPlayCommandIssuedType:
		if h.PlaybackControllerPlayCommandRequest != nil {
			req := &PlaybackControllerRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return nil, h.PlaybackControllerPlayCommandRequest(resp, req)
		}
	case playbackControllerPausedCommandIssuedType:
		if h.PlaybackControllerPausedCommandRequest != nil {
			req := &PlaybackControllerRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return nil, h.PlaybackControllerPausedCommandRequest(resp, req)
		}
	case playbackControllerPreviousCommandIssuedType:
		if h.PlaybackControllerPreviousCommandRequest != nil {
			req := &PlaybackControllerRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return nil, h.PlaybackControllerPreviousCommandRequest(resp, req)
		}
	case systemExceptionEncounteredType:
		if h.SystemExceptionRequest != nil {
			req := &SystemExceptionEncounteredRequest{}
			if err := json.Unmarshal(b.bs, req); err != nil {
				return nil, err
			}
			return nil, h.SystemExceptionRequest(req)
		}
	}
	return resp, nil
}

func parseRequestBody(r io.Reader) (*body, error) {
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}

	b := &body{bs: buf.Bytes()}

	if err := json.Unmarshal(b.bs, b); err != nil {
		return nil, err
	}

	return b, nil
}
