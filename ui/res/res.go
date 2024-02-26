package res

import (
    _ "embed"
)

//go:embed wiim_icon.png
var WiimIcon []byte

//go:embed style.css
var StyleCss string

//go:embed default_album.png
var DefaultAlbum []byte

//go:embed default_track.png
var DefaultTrack []byte

//go:embed shuffle_off.png
var ShuffleOff []byte

//go:embed shuffle_on.png
var ShuffleOn []byte

//go:embed previous.png
var Previous []byte

//go:embed play.png
var Play []byte

//go:embed next.png
var Next []byte

//go:embed loop_off.png
var LoopOff []byte

//go:embed loop_all.png
var LoopAll []byte

//go:embed loop_one.png
var LoopOne []byte

//go:embed pause.png
var Pause []byte

//go:embed volume_low.png
var VolumeLow []byte

//go:embed volume_high.png
var VolumeHigh []byte

//go:embed mute_off.png
var MuteOff []byte

//go:embed mute_on.png
var MuteOn []byte

//go:embed settings.png
var Settings []byte

//go:embed refresh.png
var Refresh []byte

//go:embed about.png
var About []byte

//go:embed wiim_about.png
var WiimAbout []byte

//go:embed eq_on.png
var EqOn []byte

//go:embed eq_off.png
var EqOff []byte

//go:embed device_controls.png
var DeviceControls []byte
