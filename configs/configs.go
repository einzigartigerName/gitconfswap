package configs

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
)

type AppConfig struct {
	Profiles  map[string][]GitValue `yaml:"profiles"`
	IconName  string                `yaml:"icon,omitempty"`
	IconAsset *[]byte
}

type GitValue struct {
	Variable string
	Value    string
}

// ConfigLoader loads the configuration for the app
type ConfigLoader interface {
	Load() (*AppConfig, error)
}

// ByteConfigLoader implements configs.ConfigLoader by loading config from
// two provided bytes slices.
type ByteConfigLoader struct {
	ConfigBytes []byte
}

// Load implements configs.ConfigLoader by using the properties of the ByteConfigLoader.
func (l *ByteConfigLoader) Load() (*AppConfig, error) {
	if l.ConfigBytes == nil {
		return nil, errors.Errorf("ConfigBytes must not be nil")
	}

	result := new(AppConfig)
	if err := yaml.Unmarshal(l.ConfigBytes, result); err != nil {
		return nil, errors.Errorf("Unable to parse config: " + err.Error())
	}

	// Validate config
	if err := result.validate(); err != nil {
		return nil, errors.Wrap(err, "Loaded invalid config")
	}

	return result, nil
}

// FileConfigLoader implements configs.ConfigLoader by loading config from
// two files located at given paths.
type FileConfigLoader struct {
	ConfigPath string
}

// NewFileConfigLoader returns a ConfigLoader with the given path to the config file
// If no path was provided, the standard location will be searched for a valid file
func NewFileConfigLoader(configPath string) ConfigLoader {
	if configPath == "" {
		configPath = getDefaultConfigPath()
	}

	return &FileConfigLoader{
		ConfigPath: configPath,
	}
}

// Load implements configs.ConfigLoader by using the properties of the FileConfigLoader.
func (l *FileConfigLoader) Load() (*AppConfig, error) {
	if l.ConfigPath == "" {
		return nil, errors.Errorf("ConfigPath must not be empty!")
	}

	// Switch dsConfigBytes from file
	var (
		ioError     error
		configBytes []byte
	)
	if configBytes, ioError = os.ReadFile(l.ConfigPath); ioError == nil {
		return (&ByteConfigLoader{
			ConfigBytes: configBytes,
		}).Load()
	}
	return nil, ioError
}

func (c AppConfig) validate() error {
	if c.IconName != IconDark && c.IconName != IconLight && c.IconName != IconColor {
		return errors.Errorf("unknown icon %s", c.IconName)
	}
	return nil
}

func getDefaultConfigPath() string {
	home := os.Getenv("HOME")

	return fmt.Sprintf("%s/.config/gitconfswap/config.yaml", home)
}
