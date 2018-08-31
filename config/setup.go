package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	// Keys to retrieve configuration via environment variables
	KeyProtocol, KeyURL, KeyClientKey, KeyClientSecret, KeyUser, KeyPassword, KeySkipVerify = "TARGET_PROTOCOL", "TARGET_URL", "TARGET_CLIENT_KEY", "TARGET_CLIENT_SECRET", "TARGET_USER", "TARGET_PASSWORD", "TARGET_SKIP_VERIFY"

	// Keys to retrieve environment variables to configure connection to Pydio Cells S3 API
	KeyS3Endpoint, KeyS3Region, KeyS3Bucket, KeyS3ApiKey, KeyS3ApiSecret, KeyS3UsePydioSpecificHeader, KeyS3IsDebug = "TARGET_S3_ENDPOINT", "TARGET_S3_REGION", "TARGET_S3_BUCKET", "TARGET_S3_API_KEY", "TARGET_S3_API_SECRET", "TARGET_S3_USE_PYDIO_SPECIFIC_HEADER", "TARGET_S3_IS_DEBUG"
)

// SetUpEnvironment retrieves parameters and stores them in the DefaultConfig of the SDK.
// configFilePath and s3ConfigFilePath can be <nil> if the parameters are defined via env variables.
func SetUpEnvironment(configFilePath, s3ConfigFilePath string) error {

	c, err := getSdkConfigFromEnv()
	if err != nil {
		return err
	}

	if c.Url == "" {
		s, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			return err
		}
		err = json.Unmarshal(s, &c)
		if err != nil {
			return err
		}
	}

	// Stores the retrieved parameters in a public static singleton
	DefaultConfig = &c

	cs3, err := getS3ConfigFromEnv()
	if err != nil {
		return err
	}

	if cs3.Bucket == "" {
		s, err := ioutil.ReadFile(s3ConfigFilePath)
		if err != nil {
			return err
		}
		err = json.Unmarshal(s, &cs3)
		if err != nil {
			return err
		}
	}

	// Stores the retrieved parameters in a public static singleton
	DefaultS3Config = &cs3

	return nil
}

func getSdkConfigFromEnv() (SdkConfig, error) {

	var c SdkConfig

	// check presence of Env variable
	protocol := os.Getenv(KeyProtocol)
	url := os.Getenv(KeyURL)
	clientKey := os.Getenv(KeyClientKey)
	clientSecret := os.Getenv(KeyClientSecret)
	user := os.Getenv(KeyUser)
	password := os.Getenv(KeyPassword)
	skipVerifyStr := os.Getenv(KeySkipVerify)
	if skipVerifyStr == "" {
		skipVerifyStr = "false"
	}
	skipVerify, err := strconv.ParseBool(skipVerifyStr)
	if err != nil {
		return c, err
	}

	if !(len(protocol) > 0 && len(url) > 0 && len(clientKey) > 0 && len(clientSecret) > 0 && len(user) > 0 && len(password) > 0) {
		return c, nil
	}

	c.Protocol = protocol
	c.Url = url
	c.ClientKey = clientKey
	c.ClientSecret = clientSecret
	c.User = user
	c.Password = password
	c.SkipVerify = skipVerify

	return c, nil
}

func getS3ConfigFromEnv() (S3Config, error) {

	var c S3Config

	// check presence of Env variable
	endpoint := os.Getenv(KeyS3Endpoint)
	region := os.Getenv(KeyS3Region)
	bucket := os.Getenv(KeyS3Bucket)
	apiKey := os.Getenv(KeyS3ApiKey)
	apiSecret := os.Getenv(KeyS3ApiSecret)
	usePSHStr := os.Getenv(KeyS3UsePydioSpecificHeader)
	if usePSHStr == "" {
		usePSHStr = "false"
	}
	usePSH, err := strconv.ParseBool(usePSHStr)
	if err != nil {
		return c, err
	}

	isDebugStr := os.Getenv(KeyS3IsDebug)
	if isDebugStr == "" {
		isDebugStr = "false"
	}
	isDebug, err := strconv.ParseBool(isDebugStr)
	if err != nil {
		return c, err
	}

	if !(len(endpoint) > 0 && len(region) > 0 && len(bucket) > 0 && len(apiKey) > 0 && len(apiSecret) > 0) {
		return c, nil
	}

	c.Endpoint = endpoint
	c.Region = region
	c.Bucket = bucket
	c.ApiKey = apiKey
	c.ApiSecret = apiSecret
	c.UsePydioSpecificHeader = usePSH
	c.IsDebug = isDebug

	return c, nil
}
