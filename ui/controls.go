package ui

import (
    "github.com/gotk3/gotk3/gtk"
)

type Controls struct {
    AudioOutput int
    AudioOutputList []string
    AudioInput int
    AudioInputList []string
    Balance float64
    FadeEffects bool
    FixedVolume bool
}

type ControlsListener interface {
    AudioOutputSelect(index int)
    AudioInputSelect(index int)
    BalanceSelect(value float64)
    FadeEffectsSelect(on bool)
    FixedVolumeSelect(on bool)
}

func ShowControlsDialog(window gtk.IWindow, controls *Controls, listener ControlsListener) error {
    flags := gtk.DIALOG_MODAL | gtk.DIALOG_DESTROY_WITH_PARENT
    close := []interface{}{"Close", gtk.RESPONSE_CLOSE}
    dialog, err := gtk.DialogNewWithButtons("Controls", window, flags, close)
    if err != nil {
        return err
    }

    box, err := dialog.GetContentArea()
    if err != nil {
        return err
    }

    grid, err := gtk.GridNew()
    if err != nil {
        return err
    }
    grid.SetRowSpacing(10)
    grid.SetColumnSpacing(10)

    label, err := gtk.LabelNew("Audio Output:")
    if err != nil {
        return err
    }
    label.SetHAlign(gtk.ALIGN_START)
    grid.Attach(label, 0, 0, 1, 1)

    combo, err := gtk.ComboBoxTextNew()
    if err != nil {
        return err
    }
    for _, entry := range controls.AudioOutputList {
        combo.AppendText(entry)
    }
    combo.SetActive(controls.AudioOutput)
    combo.Connect("changed", func(self *gtk.ComboBoxText) {
        listener.AudioOutputSelect(self.GetActive())
    })
    grid.Attach(combo, 1, 0, 1, 1)

    label, err = gtk.LabelNew("Audio Input:")
    if err != nil {
        return err
    }
    label.SetHAlign(gtk.ALIGN_START)
    grid.Attach(label, 0, 1, 1, 1)

    combo, err = gtk.ComboBoxTextNew()
    if err != nil {
        return err
    }
    for _, entry := range controls.AudioInputList {
        combo.AppendText(entry)
    }
    combo.SetActive(controls.AudioInput)
    combo.Connect("changed", func(self *gtk.ComboBoxText) {
        listener.AudioInputSelect(self.GetActive())
    })
    grid.Attach(combo, 1, 1, 1, 1)

    label, err = gtk.LabelNew("Volume Balance:")
    if err != nil {
        return err
    }
    label.SetHAlign(gtk.ALIGN_START)
    grid.Attach(label, 0, 2, 1, 1)

    adjust, err := gtk.AdjustmentNew(0.0, -1.0, 1.0, 0.05, 0.1, 0.0)
    if err != nil {
        return nil
    }
    adjust.SetValue(controls.Balance)
    adjust.Connect("value-changed", func(self *gtk.Adjustment) {
        listener.BalanceSelect(self.GetValue())
    })

    scale, err := ScaleExtNew(gtk.ORIENTATION_HORIZONTAL, adjust)
    if err != nil {
        return nil
    }
    scale.SetDrawValue(true)
    scale.SetHasOrigin(false)
    grid.Attach(scale, 1, 2, 1, 1)

    label, err = gtk.LabelNew("Fade Effects:")
    if err != nil {
        return err
    }
    label.SetHAlign(gtk.ALIGN_START)
    grid.Attach(label, 0, 3, 1, 1)

    sw, err := gtk.SwitchNew()
    if err != nil {
        return err
    }
    sw.SetActive(controls.FadeEffects)
    sw.Connect("state-set", func(self *gtk.Switch) {
        listener.FadeEffectsSelect(self.GetActive())
    })
    grid.Attach(sw, 1, 3, 1, 1)

    label, err = gtk.LabelNew("Fixed Volume:")
    if err != nil {
        return err
    }
    label.SetHAlign(gtk.ALIGN_START)
    grid.Attach(label, 0, 4, 1, 1)

    sw, err = gtk.SwitchNew()
    if err != nil {
        return err
    }
    sw.SetActive(controls.FixedVolume)
    sw.Connect("state-set", func(self *gtk.Switch) {
        listener.FixedVolumeSelect(self.GetActive())
    })
    grid.Attach(sw, 1, 4, 1, 1)

    box.Add(grid)

    box.SetSpacing(10)
    box.SetMarginTop(10)
    box.SetMarginBottom(10)
    box.SetMarginLeft(10)
    box.SetMarginRight(10)
    box.ShowAll()

    dialog.Run()
    dialog.Destroy()

    return nil
}