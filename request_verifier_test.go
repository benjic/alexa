package alexa

import (
	"crypto/x509"
	"testing"
	"time"
)

func TestURLVerifier(t *testing.T) {
	cases := []struct {
		url string
		err bool
	}{
		{"https://s3.amazonaws.com/echo.api/echo-api-cert.pem", false},
		{"https://s3.amazonaws.com:443/echo.api/echo-api-cert.pem", false},
		{"https://s3.amazonaws.com/echo.api/../echo.api/echo-api-cert.pem", false},

		{"://force-url-parse-to-fail", true},
		{"http://s3.amazonaws.com/echo.api/echo-api-cert.pem", true},
		{"https://notamazon.com/echo.api/echo-api-cert.pem", true},
		{"https://s3.amazonaws.com/EcHo.aPi/echo-api-cert.pem", true},
		{"https://s3.amazonaws.com/invalid.path/echo-api-cert.pem", true},
		{"https://s3.amazonaws.com:563/echo.api/echo-api-cert.pem", true},
	}

	for _, c := range cases {
		if err := verifyURL(c.url); (err == nil) == c.err {
			t.Errorf("Did want err %t; got %s for %s", c.err, err, c.url)
		}
	}
}

func TestVerifyTimestamp(t *testing.T) {
	cases := []struct {
		timestamp string
		err       bool
	}{
		// Future
		{time.Now().Add(100 * time.Second).Format(time.RFC3339), false},
		{time.Now().Add(500 * time.Second).Format(time.RFC3339), true},
		// Past
		{time.Now().Add(-100 * time.Second).Format(time.RFC3339), false},
		{time.Now().Add(-500 * time.Second).Format(time.RFC3339), true},
		// Wacky dates
		{"whoa", true},
	}

	for _, c := range cases {
		if err := verifyTimestamp(c.timestamp); (err == nil) == c.err {
			t.Errorf("Did want err %t; got %s for %s", c.err, err, c.timestamp)
		}
	}
}

func TestVerifySignature(t *testing.T) {
	type args struct {
		b64sig string
		cert   *x509.Certificate
		bs     []byte
	}
	cases := []struct {
		args args
		err  bool
	}{
		{args{b64sig: "this is not base64"}, true},
	}

	for _, c := range cases {
		if err := verifySignature(c.args.b64sig, c.args.cert, c.args.bs); (err == nil) == c.err {
			t.Errorf("Did want err %t; got %s\n\t%x\n\t%+v\t\n%x", c.err, err, c.args.b64sig, c.args.cert, c.args.bs)
		}
	}
}
