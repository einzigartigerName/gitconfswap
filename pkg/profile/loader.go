package profile

import (
	"os/exec"

	"github.com/einzigartigerName/gitconfswap/configs"
	"github.com/pkg/errors"
)

// Switcher is the interface for the main functionality of this application: Loading different git configs
type Switcher interface {
	// Switch the profile indicated by the name from the configuration file
	Switch(name string) error
}

func NewDefaultProfileLoader(config *configs.AppConfig) Switcher {
	return &defaultLoader{
		config: config,
	}
}

type defaultLoader struct {
	config *configs.AppConfig
}

func (d *defaultLoader) Switch(name string) error {
	profile, ok := d.config.Profiles[name]
	if !ok {
		return errors.Errorf("unknown profile: %s", name)
	}

	for _, confValue := range profile {
		if err := execGitConfig(confValue); err != nil {
			return err
		}
	}
	return nil
}

func execGitConfig(conf configs.GitValue) error {
	err := exec.Command("git", "config", "--global", conf.Variable, conf.Value).Run()
	if err != nil {
		return err
	}
	return nil
}
