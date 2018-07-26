package config

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetPreparedS3CLient creates and configure a new S3 client at each request.
// Should be optimized
func GetPreparedS3CLient(sdkConfig *SdkConfig, s3Config *S3Config) (*s3.S3, error) {

	var apiKey string
	var err error

	if s3Config.UsePydioSpecificHeader { // Legacy
		apiKey = s3Config.ApiKey
	} else {
		apiKey, err = retrieveToken(sdkConfig)
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve token from config, cause: %s", err.Error())
		}
	}

	conf := &aws.Config{
		// Static credentials are the best we've found to do the job until now. Might be enhanced, together with a better session pool management
		Credentials: credentials.NewStaticCredentials(apiKey, s3Config.ApiSecret, ""),
	}

	tr := http.DefaultTransport
	if s3Config.UsePydioSpecificHeader { // Legacy:
		tr = NewAuthTransport(tr, sdkConfig, s3Config)
	}

	client := &http.Client{Transport: tr}
	conf.WithHTTPClient(client).WithEndpoint(s3Config.Endpoint).WithRegion(s3Config.Region)

	s3Session, err := session.NewSession(conf)
	if err != nil {
		log.Fatal("cannot initialise default S3 session: " + err.Error())
	}

	return s3.New(s3Session), nil
}

// NewAuthTransport takes an http.RoundTripper and returns a new one that adds the JWT Auth
// header on each request (and optionally logs the additional cost of getting the JWT token)
func NewAuthTransport(rt http.RoundTripper, sdkConfig *SdkConfig, s3Config *S3Config) http.RoundTripper {
	return &authRoundTripper{rt: rt, sc: sdkConfig, s3c: s3Config, log: DefaultLogger{}}
}

type authRoundTripper struct {
	rt  http.RoundTripper
	log HTTPLogger
	sc  *SdkConfig
	s3c *S3Config
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
