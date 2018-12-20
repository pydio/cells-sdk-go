package transport

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"crypto/tls"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pydio/cells-sdk-go"
)

// GetPreparedS3CLient creates and configure a new S3 client at each request.
// Should be optimized
func GetS3CLient(sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (*s3.S3, error) {

	var apiKey string
	var err error

	if s3c.UsePydioSpecificHeader { // Legacy
		apiKey = s3c.ApiKey
	} else {
		apiKey, err = retrieveToken(sdc)
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve token from config, cause: %s", err.Error())
		}
	}

	conf := &aws.Config{
		// Static credentials are the best we've found to do the job until now. Might be enhanced, together with a better session pool management
		Credentials: credentials.NewStaticCredentials(apiKey, s3c.ApiSecret, ""),
	}

	tr := http.DefaultTransport
	if sdc.SkipVerify {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		}
	}
	if s3c.UsePydioSpecificHeader { // Legacy:
		tr = NewAuthTransport(tr, sdc, s3c)
	}

	client := &http.Client{Transport: tr}
	conf.WithHTTPClient(client).WithEndpoint(s3c.Endpoint).WithRegion(s3c.Region)

	s3Session, err := session.NewSession(conf)
	if err != nil {
		log.Fatal("cannot initialise default S3 session: " + err.Error())
	}

	// s3.New(s3Session).ListBuckets(nil)

	return s3.New(s3Session), nil
}

// NewAuthTransport takes an http.RoundTripper and returns a new one that adds the JWT Auth
// header on each request (and optionally logs the additional cost of getting the JWT token)
func NewAuthTransport(rt http.RoundTripper, sdkConfig *cells_sdk.SdkConfig, s3Config *cells_sdk.S3Config) http.RoundTripper {
	return &authRoundTripper{rt: rt, sc: sdkConfig, s3c: s3Config, log: DefaultLogger{}}
}

type authRoundTripper struct {
	rt  http.RoundTripper
	log HTTPLogger
	sc  *cells_sdk.SdkConfig
	s3c *cells_sdk.S3Config
}

func (c *authRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	var start time.Time
	if c.s3c.IsDebug {
		c.log.LogRequest(request)
		start = time.Now()
	}

	jwt, err := retrieveToken(c.sc)
	if err != nil {
		return nil, err
	}
	request.Header.Set(KeyS3BearerHeader, jwt)
	response, err := c.rt.RoundTrip(request)

	if c.s3c.IsDebug {
		duration := time.Since(start)
		c.log.LogResponse(request, response, err, duration)
	}
	return response, err
}

// HTTPLogger defines the interface to log http request and responses
type HTTPLogger interface {
	LogRequest(*http.Request)
	LogResponse(*http.Request, *http.Response, error, time.Duration)
}

// DefaultLogger logs basic information about the http responses.
type DefaultLogger struct {
}

// LogRequest does nothing.
func (dl DefaultLogger) LogRequest(*http.Request) {
}

// LogResponse logs path, host, status code and duration in milliseconds.
func (dl DefaultLogger) LogResponse(req *http.Request, res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	if err != nil {
		log.Printf("HTTP Request method=%s host=%s path=%s status=error durationMs=%d error=%q", req.Method, req.Host, req.URL.Path, duration, err.Error())
	} else {
		log.Printf("HTTP Request method=%s host=%s path=%s status=%d durationMs=%d", req.Method, req.Host, req.URL.Path, res.StatusCode, duration)
	}
}
