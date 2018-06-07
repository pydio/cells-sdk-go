package config

import (
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	DefaultS3Session *session.Session
)

func InitDefaultSession(s3config S3Config) {
	conf := &aws.Config{
		Credentials: credentials.NewStaticCredentials(s3config.ApiKey, s3config.ApiSecret, ""),
	}

	tr := NewAuthTransport(http.DefaultTransport, s3config.IsDebug)
	client := &http.Client{Transport: tr}
	conf.WithHTTPClient(client).WithEndpoint(s3config.Endpoint).WithRegion(s3config.Region)

	var err error
	DefaultS3Session, err = session.NewSession(conf)
	if err != nil {
		log.Fatal("cannot initialise default S3 session: " + err.Error())
	}
}

type authRoundTripper struct {
	rt      http.RoundTripper
	log     HTTPLogger
	isDebug bool
}

func (c *authRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	var start time.Time
	if c.isDebug {
		c.log.LogRequest(request)
		start = time.Now()
	}
	jwt, err := retrieveToken(DefaultConfig)
	if err != nil {
		return nil, err
	}
	request.Header.Set(KeyS3BearerHeader, jwt)
	response, err := c.rt.RoundTrip(request)

	if c.isDebug {
		duration := time.Since(start)
		c.log.LogResponse(request, response, err, duration)
	}
	return response, err
}

// NewAuthTransport takes an http.RoundTripper and returns a new one that adds authorisation header and optionally logs requests and responses
func NewAuthTransport(rt http.RoundTripper, isDebug bool) http.RoundTripper {
	return &authRoundTripper{rt: rt, isDebug: isDebug, log: DefaultLogger{}}
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
