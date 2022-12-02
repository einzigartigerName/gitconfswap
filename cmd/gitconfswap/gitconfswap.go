package main

import (
	"log"

	"github.com/einzigartigerName/gitconfswap/configs"
	"github.com/einzigartigerName/gitconfswap/pkg/gui"
	p "github.com/einzigartigerName/gitconfswap/pkg/profile"
	"github.com/getlantern/systray"
	"gopkg.in/alecthomas/kingpin.v2"
)

//nolint:gochecknoglobals,gocritic
var (
	filepath = kingpin.Flag("file", "config file to use").Short('f').ExistingFile()
	cliMode  = kingpin.Flag("cli", "run in cli mode").Bool()
	profile  = kingpin.Flag("profile", "if run in cli mode, load this profile").Short('p').String()
)

func main() {
	kingpin.Parse()
	loader := configs.NewFileConfigLoader(*filepath)

	config, err := loader.Load()
	if err != nil {
		log.Fatalf("error while loading config: %s", err.Error())
	}

	if *cliMode {
		startCLI(config)
	} else {
		startTray(config)
	}
}

func startCLI(config *configs.AppConfig) {
	if *profile == "" {
		log.Fatalf("no profile set. nothing to do")
	}

	err := p.NewDefaultProfileLoader(config).Switch(*profile)
	if err != nil {
		log.Fatalf("error while loading profile [%s]: %s", *profile, err.Error())
	}
}

func startTray(config *configs.AppConfig) {
	handler := gui.NewDefaultHandler(config)

	systray.Run(handler.OnReady, handler.OnExit)
}
