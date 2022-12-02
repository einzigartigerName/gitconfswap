package gui

import (
	"log"
	"reflect"
	"sort"

	"github.com/einzigartigerName/gitconfswap/assets"
	"github.com/einzigartigerName/gitconfswap/configs"
	p "github.com/einzigartigerName/gitconfswap/pkg/profile"
	"github.com/getlantern/systray"
)

type Handler interface {
	OnReady()
	OnExit()
}

func NewDefaultHandler(config *configs.AppConfig) Handler {
	return &defaultHandler{
		config:   config,
		switcher: p.NewDefaultProfileLoader(config),
		notifier: &defaultNotifier{icon: config.IconName},
	}
}

type defaultHandler struct {
	config   *configs.AppConfig
	switcher p.Switcher
	notifier Notifier
}

func (h *defaultHandler) OnReady() {
	switch h.config.IconName {
	case configs.IconLight:
		systray.SetIcon(assets.IconLight)
	case configs.IconColor:
		systray.SetIcon(assets.IconColor)
		fallthrough
	default:
		systray.SetIcon(assets.IconDark)
	}

	systray.SetTooltip("git config switcher")

	var profiles []string
	for profile := range h.config.Profiles {
		profiles = append(profiles, profile)
	}
	sort.Strings(profiles)

	var channels []chan struct{}
	for _, profile := range profiles {
		item := systray.AddMenuItem(profile, profile)
		channels = append(channels, item.ClickedCh)
	}

	systray.AddSeparator()
	quit := systray.AddMenuItem("Quit", "Quits this app")
	channels = append(channels, quit.ClickedCh)

	go func() {
		for {
			index := receiveFromAny(channels)

			// Quit MenuItem Clicked
			if index == len(channels)-1 {
				systray.Quit()
				return
			}

			profile := profiles[index]
			if err := h.switcher.Switch(profile); err != nil {
				h.notifier.Error(profile, err)
			} else {
				h.notifier.Success(profile)
			}
		}
	}()
}

func receiveFromAny(channels []chan struct{}) int {
	var set []reflect.SelectCase
	for _, ch := range channels {
		set = append(set, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	from, _, _ := reflect.Select(set)
	return from
}

func (h *defaultHandler) OnExit() {
	log.Print("Exiting\n")
}
