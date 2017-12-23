package alexa

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationMux_Handle(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		mux     map[string]*Handler
	}
	type args struct {
		applicationID string
		h             *Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"register handler",
			fields{sync.RWMutex{}, map[string]*Handler{}},
			args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ApplicationMux{
				RWMutex: tt.fields.RWMutex,
				mux:     tt.fields.mux,
			}
			assert.Empty(t, m.mux)
			m.Handle(tt.args.applicationID, tt.args.h)

			assert.Len(t, m.mux, 1)
			assert.Contains(t, m.mux, tt.args.applicationID)
		})
	}
}

func TestApplicationMux_ServeHTTP(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		mux     map[string]*Handler
	}
	type args struct {
		r          io.Reader
		getBodyErr error
	}
	type want struct {
		code int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			"GetBody Error",
			fields{sync.RWMutex{}, map[string]*Handler{}},
			args{
				new(bytes.Buffer),
				errors.New("get body error"),
			},
			want{http.StatusBadRequest},
		},
		{
			"bad request body",
			fields{sync.RWMutex{}, map[string]*Handler{}},
			args{
				bytes.NewBufferString("not json at all"),
				nil,
			},
			want{http.StatusBadRequest},
		},
		{
			"missing handler for application id",
			fields{sync.RWMutex{}, map[string]*Handler{}},
			args{
				bytes.NewBufferString(
					`{
						"context": {
							"system": {
								"application": {
									"applicationId": "test-id"
								}
							}
						}
					}`),
				nil,
			},
			want{http.StatusNotFound},
		},
		{
			"found application handler",
			fields{sync.RWMutex{}, map[string]*Handler{
				"test-id": &Handler{},
			}},
			args{
				bytes.NewBufferString(
					`{
						"context": {
							"system": {
								"application": {
									"applicationId": "test-id"
								}
							}
						}
					}`),
				nil,
			},
			want{http.StatusBadRequest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ApplicationMux{
				RWMutex: tt.fields.RWMutex,
				mux:     tt.fields.mux,
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", tt.args.r)
			r.GetBody = func() (io.ReadCloser, error) { return r.Body, tt.args.getBodyErr }

			m.ServeHTTP(w, r)

			assert.Equal(t, tt.want.code, w.Code)
		})
	}
}
