package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Proxy
var (
	SetDefault        = viper.SetDefault
	SetEnvKeyReplacer = viper.SetEnvKeyReplacer
)

// Option configures viper.
type Option func()

// Parse the configuration from the environment variables and a file with the specified path.
func Parse(filePath string, config interface{}, options ...Option) error {
	// Apply options
	for _, opt := range options {
		opt()
	}

	// Parse environments variables
	viper.AutomaticEnv()

	// Parse the file
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(config); err != nil {
		return errors.Wrap(err, "viper.Unmarshal")
	}
	return nil
}
