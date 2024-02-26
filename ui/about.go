package ui

import (
    "github.com/shumatech/wiimplay/ui/res"

    "github.com/gotk3/gotk3/gtk"
)

func ShowAboutDialog(parent *gtk.Window, about string) error {
    dialog, err := gtk.DialogNewWithButtons("About", parent, gtk.DIALOG_DESTROY_WITH_PARENT,
        []interface{}{ "Close", gtk.RESPONSE_CLOSE })
    if err != nil {
        return err
    }
    box, err := dialog.GetContentArea()
    if err != nil {
        return err
    }
    image, err := createImageFromData(res.WiimAbout)
    if err != nil {
        return err
    }
    box.Add(image)

    label, err := gtk.LabelNew("")
    if err != nil {
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

    box.SetSpacing(10)
    box.SetMarginTop(10)
    box.SetMarginBottom(10)
    box.SetMarginLeft(10)
    box.SetMarginRight(10)

    dialog.ShowAll()
    dialog.Run()
    dialog.Destroy()

    return nil
}