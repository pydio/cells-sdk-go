// Package cells_sdk provides a ready to use SDK to use the Cells REST API in Go language.
// It also provides some patterns that ease implementation of Go programs that use the SDK.
package cells_sdk

import (
	"fmt"
	"net/url"
	"time"
)

// SdkConfig stores parameters to talk to a running Cells instance REST API via the Go SDK.
type SdkConfig struct {

	// Auth type is a convenience flag to store the type of authentication used by this SDK config
	// in v5, we support: PAT (Personal Access Token), OAuth2 (based on a JWT retrieved interreactively via OAuth credential flow) and Basic (login/password, less secured)
	AuthType string `json:"authType,omitempty"`

	// Url stores domain name or IP & port to the server.
	Url string `json:"url"`
	// Username (login) for the currenly configured Pydio Account
	User string `json:"user"`

	// IdToken might be a personal access Token (generated on your server) or an OAuth2 Token retrieved via the OIDC code flow.
	IdToken string `json:"idToken,omitempty"`

	// OIDC Code Flow additional info
	// RefreshToken holds the token to refresh your JWT. Warning: this token can be used only **once**, it is then blocked on the server side.
	RefreshToken string `json:"refreshToken,omitempty"`
	// TokenExpiresAt holds the expiration timestamp for the current JWT.
	TokenExpiresAt int `json:"tokenExpiresAt,omitempty"`

	// Password for client credential authentification (Legacy, less secure).
	Password string `json:"password,omitempty"`

	// SkipVerify tells the transport to ignore expired or self-signed TLS certificates.
	SkipVerify bool `json:"skipVerify"`

	// UseTokenCache flags wether we should rely on a local cache to avoid retrieving a new JWT token at each request.
	// It is useful to *not* use the cache when running connection tests for instance.
	UseTokenCache bool `json:"useTokenCache"`

	// CustomHeaders holds an optional list of headers to be overriden in requests, e.g the User-Agent.
	CustomHeaders map[string]string
}

const (
	// Supported Auth types for v5+.
	// AuthTypePat relies on a Personal Access Token generated on the server for a given user.
	AuthTypePat = "PAT"
	// AuthTypeClientAuth is the legacy authentication method, based on user password: this is less secured.
	AuthTypeClientAuth = "Basic"
	// AuthTypeOAuth uses OAuth2 credential retrieval flow.
	AuthTypeOAuth = "OAuth2"
)

// Make SdkConfig implement the TokenProvider interface

// Retrieve simply returns the ID token that is currently held in the conf.
func (s *SdkConfig) Retrieve() (string, error) {
	// fmt.Println("[Debug] Retrieving token: [", s.IdToken, "], Expires at:", time.Unix(int64(s.TokenExpiresAt), 0), "\n ")
	return s.IdToken, nil
}

// Expired checks if expiration time is in less than 10 seconds or already passed.
func (s *SdkConfig) Expired() bool {
	isExpired := time.Now().After(time.Unix(int64(s.TokenExpiresAt), 0).Add(-10 * time.Second))
	// fmt.Println("[Debug] Checking validity of [", s.IdToken, "], expired =", isExpired, "\n ")
	return isExpired
}

func (s *SdkConfig) GetId() string {
	var port string
	u, _ := url.Parse(s.Url)
	port = u.Port()
	if port == "" {
		switch u.Scheme {
		case "http":
			port = "80"
		case "https":
			port = "443"
		}
	}
	id := fmt.Sprintf("%s@%s:%s", s.User, u.Hostname(), port)
	return id
}

// S3Config stores connection parameters to a running Cells instance S3 gateway via the AWS SDK for Go.
type S3Config struct {
	// Endpoint overrides the default URL generated by the S3 SDK from the bucket name.
	Endpoint string `json:"enpoint"`
	// Region param, usually us-east-1.
	Region string `json:"region"`
	// ApiKey is used by the Cells SDK to transmit the JWT token.
	ApiKey string `json:"apiKey"`
	// ApiSecret identifies this client.
	ApiSecret string `json:"apiSecret"`
	// Bucket name, usually io.
	Bucket string `json:"bucket"`

	// RequestTimout overrides default HTTP request timeout when transfering files.
	RequestTimout int `json:"requestTimout,omitempty"`

	// Legacy
	// Set to true to rather use legacy mode (JWT Auth is transmitted via custom 'X-Pydio-Bearer' header).
	UsePydioSpecificHeader bool `json:"usePydioSpecificHeader"`
	// IsDebug is a convenience legacy flag to add some logging to this S3 client.
	// Should be cleaned as soon as we defined the logging strategy for this repo.
	IsDebug bool `json:"isDebug"`
}

func NewS3Config() *S3Config {
	return &S3Config{
		Bucket:                 "io",
		ApiSecret:              "gatewaysecret",
		Region:                 "us-east-1",
		RequestTimout:          -1,
		UsePydioSpecificHeader: false,
		IsDebug:                false,
	}
}

// Log is a temporary hack to help while debugging.
func Log(format string, a ...any) {
	// TODO implement a real logger support
	//   also in between Un-comment addtional strings
	//   below to debug while having a progress bar (scp...)
	// fmt.Println("")
	// fmt.Println("")
	fmt.Println(fmt.Sprintf(format, a...))
	// fmt.Println("")
	// fmt.Println("")
}
