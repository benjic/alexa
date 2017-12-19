# alexa 

> A set of utilities for handling Amazon Alexa requests.

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
