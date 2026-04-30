package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

type Settings struct {
	SendNotifications bool
	MprisSupport      bool
	ShowStatusIcon    bool
	HideOnClose       bool
	HideOnStart       bool
}

func addSettingsCheckButton(box *gtk.Box, label string, active bool) (*gtk.CheckButton, error) {
	button, err := gtk.CheckButtonNewWithLabel(label)
	if err != nil {
		return nil, err
	}
	button.SetActive(active)
	box.Add(button)
	return button, nil
}

func ShowSettingsDialog(window gtk.IWindow, settings *Settings) (bool, error) {
	flags := gtk.DIALOG_MODAL | gtk.DIALOG_DESTROY_WITH_PARENT
	ok := []interface{}{"OK", gtk.RESPONSE_ACCEPT}
	cancel := []interface{}{"Cancel", gtk.RESPONSE_REJECT}
	dialog, err := gtk.DialogNewWithButtons("Settings", window, flags, cancel, ok)
	if err != nil {
		return false, err
	}

	box, err := dialog.GetContentArea()
	if err != nil {
		return false, err
	}

	sendNotifications, err := addSettingsCheckButton(box, "Send notifications", settings.SendNotifications)
	if err != nil {
		return false, err
	}

	mprisSupport, err := addSettingsCheckButton(box, "MPRIS support", settings.MprisSupport)
	if err != nil {
		return false, err
	}

	showStatusIcon, err := addSettingsCheckButton(box, "Show status icon", settings.ShowStatusIcon)
	if err != nil {
		return false, err
	}

	hideOnClose, err := addSettingsCheckButton(box, "Hide on close", settings.HideOnClose)
	if err != nil {
		return false, err
	}

	hideOnStart, err := addSettingsCheckButton(box, "Hide on start", settings.HideOnStart)
	if err != nil {
		return false, err
	}

	setBoxSpacingMargin(box, 10, 10)
	box.ShowAll()

	response := dialog.Run()
	dialog.Destroy()

	if response != gtk.RESPONSE_ACCEPT {
		return false, nil
	}

	settings.SendNotifications = sendNotifications.GetActive()
	settings.MprisSupport = mprisSupport.GetActive()
	settings.ShowStatusIcon = showStatusIcon.GetActive()
	settings.HideOnClose = hideOnClose.GetActive()
	settings.HideOnStart = hideOnStart.GetActive()

	return true, nil
}
