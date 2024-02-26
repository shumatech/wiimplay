package ui

import (
    "errors"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/shumatech/wiimplay/dcache"
    "github.com/shumatech/wiimplay/mpris"
    "github.com/shumatech/wiimplay/wiimdev"

    "github.com/gotk3/gotk3/gdk"
    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
    "github.com/hashicorp/golang-lru/v2"
)

type ConfigAdapter interface {
    Get(key string, defvalue interface{}) interface{}
    Set(key string, value interface{})
    Save() error
}

type PixbufRequest struct {
    url string
    width int
    height int
    complete func(*gdk.Pixbuf)
}

type Device struct {
    Name string `yaml:"Name"`
    Url string  `yaml:"Url"`
}

type Controller struct {
    application *gtk.Application
    config ConfigAdapter
    version string
    build string

    view *View
    device wiimdev.Device
    mprisServer *mpris.Server

    info *wiimdev.DeviceInfo
    tracks []wiimdev.Track
    playLists []wiimdev.PlayList
    currPlayList string

    runChan chan RunReq
    pixbufChan chan PixbufRequest
    pixbufCache *lru.Cache[string, *gdk.Pixbuf]
    downloadCache *dcache.DiskCache
    isHidden bool
    updateTimer glib.SourceHandle
    devices []Device
}

///////////////////////////////////////////////////////////////////////////////
// ViewListener Interface (called from main thread)
///////////////////////////////////////////////////////////////////////////////
func (ctl *Controller) SeekSelect(position float64) {
    target := time.Duration(position * float64(ctl.info.TrackTotal))
    device := ctl.device
    ctl.run(func() { logErr(device.Seek(target)) })
}

func (ctl *Controller) ShufflePress() {
    loopMode := ctl.info.LoopMode
    loopState, shuffle := loopModeToState(loopMode)
    shuffle = !shuffle
    ctl.view.SetShuffle(shuffle)
    loopMode = loopStateToMode(loopState, shuffle)
    ctl.info.LoopMode = loopMode
    device := ctl.device
    ctl.run(func() { logErr(device.LoopMode(loopMode)) })
}

func (ctl *Controller) PreviousPress() {
    device := ctl.device
    ctl.run(func() { logErr(device.Previous()) })
}

func (ctl *Controller) PlayPausePress() {
    device := ctl.device
    switch (ctl.info.TransportState) {
        case wiimdev.StatePlaying:
            ctl.info.TransportState = wiimdev.StatePaused
            ctl.run(func() { logErr(device.Pause()) })
        case wiimdev.StatePaused:
            ctl.info.TransportState = wiimdev.StatePlaying
            ctl.run(func() { logErr(device.Play()) })
        default:
            return
    }
    playerState := stateTransportToPlayer(ctl.info.TransportState)
    ctl.view.SetPlayerState(playerState)
    ctl.emitStatusChanged()
}

func (ctl *Controller) NextPress() {
    device := ctl.device
    ctl.run(func() { logErr(device.Next()) })
}

func (ctl *Controller) LoopPress() {
    loopMode := ctl.info.LoopMode
    loopState, shuffle := loopModeToState(loopMode)
    switch loopState {
        case LoopNone:
            loopState = LoopOne
        case LoopOne:
            loopState = LoopAll
        case LoopAll:
            loopState = LoopNone
    }
    ctl.view.SetLoopState(loopState)
    loopMode = loopStateToMode(loopState, shuffle)
    ctl.info.LoopMode = loopMode
    device := ctl.device
    ctl.run(func() { logErr(device.LoopMode(loopMode)) })
}

func (ctl *Controller) VolumeSelect(level int) {
    ctl.info.Volume = level
    device := ctl.device
    ctl.run(func() { logErr(device.Volume(level)) })
}

func (ctl *Controller) PlayListSelect(index int) {
    if index < len(ctl.playLists) {
        ctl.tracks = []wiimdev.Track{}
        id := ctl.playLists[index].Id
        device := ctl.device
        ctl.run(func() { logErr(device.SetPlayList(id)) })
    }
}

func (ctl *Controller) TrackSelect(index int) {
    ctl.info.Track = index + 1
    track := ctl.info.Track
    device := ctl.device
    ctl.run(func() { logErr(device.PlayTrack(track)) })
}

func (ctl *Controller) WindowHidden(flag bool) {
    if ctl.isHidden && !flag {
        ctl.isHidden = flag
        ctl.stopUpdate()
        ctl.startUpdate()
    } else {
        ctl.isHidden = flag
    }
}

func (ctl *Controller) MutePress() {
    ctl.info.Mute = !ctl.info.Mute
    ctl.view.SetMute(ctl.info.Mute)
    mute := ctl.info.Mute
    device := ctl.device
    ctl.run(func() { logErr(device.Mute(mute)) })
}

func (ctl *Controller) SettingsPress() {
    settings := Settings {
        SendNotifications: ctl.config.Get("SendNotifications", true).(bool),
        MprisSupport: ctl.config.Get("MprisSupport", true).(bool),
        ShowStatusIcon: ctl.config.Get("ShowStatusIcon", true).(bool),
        HideOnClose: ctl.config.Get("HideOnClose", true).(bool),
        HideOnStart: ctl.config.Get("HideOnStart", false).(bool),
    }
    ok, err := ShowSettingsDialog(ctl.application.GetActiveWindow(), &settings)
    logErr(err)
    if ok {
        ctl.config.Set("SendNotifications", settings.SendNotifications)
        ctl.config.Set("MprisSupport", settings.MprisSupport)
        ctl.config.Set("ShowStatusIcon", settings.ShowStatusIcon)
        ctl.config.Set("HideOnClose", settings.HideOnClose)
        ctl.config.Set("HideOnStart", settings.HideOnStart)
        ctl.config.Save()

        ctl.view.HideOnDelete(settings.HideOnClose)
        ctl.view.ShowStatusIcon(settings.ShowStatusIcon)
        ctl.run( func() { ctl.mprisServer.Listen(settings.MprisSupport) })
    }
}

func (ctl *Controller) DeviceSelect(index int) {
    ctl.connectToDevice(index)
}

func (ctl *Controller) RefreshPress() {
    ctl.discoverDevices()
}

func (ctl *Controller) AboutPress() {
    about := "Version: " + ctl.version +
             "\nBuild: " + ctl.build
    discovery := ctl.device.GetDeviceDiscovery()
    if discovery != nil {
        about += "\n\nModel: " + discovery.ModelName +
                 "\nFirmware: " + discovery.ModelNumber +
                 "\nURL: " + discovery.Url
    }
    logErr(ShowAboutDialog(ctl.view.GetWindow(), about))
}

func (ctl *Controller) EqualizerSelect(name string) {
    ctl.view.SetEqualizerOn(name != "Off")
    device := ctl.device
    ctl.run(func() { logErr(device.SetEqualizer(name)) })
}

func (ctl *Controller) ControlsPress() {
    var controls Controls
    device := ctl.device
    var err error
    ctl.runThen(func() {
        if controls.AudioOutput, err = device.GetAudioOutput(); err != nil {
            return
        }
        if controls.AudioOutputList, err = device.GetAudioOutputList(); err != nil {
            return
        }
        if controls.Balance, err = device.GetBalance(); err != nil {
            return
        }
        if controls.FadeEffects, err = device.GetFadeEffects(); err != nil {
            return
        }
        if controls.FixedVolume, err = device.GetFixedVolume(); err != nil {
            return
        }
    }, func() {
        logErr(err)
        if err == nil {
            ShowControlsDialog(ctl.view.GetWindow(), &controls, ctl)
        } else {
            messageDialog(ctl.view.GetWindow(),
                          fmt.Sprintf("Error getting value:\n%v", err))
        }
    })
}

///////////////////////////////////////////////////////////////////////////////
// Controls Interface (called from main thread)
///////////////////////////////////////////////////////////////////////////////

func (ctl *Controller) AudioOutputSelect(index int) {
    device := ctl.device
    ctl.run(func() { device.SetAudioOutput(index) })
}

func (ctl *Controller) BalanceSelect(value float64) {
    device := ctl.device
    ctl.run(func() { device.SetBalance(value) })
}

func (ctl *Controller) FadeEffectsSelect(on bool) {
    device := ctl.device
    ctl.run(func() { device.SetFadeEffects(on) })
}

func (ctl *Controller) FixedVolumeSelect(on bool) {
    device := ctl.device
    ctl.run(func() { device.SetFixedVolume(on) })
}

///////////////////////////////////////////////////////////////////////////////
// MprisPlayer Interface (called outside main thread)
///////////////////////////////////////////////////////////////////////////////

func (ctl *Controller) Next() {
    mainRun(func() {
        device := ctl.device
        ctl.run( func() { logErr(device.Next()) })
    })
}

func (ctl *Controller) Previous() {
    mainRun(func() {
        device := ctl.device
        ctl.run( func() { logErr(device.Previous()) })
    })
}

func (ctl *Controller) Pause() {
    mainRun(func() {
        if ctl.info.TransportState == wiimdev.StatePlaying {
            ctl.info.TransportState = wiimdev.StatePaused
            ctl.view.SetPlayerState(PlayerPaused)
            device := ctl.device
            ctl.run( func() { logErr(device.Pause()) })
        }
    })
}

func (ctl *Controller) Play() {
    mainRun(func() {
        if ctl.info.TransportState == wiimdev.StatePaused {
            ctl.info.TransportState = wiimdev.StatePlaying
            ctl.view.SetPlayerState(PlayerPlaying)
            device := ctl.device
            ctl.emitStatusChanged()
            ctl.run( func() { logErr(device.Play()) })
        }
    })
}

func (ctl *Controller) PlaybackStatus() mpris.PlaybackStatus {
    result := mainAwait(func() interface{} {
        switch ctl.info.TransportState {
        case wiimdev.StatePlaying:
            return mpris.PlaybackStatusPlaying
        case wiimdev.StatePaused:
            return mpris.PlaybackStatusPaused
        default:
            return mpris.PlaybackStatusStopped
        }
    })
    return result.(mpris.PlaybackStatus)
}

func (ctl *Controller) SetVolume(level int) {
    mainRun(func() {
        device := ctl.device
        ctl.info.Volume = level
        ctl.view.SetVolume(level)
        ctl.run( func() { logErr(device.Volume(level)) })
    })
}

func (ctl *Controller) GetVolume() int {
    result := mainAwait(func() interface{} {
        return ctl.info.Volume
    })
    return result.(int)
}

func (ctl *Controller) GetMetadata() *mpris.Metadata {
    result := mainAwait(func() interface{} {
        md := ctl.info.TrackMetadata
        return &mpris.Metadata {
            Id: md.Id,
            Title: md.Title,
            Artist: md.Artist,
            Album: md.Album,
            AlbumArt: md.AlbumArt,
            Length: ctl.info.TrackTotal,
        }
    })
    return result.(*mpris.Metadata)
}

func (ctl *Controller) GetPosition() time.Duration {
    result := mainAwait(func() interface{} {
        return ctl.info.TrackUsed
    })
    return result.(time.Duration)
}

func (ctl *Controller) SetPosition(position time.Duration) {
    mainRun(func() {
        device := ctl.device
        ctl.run( func() { logErr(device.Seek(position)) })
    })
}

func (ctl *Controller) Seek(offset time.Duration) {
    mainRun(func() {
        device := ctl.device
        position := ctl.info.TrackUsed + offset
        ctl.run( func() { logErr(device.Seek(position)) })
    })
}

func (ctl *Controller) GetLoopStatus() mpris.LoopStatus {
    result := mainAwait(func() interface{} {
        loopStatus, _ := loopModeToStatus(ctl.info.LoopMode)
        return loopStatus
    })
    return result.(mpris.LoopStatus)
}

func (ctl *Controller) SetLoopStatus(status mpris.LoopStatus) {
    mainRun(func() {
        _, shuffle := loopModeToState(ctl.info.LoopMode)
        mode := loopStatusToMode(status, shuffle)
        device := ctl.device
        ctl.run( func() { logErr(device.LoopMode(mode)) })
    })
}

func (ctl *Controller) GetShuffle() bool {
    result := mainAwait(func() interface{} {
        _, shuffle := loopModeToStatus(ctl.info.LoopMode)
        return shuffle
    })
    return result.(bool)
}

func (ctl *Controller) SetShuffle(shuffle bool) {
    mainRun(func() {
        loopStatus, _ := loopModeToStatus(ctl.info.LoopMode)
        mode := loopStatusToMode(loopStatus, shuffle)
        device := ctl.device
        ctl.run( func() { logErr(device.LoopMode(mode)) })
    })
}

///////////////////////////////////////////////////////////////////////////////
// Internal methods
///////////////////////////////////////////////////////////////////////////////

func (ctl *Controller) sendNotification(md *wiimdev.Metadata) {
    if !ctl.config.Get("SendNotifications", true).(bool) {
        return
    }
    ctl.pixbufRequest(md.AlbumArt, TrackImageWidth,
                      TrackImageHeight, func(pixbuf *gdk.Pixbuf) {
        var tmpFile string
        if pixbuf != nil {
            tmpFile = os.TempDir() + "/wiiplay_notify_icon.png"
            err := pixbuf.SavePNG(tmpFile, 5)
            logErr(err)
        }
        log.Printf("Sending notification")
        sendNotification(ctl.application, md.Title,
            md.Artist + "\n" + md.Album, tmpFile)
    })
}

func (ctl *Controller) stopUpdate() {
    if ctl.updateTimer != 0 {
        glib.SourceRemove(ctl.updateTimer)
        ctl.updateTimer = 0
    }
}

func (ctl *Controller) startUpdate() {
    if ctl.updateTimer == 0 {
        ctl.updateInfo()
    }
}

func (ctl *Controller) scheduleUpdate() {
    var delay int
    if ctl.isHidden {
        delay = ctl.config.Get("ClosedUpdatePeriod", 5000).(int)
    } else {
        delay = ctl.config.Get("OpenUpdatePeriod", 1000).(int)
    }
    ctl.updateTimer = mainDelay(uint(delay), func() {
        ctl.updateInfo()
    })
}

func (ctl *Controller) emitStatusChanged() {
    if ctl.config.Get("MprisSupport", true).(bool) {
        ctl.run( ctl.mprisServer.StatusChanged )
    }
}

func (ctl *Controller) emitMetadataChanged() {
    if ctl.config.Get("MprisSupport", true).(bool) {
        ctl.run( ctl.mprisServer.MetadataChanged )
    }
}

func (ctl *Controller) updateInfo() {
    var err error
    var newInfo *wiimdev.DeviceInfo
    device := ctl.device
    ctl.runThen(func() {
        newInfo, err = device.GetInfo()
        logErr(err)
    }, func() {
        defer ctl.scheduleUpdate()
        if newInfo == nil {
            return
        }
        if newInfo.TransportState != ctl.info.TransportState {
            playerState := stateTransportToPlayer(newInfo.TransportState)
            ctl.view.SetPlayerState(playerState)
            defer ctl.emitStatusChanged()
        }
        if (newInfo.TransportState == wiimdev.StateNoMedia) {
            ctl.view.SetStatus("No media. Select a playlist.")
            ctl.view.SetAlbumArt(nil)
            ctl.view.ClearTracks();
            ctl.view.SetPlayList(-1)
            if ctl.playLists == nil {
                ctl.updatePlayLists()
            }
            return
        }
        newMd := newInfo.TrackMetadata
        md := ctl.info.TrackMetadata
        metadataChanged := !metadataEqual(newMd, md)
        if metadataChanged {
            ctl.view.SetPlaying(newMd.Title, newMd.Artist, newMd.Album)
            if newMd.Title != "" {
                ctl.sendNotification(newMd)
                defer ctl.emitMetadataChanged()
            }
        }
        if newInfo.Track != ctl.info.Track {
            ctl.view.SetTrack(newInfo.Track - 1)
            if newInfo.Track == 0 {
                // Track index of 0 indicates the WiiM is not playing from a playlist.
                log.Println("Not playing from a playlist");
                ctl.view.SetPlayList(-1)
                ctl.view.ClearTracks();
            }
        }
        if newInfo.Track != ctl.info.Track || metadataChanged {
            ctl.updatePlayLists()
        }
        if newInfo.LoopMode != ctl.info.LoopMode {
            loopState, shuffle := loopModeToState(newInfo.LoopMode)
            ctl.view.SetLoopState(loopState)
            ctl.view.SetShuffle(shuffle)
        }
        if newMd.Quality != md.Quality || newMd.Rate != md.Rate ||
           newMd.Format != md.Format || newMd.BitRate != md.BitRate {
            ctl.view.SetInfo(newMd.Quality, newMd.Rate, newMd.Format, newMd.BitRate)
        }
        if newInfo.TrackUsed != ctl.info.TrackUsed ||
           newInfo.TrackTotal != ctl.info.TrackTotal {
            ctl.view.SetSeek(newInfo.TrackUsed, newInfo.TrackTotal)
        }
        if newInfo.Volume != ctl.info.Volume {
            ctl.view.SetVolume(newInfo.Volume)
        }
        if newInfo.Mute != ctl.info.Mute {
            ctl.view.SetMute(newInfo.Mute)
        }
        if newInfo.TrackMetadata.AlbumArt != ctl.info.TrackMetadata.AlbumArt {
            id := newMd.Id
            ctl.pixbufRequest(newInfo.TrackMetadata.AlbumArt, AlbumImageWidth,
                              AlbumImageHeight, func(pixbuf *gdk.Pixbuf) {
                if id == ctl.info.TrackMetadata.Id {
                    ctl.view.SetAlbumArt(pixbuf)
                }
            })
        }
        ctl.info = newInfo
    })
}

func (ctl *Controller) updatePlayLists() {
    var err error
    var playLists []wiimdev.PlayList
    var currPlayList string
    device := ctl.device
    ctl.runThen(func() {
        playLists, currPlayList, err = device.GetPlayLists()
        logErr(err)
    }, func() {
        if playLists == nil {
            return
        }
        if !playListsEqual(playLists, ctl.playLists) {
            log.Println("Updating playlists")
            ctl.playLists = playLists
            ctl.view.ClearPlayLists()
            for _, playList := range ctl.playLists {
                log.Println("Playlist:", playList.Id)
                ctl.view.AddPlayList(playList.Name)
            }
            ctl.currPlayList = ""
        }
        if ctl.info.Track <= 0 {
            // Don't update tracks if not playing a playlist
            return
        }
        if currPlayList != ctl.currPlayList {
            log.Println("Current Playlist:", currPlayList)
            for i := range ctl.playLists {
                if ctl.playLists[i].Id == currPlayList {
                    ctl.view.SetPlayList(i)
                    break
                }
            }
            ctl.currPlayList = currPlayList
        }
        if ctl.currPlayList != "" {
            ctl.updateTracks(currPlayList)
        }
    })
}

func (ctl *Controller) updateTracks(playList string) {
    var err error
    var tracks []wiimdev.Track
    device := ctl.device
    ctl.runThen(func() {
        tracks, _, err = device.GetTracks(playList)
        logErr(err)
    }, func() {
        if tracks == nil || tracksEqual(tracks, ctl.tracks) {
            return
        }
        log.Println("Updating tracks for", playList)
        ctl.tracks = tracks
        ctl.view.ClearTracks()
        for i, track := range ctl.tracks {
            log.Printf("Track: %d - %s - %s\n", i, track.Metadata.Title, track.Metadata.Artist)
            ctl.view.AddTrack(track.Metadata.Title, track.Metadata.Artist)
            id := track.Id
            index := i
            ctl.pixbufRequest(track.Metadata.AlbumArt, TrackImageWidth,
                              TrackImageHeight, func(pixbuf *gdk.Pixbuf) {
                if pixbuf != nil {
                    if index < len(ctl.tracks) && id == ctl.tracks[index].Id {
                        logErr(ctl.view.SetTrackAlbumArt(index, pixbuf))
                    }
                }
            })
        }
        ctl.view.SetTrack(ctl.info.Track - 1)
    })
}

func (ctl *Controller) updateEqualizer() {
    var err error
    var eqList []string
    var currEq string
    device := ctl.device
    ctl.runThen(func() {
        eqList, err = device.GetEqualizerList();
        logErr(err)
        if err == nil {
            currEq, err = device.GetEqualizer();
            logErr(err)
        }
    }, func() {
        if eqList != nil {
            ctl.view.SetEqualizers(eqList, currEq)
            ctl.view.SetEqualizerOn(currEq != "Off")
        }
    })
}

type RunReq struct {
    run func()
    then func()
}

// Schedule a closure to execute on the run loop
func (ctl *Controller) run(run func()) {
    select {
        case ctl.runChan <- RunReq{run, nil}:
    }
}

// Schedule a closure to execute on the run loop and then run a
// second closure that executes on the main thread.
func (ctl *Controller) runThen(run func(), then func()) {
    select {
        case ctl.runChan <- RunReq{run, then}:
    }
}

// The run loop serializes closures and runs them outside of the main thread
func (ctl *Controller) runLoop() {
    for {
        req := <- ctl.runChan
        req.run()
        if req.then != nil {
            mainRun(req.then)
        }
    }
}

// Download a URL and return it as a byte array. Uses a disk cache
// to improve performance.
func (ctl *Controller) download(url string) ([]byte, error) {
    if data := ctl.downloadCache.Get(url); data != nil {
        return data, nil
    }
    log.Println("Downloading", url)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    ctl.downloadCache.Add(url, data)
    return data, nil
}

// Download the URL and generate a scaled pixbuf of the specified dimensions.
// The callback for the pixbuf is executed on the main thread when complete.
func (ctl *Controller) pixbufRequest(url string, width int, height int,
                                     complete func(*gdk.Pixbuf)) {
    if !strings.HasPrefix(url, "http:") && !strings.HasPrefix(url, "https:") {
        mainRun( func() { complete(nil) })
        return
    }
    req := PixbufRequest{
        url: url,
        width: width,
        height: height,
        complete: complete,
    }
    select {
        case ctl.pixbufChan <- req:
    }
}

func (ctl *Controller) pixbufLoop() {
    for {
        req := <-ctl.pixbufChan
        key := fmt.Sprintf("%s:%d:%d", req.url, req.width, req.height)
        pixbuf, ok := ctl.pixbufCache.Get(key)
        if !ok {
            data, err := ctl.download(req.url)
            logErr(err)
            if data != nil {
                pixbuf, err = newScaledPixbufFromData(data, req.width, req.height)
                logErr(err)
                if pixbuf != nil {
                    ctl.pixbufCache.Add(key, pixbuf)
                }
            }
        }
        mainRun( func() { req.complete(pixbuf) })
    }
}

func (ctl *Controller) connectToDevice(deviceIndex int) {
    var err error
    var device *wiimdev.WiimDevice
    url := ctl.devices[deviceIndex].Url
    ctl.clearState()
    ctl.view.SetStatus("Connecting to " + urlToHost(url) + "...")
    ctl.device = wiimdev.NewFakeDevice()
    log.Println("Connecting to", url)
    ctl.runThen(func() {
        device, err = wiimdev.NewWiimDevice(url)
        logErr(err)
    }, func() {
        if device == nil {
            ctl.view.SetStatus("Cannot connect to " + urlToHost(url) + ".")
            ctl.view.SetPlayerState(PlayerPaused)
            return
        }
        ctl.view.SetStatus("")
        ctl.device = device
        ctl.config.Set("DeviceIndex", deviceIndex)
        logErr(ctl.config.Save())
        ctl.updateEqualizer()
        ctl.startUpdate()
    })
}

func (ctl *Controller) discoverDevices() {
    var err error
    var discovers []*wiimdev.DeviceDiscovery
    ctl.clearState()
    ctl.view.SetStatus("Discovering WiiM devices...")
    ctl.device = wiimdev.NewFakeDevice()
    ctl.view.ClearDevices()
    ctl.devices = nil
    log.Println("Discovering WiiM devices")
    ctl.runThen(func() {
        discovers, err = wiimdev.DiscoverWiimDevices()
        logErr(err)
    }, func() {
        ctl.view.SetPlayerState(PlayerPaused)
        if discovers == nil {
            ctl.view.SetStatus("WiiM discovery failed.")
            return
        }
        if len(discovers) == 0 {
            log.Println("No WiiM devices found")
            ctl.view.SetStatus("No WiiM devices found.")
        } else {
            plural := ""
            if len(discovers) != 1 {
                plural = "s"
            }
            str := fmt.Sprintf("Discovered %d WiiM device%s.",
                len(discovers), plural)
            log.Println(str)
            ctl.view.SetStatus(str)
            for _, discover := range discovers {
                ctl.view.AddDevice(discover.Name)
                log.Println("Device:", discover.Name, discover.Url)
                ctl.devices = append(ctl.devices, Device{
                    Name: discover.Name, Url: discover.Url })
            }
        }
        ctl.config.Set("Devices", ctl.devices)
        logErr(ctl.config.Save())
    })
}

func (ctl *Controller) clearState() {
    ctl.stopUpdate()

    ctl.playLists = nil
    ctl.tracks = nil
    ctl.info = &wiimdev.DeviceInfo{
        Track: -1,
        TrackMetadata: &wiimdev.Metadata{},
    }

    ctl.view.SetPlayerState(PlayerUnavailable)
    loopState, shuffle := loopModeToState(0)
    ctl.view.SetLoopState(loopState)
    ctl.view.SetShuffle(shuffle)
    ctl.view.SetVolume(0)
    ctl.view.ClearPlayLists()
    ctl.view.ClearTracks()
    ctl.view.SetAlbumArt(nil)
    ctl.view.SetInfo("", 0, "", 0)
    ctl.view.SetSeek(time.Duration(0), time.Duration(0))
    ctl.view.ClearEqualizers()
}

///////////////////////////////////////////////////////////////////////////////
// Public interface
///////////////////////////////////////////////////////////////////////////////

func NewController(application *gtk.Application, config ConfigAdapter,
                   version string, build string) (*Controller, error) {
    var err error

    ctl := &Controller{}
    ctl.application = application
    ctl.config = config
    ctl.version = version
    ctl.build = build

    ctl.mprisServer = mpris.NewMprisServer("wiimplay", ctl)
    ctl.mprisServer.Listen(ctl.config.Get("MprisSupport", true).(bool))

    ctl.pixbufCache, err = lru.New[string, *gdk.Pixbuf](
        config.Get("PixbufCacheSize", 50).(int))
    if err != nil {
        return nil, err
    }

    ctl.runChan = make(chan RunReq, 10)
    go ctl.runLoop()

    ctl.pixbufChan = make(chan PixbufRequest, 10)
    go ctl.pixbufLoop()

    cacheDir := filepath.Join(glib.GetUserCacheDir(), "wiimplay")
    err = os.MkdirAll(cacheDir, 0755)
    if err != nil {
        return nil, err
    }
    ctl.downloadCache, err = dcache.NewDiskCache(cacheDir,
        config.Get("DownloadCacheSize", 500).(int))
    if err != nil {
        return nil, err
    }

    ctl.view, err = NewView(application, ctl)
    if err != nil {
        return nil, err
    }

    if ctl.config.Get("HideOnStart", false).(bool) {
        ctl.view.Hide()
    }
    ctl.view.HideOnDelete(ctl.config.Get("HideOnClose", true).(bool))
    ctl.view.ShowStatusIcon(ctl.config.Get("ShowStatusIcon", true).(bool))

    devices := config.Get("Devices", []interface{}{}).([]interface{})
    if len(devices) > 0 {
        for _, item := range devices {
            deviceMap := item.(map[string]interface{})
            name := deviceMap["Name"].(string)
            url := deviceMap["Url"].(string)
            ctl.devices = append(ctl.devices, Device{ Name: name, Url: url})
            ctl.view.AddDevice(name)
        }
        deviceIndex := config.Get("DeviceIndex", 0).(int)
        if deviceIndex >= 0 && deviceIndex < len(ctl.devices) {
            ctl.view.SetDevice(deviceIndex)
            ctl.connectToDevice(deviceIndex)
        } else {
            return nil, errors.New("Invalid DeviceIndex in config")
        }
    } else {
        ctl.discoverDevices()
    }

    return ctl, nil
}
