// Client for UPnP Device Control Protocol Wiim.
//
// This DCP is documented in detail at:
// - http://upnp.org/specs/av/av1/
//
// Typically, use one of the New* functions to create clients for services.
package upnp

// ***********************************************************
// GENERATED FILE - DO NOT EDIT BY HAND. See README.md
// ***********************************************************

import (
    "context"
    "net/url"
    "time"

    "github.com/huin/goupnp"
    "github.com/huin/goupnp/soap"
)

// Hack to avoid Go complaining if time isn't used.
var _ time.Time

// Device URNs:
const ()

// Service URNs:
const (
    URN_AVTransport_1       = "urn:schemas-upnp-org:service:AVTransport:1"
    URN_ConnectionManager_1 = "urn:schemas-upnp-org:service:ConnectionManager:1"
    URN_RenderingControl_1  = "urn:schemas-upnp-org:service:RenderingControl:1"
    URN_PlayQueue_1         = "urn:schemas-wiimu-com:service:PlayQueue:1"
)

// AVTransport1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:AVTransport:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type AVTransport1 struct {
    goupnp.ServiceClient
}

// NewAVTransport1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewAVTransport1ClientsCtx(ctx context.Context) (clients []*AVTransport1, errors []error, err error) {
    var genericClients []goupnp.ServiceClient
    if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_AVTransport_1); err != nil {
        return
    }
    clients = newAVTransport1ClientsFromGenericClients(genericClients)
    return
}

// NewAVTransport1Clients is the legacy version of NewAVTransport1ClientsCtx, but uses
// context.Background() as the context.
func NewAVTransport1Clients() (clients []*AVTransport1, errors []error, err error) {
    return NewAVTransport1ClientsCtx(context.Background())
}

// NewAVTransport1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewAVTransport1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*AVTransport1, error) {
    genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_AVTransport_1)
    if err != nil {
        return nil, err
    }
    return newAVTransport1ClientsFromGenericClients(genericClients), nil
}

// NewAVTransport1ClientsByURL is the legacy version of NewAVTransport1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewAVTransport1ClientsByURL(loc *url.URL) ([]*AVTransport1, error) {
    return NewAVTransport1ClientsByURLCtx(context.Background(), loc)
}

// NewAVTransport1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewAVTransport1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*AVTransport1, error) {
    genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_AVTransport_1)
    if err != nil {
        return nil, err
    }
    return newAVTransport1ClientsFromGenericClients(genericClients), nil
}

func newAVTransport1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*AVTransport1 {
    clients := make([]*AVTransport1, len(genericClients))
    for i := range genericClients {
        clients[i] = &AVTransport1{genericClients[i]}
    }
    return clients
}

func (client *AVTransport1) GetCurrentTransportActionsCtx(
    ctx context.Context,
    InstanceID uint32,
) (Actions string, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Actions string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetCurrentTransportActions", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Actions, err = soap.UnmarshalString(response.Actions); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetCurrentTransportActions is the legacy version of GetCurrentTransportActionsCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) GetCurrentTransportActions(InstanceID uint32) (Actions string, err error) {
    return client.GetCurrentTransportActionsCtx(context.Background(),
        InstanceID,
    )
}

// Return values:
//
// * PlayMedia: allowed values: UNKNOWN, CD-DA, DVD-VIDEO, HDD, NETWORK, NONE
//
// * RecMedia: allowed values: NOT_IMPLEMENTED
//
// * RecQualityModes: allowed values: NOT_IMPLEMENTED
func (client *AVTransport1) GetDeviceCapabilitiesCtx(
    ctx context.Context,
    InstanceID uint32,
) (PlayMedia string, RecMedia string, RecQualityModes string, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        PlayMedia       string
        RecMedia        string
        RecQualityModes string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetDeviceCapabilities", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if PlayMedia, err = soap.UnmarshalString(response.PlayMedia); err != nil {
        return
    }
    if RecMedia, err = soap.UnmarshalString(response.RecMedia); err != nil {
        return
    }
    if RecQualityModes, err = soap.UnmarshalString(response.RecQualityModes); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetDeviceCapabilities is the legacy version of GetDeviceCapabilitiesCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) GetDeviceCapabilities(InstanceID uint32) (PlayMedia string, RecMedia string, RecQualityModes string, err error) {
    return client.GetDeviceCapabilitiesCtx(context.Background(),
        InstanceID,
    )
}

// Return values:
//
// * CurrentTransportState: allowed values: STOPPED, PAUSED_PLAYBACK, PLAYING, TRANSITIONING, NO_MEDIA_PRESENT
//
// * CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
//
// * CurrentSpeed: allowed values: 1
//
// * Track: allowed value range: minimum=0, maximum=65535, step=1
//
// * PlayMedium: allowed values: UNKNOWN, CD-DA, DVD-VIDEO, HDD, NETWORK, NONE
func (client *AVTransport1) GetInfoExCtx(
    ctx context.Context,
    InstanceID uint32,
) (CurrentTransportState string, CurrentTransportStatus string, CurrentSpeed string, Track uint32, TrackDuration string, TrackMetaData string, TrackURI string, RelTime string, AbsTime string, RelCount int32, AbsCount int32, LoopMode uint32, PlayType string, CurrentVolume uint32, CurrentMute uint32, CurrentChannel uint32, SlaveFlag string, MasterUUID string, SlaveList string, PlayMedium string, TrackSource string, InternetAccess uint32, VerUpdateFlag uint32, VerUpdateStatus uint32, BatteryFlag uint32, BatteryPercent uint32, AlarmFlag uint32, TimeStamp uint32, SubNum uint32, SpotifyActive uint32, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        CurrentTransportState  string
        CurrentTransportStatus string
        CurrentSpeed           string
        Track                  string
        TrackDuration          string
        TrackMetaData          string
        TrackURI               string
        RelTime                string
        AbsTime                string
        RelCount               string
        AbsCount               string
        LoopMode               string
        PlayType               string
        CurrentVolume          string
        CurrentMute            string
        CurrentChannel         string
        SlaveFlag              string
        MasterUUID             string
        SlaveList              string
        PlayMedium             string
        TrackSource            string
        InternetAccess         string
        VerUpdateFlag          string
        VerUpdateStatus        string
        BatteryFlag            string
        BatteryPercent         string
        AlarmFlag              string
        TimeStamp              string
        SubNum                 string
        SpotifyActive          string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetInfoEx", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if CurrentTransportState, err = soap.UnmarshalString(response.CurrentTransportState); err != nil {
        return
    }
    if CurrentTransportStatus, err = soap.UnmarshalString(response.CurrentTransportStatus); err != nil {
        return
    }
    if CurrentSpeed, err = soap.UnmarshalString(response.CurrentSpeed); err != nil {
        return
    }
    if Track, err = soap.UnmarshalUi4(response.Track); err != nil {
        return
    }
    if TrackDuration, err = soap.UnmarshalString(response.TrackDuration); err != nil {
        return
    }
    if TrackMetaData, err = soap.UnmarshalString(response.TrackMetaData); err != nil {
        return
    }
    if TrackURI, err = soap.UnmarshalString(response.TrackURI); err != nil {
        return
    }
    if RelTime, err = soap.UnmarshalString(response.RelTime); err != nil {
        return
    }
    if AbsTime, err = soap.UnmarshalString(response.AbsTime); err != nil {
        return
    }
    if RelCount, err = soap.UnmarshalI4(response.RelCount); err != nil {
        return
    }
    if AbsCount, err = soap.UnmarshalI4(response.AbsCount); err != nil {
        return
    }
    if LoopMode, err = soap.UnmarshalUi4(response.LoopMode); err != nil {
        return
    }
    if PlayType, err = soap.UnmarshalString(response.PlayType); err != nil {
        return
    }
    if CurrentVolume, err = soap.UnmarshalUi4(response.CurrentVolume); err != nil {
        return
    }
    if CurrentMute, err = soap.UnmarshalUi4(response.CurrentMute); err != nil {
        return
    }
    if CurrentChannel, err = soap.UnmarshalUi4(response.CurrentChannel); err != nil {
        return
    }
    if SlaveFlag, err = soap.UnmarshalString(response.SlaveFlag); err != nil {
        return
    }
    if MasterUUID, err = soap.UnmarshalString(response.MasterUUID); err != nil {
        return
    }
    if SlaveList, err = soap.UnmarshalString(response.SlaveList); err != nil {
        return
    }
    if PlayMedium, err = soap.UnmarshalString(response.PlayMedium); err != nil {
        return
    }
    if TrackSource, err = soap.UnmarshalString(response.TrackSource); err != nil {
        return
    }
    if InternetAccess, err = soap.UnmarshalUi4(response.InternetAccess); err != nil {
        return
    }
    if VerUpdateFlag, err = soap.UnmarshalUi4(response.VerUpdateFlag); err != nil {
        return
    }
    if VerUpdateStatus, err = soap.UnmarshalUi4(response.VerUpdateStatus); err != nil {
        return
    }
    if BatteryFlag, err = soap.UnmarshalUi4(response.BatteryFlag); err != nil {
        return
    }
    if BatteryPercent, err = soap.UnmarshalUi4(response.BatteryPercent); err != nil {
        return
    }
    if AlarmFlag, err = soap.UnmarshalUi4(response.AlarmFlag); err != nil {
        return
    }
    if TimeStamp, err = soap.UnmarshalUi4(response.TimeStamp); err != nil {
        return
    }
    if SubNum, err = soap.UnmarshalUi4(response.SubNum); err != nil {
        return
    }
    if SpotifyActive, err = soap.UnmarshalUi4(response.SpotifyActive); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetInfoEx is the legacy version of GetInfoExCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) GetInfoEx(InstanceID uint32) (CurrentTransportState string, CurrentTransportStatus string, CurrentSpeed string, Track uint32, TrackDuration string, TrackMetaData string, TrackURI string, RelTime string, AbsTime string, RelCount int32, AbsCount int32, LoopMode uint32, PlayType string, CurrentVolume uint32, CurrentMute uint32, CurrentChannel uint32, SlaveFlag string, MasterUUID string, SlaveList string, PlayMedium string, TrackSource string, InternetAccess uint32, VerUpdateFlag uint32, VerUpdateStatus uint32, BatteryFlag uint32, BatteryPercent uint32, AlarmFlag uint32, TimeStamp uint32, SubNum uint32, SpotifyActive uint32, err error) {
    return client.GetInfoExCtx(context.Background(),
        InstanceID,
    )
}

// Return values:
//
// * NrTracks: allowed value range: minimum=0, maximum=65535
//
// * PlayMedium: allowed values: UNKNOWN, CD-DA, DVD-VIDEO, HDD, NETWORK, NONE
//
// * RecordMedium: allowed values: NOT_IMPLEMENTED
//
// * WriteStatus: allowed values: NOT_IMPLEMENTED
func (client *AVTransport1) GetMediaInfoCtx(
    ctx context.Context,
    InstanceID uint32,
) (NrTracks uint32, MediaDuration string, CurrentURI string, CurrentURIMetaData string, NextURI string, NextURIMetaData string, TrackSource string, PlayMedium string, RecordMedium string, WriteStatus string, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        NrTracks           string
        MediaDuration      string
        CurrentURI         string
        CurrentURIMetaData string
        NextURI            string
        NextURIMetaData    string
        TrackSource        string
        PlayMedium         string
        RecordMedium       string
        WriteStatus        string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetMediaInfo", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if NrTracks, err = soap.UnmarshalUi4(response.NrTracks); err != nil {
        return
    }
    if MediaDuration, err = soap.UnmarshalString(response.MediaDuration); err != nil {
        return
    }
    if CurrentURI, err = soap.UnmarshalString(response.CurrentURI); err != nil {
        return
    }
    if CurrentURIMetaData, err = soap.UnmarshalString(response.CurrentURIMetaData); err != nil {
        return
    }
    if NextURI, err = soap.UnmarshalString(response.NextURI); err != nil {
        return
    }
    if NextURIMetaData, err = soap.UnmarshalString(response.NextURIMetaData); err != nil {
        return
    }
    if TrackSource, err = soap.UnmarshalString(response.TrackSource); err != nil {
        return
    }
    if PlayMedium, err = soap.UnmarshalString(response.PlayMedium); err != nil {
        return
    }
    if RecordMedium, err = soap.UnmarshalString(response.RecordMedium); err != nil {
        return
    }
    if WriteStatus, err = soap.UnmarshalString(response.WriteStatus); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetMediaInfo is the legacy version of GetMediaInfoCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) GetMediaInfo(InstanceID uint32) (NrTracks uint32, MediaDuration string, CurrentURI string, CurrentURIMetaData string, NextURI string, NextURIMetaData string, TrackSource string, PlayMedium string, RecordMedium string, WriteStatus string, err error) {
    return client.GetMediaInfoCtx(context.Background(),
        InstanceID,
    )
}

func (client *AVTransport1) GetPlayTypeCtx(
    ctx context.Context,
    InstanceID uint32,
) (PlayType string, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        PlayType string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetPlayType", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if PlayType, err = soap.UnmarshalString(response.PlayType); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetPlayType is the legacy version of GetPlayTypeCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) GetPlayType(InstanceID uint32) (PlayType string, err error) {
    return client.GetPlayTypeCtx(context.Background(),
        InstanceID,
    )
}

// Return values:
//
// * Track: allowed value range: minimum=0, maximum=65535, step=1
func (client *AVTransport1) GetPositionInfoCtx(
    ctx context.Context,
    InstanceID uint32,
) (Track uint32, TrackDuration string, TrackMetaData string, TrackURI string, RelTime string, AbsTime string, RelCount int32, AbsCount int32, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Track         string
        TrackDuration string
        TrackMetaData string
        TrackURI      string
        RelTime       string
        AbsTime       string
        RelCount      string
        AbsCount      string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetPositionInfo", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Track, err = soap.UnmarshalUi4(response.Track); err != nil {
        return
    }
    if TrackDuration, err = soap.UnmarshalString(response.TrackDuration); err != nil {
        return
    }
    if TrackMetaData, err = soap.UnmarshalString(response.TrackMetaData); err != nil {
        return
    }
    if TrackURI, err = soap.UnmarshalString(response.TrackURI); err != nil {
        return
    }
    if RelTime, err = soap.UnmarshalString(response.RelTime); err != nil {
        return
    }
    if AbsTime, err = soap.UnmarshalString(response.AbsTime); err != nil {
        return
    }
    if RelCount, err = soap.UnmarshalI4(response.RelCount); err != nil {
        return
    }
    if AbsCount, err = soap.UnmarshalI4(response.AbsCount); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetPositionInfo is the legacy version of GetPositionInfoCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) GetPositionInfo(InstanceID uint32) (Track uint32, TrackDuration string, TrackMetaData string, TrackURI string, RelTime string, AbsTime string, RelCount int32, AbsCount int32, err error) {
    return client.GetPositionInfoCtx(context.Background(),
        InstanceID,
    )
}

// Return values:
//
// * CurrentTransportState: allowed values: STOPPED, PAUSED_PLAYBACK, PLAYING, TRANSITIONING, NO_MEDIA_PRESENT
//
// * CurrentTransportStatus: allowed values: OK, ERROR_OCCURRED
//
// * CurrentSpeed: allowed values: 1
func (client *AVTransport1) GetTransportInfoCtx(
    ctx context.Context,
    InstanceID uint32,
) (CurrentTransportState string, CurrentTransportStatus string, CurrentSpeed string, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        CurrentTransportState  string
        CurrentTransportStatus string
        CurrentSpeed           string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetTransportInfo", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if CurrentTransportState, err = soap.UnmarshalString(response.CurrentTransportState); err != nil {
        return
    }
    if CurrentTransportStatus, err = soap.UnmarshalString(response.CurrentTransportStatus); err != nil {
        return
    }
    if CurrentSpeed, err = soap.UnmarshalString(response.CurrentSpeed); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetTransportInfo is the legacy version of GetTransportInfoCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) GetTransportInfo(InstanceID uint32) (CurrentTransportState string, CurrentTransportStatus string, CurrentSpeed string, err error) {
    return client.GetTransportInfoCtx(context.Background(),
        InstanceID,
    )
}

// Return values:
//
// * PlayMode: allowed values: NORMAL, REPEAT_TRACK, REPEAT_ALL
//
// * RecQualityMode: allowed values: NOT_IMPLEMENTED
func (client *AVTransport1) GetTransportSettingsCtx(
    ctx context.Context,
    InstanceID uint32,
) (PlayMode string, RecQualityMode string, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        PlayMode       string
        RecQualityMode string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "GetTransportSettings", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if PlayMode, err = soap.UnmarshalString(response.PlayMode); err != nil {
        return
    }
    if RecQualityMode, err = soap.UnmarshalString(response.RecQualityMode); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetTransportSettings is the legacy version of GetTransportSettingsCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) GetTransportSettings(InstanceID uint32) (PlayMode string, RecQualityMode string, err error) {
    return client.GetTransportSettingsCtx(context.Background(),
        InstanceID,
    )
}

func (client *AVTransport1) NextCtx(
    ctx context.Context,
    InstanceID uint32,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Next", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// Next is the legacy version of NextCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Next(InstanceID uint32) (err error) {
    return client.NextCtx(context.Background(),
        InstanceID,
    )
}

func (client *AVTransport1) PauseCtx(
    ctx context.Context,
    InstanceID uint32,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Pause", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// Pause is the legacy version of PauseCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Pause(InstanceID uint32) (err error) {
    return client.PauseCtx(context.Background(),
        InstanceID,
    )
}

//
// Arguments:
//
// * Speed: allowed values: 1

func (client *AVTransport1) PlayCtx(
    ctx context.Context,
    InstanceID uint32,
    Speed string,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
        Speed      string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.Speed, err = soap.MarshalString(Speed); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Play", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// Play is the legacy version of PlayCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Play(InstanceID uint32, Speed string) (err error) {
    return client.PlayCtx(context.Background(),
        InstanceID,
        Speed,
    )
}

func (client *AVTransport1) PreviousCtx(
    ctx context.Context,
    InstanceID uint32,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Previous", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// Previous is the legacy version of PreviousCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Previous(InstanceID uint32) (err error) {
    return client.PreviousCtx(context.Background(),
        InstanceID,
    )
}

//
// Arguments:
//
// * Unit: allowed values: REL_TIME, TRACK_NR

func (client *AVTransport1) SeekCtx(
    ctx context.Context,
    InstanceID uint32,
    Unit string,
    Target string,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
        Unit       string
        Target     string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.Unit, err = soap.MarshalString(Unit); err != nil {
        return
    }
    if request.Target, err = soap.MarshalString(Target); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Seek", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// Seek is the legacy version of SeekCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Seek(InstanceID uint32, Unit string, Target string) (err error) {
    return client.SeekCtx(context.Background(),
        InstanceID,
        Unit,
        Target,
    )
}

func (client *AVTransport1) SeekBackwardCtx(
    ctx context.Context,
    InstanceID uint32,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "SeekBackward", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SeekBackward is the legacy version of SeekBackwardCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) SeekBackward(InstanceID uint32) (err error) {
    return client.SeekBackwardCtx(context.Background(),
        InstanceID,
    )
}

func (client *AVTransport1) SeekForwardCtx(
    ctx context.Context,
    InstanceID uint32,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "SeekForward", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SeekForward is the legacy version of SeekForwardCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) SeekForward(InstanceID uint32) (err error) {
    return client.SeekForwardCtx(context.Background(),
        InstanceID,
    )
}

func (client *AVTransport1) SetAVTransportURICtx(
    ctx context.Context,
    InstanceID uint32,
    CurrentURI string,
    CurrentURIMetaData string,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID         string
        CurrentURI         string
        CurrentURIMetaData string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.CurrentURI, err = soap.MarshalString(CurrentURI); err != nil {
        return
    }
    if request.CurrentURIMetaData, err = soap.MarshalString(CurrentURIMetaData); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "SetAVTransportURI", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetAVTransportURI is the legacy version of SetAVTransportURICtx, but uses
// context.Background() as the context.
func (client *AVTransport1) SetAVTransportURI(InstanceID uint32, CurrentURI string, CurrentURIMetaData string) (err error) {
    return client.SetAVTransportURICtx(context.Background(),
        InstanceID,
        CurrentURI,
        CurrentURIMetaData,
    )
}

func (client *AVTransport1) SetNextAVTransportURICtx(
    ctx context.Context,
    InstanceID uint32,
    NextURI string,
    NextURIMetaData string,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID      string
        NextURI         string
        NextURIMetaData string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.NextURI, err = soap.MarshalString(NextURI); err != nil {
        return
    }
    if request.NextURIMetaData, err = soap.MarshalString(NextURIMetaData); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "SetNextAVTransportURI", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetNextAVTransportURI is the legacy version of SetNextAVTransportURICtx, but uses
// context.Background() as the context.
func (client *AVTransport1) SetNextAVTransportURI(InstanceID uint32, NextURI string, NextURIMetaData string) (err error) {
    return client.SetNextAVTransportURICtx(context.Background(),
        InstanceID,
        NextURI,
        NextURIMetaData,
    )
}

//
// Arguments:
//
// * NewPlayMode: allowed values: NORMAL, REPEAT_TRACK, REPEAT_ALL

func (client *AVTransport1) SetPlayModeCtx(
    ctx context.Context,
    InstanceID uint32,
    NewPlayMode string,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID  string
        NewPlayMode string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.NewPlayMode, err = soap.MarshalString(NewPlayMode); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "SetPlayMode", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetPlayMode is the legacy version of SetPlayModeCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) SetPlayMode(InstanceID uint32, NewPlayMode string) (err error) {
    return client.SetPlayModeCtx(context.Background(),
        InstanceID,
        NewPlayMode,
    )
}

func (client *AVTransport1) StopCtx(
    ctx context.Context,
    InstanceID uint32,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_AVTransport_1, "Stop", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// Stop is the legacy version of StopCtx, but uses
// context.Background() as the context.
func (client *AVTransport1) Stop(InstanceID uint32) (err error) {
    return client.StopCtx(context.Background(),
        InstanceID,
    )
}

// ConnectionManager1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:ConnectionManager:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type ConnectionManager1 struct {
    goupnp.ServiceClient
}

// NewConnectionManager1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewConnectionManager1ClientsCtx(ctx context.Context) (clients []*ConnectionManager1, errors []error, err error) {
    var genericClients []goupnp.ServiceClient
    if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_ConnectionManager_1); err != nil {
        return
    }
    clients = newConnectionManager1ClientsFromGenericClients(genericClients)
    return
}

// NewConnectionManager1Clients is the legacy version of NewConnectionManager1ClientsCtx, but uses
// context.Background() as the context.
func NewConnectionManager1Clients() (clients []*ConnectionManager1, errors []error, err error) {
    return NewConnectionManager1ClientsCtx(context.Background())
}

// NewConnectionManager1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewConnectionManager1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*ConnectionManager1, error) {
    genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_ConnectionManager_1)
    if err != nil {
        return nil, err
    }
    return newConnectionManager1ClientsFromGenericClients(genericClients), nil
}

// NewConnectionManager1ClientsByURL is the legacy version of NewConnectionManager1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewConnectionManager1ClientsByURL(loc *url.URL) ([]*ConnectionManager1, error) {
    return NewConnectionManager1ClientsByURLCtx(context.Background(), loc)
}

// NewConnectionManager1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewConnectionManager1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*ConnectionManager1, error) {
    genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_ConnectionManager_1)
    if err != nil {
        return nil, err
    }
    return newConnectionManager1ClientsFromGenericClients(genericClients), nil
}

func newConnectionManager1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*ConnectionManager1 {
    clients := make([]*ConnectionManager1, len(genericClients))
    for i := range genericClients {
        clients[i] = &ConnectionManager1{genericClients[i]}
    }
    return clients
}

func (client *ConnectionManager1) GetCurrentConnectionIDsCtx(
    ctx context.Context,
) (ConnectionIDs string, err error) {
    // Request structure.
    request := interface{}(nil)
    // BEGIN Marshal arguments into request.

    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        ConnectionIDs string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_1, "GetCurrentConnectionIDs", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if ConnectionIDs, err = soap.UnmarshalString(response.ConnectionIDs); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetCurrentConnectionIDs is the legacy version of GetCurrentConnectionIDsCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager1) GetCurrentConnectionIDs() (ConnectionIDs string, err error) {
    return client.GetCurrentConnectionIDsCtx(context.Background())
}

// Return values:
//
// * Direction: allowed values: Input, Output
//
// * Status: allowed values: OK, ContentFormatMismatch, InsufficientBandwidth, UnreliableChannel, Unknown
func (client *ConnectionManager1) GetCurrentConnectionInfoCtx(
    ctx context.Context,
    ConnectionID int32,
) (RcsID int32, AVTransportID int32, ProtocolInfo string, PeerConnectionManager string, PeerConnectionID int32, Direction string, Status string, err error) {
    // Request structure.
    request := &struct {
        ConnectionID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.ConnectionID, err = soap.MarshalI4(ConnectionID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        RcsID                 string
        AVTransportID         string
        ProtocolInfo          string
        PeerConnectionManager string
        PeerConnectionID      string
        Direction             string
        Status                string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_1, "GetCurrentConnectionInfo", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if RcsID, err = soap.UnmarshalI4(response.RcsID); err != nil {
        return
    }
    if AVTransportID, err = soap.UnmarshalI4(response.AVTransportID); err != nil {
        return
    }
    if ProtocolInfo, err = soap.UnmarshalString(response.ProtocolInfo); err != nil {
        return
    }
    if PeerConnectionManager, err = soap.UnmarshalString(response.PeerConnectionManager); err != nil {
        return
    }
    if PeerConnectionID, err = soap.UnmarshalI4(response.PeerConnectionID); err != nil {
        return
    }
    if Direction, err = soap.UnmarshalString(response.Direction); err != nil {
        return
    }
    if Status, err = soap.UnmarshalString(response.Status); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetCurrentConnectionInfo is the legacy version of GetCurrentConnectionInfoCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager1) GetCurrentConnectionInfo(ConnectionID int32) (RcsID int32, AVTransportID int32, ProtocolInfo string, PeerConnectionManager string, PeerConnectionID int32, Direction string, Status string, err error) {
    return client.GetCurrentConnectionInfoCtx(context.Background(),
        ConnectionID,
    )
}

func (client *ConnectionManager1) GetProtocolInfoCtx(
    ctx context.Context,
) (Source string, Sink string, err error) {
    // Request structure.
    request := interface{}(nil)
    // BEGIN Marshal arguments into request.

    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Source string
        Sink   string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_ConnectionManager_1, "GetProtocolInfo", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Source, err = soap.UnmarshalString(response.Source); err != nil {
        return
    }
    if Sink, err = soap.UnmarshalString(response.Sink); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetProtocolInfo is the legacy version of GetProtocolInfoCtx, but uses
// context.Background() as the context.
func (client *ConnectionManager1) GetProtocolInfo() (Source string, Sink string, err error) {
    return client.GetProtocolInfoCtx(context.Background())
}

// PlayQueue1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:PlayQueue:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type PlayQueue1 struct {
    goupnp.ServiceClient
}

// NewPlayQueue1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewPlayQueue1ClientsCtx(ctx context.Context) (clients []*PlayQueue1, errors []error, err error) {
    var genericClients []goupnp.ServiceClient
    if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_PlayQueue_1); err != nil {
        return
    }
    clients = newPlayQueue1ClientsFromGenericClients(genericClients)
    return
}

// NewPlayQueue1Clients is the legacy version of NewPlayQueue1ClientsCtx, but uses
// context.Background() as the context.
func NewPlayQueue1Clients() (clients []*PlayQueue1, errors []error, err error) {
    return NewPlayQueue1ClientsCtx(context.Background())
}

// NewPlayQueue1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewPlayQueue1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*PlayQueue1, error) {
    genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_PlayQueue_1)
    if err != nil {
        return nil, err
    }
    return newPlayQueue1ClientsFromGenericClients(genericClients), nil
}

// NewPlayQueue1ClientsByURL is the legacy version of NewPlayQueue1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewPlayQueue1ClientsByURL(loc *url.URL) ([]*PlayQueue1, error) {
    return NewPlayQueue1ClientsByURLCtx(context.Background(), loc)
}

// NewPlayQueue1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewPlayQueue1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*PlayQueue1, error) {
    genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_PlayQueue_1)
    if err != nil {
        return nil, err
    }
    return newPlayQueue1ClientsFromGenericClients(genericClients), nil
}

func newPlayQueue1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*PlayQueue1 {
    clients := make([]*PlayQueue1, len(genericClients))
    for i := range genericClients {
        clients[i] = &PlayQueue1{genericClients[i]}
    }
    return clients
}

func (client *PlayQueue1) AppendQueueCtx(
    ctx context.Context,
    QueueContext string,
) (err error) {
    // Request structure.
    request := &struct {
        QueueContext string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueContext, err = soap.MarshalString(QueueContext); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "AppendQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// AppendQueue is the legacy version of AppendQueueCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) AppendQueue(QueueContext string) (err error) {
    return client.AppendQueueCtx(context.Background(),
        QueueContext,
    )
}

func (client *PlayQueue1) AppendTracksInQueueCtx(
    ctx context.Context,
    QueueContext string,
) (err error) {
    // Request structure.
    request := &struct {
        QueueContext string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueContext, err = soap.MarshalString(QueueContext); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "AppendTracksInQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// AppendTracksInQueue is the legacy version of AppendTracksInQueueCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) AppendTracksInQueue(QueueContext string) (err error) {
    return client.AppendTracksInQueueCtx(context.Background(),
        QueueContext,
    )
}

func (client *PlayQueue1) AppendTracksInQueueExCtx(
    ctx context.Context,
    QueueContext string,
    Direction uint32,
    StartIndex uint32,
    Play uint32,
) (err error) {
    // Request structure.
    request := &struct {
        QueueContext string
        Direction    string
        StartIndex   string
        Play         string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueContext, err = soap.MarshalString(QueueContext); err != nil {
        return
    }
    if request.Direction, err = soap.MarshalUi4(Direction); err != nil {
        return
    }
    if request.StartIndex, err = soap.MarshalUi4(StartIndex); err != nil {
        return
    }
    if request.Play, err = soap.MarshalUi4(Play); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "AppendTracksInQueueEx", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// AppendTracksInQueueEx is the legacy version of AppendTracksInQueueExCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) AppendTracksInQueueEx(QueueContext string, Direction uint32, StartIndex uint32, Play uint32) (err error) {
    return client.AppendTracksInQueueExCtx(context.Background(),
        QueueContext,
        Direction,
        StartIndex,
        Play,
    )
}

func (client *PlayQueue1) BackUpQueueCtx(
    ctx context.Context,
    QueueContext string,
) (err error) {
    // Request structure.
    request := &struct {
        QueueContext string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueContext, err = soap.MarshalString(QueueContext); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "BackUpQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// BackUpQueue is the legacy version of BackUpQueueCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) BackUpQueue(QueueContext string) (err error) {
    return client.BackUpQueueCtx(context.Background(),
        QueueContext,
    )
}

func (client *PlayQueue1) BrowseQueueCtx(
    ctx context.Context,
    QueueName string,
) (QueueContext string, err error) {
    // Request structure.
    request := &struct {
        QueueName string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        QueueContext string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "BrowseQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if QueueContext, err = soap.UnmarshalString(response.QueueContext); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// BrowseQueue is the legacy version of BrowseQueueCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) BrowseQueue(QueueName string) (QueueContext string, err error) {
    return client.BrowseQueueCtx(context.Background(),
        QueueName,
    )
}

func (client *PlayQueue1) CreateQueueCtx(
    ctx context.Context,
    QueueContext string,
) (err error) {
    // Request structure.
    request := &struct {
        QueueContext string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueContext, err = soap.MarshalString(QueueContext); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "CreateQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// CreateQueue is the legacy version of CreateQueueCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) CreateQueue(QueueContext string) (err error) {
    return client.CreateQueueCtx(context.Background(),
        QueueContext,
    )
}

func (client *PlayQueue1) DeleteActionQueueCtx(
    ctx context.Context,
    PressType uint32,
) (err error) {
    // Request structure.
    request := &struct {
        PressType string
    }{}
    // BEGIN Marshal arguments into request.

    if request.PressType, err = soap.MarshalUi4(PressType); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "DeleteActionQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// DeleteActionQueue is the legacy version of DeleteActionQueueCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) DeleteActionQueue(PressType uint32) (err error) {
    return client.DeleteActionQueueCtx(context.Background(),
        PressType,
    )
}

func (client *PlayQueue1) DeleteQueueCtx(
    ctx context.Context,
    QueueName string,
) (err error) {
    // Request structure.
    request := &struct {
        QueueName string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "DeleteQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// DeleteQueue is the legacy version of DeleteQueueCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) DeleteQueue(QueueName string) (err error) {
    return client.DeleteQueueCtx(context.Background(),
        QueueName,
    )
}

func (client *PlayQueue1) GetKeyMappingCtx(
    ctx context.Context,
) (QueueContext string, err error) {
    // Request structure.
    request := interface{}(nil)
    // BEGIN Marshal arguments into request.

    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        QueueContext string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "GetKeyMapping", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if QueueContext, err = soap.UnmarshalString(response.QueueContext); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetKeyMapping is the legacy version of GetKeyMappingCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) GetKeyMapping() (QueueContext string, err error) {
    return client.GetKeyMappingCtx(context.Background())
}

func (client *PlayQueue1) GetQueueIndexCtx(
    ctx context.Context,
    QueueName string,
) (CurrentIndex uint32, CurrentPage uint32, err error) {
    // Request structure.
    request := &struct {
        QueueName string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        CurrentIndex string
        CurrentPage  string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "GetQueueIndex", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if CurrentIndex, err = soap.UnmarshalUi4(response.CurrentIndex); err != nil {
        return
    }
    if CurrentPage, err = soap.UnmarshalUi4(response.CurrentPage); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetQueueIndex is the legacy version of GetQueueIndexCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) GetQueueIndex(QueueName string) (CurrentIndex uint32, CurrentPage uint32, err error) {
    return client.GetQueueIndexCtx(context.Background(),
        QueueName,
    )
}

func (client *PlayQueue1) GetQueueLoopModeCtx(
    ctx context.Context,
) (LoopMode uint32, err error) {
    // Request structure.
    request := interface{}(nil)
    // BEGIN Marshal arguments into request.

    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        LoopMode string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "GetQueueLoopMode", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if LoopMode, err = soap.UnmarshalUi4(response.LoopMode); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetQueueLoopMode is the legacy version of GetQueueLoopModeCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) GetQueueLoopMode() (LoopMode uint32, err error) {
    return client.GetQueueLoopModeCtx(context.Background())
}

func (client *PlayQueue1) GetQueueOnlineCtx(
    ctx context.Context,
    QueueName string,
    QueueID string,
    QueueType string,
    Queuelimit uint32,
    QueueAutoInsert string,
) (QueueContext string, err error) {
    // Request structure.
    request := &struct {
        QueueName       string
        QueueID         string
        QueueType       string
        Queuelimit      string
        QueueAutoInsert string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    if request.QueueID, err = soap.MarshalString(QueueID); err != nil {
        return
    }
    if request.QueueType, err = soap.MarshalString(QueueType); err != nil {
        return
    }
    if request.Queuelimit, err = soap.MarshalUi4(Queuelimit); err != nil {
        return
    }
    if request.QueueAutoInsert, err = soap.MarshalString(QueueAutoInsert); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        QueueContext string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "GetQueueOnline", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if QueueContext, err = soap.UnmarshalString(response.QueueContext); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetQueueOnline is the legacy version of GetQueueOnlineCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) GetQueueOnline(QueueName string, QueueID string, QueueType string, Queuelimit uint32, QueueAutoInsert string) (QueueContext string, err error) {
    return client.GetQueueOnlineCtx(context.Background(),
        QueueName,
        QueueID,
        QueueType,
        Queuelimit,
        QueueAutoInsert,
    )
}

func (client *PlayQueue1) GetUserAccountHistoryCtx(
    ctx context.Context,
    AccountSource string,
    Number uint32,
) (Result string, err error) {
    // Request structure.
    request := &struct {
        AccountSource string
        Number        string
    }{}
    // BEGIN Marshal arguments into request.

    if request.AccountSource, err = soap.MarshalString(AccountSource); err != nil {
        return
    }
    if request.Number, err = soap.MarshalUi4(Number); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Result string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "GetUserAccountHistory", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Result, err = soap.UnmarshalString(response.Result); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetUserAccountHistory is the legacy version of GetUserAccountHistoryCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) GetUserAccountHistory(AccountSource string, Number uint32) (Result string, err error) {
    return client.GetUserAccountHistoryCtx(context.Background(),
        AccountSource,
        Number,
    )
}

func (client *PlayQueue1) GetUserFavoritesCtx(
    ctx context.Context,
    AccountSource string,
    MediaType string,
    Filter string,
) (Result string, err error) {
    // Request structure.
    request := &struct {
        AccountSource string
        MediaType     string
        Filter        string
    }{}
    // BEGIN Marshal arguments into request.

    if request.AccountSource, err = soap.MarshalString(AccountSource); err != nil {
        return
    }
    if request.MediaType, err = soap.MarshalString(MediaType); err != nil {
        return
    }
    if request.Filter, err = soap.MarshalString(Filter); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Result string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "GetUserFavorites", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Result, err = soap.UnmarshalString(response.Result); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetUserFavorites is the legacy version of GetUserFavoritesCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) GetUserFavorites(AccountSource string, MediaType string, Filter string) (Result string, err error) {
    return client.GetUserFavoritesCtx(context.Background(),
        AccountSource,
        MediaType,
        Filter,
    )
}

func (client *PlayQueue1) GetUserInfoCtx(
    ctx context.Context,
    AccountSource string,
) (Result string, err error) {
    // Request structure.
    request := &struct {
        AccountSource string
    }{}
    // BEGIN Marshal arguments into request.

    if request.AccountSource, err = soap.MarshalString(AccountSource); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Result string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "GetUserInfo", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Result, err = soap.UnmarshalString(response.Result); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetUserInfo is the legacy version of GetUserInfoCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) GetUserInfo(AccountSource string) (Result string, err error) {
    return client.GetUserInfoCtx(context.Background(),
        AccountSource,
    )
}

func (client *PlayQueue1) PlayQueueWithIndexCtx(
    ctx context.Context,
    QueueName string,
    Index uint32,
) (err error) {
    // Request structure.
    request := &struct {
        QueueName string
        Index     string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    if request.Index, err = soap.MarshalUi4(Index); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "PlayQueueWithIndex", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// PlayQueueWithIndex is the legacy version of PlayQueueWithIndexCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) PlayQueueWithIndex(QueueName string, Index uint32) (err error) {
    return client.PlayQueueWithIndexCtx(context.Background(),
        QueueName,
        Index,
    )
}

func (client *PlayQueue1) RemoveTracksInQueueCtx(
    ctx context.Context,
    QueueName string,
    RangStart uint32,
    RangEnd uint32,
) (err error) {
    // Request structure.
    request := &struct {
        QueueName string
        RangStart string
        RangEnd   string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    if request.RangStart, err = soap.MarshalUi4(RangStart); err != nil {
        return
    }
    if request.RangEnd, err = soap.MarshalUi4(RangEnd); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "RemoveTracksInQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// RemoveTracksInQueue is the legacy version of RemoveTracksInQueueCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) RemoveTracksInQueue(QueueName string, RangStart uint32, RangEnd uint32) (err error) {
    return client.RemoveTracksInQueueCtx(context.Background(),
        QueueName,
        RangStart,
        RangEnd,
    )
}

func (client *PlayQueue1) ReplaceQueueCtx(
    ctx context.Context,
    QueueContext string,
) (err error) {
    // Request structure.
    request := &struct {
        QueueContext string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueContext, err = soap.MarshalString(QueueContext); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "ReplaceQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// ReplaceQueue is the legacy version of ReplaceQueueCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) ReplaceQueue(QueueContext string) (err error) {
    return client.ReplaceQueueCtx(context.Background(),
        QueueContext,
    )
}

func (client *PlayQueue1) SearchQueueOnlineCtx(
    ctx context.Context,
    QueueName string,
    SearchKey string,
    Queuelimit uint32,
) (QueueContext string, err error) {
    // Request structure.
    request := &struct {
        QueueName  string
        SearchKey  string
        Queuelimit string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    if request.SearchKey, err = soap.MarshalString(SearchKey); err != nil {
        return
    }
    if request.Queuelimit, err = soap.MarshalUi4(Queuelimit); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        QueueContext string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "SearchQueueOnline", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if QueueContext, err = soap.UnmarshalString(response.QueueContext); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// SearchQueueOnline is the legacy version of SearchQueueOnlineCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) SearchQueueOnline(QueueName string, SearchKey string, Queuelimit uint32) (QueueContext string, err error) {
    return client.SearchQueueOnlineCtx(context.Background(),
        QueueName,
        SearchKey,
        Queuelimit,
    )
}

func (client *PlayQueue1) SetKeyMappingCtx(
    ctx context.Context,
    QueueContext string,
) (err error) {
    // Request structure.
    request := &struct {
        QueueContext string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueContext, err = soap.MarshalString(QueueContext); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "SetKeyMapping", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetKeyMapping is the legacy version of SetKeyMappingCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) SetKeyMapping(QueueContext string) (err error) {
    return client.SetKeyMappingCtx(context.Background(),
        QueueContext,
    )
}

func (client *PlayQueue1) SetQueueLoopModeCtx(
    ctx context.Context,
    LoopMode uint32,
) (err error) {
    // Request structure.
    request := &struct {
        LoopMode string
    }{}
    // BEGIN Marshal arguments into request.

    if request.LoopMode, err = soap.MarshalUi4(LoopMode); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "SetQueueLoopMode", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetQueueLoopMode is the legacy version of SetQueueLoopModeCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) SetQueueLoopMode(LoopMode uint32) (err error) {
    return client.SetQueueLoopModeCtx(context.Background(),
        LoopMode,
    )
}

func (client *PlayQueue1) SetQueuePolicyCtx(
    ctx context.Context,
    QueueName string,
) (err error) {
    // Request structure.
    request := &struct {
        QueueName string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "SetQueuePolicy", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetQueuePolicy is the legacy version of SetQueuePolicyCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) SetQueuePolicy(QueueName string) (err error) {
    return client.SetQueuePolicyCtx(context.Background(),
        QueueName,
    )
}

func (client *PlayQueue1) SetQueueRecordCtx(
    ctx context.Context,
    QueueName string,
    QueueID string,
    Action string,
) (err error) {
    // Request structure.
    request := &struct {
        QueueName string
        QueueID   string
        Action    string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    if request.QueueID, err = soap.MarshalString(QueueID); err != nil {
        return
    }
    if request.Action, err = soap.MarshalString(Action); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "SetQueueRecord", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetQueueRecord is the legacy version of SetQueueRecordCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) SetQueueRecord(QueueName string, QueueID string, Action string) (err error) {
    return client.SetQueueRecordCtx(context.Background(),
        QueueName,
        QueueID,
        Action,
    )
}

func (client *PlayQueue1) SetSongsRecordCtx(
    ctx context.Context,
    QueueName string,
    SongID string,
    Action string,
) (err error) {
    // Request structure.
    request := &struct {
        QueueName string
        SongID    string
        Action    string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    if request.SongID, err = soap.MarshalString(SongID); err != nil {
        return
    }
    if request.Action, err = soap.MarshalString(Action); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "SetSongsRecord", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetSongsRecord is the legacy version of SetSongsRecordCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) SetSongsRecord(QueueName string, SongID string, Action string) (err error) {
    return client.SetSongsRecordCtx(context.Background(),
        QueueName,
        SongID,
        Action,
    )
}

func (client *PlayQueue1) SetSpotifyPresetCtx(
    ctx context.Context,
    KeyIndex uint32,
) (Result string, err error) {
    // Request structure.
    request := &struct {
        KeyIndex string
    }{}
    // BEGIN Marshal arguments into request.

    if request.KeyIndex, err = soap.MarshalUi4(KeyIndex); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Result string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "SetSpotifyPreset", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Result, err = soap.UnmarshalString(response.Result); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// SetSpotifyPreset is the legacy version of SetSpotifyPresetCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) SetSpotifyPreset(KeyIndex uint32) (Result string, err error) {
    return client.SetSpotifyPresetCtx(context.Background(),
        KeyIndex,
    )
}

func (client *PlayQueue1) SetUserFavoritesCtx(
    ctx context.Context,
    AccountSource string,
    Action string,
    MediaType string,
    MediaID string,
    MediaContext string,
) (Result string, err error) {
    // Request structure.
    request := &struct {
        AccountSource string
        Action        string
        MediaType     string
        MediaID       string
        MediaContext  string
    }{}
    // BEGIN Marshal arguments into request.

    if request.AccountSource, err = soap.MarshalString(AccountSource); err != nil {
        return
    }
    if request.Action, err = soap.MarshalString(Action); err != nil {
        return
    }
    if request.MediaType, err = soap.MarshalString(MediaType); err != nil {
        return
    }
    if request.MediaID, err = soap.MarshalString(MediaID); err != nil {
        return
    }
    if request.MediaContext, err = soap.MarshalString(MediaContext); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Result string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "SetUserFavorites", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Result, err = soap.UnmarshalString(response.Result); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// SetUserFavorites is the legacy version of SetUserFavoritesCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) SetUserFavorites(AccountSource string, Action string, MediaType string, MediaID string, MediaContext string) (Result string, err error) {
    return client.SetUserFavoritesCtx(context.Background(),
        AccountSource,
        Action,
        MediaType,
        MediaID,
        MediaContext,
    )
}

func (client *PlayQueue1) UserLoginCtx(
    ctx context.Context,
    AccountSource string,
    UserName string,
    PassWord string,
    SavePass uint32,
    Code uint32,
    Proxy string,
) (Result string, err error) {
    // Request structure.
    request := &struct {
        AccountSource string
        UserName      string
        PassWord      string
        SavePass      string
        Code          string
        Proxy         string
    }{}
    // BEGIN Marshal arguments into request.

    if request.AccountSource, err = soap.MarshalString(AccountSource); err != nil {
        return
    }
    if request.UserName, err = soap.MarshalString(UserName); err != nil {
        return
    }
    if request.PassWord, err = soap.MarshalString(PassWord); err != nil {
        return
    }
    if request.SavePass, err = soap.MarshalUi4(SavePass); err != nil {
        return
    }
    if request.Code, err = soap.MarshalUi4(Code); err != nil {
        return
    }
    if request.Proxy, err = soap.MarshalString(Proxy); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Result string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "UserLogin", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Result, err = soap.UnmarshalString(response.Result); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// UserLogin is the legacy version of UserLoginCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) UserLogin(AccountSource string, UserName string, PassWord string, SavePass uint32, Code uint32, Proxy string) (Result string, err error) {
    return client.UserLoginCtx(context.Background(),
        AccountSource,
        UserName,
        PassWord,
        SavePass,
        Code,
        Proxy,
    )
}

func (client *PlayQueue1) UserLogoutCtx(
    ctx context.Context,
    AccountSource string,
) (Result string, err error) {
    // Request structure.
    request := &struct {
        AccountSource string
    }{}
    // BEGIN Marshal arguments into request.

    if request.AccountSource, err = soap.MarshalString(AccountSource); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Result string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "UserLogout", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Result, err = soap.UnmarshalString(response.Result); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// UserLogout is the legacy version of UserLogoutCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) UserLogout(AccountSource string) (Result string, err error) {
    return client.UserLogoutCtx(context.Background(),
        AccountSource,
    )
}

func (client *PlayQueue1) UserRegisterCtx(
    ctx context.Context,
    QueueName string,
    UserName string,
    PassWord string,
) (Result string, err error) {
    // Request structure.
    request := &struct {
        QueueName string
        UserName  string
        PassWord  string
    }{}
    // BEGIN Marshal arguments into request.

    if request.QueueName, err = soap.MarshalString(QueueName); err != nil {
        return
    }
    if request.UserName, err = soap.MarshalString(UserName); err != nil {
        return
    }
    if request.PassWord, err = soap.MarshalString(PassWord); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        Result string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_PlayQueue_1, "UserRegister", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if Result, err = soap.UnmarshalString(response.Result); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// UserRegister is the legacy version of UserRegisterCtx, but uses
// context.Background() as the context.
func (client *PlayQueue1) UserRegister(QueueName string, UserName string, PassWord string) (Result string, err error) {
    return client.UserRegisterCtx(context.Background(),
        QueueName,
        UserName,
        PassWord,
    )
}

// RenderingControl1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:RenderingControl:1". See
// goupnp.ServiceClient, which contains RootDevice and Service attributes which
// are provided for informational value.
type RenderingControl1 struct {
    goupnp.ServiceClient
}

// NewRenderingControl1ClientsCtx discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewRenderingControl1ClientsCtx(ctx context.Context) (clients []*RenderingControl1, errors []error, err error) {
    var genericClients []goupnp.ServiceClient
    if genericClients, errors, err = goupnp.NewServiceClientsCtx(ctx, URN_RenderingControl_1); err != nil {
        return
    }
    clients = newRenderingControl1ClientsFromGenericClients(genericClients)
    return
}

// NewRenderingControl1Clients is the legacy version of NewRenderingControl1ClientsCtx, but uses
// context.Background() as the context.
func NewRenderingControl1Clients() (clients []*RenderingControl1, errors []error, err error) {
    return NewRenderingControl1ClientsCtx(context.Background())
}

// NewRenderingControl1ClientsByURLCtx discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewRenderingControl1ClientsByURLCtx(ctx context.Context, loc *url.URL) ([]*RenderingControl1, error) {
    genericClients, err := goupnp.NewServiceClientsByURLCtx(ctx, loc, URN_RenderingControl_1)
    if err != nil {
        return nil, err
    }
    return newRenderingControl1ClientsFromGenericClients(genericClients), nil
}

// NewRenderingControl1ClientsByURL is the legacy version of NewRenderingControl1ClientsByURLCtx, but uses
// context.Background() as the context.
func NewRenderingControl1ClientsByURL(loc *url.URL) ([]*RenderingControl1, error) {
    return NewRenderingControl1ClientsByURLCtx(context.Background(), loc)
}

// NewRenderingControl1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewRenderingControl1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*RenderingControl1, error) {
    genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, URN_RenderingControl_1)
    if err != nil {
        return nil, err
    }
    return newRenderingControl1ClientsFromGenericClients(genericClients), nil
}

func newRenderingControl1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*RenderingControl1 {
    clients := make([]*RenderingControl1, len(genericClients))
    for i := range genericClients {
        clients[i] = &RenderingControl1{genericClients[i]}
    }
    return clients
}

//
// Arguments:
//
// * Name: allowed values: Start, Stop

func (client *RenderingControl1) AirplayAutoSyncDelayCtx(
    ctx context.Context,
    Name string,
) (err error) {
    // Request structure.
    request := &struct {
        Name string
    }{}
    // BEGIN Marshal arguments into request.

    if request.Name, err = soap.MarshalString(Name); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "AirplayAutoSyncDelay", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// AirplayAutoSyncDelay is the legacy version of AirplayAutoSyncDelayCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) AirplayAutoSyncDelay(Name string) (err error) {
    return client.AirplayAutoSyncDelayCtx(context.Background(),
        Name,
    )
}

func (client *RenderingControl1) DeleteAlarmQueueCtx(
    ctx context.Context,
    AlarmName string,
) (err error) {
    // Request structure.
    request := &struct {
        AlarmName string
    }{}
    // BEGIN Marshal arguments into request.

    if request.AlarmName, err = soap.MarshalString(AlarmName); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "DeleteAlarmQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// DeleteAlarmQueue is the legacy version of DeleteAlarmQueueCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) DeleteAlarmQueue(AlarmName string) (err error) {
    return client.DeleteAlarmQueueCtx(context.Background(),
        AlarmName,
    )
}

func (client *RenderingControl1) GetAlarmQueueCtx(
    ctx context.Context,
    AlarmName string,
) (AlarmContext string, err error) {
    // Request structure.
    request := &struct {
        AlarmName string
    }{}
    // BEGIN Marshal arguments into request.

    if request.AlarmName, err = soap.MarshalString(AlarmName); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        AlarmContext string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetAlarmQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if AlarmContext, err = soap.UnmarshalString(response.AlarmContext); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetAlarmQueue is the legacy version of GetAlarmQueueCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) GetAlarmQueue(AlarmName string) (AlarmContext string, err error) {
    return client.GetAlarmQueueCtx(context.Background(),
        AlarmName,
    )
}

//
// Arguments:
//
// * Channel: allowed values: Master, Single

func (client *RenderingControl1) GetChannelCtx(
    ctx context.Context,
    InstanceID uint32,
    Channel string,
) (CurrentChannel uint16, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
        Channel    string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.Channel, err = soap.MarshalString(Channel); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        CurrentChannel string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetChannel", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if CurrentChannel, err = soap.UnmarshalUi2(response.CurrentChannel); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetChannel is the legacy version of GetChannelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) GetChannel(InstanceID uint32, Channel string) (CurrentChannel uint16, err error) {
    return client.GetChannelCtx(context.Background(),
        InstanceID,
        Channel,
    )
}

// Return values:
//
// * CurrentVolume: allowed value range: minimum=0, maximum=100, step=1
func (client *RenderingControl1) GetControlDeviceInfoCtx(
    ctx context.Context,
    InstanceID uint32,
) (MultiType uint16, PlayMode uint16, Router string, Ssid string, SlaveMask uint16, CurrentVolume uint16, CurrentMute bool, CurrentChannel uint16, SlaveList string, Status string, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        MultiType      string
        PlayMode       string
        Router         string
        Ssid           string
        SlaveMask      string
        CurrentVolume  string
        CurrentMute    string
        CurrentChannel string
        SlaveList      string
        Status         string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetControlDeviceInfo", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if MultiType, err = soap.UnmarshalUi2(response.MultiType); err != nil {
        return
    }
    if PlayMode, err = soap.UnmarshalUi2(response.PlayMode); err != nil {
        return
    }
    if Router, err = soap.UnmarshalString(response.Router); err != nil {
        return
    }
    if Ssid, err = soap.UnmarshalString(response.Ssid); err != nil {
        return
    }
    if SlaveMask, err = soap.UnmarshalUi2(response.SlaveMask); err != nil {
        return
    }
    if CurrentVolume, err = soap.UnmarshalUi2(response.CurrentVolume); err != nil {
        return
    }
    if CurrentMute, err = soap.UnmarshalBoolean(response.CurrentMute); err != nil {
        return
    }
    if CurrentChannel, err = soap.UnmarshalUi2(response.CurrentChannel); err != nil {
        return
    }
    if SlaveList, err = soap.UnmarshalString(response.SlaveList); err != nil {
        return
    }
    if Status, err = soap.UnmarshalString(response.Status); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetControlDeviceInfo is the legacy version of GetControlDeviceInfoCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) GetControlDeviceInfo(InstanceID uint32) (MultiType uint16, PlayMode uint16, Router string, Ssid string, SlaveMask uint16, CurrentVolume uint16, CurrentMute bool, CurrentChannel uint16, SlaveList string, Status string, err error) {
    return client.GetControlDeviceInfoCtx(context.Background(),
        InstanceID,
    )
}

//
// Arguments:
//
// * Channel: allowed values: Master, Single

func (client *RenderingControl1) GetEqualizerCtx(
    ctx context.Context,
    InstanceID uint32,
    Channel string,
) (CurrentEqualizer uint16, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
        Channel    string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.Channel, err = soap.MarshalString(Channel); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        CurrentEqualizer string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetEqualizer", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if CurrentEqualizer, err = soap.UnmarshalUi2(response.CurrentEqualizer); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetEqualizer is the legacy version of GetEqualizerCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) GetEqualizer(InstanceID uint32, Channel string) (CurrentEqualizer uint16, err error) {
    return client.GetEqualizerCtx(context.Background(),
        InstanceID,
        Channel,
    )
}

//
// Arguments:
//
// * Channel: allowed values: Master, Single

func (client *RenderingControl1) GetMuteCtx(
    ctx context.Context,
    InstanceID uint32,
    Channel string,
) (CurrentMute bool, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
        Channel    string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.Channel, err = soap.MarshalString(Channel); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        CurrentMute string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetMute", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if CurrentMute, err = soap.UnmarshalBoolean(response.CurrentMute); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetMute is the legacy version of GetMuteCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) GetMute(InstanceID uint32, Channel string) (CurrentMute bool, err error) {
    return client.GetMuteCtx(context.Background(),
        InstanceID,
        Channel,
    )
}

// Return values:
//
// * CurrentVolume: allowed value range: minimum=0, maximum=100, step=1
func (client *RenderingControl1) GetSimpleDeviceInfoCtx(
    ctx context.Context,
    InstanceID uint32,
) (MultiType uint16, SlaveMask uint16, PlayMode uint16, Name string, CurrentVolume uint16, CurrentChannel uint16, SlaveList string, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        MultiType      string
        SlaveMask      string
        PlayMode       string
        Name           string
        CurrentVolume  string
        CurrentChannel string
        SlaveList      string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetSimpleDeviceInfo", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if MultiType, err = soap.UnmarshalUi2(response.MultiType); err != nil {
        return
    }
    if SlaveMask, err = soap.UnmarshalUi2(response.SlaveMask); err != nil {
        return
    }
    if PlayMode, err = soap.UnmarshalUi2(response.PlayMode); err != nil {
        return
    }
    if Name, err = soap.UnmarshalString(response.Name); err != nil {
        return
    }
    if CurrentVolume, err = soap.UnmarshalUi2(response.CurrentVolume); err != nil {
        return
    }
    if CurrentChannel, err = soap.UnmarshalUi2(response.CurrentChannel); err != nil {
        return
    }
    if SlaveList, err = soap.UnmarshalString(response.SlaveList); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetSimpleDeviceInfo is the legacy version of GetSimpleDeviceInfoCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) GetSimpleDeviceInfo(InstanceID uint32) (MultiType uint16, SlaveMask uint16, PlayMode uint16, Name string, CurrentVolume uint16, CurrentChannel uint16, SlaveList string, err error) {
    return client.GetSimpleDeviceInfoCtx(context.Background(),
        InstanceID,
    )
}

//
// Arguments:
//
// * Channel: allowed values: Master, Single

// Return values:
//
// * CurrentVolume: allowed value range: minimum=0, maximum=100, step=1
func (client *RenderingControl1) GetVolumeCtx(
    ctx context.Context,
    InstanceID uint32,
    Channel string,
) (CurrentVolume uint16, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
        Channel    string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.Channel, err = soap.MarshalString(Channel); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        CurrentVolume string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "GetVolume", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if CurrentVolume, err = soap.UnmarshalUi2(response.CurrentVolume); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// GetVolume is the legacy version of GetVolumeCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) GetVolume(InstanceID uint32, Channel string) (CurrentVolume uint16, err error) {
    return client.GetVolumeCtx(context.Background(),
        InstanceID,
        Channel,
    )
}

func (client *RenderingControl1) ListPresetsCtx(
    ctx context.Context,
    InstanceID uint32,
) (CurrentPresetNameList string, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        CurrentPresetNameList string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "ListPresets", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if CurrentPresetNameList, err = soap.UnmarshalString(response.CurrentPresetNameList); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// ListPresets is the legacy version of ListPresetsCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) ListPresets(InstanceID uint32) (CurrentPresetNameList string, err error) {
    return client.ListPresetsCtx(context.Background(),
        InstanceID,
    )
}

func (client *RenderingControl1) MultiRoomJoinGroupCtx(
    ctx context.Context,
    InstanceID uint32,
    MasterInfo uint16,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
        MasterInfo string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.MasterInfo, err = soap.MarshalUi2(MasterInfo); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "MultiRoomJoinGroup", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// MultiRoomJoinGroup is the legacy version of MultiRoomJoinGroupCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) MultiRoomJoinGroup(InstanceID uint32, MasterInfo uint16) (err error) {
    return client.MultiRoomJoinGroupCtx(context.Background(),
        InstanceID,
        MasterInfo,
    )
}

func (client *RenderingControl1) MultiRoomLeaveGroupCtx(
    ctx context.Context,
    InstanceID uint32,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "MultiRoomLeaveGroup", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// MultiRoomLeaveGroup is the legacy version of MultiRoomLeaveGroupCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) MultiRoomLeaveGroup(InstanceID uint32) (err error) {
    return client.MultiRoomLeaveGroupCtx(context.Background(),
        InstanceID,
    )
}

//
// Arguments:
//
// * PresetName: allowed values: FactoryDefaults, InstallationDefaults

func (client *RenderingControl1) SelectPresetCtx(
    ctx context.Context,
    InstanceID uint32,
    PresetName string,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID string
        PresetName string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.PresetName, err = soap.MarshalString(PresetName); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SelectPreset", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SelectPreset is the legacy version of SelectPresetCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SelectPreset(InstanceID uint32, PresetName string) (err error) {
    return client.SelectPresetCtx(context.Background(),
        InstanceID,
        PresetName,
    )
}

func (client *RenderingControl1) SetAlarmQueueCtx(
    ctx context.Context,
    AlarmContext string,
) (err error) {
    // Request structure.
    request := &struct {
        AlarmContext string
    }{}
    // BEGIN Marshal arguments into request.

    if request.AlarmContext, err = soap.MarshalString(AlarmContext); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetAlarmQueue", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetAlarmQueue is the legacy version of SetAlarmQueueCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetAlarmQueue(AlarmContext string) (err error) {
    return client.SetAlarmQueueCtx(context.Background(),
        AlarmContext,
    )
}

//
// Arguments:
//
// * Channel: allowed values: Master, Single

func (client *RenderingControl1) SetChannelCtx(
    ctx context.Context,
    InstanceID uint32,
    Channel string,
    DesiredChannel uint16,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID     string
        Channel        string
        DesiredChannel string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.Channel, err = soap.MarshalString(Channel); err != nil {
        return
    }
    if request.DesiredChannel, err = soap.MarshalUi2(DesiredChannel); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetChannel", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetChannel is the legacy version of SetChannelCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetChannel(InstanceID uint32, Channel string, DesiredChannel uint16) (err error) {
    return client.SetChannelCtx(context.Background(),
        InstanceID,
        Channel,
        DesiredChannel,
    )
}

func (client *RenderingControl1) SetDeviceNameCtx(
    ctx context.Context,
    Name string,
) (err error) {
    // Request structure.
    request := &struct {
        Name string
    }{}
    // BEGIN Marshal arguments into request.

    if request.Name, err = soap.MarshalString(Name); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetDeviceName", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetDeviceName is the legacy version of SetDeviceNameCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetDeviceName(Name string) (err error) {
    return client.SetDeviceNameCtx(context.Background(),
        Name,
    )
}

//
// Arguments:
//
// * Channel: allowed values: Master, Single

func (client *RenderingControl1) SetEqualizerCtx(
    ctx context.Context,
    InstanceID uint32,
    Channel string,
    DesiredEqualizer uint16,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID       string
        Channel          string
        DesiredEqualizer string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.Channel, err = soap.MarshalString(Channel); err != nil {
        return
    }
    if request.DesiredEqualizer, err = soap.MarshalUi2(DesiredEqualizer); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetEqualizer", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetEqualizer is the legacy version of SetEqualizerCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetEqualizer(InstanceID uint32, Channel string, DesiredEqualizer uint16) (err error) {
    return client.SetEqualizerCtx(context.Background(),
        InstanceID,
        Channel,
        DesiredEqualizer,
    )
}

//
// Arguments:
//
// * Channel: allowed values: Master, Single

func (client *RenderingControl1) SetMuteCtx(
    ctx context.Context,
    InstanceID uint32,
    Channel string,
    DesiredMute bool,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID  string
        Channel     string
        DesiredMute string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.Channel, err = soap.MarshalString(Channel); err != nil {
        return
    }
    if request.DesiredMute, err = soap.MarshalBoolean(DesiredMute); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetMute", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetMute is the legacy version of SetMuteCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetMute(InstanceID uint32, Channel string, DesiredMute bool) (err error) {
    return client.SetMuteCtx(context.Background(),
        InstanceID,
        Channel,
        DesiredMute,
    )
}

func (client *RenderingControl1) SetStreamServicesCapabilityCtx(
    ctx context.Context,
    Command string,
) (err error) {
    // Request structure.
    request := &struct {
        Command string
    }{}
    // BEGIN Marshal arguments into request.

    if request.Command, err = soap.MarshalString(Command); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetStreamServicesCapability", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetStreamServicesCapability is the legacy version of SetStreamServicesCapabilityCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetStreamServicesCapability(Command string) (err error) {
    return client.SetStreamServicesCapabilityCtx(context.Background(),
        Command,
    )
}

//
// Arguments:
//
// * Channel: allowed values: Master, Single
//
// * DesiredVolume: allowed value range: minimum=0, maximum=100, step=1

func (client *RenderingControl1) SetVolumeCtx(
    ctx context.Context,
    InstanceID uint32,
    Channel string,
    DesiredVolume uint16,
) (err error) {
    // Request structure.
    request := &struct {
        InstanceID    string
        Channel       string
        DesiredVolume string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.Channel, err = soap.MarshalString(Channel); err != nil {
        return
    }
    if request.DesiredVolume, err = soap.MarshalUi2(DesiredVolume); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := interface{}(nil)

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "SetVolume", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    // END Unmarshal arguments from response.
    return
}

// SetVolume is the legacy version of SetVolumeCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) SetVolume(InstanceID uint32, Channel string, DesiredVolume uint16) (err error) {
    return client.SetVolumeCtx(context.Background(),
        InstanceID,
        Channel,
        DesiredVolume,
    )
}

func (client *RenderingControl1) StreamServicesCapabilityCtx(
    ctx context.Context,
    InstanceID uint32,
    AppVersion string,
) (StreamCapability string, err error) {
    // Request structure.
    request := &struct {
        InstanceID string
        AppVersion string
    }{}
    // BEGIN Marshal arguments into request.

    if request.InstanceID, err = soap.MarshalUi4(InstanceID); err != nil {
        return
    }
    if request.AppVersion, err = soap.MarshalString(AppVersion); err != nil {
        return
    }
    // END Marshal arguments into request.

    // Response structure.
    response := &struct {
        StreamCapability string
    }{}

    // Perform the SOAP call.
    if err = client.SOAPClient.PerformActionCtx(ctx, URN_RenderingControl_1, "StreamServicesCapability", request, response); err != nil {
        return
    }

    // BEGIN Unmarshal arguments from response.

    if StreamCapability, err = soap.UnmarshalString(response.StreamCapability); err != nil {
        return
    }
    // END Unmarshal arguments from response.
    return
}

// StreamServicesCapability is the legacy version of StreamServicesCapabilityCtx, but uses
// context.Background() as the context.
func (client *RenderingControl1) StreamServicesCapability(InstanceID uint32, AppVersion string) (StreamCapability string, err error) {
    return client.StreamServicesCapabilityCtx(context.Background(),
        InstanceID,
        AppVersion,
    )
}
