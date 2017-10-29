package alexa

import "testing"

func TestURLVerifier(t *testing.T) {
	cases := []struct {
		url string
		err bool
	}{
		{"https://s3.amazonaws.com/echo.api/echo-api-cert.pem", false},
		{"https://s3.amazonaws.com:443/echo.api/echo-api-cert.pem", false},
		{"https://s3.amazonaws.com/echo.api/../echo.api/echo-api-cert.pem", false},

		{"http://s3.amazonaws.com/echo.api/echo-api-cert.pem", true},
		{"https://notamazon.com/echo.api/echo-api-cert.pem", true},
		{"https://s3.amazonaws.com/EcHo.aPi/echo-api-cert.pem", true},
		{"https://s3.amazonaws.com/invalid.path/echo-api-cert.pem", true},
		{"https://s3.amazonaws.com:563/echo.api/echo-api-cert.pem", true},
	}

	for _, c := range cases {
		if _, err := verifyURL(c.url); (err == nil) == c.err {
			t.Errorf("Did want err %t; got %s for %s", c.err, err, c.url)
		}
	}
}
