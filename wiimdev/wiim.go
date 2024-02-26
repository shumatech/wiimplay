package wiimdev

import (
    "crypto/tls"
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "time"

    "github.com/shumatech/wiimplay/upnp"

    "github.com/huin/goupnp"
)

type WiimDevice struct {
    transport *upnp.AVTransport1
    control *upnp.RenderingControl1
    playQueue *upnp.PlayQueue1
    url string
    client *http.Client
}

func parseDuration(str string) time.Duration {
    var h, m, s int
    _, err := fmt.Sscanf(str, "%d:%d:%d", &h, &m, &s)
    if err != nil {
        return time.Duration(0)
    }
    dur := time.Duration(h) * time.Hour + time.Duration(m) * time.Minute + time.Duration(s) * time.Second
    if dur.Seconds() < 0 {
        dur = time.Duration(0)
    }
    return dur
}

func getMetadata(metadata string) *Metadata {
    didl := &upnp.DidlLiteXml{}
    err := xml.Unmarshal([]byte(metadata), &didl)
    if err != nil {
        return &Metadata{Error: err}
    }
    md := didl.Items[0]
    return &Metadata{
        Id: md.Id,
        Title: md.Title,
        Album: md.Album,
        Artist: md.Artist,
        AlbumArt: md.AlbumArt,
        Rate: md.Rate,
        Format: md.Format,
        Quality: md.Quality,
        BitRate: md.BitRate,
        Error: nil,
    }
}

func newDeviceDiscovery(root *goupnp.RootDevice) *DeviceDiscovery {
    return &DeviceDiscovery{
        Url: strings.TrimSpace(root.URLBaseStr),
        Name: strings.TrimSpace(root.Device.FriendlyName),
        ModelName: strings.TrimSpace(root.Device.ModelName),
        ModelNumber: strings.TrimSpace(root.Device.ModelNumber),
        Uuid: strings.TrimSpace(root.Device.UDN),
    }
}

func (device *WiimDevice) command(command string, result interface{}) error {
    response, err := device.client.Get(device.url + command)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    data, err := io.ReadAll(response.Body)
    if err != nil {
        return err
    }

    switch result.(type) {
    case string:
        result = string(data)
    case int:
        result, err = strconv.Atoi(string(data))
    case float64:
        result, err = strconv.ParseFloat(string(data), 64)
    default:
        err = json.Unmarshal(data, &result)
    }
    return err
}

func (device *WiimDevice) GetInfo() (*DeviceInfo, error) {
    var err error
    var metadata string

    // Auto-generated return ugliness
    stateString, _, _, track, trackTotal, metadata, _, trackUsed,
        _, _, _, loopMode, _, volume, mute,
        _, _, _, _, _, _, _, _, _, _, _, _, _, _, _,
        err := device.transport.GetInfoEx(0)
    if err != nil {
        return nil, err
    }
    var transportState TransportState
    switch {
    case stateString == "STOPPED":
        transportState = StateStopped
    case stateString == "PAUSED_PLAYBACK":
        transportState = StatePaused
    case stateString == "PLAYING":
        transportState = StatePlaying
    case stateString == "TRANSITIONING":
        transportState = StateTransitioning
    case stateString == "NO_MEDIA_PRESENT":
        transportState = StateNoMedia
    default:
        transportState = StateUnknown
    }

    info := &DeviceInfo{
        TransportState: transportState,
        Track: int(track),
        TrackTotal: parseDuration(trackTotal),
        TrackUsed: parseDuration(trackUsed),
        Volume: int(volume),
        Mute: mute != 0,
        LoopMode: LoopMode(loopMode),
        TrackMetadata: getMetadata(metadata),
    }

    return info, nil
}

func (device *WiimDevice) Play() error {
    return device.transport.Play(0, "1")
}

func (device *WiimDevice) Pause() error {
    return device.transport.Pause(0)
}

func (device *WiimDevice) Previous() error {
    return device.transport.Previous(0)
}

func (device *WiimDevice) Next() error {
    return device.transport.Next(0)
}

func (device *WiimDevice) LoopMode(loopMode LoopMode) error {
    return device.playQueue.SetQueueLoopMode(uint32(loopMode))
}

func (device *WiimDevice) Volume(level int) error {
    return device.control.SetVolume(0, "Master", uint16(level))
}

func (device *WiimDevice) Seek(position time.Duration) error {
    h := int(position.Hours())
    m := int(position.Minutes()) % 60
    s := int(position.Seconds()) % 60
    target := fmt.Sprintf("%02d:%02d:%02d", h, m, s)
    return device.transport.Seek(0, "REL_TIME", target)
}

func (device *WiimDevice) Mute(mute bool) error {
    return device.control.SetMute(0, "Master", mute)
}

func (device *WiimDevice) GetPlayLists() ([]PlayList, string, error) {
    queue, err := device.playQueue.BrowseQueue("TotalQueue")
    if err != nil {
        return nil, "", err
    }

    tpq := &upnp.TotalPlayQueueXml{}
    err = xml.Unmarshal([]byte(queue), &tpq)
    if err != nil {
        return nil, "", err
    }

    playLists := []PlayList{}
    reader := strings.NewReader(tpq.PlayListInfo.InnerXml)
    decoder := xml.NewDecoder(reader)
    for i := 0; i < tpq.TotalQueue; i++ {
        var tpl upnp.TotalPlayListXml
        err = decoder.Decode(&tpl)
        if err == nil {
            name, _, _ := strings.Cut(tpl.Name, "_")
            playLists = append(playLists, PlayList{
                Id: tpl.Name,
                Name: name,
                Source: tpl.ListInfo.Source,
            })
        }
    }
    return playLists, tpq.CurrentPlayList, nil
}

func (device *WiimDevice) GetTracks(id string) ([]Track, int, error) {
    queue, err := device.playQueue.BrowseQueue(id)
    if err != nil {
        return nil, 0, err
    }

    pl := &upnp.PlayListXml{}
    err = xml.Unmarshal([]byte(queue), &pl)
    if err != nil {
        return nil, 0, err
    }

    tracks := []Track{}
    reader := strings.NewReader(pl.Tracks.InnerXml)
    decoder := xml.NewDecoder(reader)
    for i := 0; i < pl.ListInfo.TrackNumber; i++ {
        var plt upnp.PlayListTrackXml
        err = decoder.Decode(&plt)
        if err == nil {
            tracks = append(tracks, Track{
                Id: plt.Id,
                Source: plt.Source,
                Metadata: getMetadata(plt.Metadata),
            })
        }
    }
    return tracks, pl.ListInfo.LastPlayIndex, nil
}

func (device *WiimDevice) PlayTrack(track int) error {
    return device.playQueue.PlayQueueWithIndex("CurrentQueue", uint32(track))
}

func (device *WiimDevice) SetPlayList(id string) error {
    return device.playQueue.PlayQueueWithIndex(id, 1)
}

func (device *WiimDevice) GetDeviceDiscovery() *DeviceDiscovery {
    return newDeviceDiscovery(device.transport.RootDevice)
}

func (device *WiimDevice) SetEqualizer(setting string) error {
    result := &struct {
        Status string
    }{}
    var command string
    if setting == "Off" {
        command = "EQOff"
    } else {
        command = "EQLoad:" + setting
    }
    err := device.command(command, &result)
    if err != nil {
        return err
    }
    if result.Status != "OK" {
        return fmt.Errorf("EQLoad returned %s", result.Status)
    }
    return nil
}

func (device *WiimDevice) GetEqualizer() (string, error) {
    result := &struct {
        Status    string `json:"status"`
        EQStat    string `json:"EQStat"`
        Name      string `json:"Name"`
    }{}
    err := device.command("EQGetBand", &result)
    if err != nil {
        return "", err
    }
    if result.Status != "OK" {
        return "", fmt.Errorf("EQGetBand returned %s", result.Status)
    }
    if result.EQStat == "Off" {
        return "Off", nil
    }
    return result.Name, nil
}

func (device *WiimDevice) GetEqualizerList() ([]string, error) {
    result := []string{}
    err := device.command("EQGetList", &result)
    if err != nil {
        return nil, err
    }
    result = append(result, "Off")
    return result, nil
}

func (device *WiimDevice) GetAudioOutput() (int, error) {
    result := &struct {
        Hardware string
    }{}
    err := device.command("getNewAudioOutputHardwareMode", &result)
    if err != nil {
        return 0, err
    }
    hardware, err := strconv.Atoi(result.Hardware)
    if err != nil {
        return 0, err
    }
    return hardware - 1, nil
}

func (device *WiimDevice) GetAudioOutputList() ([]string, error) {
    return []string{"Optical Out", "Line Out", "Coax Out"}, nil
}

func (device *WiimDevice) SetAudioOutput(output int) error {
    var result string
    err := device.command(fmt.Sprintf("setAudioOutputHardwareMode:%d", output + 1), &result)
    if err != nil {
        return err
    }
    if result != "OK" {
        return fmt.Errorf("setAudioOutputHardwareMode failed")
    }
    return nil
}

func (device *WiimDevice) GetBalance() (float64, error) {
    result := 0.0
    err := device.command("getChannelBalance", &result)
    return result, err
}

func (device *WiimDevice) SetBalance(level float64) error {
    var result string
    err := device.command(fmt.Sprintf("setChannelBalance:%f", level), &result)
    if err != nil {
        return err
    }
    if result != "OK" {
        return fmt.Errorf("setChannelBalance failed")
    }
    return nil
}

func (device *WiimDevice) GetFadeEffects() (bool, error) {
    result := &struct {
        FadeFeature int `json:"FadeFeature"`
    }{}
    err := device.command("GetFadeFeature", &result)
    if err != nil {
        return false, err
    }
    return result.FadeFeature != 0, nil
}

func (device *WiimDevice) SetFadeEffects(on bool) error {
    var result string
    value := 0
    if on {
        value = 1
    }
    err := device.command(fmt.Sprintf("SetFadeFeature:%d", value), &result)
    if err != nil {
        return err
    }
    if result != "OK" {
        return fmt.Errorf("setChannelBalance failed")
    }
    return nil
}

func (device *WiimDevice) GetFixedVolume() (bool, error) {
    result := &struct {
        FixedVolume string `json:"volume_control"`
    }{}
    err := device.command("getStatusEx", &result)
    if err != nil {
        return false, err
    }
    return result.FixedVolume != "0", nil
}

func (device *WiimDevice) SetFixedVolume(on bool) error {
    var result string
    value := 0
    if on {
        value = 1
    }
    err := device.command(fmt.Sprintf("setVolumeControl:%d", value), &result)
    if err != nil {
        return err
    }
    if result != "OK" {
        return fmt.Errorf("setVolumeControl failed")
    }
    return nil
}

func NewWiimDevice(urlString string) (*WiimDevice, error) {
    device := &WiimDevice{}

    url, err := url.Parse(urlString)
    if err != nil {
        return nil, err
    }

    atClients, err := upnp.NewAVTransport1ClientsByURL(url)
    if err != nil {
        return nil, err
    }
    device.transport = atClients[0]

    rcClients, err := upnp.NewRenderingControl1ClientsByURL(url)
    if err != nil {
        return nil, err
    }
    device.control = rcClients[0]

    pqClients, err := upnp.NewPlayQueue1ClientsByURL(url)
    if err != nil {
        return nil, err
    }
    device.playQueue = pqClients[0]

    hostname := strings.Split(url.Host, ":")[0]
    device.url = "https://" + hostname + "/httpapi.asp?command="
    device.client = &http.Client{
        Timeout: 3 * time.Second,
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
        },
    }

    return device, nil
}

func DiscoverWiimDevices() ([]*DeviceDiscovery, error) {
    devices, err := goupnp.DiscoverDevices("urn:schemas-upnp-org:device:MediaRenderer:1")
    if err != nil {
        return nil, err
    }
    discovers := []*DeviceDiscovery{}
    for _, device := range devices {
        name := strings.TrimSpace(device.Root.Device.ModelName)
        if strings.HasPrefix(name, "WiiM") {
            discovers = append(discovers, newDeviceDiscovery(device.Root))
        }
    }
    return discovers, nil
}

