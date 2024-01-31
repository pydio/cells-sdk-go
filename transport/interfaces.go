package transport

import cells_sdk "github.com/pydio/cells-sdk-go/v5"

type ConfigStore interface {
	RefreshIfRequired(*cells_sdk.SdkConfig) (bool, error)
}

type TokenProvider interface {
	Retrieve() (string, error)
	Expired() bool
}
