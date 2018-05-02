package util

import (
	"fmt"
	"sync"

	"github.com/pydio/go-os/config"
	"github.com/pydio/go-os/config/source/file"

	pydioconfig "github.com/pydio/cells/common/config"
)

var (
	// TODO provision json file on test suite startup
	pydioConfigFilePath = "/home/bsinou/.config/pydio/cells/pydio.json"

	defaultConfig *pydioconfig.Config

	once sync.Once
)

func newLocalSource() config.Source {
	fmt.Println("config path: " + pydioConfigFilePath)
	return file.NewSource(
		config.SourceName(pydioConfigFilePath),
	)
}

// getServerConfig returns a deserialised config with info about the server we want to test
func getServerConfig() *pydioconfig.Config {
	once.Do(func() {
		defaultConfig = &pydioconfig.Config{config.NewConfig(
			// config.WithSource(newEnvSource()),
			config.WithSource(newLocalSource()),
		)}
	})

	fmt.Println("Found default config: " + defaultConfig.Get("os").String("test"))

	return defaultConfig
}
