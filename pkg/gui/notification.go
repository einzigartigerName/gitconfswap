package gui

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

type Notifier interface {
	Success(profile string)
	Error(profile string, err error)
}

type defaultNotifier struct {
	icon string
}

func (n *defaultNotifier) Success(profile string) {
	title := fmt.Sprintf("Loaded Profile %s", profile)
	_ = beeep.Notify(title, "", "")
}

func (n *defaultNotifier) Error(profile string, err error) {
	title := fmt.Sprintf("Error Loading Profile %s", profile)
	_ = beeep.Notify(title, err.Error(), n.icon)
}
