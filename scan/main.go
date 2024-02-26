package main

import (
    "encoding/xml"
    "fmt"
    "net/url"
    "strings"
    "github.com/shumatech/wiimplay/upnp"

    "github.com/huin/goupnp"
)

func main() {
    devices, err := goupnp.DiscoverDevices("urn:schemas-upnp-org:device:MediaRenderer:1")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    for _, device := range devices {
        if strings.HasPrefix(device.Root.Device.ModelDescription, "WiiM") {
            err := dumpDevice(device)
            if err != nil {
                fmt.Println(err)
            } else {
                baseUrl, err := url.Parse(device.Root.URLBaseStr)
                if err != nil {
                    fmt.Println(err)
                } else {
                    err = dumpGetInfoEx(device.Root, baseUrl)
                    if err != nil {
                        fmt.Println(err)
                    }
                    err = dumpTotalQueue(device.Root, baseUrl)
                    if err != nil {
                        fmt.Println(err)
                    }
                    err = dumpCurrentQueue(device.Root, baseUrl)
                    if err != nil {
                        fmt.Println(err)
                    }
                }
            }
        }
    }
}

func printMetadata(metadata string) {
    didl := &upnp.DidlLiteXml{}
    err := xml.Unmarshal([]byte(metadata), &didl)
    if err != nil {
        fmt.Printf("Could not decode metadata: %v\n", err)
    } else {
        item := didl.Items[0]
        fmt.Printf("Metadata:\n")
        fmt.Printf("  Id:          %s\n", item.Id)
        fmt.Printf("  SubId:       %s\n", item.SubId)
        fmt.Printf("  Description: %s\n", item.Description)
        fmt.Printf("  SkipLimit:   %d\n", item.SkipLimit)
        fmt.Printf("  Like:        %d\n", item.Like)
        fmt.Printf("  Res:         %s\n", item.Res)
        fmt.Printf("  Title:       %s\n", item.Title)
        fmt.Printf("  Album:       %s\n", item.Album)
        fmt.Printf("  Artist:      %s\n", item.Artist)
        fmt.Printf("  Creator:     %s\n", item.Creator)
        fmt.Printf("  ThumbRating: %s\n", item.ThumbRating)
        fmt.Printf("  RatingUri:   %s\n", item.RatingUri)
        fmt.Printf("  AlbumArt:    %s\n", item.AlbumArt)
        fmt.Printf("  Rate:        %d\n", item.Rate)
        fmt.Printf("  Format:      %s\n", item.Format)
        fmt.Printf("  Quality:     %s\n", item.Quality)
        fmt.Printf("  BitRate:     %d\n", item.BitRate)
    }
}

func dumpDevice(device goupnp.MaybeRootDevice) error{
    fmt.Printf("\n=== Device ===\n")
    fmt.Printf("Location:      %v\n", device.Location)
    fmt.Printf("USN:           %v\n", device.USN)
    if device.Err != nil {
        return fmt.Errorf("%s", device.Err);
    }

    fmt.Printf("Root:          v%d.%d @ %s\n",
        device.Root.SpecVersion.Major, device.Root.SpecVersion.Minor,
        device.Root.URLBaseStr)
    fmt.Printf("Type:          %s\n", device.Root.Device.DeviceType)
    fmt.Printf("Friendly name: %s\n", device.Root.Device.FriendlyName)
    fmt.Printf("Manufacturer:  %s\n", device.Root.Device.Manufacturer)
    fmt.Printf("Model desc:    %s\n", device.Root.Device.ModelDescription)
    fmt.Printf("Model name:    %s\n", device.Root.Device.ModelName)
    fmt.Printf("Model number:  %s\n", device.Root.Device.ModelNumber)
    fmt.Printf("Model type:    %s\n", device.Root.Device.ModelType)
    fmt.Printf("Serial number: %s\n", device.Root.Device.SerialNumber)
    fmt.Printf("UDN:           %s\n", device.Root.Device.UDN)
    fmt.Printf("UPC:           %s\n", device.Root.Device.UPC)

    fmt.Printf("Devices:\n")
    for _, device := range device.Root.Device.Devices {
        fmt.Printf("  Device type: %s\n", device.DeviceType)
    }
    fmt.Printf("Services:\n")
    for _, service := range device.Root.Device.Services {
        fmt.Printf("  Service type: %s\n", service.ServiceType)
    }
    return nil;
}

func dumpGetInfoEx(device *goupnp.RootDevice, baseUrl *url.URL) error {
    av1clients, err := upnp.NewAVTransport1ClientsFromRootDevice(device, baseUrl)
    if err != nil {
        return err
    }
    av1client := av1clients[0]

    currentTransportState, currentTransportStatus, currentSpeed, track, trackDuration,
        trackMetaData, trackURI, relTime, absTime, relCount, absCount, loopMode,
        playType, currentVolume, currentMute, currentChannel, slaveFlag, masterUUID,
        slaveList, playMedium, trackSource, internetAccess, verUpdateFlag,
        verUpdateStatus, batteryFlag, batteryPercent, alarmFlag, timeStamp, subNum,
        spotifyActive, err := av1client.GetInfoEx(0)

    if err != nil {
        return err
    }

    fmt.Printf("\n=== GetInfoEx ===\n")
    fmt.Printf("CurrentTransportState:  %s\n", currentTransportState)
    fmt.Printf("CurrentTransportStatus: %s\n", currentTransportStatus)
    fmt.Printf("CurrentSpeed:           %s\n", currentSpeed)
    fmt.Printf("Track:                  %d\n", track)
    fmt.Printf("TrackDuration:          %s\n", trackDuration)
    fmt.Printf("TrackURI:               %s\n", trackURI)
    fmt.Printf("RelTime:                %s\n", relTime)
    fmt.Printf("AbsTime:                %s\n", absTime)
    fmt.Printf("RelCount:               %d\n", relCount)
    fmt.Printf("AbsCount:               %d\n", absCount)
    fmt.Printf("LoopMode:               %d\n", loopMode)
    fmt.Printf("PlayType:               %s\n", playType)
    fmt.Printf("CurrentVolume:          %d\n", currentVolume)
    fmt.Printf("CurrentMute:            %d\n", currentMute)
    fmt.Printf("CurrentChannel:         %d\n", currentChannel)
    fmt.Printf("SlaveFlag:              %s\n", slaveFlag)
    fmt.Printf("MasterUUID:             %s\n", masterUUID)
    fmt.Printf("SlaveList:              %s\n", slaveList)
    fmt.Printf("PlayMedium:             %s\n", playMedium)
    fmt.Printf("TrackSource:            %s\n", trackSource)
    fmt.Printf("InternetAccess:         %d\n", internetAccess)
    fmt.Printf("VerUpdateFlag:          %d\n", verUpdateFlag)
    fmt.Printf("VerUpdateStatus:        %d\n", verUpdateStatus)
    fmt.Printf("BatteryFlag:            %d\n", batteryFlag)
    fmt.Printf("BatteryPercent:         %d\n", batteryPercent)
    fmt.Printf("AlarmFlag:              %d\n", alarmFlag)
    fmt.Printf("TimeStamp:              %d\n", timeStamp)
    fmt.Printf("SubNum:                 %d\n", subNum)
    fmt.Printf("SpotifyActive:          %d\n", spotifyActive)

    printMetadata(trackMetaData)

    return nil
}

func dumpTotalQueue(device *goupnp.RootDevice, baseUrl *url.URL) error {
    pqclients, err := upnp.NewPlayQueue1ClientsFromRootDevice(device, baseUrl)
    if err != nil {
        return err
    }
    pqclient := pqclients[0]

    playlists := []string{}

    context, err := pqclient.BrowseQueue("TotalQueue")
    if err != nil {
        return err
    }

    tpq := &upnp.TotalPlayQueueXml{}
    err = xml.Unmarshal([]byte(context), &tpq)
    if err != nil {
        return err
    }

    fmt.Printf("\n=== BrowseQueue(TotalQueue) ===\n")
    fmt.Printf("TotalQueue:      %d\n", tpq.TotalQueue)
    fmt.Printf("CurrentPlayList: %s\n", tpq.CurrentPlayList)

    reader := strings.NewReader(tpq.PlayListInfo.InnerXml)
    decoder := xml.NewDecoder(reader)
    for i := 0; i < tpq.TotalQueue; i++ {
        var tpl upnp.TotalPlayListXml
        err = decoder.Decode(&tpl)
        if err != nil {
            fmt.Printf("Could not decode TotalPlayListItem: %v\n", err)
        } else {
            playlists = append(playlists, tpl.Name)
            fmt.Printf("\nName:           %s\n", tpl.Name)
            listinfo := tpl.ListInfo
            fmt.Printf("Source:         %s\n", listinfo.Source)
            fmt.Printf("SearchUrl:      %s\n", listinfo.SearchUrl)
            fmt.Printf("AutoGenerate:   %d\n", listinfo.AutoGenerate)
            fmt.Printf("StationLimit:   %d\n", listinfo.StationLimit)
            fmt.Printf("MarkSearch:     %d\n", listinfo.MarkSearch)
            fmt.Printf("Quality:        %d\n", listinfo.Quality)
            fmt.Printf("UpdateTime:     %d\n", listinfo.UpdateTime)
            fmt.Printf("LastPlayIndex:  %d\n", listinfo.LastPlayIndex)
            fmt.Printf("AlarmPlayIndex: %d\n", listinfo.AlarmPlayIndex)
            fmt.Printf("RealIndex:      %d\n", listinfo.RealIndex)
            fmt.Printf("TrackNumber:    %d\n", listinfo.TrackNumber)
            fmt.Printf("SwitchPageMode: %d\n", listinfo.SwitchPageMode)
            fmt.Printf("PressType:      %d\n", listinfo.PressType)
            fmt.Printf("Volume:         %d\n", listinfo.Volume)
        }
    }
    return nil
}

func dumpCurrentQueue(device *goupnp.RootDevice, baseUrl *url.URL) error {
    pqclients, err := upnp.NewPlayQueue1ClientsFromRootDevice(device, baseUrl)
    if err != nil {
        return err
    }
    pqclient := pqclients[0]

    context, err := pqclient.BrowseQueue("CurrentQueue")
    if err != nil {
        return err
    }

    pl := &upnp.PlayListXml{}
    err = xml.Unmarshal([]byte(context), &pl)
    if err != nil {
        return err
    }

    fmt.Printf("\n=== BrowseQueue(CurrentQueue) ===\n")
    fmt.Printf("ListName:       %s\n", pl.ListName)
    listinfo := pl.ListInfo
    fmt.Printf("Source:         %s\n", listinfo.SourceName)
    fmt.Printf("MarkSearch:     %d\n", listinfo.MarkSearch)
    fmt.Printf("TrackNumber:    %d\n", listinfo.TrackNumber)
    fmt.Printf("Quality:        %d\n", listinfo.Quality)
    fmt.Printf("UpdateTime:     %d\n", listinfo.UpdateTime)
    fmt.Printf("LastPlayIndex:  %d\n", listinfo.LastPlayIndex)
    fmt.Printf("AlarmPlayIndex: %d\n", listinfo.AlarmPlayIndex)
    fmt.Printf("RealIndex:      %d\n", listinfo.RealIndex)
    fmt.Printf("Type:           %d\n", listinfo.Type)
    fmt.Printf("SwitchPageMode: %d\n", listinfo.SwitchPageMode)
    fmt.Printf("CurrentPage:    %d\n", listinfo.CurrentPage)
    fmt.Printf("TotalPages:     %d\n", listinfo.TotalPages)
    fmt.Printf("Searching:      %d\n", listinfo.Searching)
    fmt.Printf("PressType:      %d\n", listinfo.PressType)
    fmt.Printf("Volume:         %d\n", listinfo.Volume)
    fmt.Printf("FadeEnable:     %d\n", listinfo.FadeEnable)
    fmt.Printf("FadeInMS:       %d\n", listinfo.FadeInMS)
    fmt.Printf("FadeOutMS:      %d\n", listinfo.FadeOutMS)

    reader := strings.NewReader(pl.Tracks.InnerXml)
    decoder := xml.NewDecoder(reader)
    for i := 0; i < listinfo.TrackNumber; i++ {
        var plt upnp.PlayListTrackXml
        err = decoder.Decode(&plt)
        if err != nil {
            fmt.Printf("Could not decode PlayListTrack: %v\n", err)
        } else {
            fmt.Printf("\nTrack %d\n", i + 1)
            fmt.Printf("URL:           %s\n", plt.URL)
            fmt.Printf("Id:            %s\n", plt.Id)
            fmt.Printf("Source:        %s\n", plt.Source)
            fmt.Printf("RefreshUrl:    %s\n", plt.RefreshUrl)
            fmt.Printf("PlayEventUrl:  %s\n", plt.PlayEventUrl)
            fmt.Printf("Expires:       %s\n", plt.Expires)
            fmt.Printf("Key:           %s\n", plt.Key)
            fmt.Printf("ChapterNumber: %d\n", plt.ChapterNumber)

            printMetadata(plt.Metadata)

            if i == 19 {
                break
            }
        }
    }
    return nil
}
