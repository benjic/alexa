package alexa

import (
	"net/http"
	"sync"
)

// An ApplicationMux allows many handlers to be added by applicationId.
type ApplicationMux struct {
	sync.RWMutex
	mux map[string]Handler
}

// Handle associates the given applicationID with a handler.
func (m *ApplicationMux) Handle(applicationID string, h Handler) {
	m.Lock()
	defer m.Unlock()
	m.mux[applicationID] = h
}

// ServeHTTP routes a request to the appropriate Handler registered for the
// applicationId provided in the request.
func (m *ApplicationMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	p, err := parseRequestBody(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m.RLock()
	defer m.RUnlock()
	if h, ok := m.mux[p.Context.System.Application.ApplicationID]; ok {
		h.ServeHTTP(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
