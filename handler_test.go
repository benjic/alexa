package alexa_test

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/benjic/alexa"
)

const (
	launchRequest = `{
		"version": "1.0",
		"session": {
		  "new": true,
		  "sessionId": "amzn1.echo-api.session.0000000-0000-0000-0000-00000000000",
		  "application": {
			"applicationId": "amzn1.echo-sdk-ams.app.000000-d0ed-0000-ad00-000000d00ebe"
		  },
		  "attributes": {},
		  "user": {
			"userId": "amzn1.account.AM3B00000000000000000000000"
		  }
		},
		"context": {
		  "System": {
			"application": {
			  "applicationId": "amzn1.echo-sdk-ams.app.000000-d0ed-0000-ad00-000000d00ebe"
			},
			"user": {
			  "userId": "amzn1.account.AM3B00000000000000000000000"
			},
			"device": {
			  "supportedInterfaces": {
				"AudioPlayer": {}
			  }
			}
		  },
		  "AudioPlayer": {
			"offsetInMilliseconds": 0,
			"playerActivity": "IDLE"
		  }
		},
		"request": {
		  "type": "LaunchRequest",
		  "requestId": "amzn1.echo-api.request.0000000-0000-0000-0000-00000000000",
		  "timestamp": "2015-05-13T12:34:56Z",
		  "locale": "string"
		}
	  }`

	intentRequest = `{
		"version": "1.0",
		"session": {
		  "new": false,
		  "sessionId": "amzn1.echo-api.session.0000000-0000-0000-0000-00000000000",
		  "application": {
			"applicationId": "amzn1.echo-sdk-ams.app.000000-d0ed-0000-ad00-000000d00ebe"
		  },
		  "attributes": {
			"supportedHoroscopePeriods": {
			  "daily": true,
			  "weekly": false,
			  "monthly": false
			}
		  },
		  "user": {
			"userId": "amzn1.account.AM3B00000000000000000000000"
		  }
		},
		"context": {
		  "System": {
			"application": {
			  "applicationId": "amzn1.echo-sdk-ams.app.000000-d0ed-0000-ad00-000000d00ebe"
			},
			"user": {
			  "userId": "amzn1.account.AM3B00000000000000000000000"
			},
			"device": {
			  "supportedInterfaces": {
				"AudioPlayer": {}
			  }
			}
		  },
		  "AudioPlayer": {
			"offsetInMilliseconds": 0,
			"playerActivity": "IDLE"
		  }
		},
		"request": {
		  "type": "IntentRequest",
		  "requestId": " amzn1.echo-api.request.0000000-0000-0000-0000-00000000000",
		  "timestamp": "2015-05-13T12:34:56Z",
		  "dialogState": "COMPLETED",
		  "locale": "string",
		  "intent": {
			"name": "GetZodiacHoroscopeIntent",
			"confirmationStatus": "NONE",
			"slots": {
			  "ZodiacSign": {
				"name": "ZodiacSign",
				"value": "virgo",
				"confirmationStatus": "NONE"
			  }
			}
		  }
		}
	  }`

	sessionEndedRequest = `{
	"version": "1.0",
	"session": {
		"new": false,
		"sessionId": "amzn1.echo-api.session.0000000-0000-0000-0000-00000000000",
		"application": {
		"applicationId": "amzn1.echo-sdk-ams.app.000000-d0ed-0000-ad00-000000d00ebe"
		},
		"attributes": {
		"supportedHoroscopePeriods": {
			"daily": true,
			"weekly": false,
			"monthly": false
		}
		},
		"user": {
		"userId": "amzn1.account.AM3B00000000000000000000000"
		}
	},
	"context": {
		"System": {
		"application": {
			"applicationId": "amzn1.echo-sdk-ams.app.000000-d0ed-0000-ad00-000000d00ebe"
		},
		"user": {
			"userId": "amzn1.account.AM3B00000000000000000000000"
		},
		"device": {
			"supportedInterfaces": {
			"AudioPlayer": {}
			}
		}
		},
		"AudioPlayer": {
		"offsetInMilliseconds": 0,
		"playerActivity": "IDLE"
		}
	},
	"request": {
		"type": "SessionEndedRequest",
		"requestId": "amzn1.echo-api.request.0000000-0000-0000-0000-00000000000",
		"timestamp": "2015-05-13T12:34:56Z",
		"reason": "USER_INITIATED",
		"locale": "en-US"
	}
	}`
)

func launchRequestHandler(resp alexa.Response, req *alexa.LaunchRequest) error {
	return nil
}

func intentRequestHandler(resp alexa.Response, req *alexa.IntentRequest) error {
	return nil
}

func sessionEndedRequestHandler(req *alexa.SessionEndedRequest) error {
	return nil
}

func newTestHandler() *alexa.Handler {
	return &alexa.Handler{
		IntentRequest:       intentRequestHandler,
		LaunchRequest:       launchRequestHandler,
		SessionEndedRequest: sessionEndedRequestHandler,
	}
}

func TestLaunchRequest(t *testing.T) {
	h := newTestHandler()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(launchRequest)))
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	t.Logf("%+v", resp.StatusCode)
	t.Logf("%+v", resp.Header)
	t.Logf("%+v", string(body))
}

func TestIntentRequest(t *testing.T) {
	h := newTestHandler()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(intentRequest)))
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	t.Logf("%+v", resp.StatusCode)
	t.Logf("%+v", resp.Header)
	t.Logf("%+v", string(body))
}

func TestSessionEndedRequest(t *testing.T) {
	h := newTestHandler()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(sessionEndedRequest)))
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	t.Logf("%+v", resp.StatusCode)
	t.Logf("%+v", resp.Header)
	t.Logf("%+v", string(body))
}
