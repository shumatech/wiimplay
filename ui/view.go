package ui

import (
    "fmt"
    "strconv"
    "time"

    "github.com/shumatech/wiimplay/ui/res"

    "github.com/gotk3/gotk3/gdk"
    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
    "github.com/gotk3/gotk3/pango"
)

type ViewListener interface {
    SeekSelect(position float64)
    ShufflePress()
    PreviousPress()
    PlayPausePress()
    NextPress()
    LoopPress()
    VolumeSelect(level int)
    PlayListSelect(index int)
    TrackSelect(index int)
    WindowHidden(flag bool)
    MutePress()
    SettingsPress()
    DeviceSelect(index int)
    RefreshPress()
    AboutPress()
    EqualizerSelect(name string)
    ControlsPress()
}

type View struct {
    listener ViewListener

    application *gtk.Application
    window *gtk.ApplicationWindow
    statusMenu *gtk.Menu
    statusIcon *gtk.StatusIcon
    deleteEvent glib.SignalHandle

    devicesCombo *gtk.ComboBoxText
    devicesSignal glib.SignalHandle

    albumImage *gtk.Image

    infoLabel *gtk.Label
    metaDataLabel *gtk.Label

    timeUsed *gtk.Label
    timeTotal *gtk.Label
    seekBar *gtk.ProgressBar

    shuffleButton *gtk.EventBox
    previousButton *gtk.EventBox
    playButton *gtk.EventBox
    nextButton *gtk.EventBox
    loopButton *gtk.EventBox

    muteButton *gtk.EventBox
    volumeAdjust *gtk.Adjustment
    volumeScale *gtk.Scale
    volumeSignal glib.SignalHandle
    eqButton *gtk.EventBox
    eqMenu *gtk.Menu
    eqGroup *glib.SList

    playListBox *gtk.ListBox
    playListCombo *gtk.ComboBoxText
    playListSignal glib.SignalHandle

    defaultAlbumPixbuf *gdk.Pixbuf
    defaultTrackPixbuf *gdk.Pixbuf
}

const (
    AlbumImageWidth = 256
    AlbumImageHeight = 256
    TrackImageWidth = 48
    TrackImageHeight = 48

    VolumePageLevel = 5
)

func (view *View) createStatusMenu() error {
    var err error
    view.statusMenu, err = createMenu([]menuItem {
        { "Previous",   view.listener.PreviousPress  },
        { "Play/Pause", view.listener.PlayPausePress },
        { "Next",       view.listener.NextPress      },
        { "Mute",       view.listener.MutePress      },
        { "Quit",       view.application.Quit        },
    })
    return err
}

func (view *View) createStatusIcon() error {
    icon, err := gdk.PixbufNewFromDataOnly(res.WiimIcon)
    if err != nil {
        return err
    }
    statusIcon, err := gtk.StatusIconNewFromPixbuf(icon)
    if err != nil {
        return err
    }
    view.statusIcon = statusIcon

    statusIcon.Connect("popup-menu", func(self *gtk.StatusIcon,
                                          button gdk.Button, activateTime uint) {
        view.statusMenu.PopupAtStatusIcon(self, button, uint32(activateTime))
    })
    statusIcon.Connect("activate", func() {
        if (view.window.IsVisible()) {
            view.window.Hide()
        } else {
            view.window.ShowAll()
        }
    })
    statusIcon.Connect("scroll-event", func(self *gtk.StatusIcon,
                                            event *gdk.Event) {
        scroll := gdk.EventScrollNewFromEvent(event)
        dir := scroll.Direction()
        if dir == gdk.SCROLL_UP {
            level := view.volumeAdjust.GetValue()
            view.volumeAdjust.SetValue(level + VolumePageLevel)
        } else if dir == gdk.SCROLL_DOWN {
            level := view.volumeAdjust.GetValue()
            view.volumeAdjust.SetValue(level - VolumePageLevel)
        }
    })
    statusIcon.Connect("button-press-event", func(self *gtk.StatusIcon,
                                                  event *gdk.Event) {
        button := gdk.EventButtonNewFromEvent(event)
        num := button.Button()
        if num == 2 {
            view.listener.PlayPausePress()
        } else if num == 9 {
            view.listener.NextPress()
        } else if num == 8 {
            view.listener.PreviousPress()
        }
    })

    statusIcon.SetVisible(true)

    return nil
}

func (view *View) addPlayListItem(title string, artist string,
                                  album *gdk.Pixbuf) error {
    title = Sanitize(title)
    artist = Sanitize(artist)

    grid, err := gtk.GridNew()
    if err != nil {
        return err
    }

    if album != nil {
        albumImage, err := gtk.ImageNewFromPixbuf(album)
        if err != nil {
            return err
        }
        grid.Attach(albumImage, 0, 0, 1, 2)
    }

    titleLabel, err := gtk.LabelNew(title)
    if err != nil {
        return err
    }
    titleLabel.SetMarkup("<span size=\"larger\">" + title + "</span>")
    titleLabel.SetHAlign(gtk.ALIGN_START)
    titleLabel.SetEllipsize(pango.ELLIPSIZE_END)
    grid.Attach(titleLabel, 1, 0, 1, 1)

    artistLabel, err := gtk.LabelNew("")
    if err != nil {
        return err
    }
    artistLabel.SetMarkup("<span size=\"small\">" + artist + "</span>")
    artistLabel.SetHAlign(gtk.ALIGN_START)
    artistLabel.SetEllipsize(pango.ELLIPSIZE_END)
    grid.Attach(artistLabel, 1, 1, 1, 1)

    grid.SetColumnSpacing(5)

    view.playListBox.Add(grid)
    view.playListBox.ShowAll()

    return nil
}

func (view *View) addPlayList(parent *gtk.Box) error {
    vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
    if err != nil {
        return err
    }
    vbox.SetMarginTop(10)
    vbox.SetMarginBottom(10)
    vbox.SetMarginLeft(10)
    vbox.SetMarginRight(10)
    vbox.SetSpacing(10)

    hbox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
    if err != nil {
        return err
    }
    vbox.Add(hbox)

    view.playListCombo, err = gtk.ComboBoxTextNewWithEntry()
    if err != nil {
        return err
    }
    entry, err := view.playListCombo.GetEntry()
    if err != nil {
        return err
    }
    entry.SetPlaceholderText("Select playlist...")
    entry.SetEditable(false)
    entry.SetCanFocus(false)

    view.playListSignal = view.playListCombo.Connect("changed", func() {
        index := view.playListCombo.GetActive()
        if index >= 0 {
            view.listener.PlayListSelect(index)
        }
    })
    hbox.Add(view.playListCombo)
    hbox.SetChildPacking(view.playListCombo, true, true, 0, gtk.PACK_START)

    button, err := createImageButton(res.Settings, func() {
        view.listener.SettingsPress()
    })
    if err != nil {
        return err
    }
    button.SetHAlign(gtk.ALIGN_END)
    hbox.Add(button)
    hbox.SetChildPacking(button, true, true, 5, gtk.PACK_START)

    button, err = createImageButton(res.About, func() {
        view.listener.AboutPress()
    })
    if err != nil {
        return err
    }
    button.SetHAlign(gtk.ALIGN_END)
    hbox.Add(button)

    swin, err := gtk.ScrolledWindowNew(nil, nil)
    if err != nil {
        return err
    }

    view.playListBox, err = gtk.ListBoxNew()
    if err != nil {
        return err
    }
    view.playListBox.Connect("row-activated", func() {
        row := view.playListBox.GetSelectedRow()
        if row != nil {
            view.listener.TrackSelect(row.GetIndex())
        }
    })
    swin.Add(view.playListBox)
    vbox.Add(swin)
    vbox.SetChildPacking(swin, true, true, 0, gtk.PACK_START)

    parent.Add(vbox)
    parent.SetChildPacking(vbox, true, true, 0, gtk.PACK_START)

    return nil
}

func (view *View) addPlayerDevices(parent *gtk.Box) error {
    box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
    if err != nil {
        return err
    }

    view.devicesCombo, err = gtk.ComboBoxTextNewWithEntry()
    if err != nil {
        return err
    }
    entry, err := view.devicesCombo.GetEntry()
    if err != nil {
        return err
    }
    entry.SetPlaceholderText("Select device...")
    entry.SetEditable(false)
    entry.SetCanFocus(false)

    view.devicesSignal = view.devicesCombo.Connect("changed", func() {
        index := view.devicesCombo.GetActive()
        if index >= 0 {
            view.listener.DeviceSelect(index)
        }
    })
    box.Add(view.devicesCombo)
    box.SetChildPacking(view.devicesCombo, true, true, 0, gtk.PACK_START)

    button, err := createImageButton(res.Refresh, func() {
        view.listener.RefreshPress()
    })
    if err != nil {
        return err
    }
    box.Add(button)

    button, err = createImageButton(res.DeviceControls, func() {
        view.listener.ControlsPress()
    })
    if err != nil {
        return err
    }
    box.Add(button)

    parent.Add(box)
    return nil
}

func (view *View) addPlayerControls(parent *gtk.Box) error {
    box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
    if err != nil {
        return err
    }
    box.SetHomogeneous(true)

    view.shuffleButton, err = createMultiImageButton(
        [][]byte{res.ShuffleOff, res.ShuffleOn},
        view.listener.ShufflePress)
    if err != nil {
        return err
    }
    box.Add(view.shuffleButton)

    view.previousButton, err = createImageButton(
        res.Previous, view.listener.PreviousPress)
    if err != nil {
        return err
    }
    box.Add(view.previousButton)

    view.playButton, err = createMultiImageButton(
        [][]byte{res.Play, res.Pause},
        view.listener.PlayPausePress)
    if err != nil {
        return err
    }
    box.Add(view.playButton)

    view.nextButton, err = createImageButton(
        res.Next, view.listener.NextPress)
    if err != nil {
        return err
    }
    box.Add(view.nextButton)

    view.loopButton, err = createMultiImageButton(
        [][]byte{res.LoopOff, res.LoopOne, res.LoopAll},
        view.listener.LoopPress)
    if err != nil {
        return err
    }
    box.Add(view.loopButton)

    parent.Add(box)
    return nil
}

func (view *View) addSeekBar(parent *gtk.Box) error {
    eventBox, err := gtk.EventBoxNew()
    if err != nil {
        return err
    }

    grid, err := gtk.GridNew()
    if err != nil {
        return err
    }
    eventBox.Add(grid)

    view.timeUsed, err = gtk.LabelNew("")
    if err != nil {
        return err
    }
    view.timeUsed.SetMarkup("<span size=\"small\">0:00</span>")
    view.timeUsed.SetHAlign(gtk.ALIGN_START)
    grid.Attach(view.timeUsed, 0, 0, 1, 1)

    view.timeTotal, err = gtk.LabelNew("0:00")
    if err != nil {
        return err
    }
    view.timeTotal.SetMarkup("<span size=\"small\">0:00</span>")
    view.timeTotal.SetHAlign(gtk.ALIGN_END)
    grid.Attach(view.timeTotal, 1, 0, 1, 1)

    view.seekBar, err = gtk.ProgressBarNew()
    if err != nil {
        return err
    }
    grid.Attach(view.seekBar, 0, 1, 2, 1)

    grid.SetColumnHomogeneous(true)

    eventBox.Connect("button-press-event", func(self *gtk.EventBox, event *gdk.Event){
        button := gdk.EventButtonNewFromEvent(event)
        width := float64(self.GetAllocation().GetWidth())
        view.listener.SeekSelect(button.X() / width)
    })

    parent.Add(eventBox)

    return nil
}

func (view *View) addVolumeBar(parent *gtk.Box) error {
    box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
    if err != nil {
        return err
    }

    view.muteButton, err = createMultiImageButton(
        [][]byte{res.MuteOff, res.MuteOn},
        view.listener.MutePress)
    if err != nil {
        return err
    }
    box.Add(view.muteButton)

    image, err := createImageFromData(res.VolumeLow)
    if err != nil {
        return nil
    }
    box.Add(image)

    view.volumeAdjust, err = gtk.AdjustmentNew(0.0, 0.0, 100.0, 1.0, VolumePageLevel, 0.0)
    if err != nil {
        return nil
    }
    view.volumeSignal = view.volumeAdjust.Connect("value-changed",
                                                  func(self *gtk.Adjustment) {
        view.listener.VolumeSelect(int(self.GetValue() + 0.5))
    })

    view.volumeScale, err = gtk.ScaleNew(gtk.ORIENTATION_HORIZONTAL, view.volumeAdjust)
    if err != nil {
        return nil
    }
    view.volumeScale.SetDrawValue(false)
    box.Add(view.volumeScale)
    box.SetChildPacking(view.volumeScale, true, true, 0, gtk.PACK_START)

    image, err = createImageFromData(res.VolumeHigh)
    if err != nil {
        return nil
    }
    box.Add(image)

    view.eqButton, err = createMultiImageButton(
        [][]byte{res.EqOff, res.EqOn},
        func() {
            if view.eqMenu != nil {
                view.eqMenu.PopupAtPointer(nil)
            }
        })
    if err != nil {
        return err
    }
    box.Add(view.eqButton)

    box.SetSpacing(5)
    parent.Add(box)

    return nil
}

func formatInfo(quality string, rate int, format string, bitrate int) string{
    str := "<span size=\"small\">"
    sep := ""
    if quality != "" {
        str += Sanitize(quality)
        sep = " "
    }
    if rate != 0 {
        if rate % 1000 == 0 {
            str += sep + strconv.Itoa(rate / 1000) + " kHz"
        } else if rate % 100 == 0 {
            str += fmt.Sprintf("%s%.1f kHz", sep, float32(rate) / 1000)
        } else {
            str += sep + strconv.Itoa(rate) + " Hz"
        }
        sep = ", "
    }
    if format != "" {
        str += sep + Sanitize(format) + " bit"
        sep = ", "
    }
    if bitrate != 0 {
        str += sep + strconv.Itoa(bitrate) + " kbps"
    }
    str += "</span>"
    return str
}

func formatMetadata(title string, artist string, album string) string{
    return fmt.Sprintf(
        "<span size=\"large\" weight=\"bold\">" + Sanitize(title) + "</span>\n" +
        "<span size=\"medium\">" + Sanitize(artist) + "</span>\n" +
        "<span size=\"medium\">" + Sanitize(album) + "</span>")
}

func formatTooltip(title string, artist string, album string) string{
    return fmt.Sprintf(
        "<span size=\"medium\">" + Sanitize(title) + "</span>\n" +
        "<span size=\"small\">" + Sanitize(artist) + "</span>\n" +
        "<span size=\"small\">" + Sanitize(album) + "</span>")
}

func formatDuration(dur time.Duration) string {
    h := int(dur.Hours())
    m := int(dur.Minutes()) % 60
    s := int(dur.Seconds()) % 60
    if h > 0 {
        return fmt.Sprintf("%d:%02d:%02d", h, m, s)
    } else {
        return fmt.Sprintf("%d:%02d", m, s)
    }
}

func (view *View) addPlayer(parent *gtk.Box) error {
    box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
    if err != nil {
        return err
    }
    box.SetMarginTop(10)
    box.SetMarginBottom(10)
    box.SetMarginLeft(10)
    box.SetMarginRight(10)
    box.SetSpacing(10)

    err = view.addPlayerDevices(box)
    if err != nil {
        return err
    }

    view.defaultTrackPixbuf, err = gdk.PixbufNewFromDataOnly(res.DefaultTrack)
    if err != nil {
        return err
    }
    view.defaultAlbumPixbuf, err = gdk.PixbufNewFromDataOnly(res.DefaultAlbum)
    if err != nil {
        return err
    }
    view.albumImage, err = gtk.ImageNewFromPixbuf(view.defaultAlbumPixbuf)
    if err != nil {
        return err
    }
    box.Add(view.albumImage)

    if err := view.addSeekBar(box); err != nil {
        return err
    }

    view.infoLabel, err = gtk.LabelNew("")
    if err != nil {
        return err
    }
    view.infoLabel.SetHAlign(gtk.ALIGN_START)
    box.Add(view.infoLabel)

    view.metaDataLabel, err = gtk.LabelNew("")
    if err != nil {
        return err
    }
    view.metaDataLabel.SetLineWrapMode(pango.WRAP_WORD)
    view.metaDataLabel.SetLineWrap(true)
    view.metaDataLabel.SetHAlign(gtk.ALIGN_START)
    view.metaDataLabel.SetVAlign(gtk.ALIGN_START)
    view.metaDataLabel.SetMaxWidthChars(32)

    box.Add(view.metaDataLabel)
    box.SetChildPacking(view.metaDataLabel, true, true, 0, gtk.PACK_START)

    if err = view.addPlayerControls(box); err != nil {
        return err
    }

    if err = view.addVolumeBar(box); err != nil {
        return err
    }

    parent.Add(box)
    return nil
}

///////////////////////////////////////////////////////////////////////////////
// Public interface
///////////////////////////////////////////////////////////////////////////////

func NewView(application *gtk.Application, listener ViewListener) (*View, error) {
    view := &View{}
    view.application = application
    view.listener = listener

    window, err := gtk.ApplicationWindowNew(view.application)
    if err != nil {
        return nil, err
    }
    view.window = window
    view.window.Connect("window-state-event", func (self *gtk.ApplicationWindow,
                                                    event *gdk.Event) {
        state := gdk.EventWindowStateNewFromEvent(event)
        view.listener.WindowHidden((state.NewWindowState() & gdk.WINDOW_STATE_WITHDRAWN) != 0)

    })

    cssProvider, err := gtk.CssProviderNew()
    if err != nil {
        return nil, err
    }
    err = cssProvider.LoadFromData(res.StyleCss)
    if err != nil {
        return nil, err
    }
    display, err := gdk.DisplayGetDefault()
    if err != nil {
        return nil, err
    }
    screen, err := display.GetDefaultScreen()
    if err != nil {
        return nil, err
    }
    gtk.AddProviderForScreen(screen, cssProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

    box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
    if err != nil {
        return nil, err
    }

    if err = view.addPlayer(box); err != nil {
        return nil, err
    }
    sep, err := gtk.SeparatorNew(gtk.ORIENTATION_VERTICAL)
    if err != nil {
        return nil, err
    }
    box.Add(sep)
    if err = view.addPlayList(box); err != nil {
        return nil, err
    }
    if err := view.createStatusMenu(); err != nil {
        return nil, err
    }
    if err := view.createStatusIcon(); err != nil {
        return nil, err
    }

    window.Add(box)

    window.SetTitle("WiiM Play")
    window.SetTypeHint(gdk.WINDOW_TYPE_HINT_DIALOG)
    window.SetDefaultSize(800, 600)
    window.ShowAll()
    window.SetIcon(view.defaultAlbumPixbuf)

    return view, nil
}

type PlayerState uint
const (
    PlayerUnavailable PlayerState = iota
    PlayerPlaying
    PlayerPaused
)

func (view *View) SetPlayerState(state PlayerState) {
    var playIndex int
    switch state {
        case PlayerUnavailable:
            playIndex = 0
        case PlayerPlaying:
            playIndex = 1
        case PlayerPaused:
            playIndex = 0
    }
    setMultiImageButton(view.playButton, playIndex)
}

func (view *View) SetInfo(quality string, rate int, format string, bitrate int) {
    view.infoLabel.SetMarkup(formatInfo(quality, rate, format, bitrate))
}

func (view *View) SetPlaying(title string, artist string, album string) {
    view.metaDataLabel.SetMarkup(formatMetadata(title, artist, album))
    view.statusIcon.SetTooltipMarkup(formatTooltip(title, artist, album))
}

func (view *View) SetStatus(status string) {
    view.SetPlaying(status, "", "")
}

func (view *View) SetAlbumArt(album *gdk.Pixbuf) {
    if album == nil {
        album = view.defaultAlbumPixbuf
    }
    view.albumImage.SetFromPixbuf(album)
}

func (view *View) SetSeek(used time.Duration, total time.Duration) {
    position := 0.0
    if total != time.Duration(0) {
        position = used.Seconds() / total.Seconds()
    }
    view.timeUsed.SetText(formatDuration(used))
    view.timeTotal.SetText(formatDuration(total))
    view.seekBar.SetFraction(position)
}

func (view *View) SetVolume(level int) {
    view.volumeAdjust.HandlerBlock(view.volumeSignal)
    view.volumeAdjust.SetValue(float64(level))
    view.volumeAdjust.HandlerUnblock(view.volumeSignal)
}

func (view *View) SetMute(mute bool) {
    index := 0
    if mute {
        index = 1
    }
    setMultiImageButton(view.muteButton, index)
    view.volumeScale.SetSensitive(!mute)
}

func (view *View) SetShuffle(on bool) {
    index := 0
    if on {
        index = 1
    }
    setMultiImageButton(view.shuffleButton, index)
}

type LoopState uint
const (
    LoopNone LoopState = iota
    LoopOne
    LoopAll
)

func (view *View) SetLoopState(state LoopState) {
    setMultiImageButton(view.loopButton, int(state))
}

func (view *View) SetPlayList(index int) {
    view.playListCombo.HandlerBlock(view.playListSignal)
    view.playListCombo.SetActive(index)
    if index == -1 {
        if entry, err := view.playListCombo.GetEntry(); err == nil {
            entry.SetText("")
        }
    }
    view.playListCombo.HandlerUnblock(view.playListSignal)
}

func (view *View) AddPlayList(name string) {
    view.playListCombo.AppendText(name)
}

func (view *View) ClearPlayLists() {
    view.playListCombo.RemoveAll()
    if entry, err := view.playListCombo.GetEntry(); err == nil {
        entry.SetText("")
    }
}

func (view *View) SetTrack(index int) {
    row := view.playListBox.GetRowAtIndex(index)
    view.playListBox.SelectRow(row)
    if row != nil {
        glib.TimeoutAdd(100, func() bool {
            row := view.playListBox.GetRowAtIndex(index)
            if row != nil {
                row.GrabFocus()
            }
            return false
        })
    }
}

func (view *View) AddTrack(title string, artist string) error {
    return view.addPlayListItem(title, artist, view.defaultTrackPixbuf)
}

func (view *View) SetTrackAlbumArt(index int, album *gdk.Pixbuf) error {
    row := view.playListBox.GetRowAtIndex(index)
    if row != nil {
        if album == nil {
            album = view.defaultTrackPixbuf
        }
        image, err := gtk.ImageNewFromPixbuf(album)
        if err != nil {
            return err
        }
        widget, err := row.GetChild()
        if err != nil {
            return err
        }
        grid := widget.(*gtk.Grid)
        widget, err = grid.GetChildAt(0, 0)
        if err != nil {
            return err
        }
        grid.Remove(widget)
        grid.Attach(image, 0, 0, 1, 2)
        image.Show()
    }
    return nil
}

func (view *View) ClearTracks() {
    children := view.playListBox.GetChildren()
    children.Foreach(func(item interface{}) {
        view.playListBox.Remove(item.(*gtk.Widget))
    })
}

func (view *View) GetWindow() *gtk.Window {
    return &view.window.Window
}

func (view *View) Hide() {
    view.window.Hide()
}

func (view *View) HideOnDelete(enable bool) {
    if (enable && view.deleteEvent == 0) {
        view.deleteEvent = view.window.Connect("delete-event", func() bool {
            view.window.Hide()
            return true
        })
    } else if (!enable && view.deleteEvent != 0) {
        view.window.HandlerDisconnect(view.deleteEvent)
        view.deleteEvent = 0
    }
}

func (view *View) ShowStatusIcon(enable bool) {
    view.statusIcon.SetVisible(enable)
}

func (view *View) AddDevice(name string) {
    view.devicesCombo.AppendText(name)
}

func (view *View) SetDevice(index int) {
    view.devicesCombo.HandlerBlock(view.devicesSignal)
    view.devicesCombo.SetActive(index)
    view.devicesCombo.HandlerUnblock(view.devicesSignal)
}

func (view *View) ClearDevices() {
    view.devicesCombo.RemoveAll()
    if entry, err := view.devicesCombo.GetEntry(); err == nil {
        entry.SetText("")
    }
}

func (view *View) ClearEqualizers() {
    if view.eqMenu != nil {
        view.eqMenu.Destroy()
        view.eqMenu = nil
        view.eqGroup = nil
    }
    setMultiImageButton(view.eqButton, 0)
}

func (view *View) SetEqualizers(names []string, selected string) error {
    var err error
    view.ClearEqualizers()
    view.eqMenu, err = gtk.MenuNew()
    if err != nil {
        return err
    }

    for _, name := range(names) {
        item, err := gtk.RadioMenuItemNewWithLabel(view.eqGroup, name)
        if err != nil {
            return err
        }
        view.eqGroup, err = item.GetGroup()
        if err != nil {
            return err
        }
        if name == selected {
            item.SetActive(true)
        }
        item.Connect("activate", func(self *gtk.RadioMenuItem) {
            if self.GetActive() {
                view.listener.EqualizerSelect(self.GetLabel())
            }
        })
        item.Show()
        view.eqMenu.Add(item)
    }

    return nil
}

func (view *View) SetEqualizerOn(on bool) {
    if on {
        setMultiImageButton(view.eqButton, 1)
    } else {
        setMultiImageButton(view.eqButton, 0)
    }
}
