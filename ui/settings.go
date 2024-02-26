package ui

import (
    "github.com/gotk3/gotk3/gtk"
)

type Settings struct {
    SendNotifications bool
    MprisSupport bool
    ShowStatusIcon bool
    HideOnClose bool
    HideOnStart bool
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

    sendNotifications, err := gtk.CheckButtonNewWithLabel("Send notifications")
    if err != nil {
        return false, err
    }
    sendNotifications.SetActive(settings.SendNotifications)
    box.Add(sendNotifications)

    mprisSupport, err := gtk.CheckButtonNewWithLabel("MPRIS support")
    if err != nil {
        return false, err
    }
    mprisSupport.SetActive(settings.MprisSupport)
    box.Add(mprisSupport)

    showStatusIcon, err := gtk.CheckButtonNewWithLabel("Show status icon")
    if err != nil {
        return false, err
    }
    showStatusIcon.SetActive(settings.ShowStatusIcon)
    box.Add(showStatusIcon)

    hideOnClose, err := gtk.CheckButtonNewWithLabel("Hide on close")
    if err != nil {
        return false, err
    }
    hideOnClose.SetActive(settings.HideOnClose)
    box.Add(hideOnClose)

    hideOnStart, err := gtk.CheckButtonNewWithLabel("Hide on start")
    if err != nil {
        return false, err
    }
    hideOnStart.SetActive(settings.HideOnStart)
    box.Add(hideOnStart)

    box.SetSpacing(10)
    box.SetMarginTop(10)
    box.SetMarginBottom(10)
    box.SetMarginLeft(10)
    box.SetMarginRight(10)
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