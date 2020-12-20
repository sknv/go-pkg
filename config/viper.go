package config

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Options struct {
	EnvPrefix      string
	EnvKeyReplacer *strings.Replacer
}

// SetDefault sets the default value for this key.
func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

// Parse will parse the configuration from the environment variables and a file with the specified path.
func Parse(filePath string, config interface{}, options Options) error {
	// Parse environments variables
	viper.SetEnvPrefix(options.EnvPrefix)
	viper.SetEnvKeyReplacer(options.EnvKeyReplacer)
	viper.AutomaticEnv()

	// Parse the file
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "read config")
	}

	if err := viper.Unmarshal(config); err != nil {
		return errors.Wrap(err, "unmarshal the config")
	}
	return nil
}
