package alexa

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	certificateDNSName          = "echo-api.amazon.com"
	maxTimeDrift                = 150 * time.Second
	signatureCertChainURLHeader = "SignatureCertChainUrl"
	signatureHeader             = "Signature"
)

// verifyRequest takes an Alexa request body and ensures it meets the conditions
// specified in the amazon documentation for request verification.
//
// https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-a-web-service.html#verifying-that-the-request-was-sent-by-alexa
func verifyRequest(r *http.Request, b *body) error {
	if err := verifyTimestamp(b.Request.Timestamp); err != nil {
		return err
	}

	cert, err := getAndVerifyCert(r.Header.Get(signatureCertChainURLHeader))
	if err != nil {
		return fmt.Errorf("failed to obtain cert: %s", err)
	}

	return verifySignature(r.Header.Get(signatureHeader), cert, b.bs)
}

// getAndVerifyCert obtains a certificate from the given URL and verifies it's
// has a valid root authority and belongs to the correct domain.
func getAndVerifyCert(url string) (*x509.Certificate, error) {
	err := verifyURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to validate url: %s", err)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected response: %s", err)
	}

	var buf bytes.Buffer
	defer resp.Body.Close()
	io.Copy(&buf, resp.Body)
	bs := buf.Bytes()

	block, rest := pem.Decode(bs)
	cp := x509.NewCertPool()
	cp.AppendCertsFromPEM(rest)

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	opts := x509.VerifyOptions{DNSName: certificateDNSName, Intermediates: cp}

	if _, err := cert.Verify(opts); err != nil {
		return nil, err
	}

	return cert, nil
}

// verifySignature ensures the request data and signature are valid.
func verifySignature(b64Sig string, cert *x509.Certificate, bs []byte) error {
	sig, err := base64.RawStdEncoding.WithPadding(base64.StdPadding).DecodeString(b64Sig)
	if err != nil {
		return fmt.Errorf("failed to decode signature: %s", err)
	}

	// TODO(benjic): Determine if cert can be used to get SignatureAlgorithm.
	// The cert.SignatureAlgorithm value was indicating that sha256 was used
	// which resulted in bad verifications.
	return cert.CheckSignature(x509.SHA1WithRSA, bs, sig)
}

// verifyTimestamp ensures the request timestamp is within temporal tolerance.
func verifyTimestamp(timestamp string) error {
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return err
	}

	if time.Since(t) > maxTimeDrift || time.Since(t) < -maxTimeDrift {
		return fmt.Errorf("request timestamp out of tolerance")
	}

	return nil
}

// verifyURL will validate the signature certificate URL.
func verifyURL(urlStr string) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	if strings.ToLower(u.Scheme) != "https" {
		return fmt.Errorf("scheme mismatch")
	}
	if strings.ToLower(u.Hostname()) != "s3.amazonaws.com" {
		return fmt.Errorf("hostname mismatch")
	}
	if u.Port() != "" && u.Port() != "443" {
		return fmt.Errorf("port mismatch")
	}
	if !strings.HasPrefix(u.Path, "/echo.api/") {
		return fmt.Errorf("path prefix mismatched")
	}
	return nil
}
