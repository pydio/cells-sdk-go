package config

// SdkConfig contains necessary data to connect to API
type SdkConfig struct {
	//  http, https or wss
	Protocol string `json:"protocol"`
	// server name or IP & port   to the server
	Url        string `json:"url"`
	SkipVerify bool   `json:"skipVerify"`
	// OIDC ClientKey / ClientSecret
	ClientKey    string `json:"clientKey"`
	ClientSecret string `json:"clientSecret"`
	// Pydio User Authentication
	User     string `json:"user"`
	Password string `json:"password"`
}

// S3Config contains necessary data to connect to S3 gateway
type S3Config struct {
	Bucket string `json:"bucket"`
	// Endpoint override default URL generated from S3 SDK from the bucket name
	Endpoint  string `json:"enpoint"`
	Region    string `json:"region"`
	ApiKey    string `json:"apiKey"`
	ApiSecret string `json:"apiSecret"`
	IsDebug   bool   `json:"isDebug"`
}

var (
	DefaultConfig   *SdkConfig
	DefaultS3Config *S3Config
)
