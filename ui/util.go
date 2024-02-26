package ui

import (
    "log"
    "fmt"
    "net/url"
    "strings"

    "github.com/shumatech/wiimplay/mpris"
    "github.com/shumatech/wiimplay/wiimdev"

    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gdk"
    "github.com/gotk3/gotk3/gtk"
)

func Sanitize(s string) string {
    reserved := [...][2]string{
        {"&", "&amp;"  },
        {"\"", "&quot;"},
        {"'", "&apos;" },
        {"<", "&lt;"   },
        {">", "&gt;"   },
    }
    for _, r := range reserved {
        s = strings.ReplaceAll(s, r[0], r[1])
    }
    return s
}

type menuItem struct {
    name string
    action func()
}

func createMenu(items []menuItem) (*gtk.Menu, error) {
    menu, err := gtk.MenuNew()
    if err != nil {
        return nil, err
    }

    for _, mi := range items {
        item, err := gtk.MenuItemNewWithLabel(mi.name)
        if err != nil {
            return nil, err
        }
        item.Show()
        if mi.action != nil {
            action := mi.action
            item.Connect("activate", func() {
                action()
            })
        }
        menu.Append(item)
    }

    return menu, nil
}

func createImageFromData(data []byte) (*gtk.Image, error) {
    pixbuf, err := gdk.PixbufNewFromDataOnly(data)
    if err != nil {
        return nil, err
    }
    image, err := gtk.ImageNewFromPixbuf(pixbuf)
    if err != nil {
        return nil, err
    }
    return image, nil
}

func createImageButton(data []byte, pressed func()) (*gtk.EventBox, error) {
    box, err := gtk.EventBoxNew()
    if err != nil {
        return nil, err
    }
    image, err := createImageFromData(data)
    if err != nil {
        return nil, err
    }

    styleContext, err := box.GetStyleContext()
    if err != nil {
        return nil, err
    }
    box.Connect("button-press-event", func() {
        styleContext.AddClass("button-press");
        pressed();
    })
    box.Connect("button-release-event", func() {
        styleContext.RemoveClass("button-press");
    })
    box.Connect("enter-notify-event", func() {
        styleContext.AddClass("button-hover");
    })
    box.Connect("leave-notify-event", func() {
        styleContext.RemoveClass("button-hover");
        styleContext.RemoveClass("button-press");
    })

    box.Add(image)
    return box, nil
}

func createMultiImageButton(dataList [][]byte, pressed func()) (*gtk.EventBox, error) {
    box, err := gtk.EventBoxNew()
    if err != nil {
        return nil, err
    }
    stack, err := gtk.StackNew()
    if err != nil {
        return nil,err
    }
    for _, data := range dataList {
        image, err := createImageFromData(data)
        if err != nil {
            return nil, err
        }
        stack.Add(image)
    }
    styleContext, err := box.GetStyleContext()
    if err != nil {
        return nil, err
    }
    box.Connect("button-press-event", func() {
        styleContext.AddClass("button-press");
        pressed();
    })
    box.Connect("button-release-event", func() {
        styleContext.RemoveClass("button-press");
    })
    box.Connect("enter-notify-event", func() {
        styleContext.AddClass("button-hover");
    })
    box.Connect("leave-notify-event", func() {
        styleContext.RemoveClass("button-hover");
        styleContext.RemoveClass("button-press");
    })

    box.Add(stack)
    return box, nil
}

func setMultiImageButton(button *gtk.EventBox, index int) {
    child, err := button.GetChild()
    if err != nil {
        return
    }
    stack := child.(*gtk.Stack)
    item := stack.GetChildren().Nth(uint(index)).Data()
    stack.SetVisibleChild(item.(*gtk.Widget))
}

func newScaledPixbufFromData(data []byte, w int, h int) (*gdk.Pixbuf, error) {
    pixbuf, err := gdk.PixbufNewFromDataOnly(data)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    if (pixbuf.GetWidth() != w || pixbuf.GetHeight() != h) {
        var err error
        pixbuf, err = pixbuf.ScaleSimple(w, h, gdk.INTERP_BILINEAR)
        if err != nil {
            return nil, err
        }
    }
    return pixbuf, nil
}

func messageDialog(window *gtk.Window, message string) {
    dialog := gtk.MessageDialogNew(window,
        gtk.DIALOG_MODAL | gtk.DIALOG_DESTROY_WITH_PARENT,
        gtk.MESSAGE_INFO, gtk.BUTTONS_CLOSE, "")
    dialog.SetMarkup(message)
    dialog.Run()
    dialog.Destroy()
}

func logErr(err error) {
    if err != nil {
        log.Println(err)
    }
}

func loopModeToStatus(loopMode wiimdev.LoopMode) (loopStatus mpris.LoopStatus, shuffle bool) {
    switch loopMode {
        case wiimdev.LoopModeAll:
            loopStatus = mpris.LoopStatusPlaylist
            shuffle = false
        case wiimdev.LoopModeOne:
            loopStatus = mpris.LoopStatusTrack
            shuffle = false
        case wiimdev.LoopModeAllShuffle:
            loopStatus = mpris.LoopStatusPlaylist
            shuffle = true
        case wiimdev.LoopModeNoneShuffle:
            loopStatus = mpris.LoopStatusNone
            shuffle = true
        case wiimdev.LoopModeNone:
            loopStatus = mpris.LoopStatusNone
            shuffle = false
        case wiimdev.LoopModeOneShuffle:
            loopStatus = mpris.LoopStatusTrack
            shuffle = true
    }
    return
}

func loopStatusToMode(loopStatus mpris.LoopStatus, shuffle bool) wiimdev.LoopMode {
    if shuffle {
        switch loopStatus {
            case mpris.LoopStatusNone:
                return wiimdev.LoopModeNoneShuffle
            case mpris.LoopStatusTrack:
                return wiimdev.LoopModeOneShuffle
            case mpris.LoopStatusPlaylist:
                return wiimdev.LoopModeAllShuffle
        }
    } else {
        switch loopStatus {
            case mpris.LoopStatusNone:
                return wiimdev.LoopModeNone
            case mpris.LoopStatusTrack:
                return wiimdev.LoopModeOne
            case mpris.LoopStatusPlaylist:
                return wiimdev.LoopModeAll
        }
    }
    return wiimdev.LoopModeNone
}

func loopModeToState(loopMode wiimdev.LoopMode) (loopState LoopState, shuffle bool) {
    switch loopMode {
        case wiimdev.LoopModeAll:
            loopState = LoopAll
            shuffle = false
        case wiimdev.LoopModeOne:
            loopState = LoopOne
            shuffle = false
        case wiimdev.LoopModeAllShuffle:
            loopState = LoopAll
            shuffle = true
        case wiimdev.LoopModeNoneShuffle:
            loopState = LoopNone
            shuffle = true
        case wiimdev.LoopModeNone:
            loopState = LoopNone
            shuffle = false
        case wiimdev.LoopModeOneShuffle:
            loopState = LoopOne
            shuffle = true
    }
    return
}

func loopStateToMode(loopState LoopState, shuffle bool) wiimdev.LoopMode {
    if shuffle {
        switch loopState {
            case LoopNone:
                return wiimdev.LoopModeNoneShuffle
            case LoopOne:
                return wiimdev.LoopModeOneShuffle
            case LoopAll:
                return wiimdev.LoopModeAllShuffle
        }
    } else {
        switch loopState {
            case LoopNone:
                return wiimdev.LoopModeNone
            case LoopOne:
                return wiimdev.LoopModeOne
            case LoopAll:
                return wiimdev.LoopModeAll
        }
    }
    return wiimdev.LoopModeNone
}

func stateTransportToPlayer(state wiimdev.TransportState) PlayerState {
    switch state {
    case wiimdev.StateStopped:
        return PlayerPaused
    case wiimdev.StatePaused:
        return PlayerPaused
    case wiimdev.StatePlaying:
        return PlayerPlaying
    case wiimdev.StateTransitioning:
        return PlayerPlaying
    case wiimdev.StateNoMedia:
        return PlayerPaused
    default:
        return PlayerUnavailable
    }
}

func statePlayerToTransport(state PlayerState) wiimdev.TransportState {
    switch state {
    case PlayerPaused:
        return wiimdev.StatePaused
    case PlayerPlaying:
        return wiimdev.StatePlaying
    default:
        return wiimdev.StateUnknown
    }
}

func urlToHost(urlStr string) string {
    urlParsed, err := url.Parse(urlStr)
    if err != nil {
        return urlStr
    }
    return strings.Split(urlParsed.Host, ":")[0]
}

// Schedule a closure to execute on the main thread after a delay.
func mainDelay(millis uint, fn func()) (glib.SourceHandle) {
    return glib.TimeoutAdd(millis, func() bool {
        fn()
        return false
    })
}

// Schedule a closure to run on the main thread
func mainRun(fn func()) {
    glib.IdleAddPriority(glib.PRIORITY_DEFAULT, fn)
}

// Schedule a closure to execute on the main thread and wait for its return.
// Warning! Do not call from the main thread or it will hang.
func mainAwait(fn func() interface{}) interface{} {
    resultChan := make(chan interface{})
    glib.IdleAddPriority(glib.PRIORITY_DEFAULT, func() {
        resultChan <- fn()
    })
    return <- resultChan
}

func sendNotification(application *gtk.Application, title string,
                      body string, icon string) {
    notification := glib.NotificationNew(title)
    notification.SetBody(body)
    if icon != "" {
        notification.SetIcon(icon)
    }
    application.SendNotification("notify", notification)
}

func metadataEqual(mdA *wiimdev.Metadata, mdB *wiimdev.Metadata) bool {
    return (mdA.Id == mdB.Id &&
            mdA.Title == mdB.Title &&
            mdA.Artist == mdB.Artist &&
            mdA.Album == mdB.Album)
}

func playListsEqual(plA []wiimdev.PlayList, plB []wiimdev.PlayList) bool {
    if len(plA) != len(plB) {
        return false
    }
    for i, pl := range(plA) {
        if pl.Id != plB[i].Id {
            return false
        }
    }
    return true
}

func tracksEqual(trA []wiimdev.Track, trB []wiimdev.Track) bool {
    if len(trA) != len(trB) {
        return false
    }
    for i, tr := range(trA) {
        if !metadataEqual(tr.Metadata, trB[i].Metadata) {
            return false
        }
    }
    return true
}
