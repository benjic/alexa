# alexa 

> A set of utilities for handling Amazon Alexa requests.

[![Build Status](https://travis-ci.org/benjic/alexa.svg?branch=master)](https://travis-ci.org/benjic/alexa) [![codecov](https://codecov.io/gh/benjic/alexa/branch/master/graph/badge.svg)](https://codecov.io/gh/benjic/alexa) [![Go Report Card](https://goreportcard.com/badge/github.com/benjic/alexa)](https://goreportcard.com/report/github.com/benjic/alexa) [![GoDoc](https://godoc.org/github.com/benjic/alexa?status.svg)](https://godoc.org/github.com/benjic/alexa)


This library provides a way of defining handlers for the various Alexa requests.
A handler is required to accept typed response and request objects to ensure
the response respects the amazon expected values.

```go
func main() {
  http.ListenAndServe(":8080", &alexa.Handler{
    LaunchRequest: func(resp alexa.Response, req *alexa.LaunchRequest) error {
      resp.PlainText("Hello world!")
      return nil
    },
  })
}
```
