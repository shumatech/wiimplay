package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

type Controls struct {
	AudioOutput     int
	AudioOutputList []string
	AudioInput      int
	AudioInputList  []string
	Balance         float64
	FadeEffects     bool
	FixedVolume     bool
}

type ControlsListener interface {
	AudioOutputSelect(index int)
	AudioInputSelect(index int)
	BalanceSelect(value float64)
	FadeEffectsSelect(on bool)
	FixedVolumeSelect(on bool)
}

func addControlRow(grid *gtk.Grid, labelText string, row int, widget gtk.IWidget) error {
	label, err := gtk.LabelNew(labelText)
	if err != nil {
		return err
	}
	label.SetHAlign(gtk.ALIGN_START)
	grid.Attach(label, 0, row, 1, 1)
	grid.Attach(widget, 1, row, 1, 1)
	return nil
}

func addComboControl(grid *gtk.Grid, label string, row int, entries []string, active int, changed func(int)) error {
	combo, err := gtk.ComboBoxTextNew()
	if err != nil {
		return err
	}
	for _, entry := range entries {
		combo.AppendText(entry)
	}
	combo.SetActive(active)
	combo.Connect("changed", func(self *gtk.ComboBoxText) {
		changed(self.GetActive())
	})
	return addControlRow(grid, label, row, combo)
}

func addSwitchControl(grid *gtk.Grid, label string, row int, active bool, changed func(bool)) error {
	sw, err := gtk.SwitchNew()
	if err != nil {
		return err
	}
	sw.SetActive(active)
	sw.Connect("state-set", func(self *gtk.Switch) {
		changed(self.GetActive())
	})
	return addControlRow(grid, label, row, sw)
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

	err = addComboControl(grid, "Audio Output:", 0, controls.AudioOutputList, controls.AudioOutput,
		listener.AudioOutputSelect)
	if err != nil {
		return err
	}

	err = addComboControl(grid, "Audio Input:", 1, controls.AudioInputList, controls.AudioInput,
		listener.AudioInputSelect)
	if err != nil {
		return err
	}

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

	err = addControlRow(grid, "Volume Balance:", 2, scale)
	if err != nil {
		return err
	}

	err = addSwitchControl(grid, "Fade Effects:", 3, controls.FadeEffects,
		listener.FadeEffectsSelect)
	if err != nil {
		return err
	}

	err = addSwitchControl(grid, "Fixed Volume:", 4, controls.FixedVolume,
		listener.FixedVolumeSelect)
	if err != nil {
		return err
	}

	box.Add(grid)

	setBoxSpacingMargin(box, 10, 10)
	box.ShowAll()

	dialog.Run()
	dialog.Destroy()

	return nil
}
