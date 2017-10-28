package alexa

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/benjic/alexa/request"
)

type IntentRequestHandlerFunc func(Response, *request.Intent) error
type LaunchRequestHandlerFunc func(Response, *request.Launch) error
type SessionEndedRequestHandlerFunc func(*request.SessionEnded) error

// Handler allows for custom behavior to be attributed to specific request
// types.
type Handler struct {
	IntentRequest       IntentRequestHandlerFunc
	LaunchRequest       LaunchRequestHandlerFunc
	SessionEndedRequest SessionEndedRequestHandlerFunc
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

	resp := &responseBuilder{Response: &response{}, Version: version}

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
	}

	return nil, nil
}

func validateRequest(r *http.Request) error {
	// TODO(benjic): validate request headers
	return nil
}
