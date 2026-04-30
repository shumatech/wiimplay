package ui

import (
	"github.com/shumatech/wiimplay/ui/res"

	"github.com/gotk3/gotk3/gtk"
)

func ShowAboutDialog(parent *gtk.Window, about string) error {
	dialog, box, err := newDialogWithContent("About", parent, gtk.DIALOG_DESTROY_WITH_PARENT,
		[]interface{}{"Close", gtk.RESPONSE_CLOSE})
	if err != nil {
		return err
	}
	image, err := createImageFromData(res.WiimAbout)
	if err != nil {
		dialog.Destroy()
		return err
	}
	box.Add(image)

	label, err := gtk.LabelNew("")
	if err != nil {
		dialog.Destroy()
		return err
	}
	label.SetMarkup(
		about + "<small>" +
			"\n\nCopyright (C) 2024 ShumaTech\n\n" +
			"This program is free software: you can redistribute it and/or modify\n" +
			"it under the terms of the GNU General Public License as published by\n" +
			"the Free Software Foundation, either version 3 of the License, or\n" +
			"(at your option) any later version.\n\n" +
			"This program is distributed in the hope that it will be useful,\n" +
			"but WITHOUT ANY WARRANTY; without even the implied warranty of\n" +
			"MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the\n" +
			"GNU General Public License for more details.\n" +
			"</small>")
	box.Add(label)

	setBoxSpacingMargin(box, 10, 10)

	dialog.ShowAll()
	dialog.Run()
	dialog.Destroy()

	return nil
}
