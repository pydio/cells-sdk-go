package config

// SdkConfig contains necessary data to connect to API
type SdkConfig struct {
	// Full Http(s) URL to the server
	Url        string `json:"url"`
	SkipVerify bool   `json:"skipVerify"`
	// OIDC ClientKey / ClientSecret
	ClientKey    string `json:"clientKey"`
	ClientSecret string `json:"clientSecret"`
	// Pydio User Authentication
	User     string `json:"user"`
	Password string `json:"password"`
}

var (
	DefaultConfig *SdkConfig
)
