package mpris

import (
    "regexp"
    "time"

    "github.com/godbus/dbus/v5"
    mpris_events "github.com/quarckster/go-mpris-server/pkg/events"
    mpris_server "github.com/quarckster/go-mpris-server/pkg/server"
    mpris_types "github.com/quarckster/go-mpris-server/pkg/types"
)

type Player interface {
    Next()
    Previous()
    Pause()
    Play()
    PlaybackStatus() PlaybackStatus
    GetVolume() int
    SetVolume(level int)
    GetMetadata() *Metadata
    GetPosition() time.Duration
    SetPosition(position time.Duration)
    Seek(offset time.Duration)
    GetLoopStatus() LoopStatus
    SetLoopStatus(LoopStatus)
    GetShuffle() bool
    SetShuffle(bool)
}

type Metadata struct {
    Id string
    Title string
    Artist string
    Album string
    AlbumArt string
    Length time.Duration
}

type PlaybackStatus string
const (
    PlaybackStatusPlaying PlaybackStatus = "Playing"
    PlaybackStatusPaused PlaybackStatus = "Paused"
    PlaybackStatusStopped PlaybackStatus = "Stopped"
)

type LoopStatus string
const (
    LoopStatusNone LoopStatus = "None"
    LoopStatusTrack LoopStatus = "Track"
    LoopStatusPlaylist LoopStatus = "Playlist"
)

type Server struct {
    mprisServer *mpris_server.Server
    mprisEvents *mpris_events.EventHandler
    player Player
    name string
    listening bool
}

var sanitizeRegExp = regexp.MustCompile(`[^a-zA-Z0-9]+`)

///////////////////////////////////////////////////////////////////////////////
// OrgMprisMediaMprisServer2Adapter interface
///////////////////////////////////////////////////////////////////////////////
func (server *Server) Raise() error {
    return nil
}
func (server *Server) Quit() error {
    return nil
}
func (server *Server) CanQuit() (bool, error) {
    return false, nil
}
func (server *Server) CanRaise() (bool, error) {
    return false, nil
}
func (server *Server) HasTrackList() (bool, error) {
    return false, nil
}
func (server *Server) Identity() (string, error) {
    return server.name, nil
}
func (server *Server) SupportedUriSchemes() ([]string, error) {
    return nil, nil
}
func (server *Server) SupportedMimeTypes() ([]string, error) {
    return nil, nil
}

///////////////////////////////////////////////////////////////////////////////
// OrgMprisMediaMprisServer2MprisServerAdapter interface
///////////////////////////////////////////////////////////////////////////////
func (server *Server) Next() error {
    server.player.Next()
    return nil
}
func (server *Server) Previous() error {
    server.player.Previous()
    return nil
}
func (server *Server) Pause() error {
    server.player.Pause()
    return nil
}
func (server *Server) PlayPause() error {
    if server.player.PlaybackStatus() == PlaybackStatusPlaying {
        server.player.Pause()
    } else {
        server.player.Play()
    }
    return nil
}
func (server *Server) Stop() error {
    return nil
}
func (server *Server) Play() error {
    server.player.Play()
    return nil
}
func (server *Server) Seek(offset mpris_types.Microseconds) error {
    server.player.Seek(time.Duration(offset) * time.Microsecond)
    return nil
}
func (server *Server) SetPosition(trackId string, position mpris_types.Microseconds) error {
    server.player.SetPosition(time.Duration(position) * time.Microsecond)
    return nil
}
func (server *Server) OpenUri(uri string) error {
    return nil
}
func (server *Server) PlaybackStatus() (mpris_types.PlaybackStatus, error) {
    status := server.player.PlaybackStatus()
    return mpris_types.PlaybackStatus(status), nil
}
func (server *Server) Rate() (float64, error) {
    return 1.0, nil
}
func (server *Server) SetRate(float64) error {
    return nil
}
func (server *Server) Metadata() (mpris_types.Metadata, error) {
    md := server.player.GetMetadata()
    mdId := sanitizeRegExp.ReplaceAllString(md.Id, "")
    if mdId == "" {
        mdId = "unknown"
    }
    id := dbus.ObjectPath("/com/shumatech/" + server.name + "/" + mdId)
    return mpris_types.Metadata{
        TrackId: id,
        Title: md.Title,
        Artist: []string{md.Artist},
        Album: md.Album,
        ArtUrl: md.AlbumArt,
        Length: mpris_types.Microseconds(md.Length.Microseconds()),
    }, nil
}
func (server *Server) Volume() (float64, error) {
    return float64(server.player.GetVolume()) / 100.0, nil
}
func (server *Server) SetVolume(level float64) error {
    server.player.SetVolume(int(level * 100.0 + 0.5))
    return nil
}
func (server *Server) Position() (int64, error) {
    return server.player.GetPosition().Microseconds(), nil
}
func (server *Server) MinimumRate() (float64, error) {
    return 1.0, nil
}
func (server *Server) MaximumRate() (float64, error) {
    return 1.0, nil
}
func (server *Server) CanGoNext() (bool, error) {
    return true, nil
}
func (server *Server) CanGoPrevious() (bool, error) {
    return true, nil
}
func (server *Server) CanPlay() (bool, error) {
    return true, nil
}
func (server *Server) CanPause() (bool, error) {
    return true, nil
}
func (server *Server) CanSeek() (bool, error) {
    return true, nil
}
func (server *Server) CanControl() (bool, error) {
    return true, nil
}
func (server *Server) LoopStatus() (mpris_types.LoopStatus, error) {
    status := server.player.GetLoopStatus()
    return mpris_types.LoopStatus(status), nil
}
func (server *Server) SetLoopStatus(status mpris_types.LoopStatus) error {
    server.player.SetLoopStatus(LoopStatus(status))
    return nil
}
func (server *Server) Shuffle() (bool, error) {
    return server.player.GetShuffle(), nil
}
func (server *Server) SetShuffle(shuffle bool) error {
    server.player.SetShuffle(shuffle)
    return nil
}

func (server *Server) Listen(enable bool) {
    if !server.listening && enable {
        go server.mprisServer.Listen()
        server.listening = true
    } else if server.listening && !enable {
        server.mprisServer.Stop()
        server.listening = false
    }
}

func (server *Server) StatusChanged() {
    server.mprisEvents.Player.OnPlayPause()
}

func (server *Server) MetadataChanged() {
    server.mprisEvents.Player.OnTitle()
}

func NewMprisServer(name string, player Player) *Server {
    server := &Server{
        name: name,
        player: player,
        listening: false,
    }
    server.mprisServer = mpris_server.NewServer(name, server, server)
    server.mprisEvents = mpris_events.NewEventHandler(server.mprisServer)

    return server
}
