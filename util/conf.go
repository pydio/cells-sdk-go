package util

import (
	"sync"

	"github.com/micro/go-config/reader"
	"github.com/pydio/go-os/config"
	"github.com/pydio/go-os/config/source/file"
)

var (
	// TODO provision json file on test suite startup
	pydioConfigFilePath = ""

	defaultConfig *Config

	once sync.Once
)

// Config wrapper around micro Config
type Config struct {
	config.Config
}

// getServerConfig returns a deserialised config with info about the server we want to test
func getServerConfig() *Config {
	once.Do(func() {
		defaultConfig = &Config{
			config.NewConfig(
				config.WithSource(newLocalSource()),
			)}
	})

	return defaultConfig
}

func newLocalSource() config.Source {
	// fmt.Println("config path: " + pydioConfigFilePath)
	return file.NewSource(
		config.SourceName(pydioConfigFilePath),
	)
}

func GetConfigValue(path ...string) reader.Value {
	return getServerConfig().Get(path...)
}

func Set(val interface{}, path ...string) {
	getServerConfig().Set(val, path...)
}

func Bytes() []byte {
	return getServerConfig().Bytes()
}
