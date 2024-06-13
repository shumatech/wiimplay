package wiimdev

import (
    "time"
)

type FakeDevice struct {
}

func (device *FakeDevice) GetInfo() (*DeviceInfo, error) { return nil, nil }
func (device *FakeDevice) Play() error { return nil }
func (device *FakeDevice) Pause() error { return nil }
func (device *FakeDevice) Previous() error { return nil }
func (device *FakeDevice) Next() error { return nil }
func (device *FakeDevice) LoopMode(loopMode LoopMode) error { return nil }
func (device *FakeDevice) Volume(level int) error { return nil }
func (device *FakeDevice) Seek(position time.Duration) error { return nil }
func (device *FakeDevice) Mute(mute bool) error { return nil }
func (device *FakeDevice) GetPlayLists() ([]PlayList, string, error) { return nil, "", nil }
func (device *FakeDevice) GetTracks(id string) ([]Track, int, error) { return nil, 0, nil }
func (device *FakeDevice) PlayTrack(track int) error { return nil }
func (device *FakeDevice) SetPlayList(id string) error { return nil }
func (device *FakeDevice) GetDeviceDiscovery() *DeviceDiscovery { return nil }
func (device *FakeDevice) SetEqualizer(setting string) error { return nil }
func (device *FakeDevice) GetEqualizer() (string, error) { return "", nil }
func (device *FakeDevice) GetEqualizerList() ([]string, error) { return nil, nil }
func (device *FakeDevice) GetAudioOutput() (int, error) { return 0, nil }
func (device *FakeDevice) GetAudioOutputList() ([]string, error) { return nil, nil }
func (device *FakeDevice) SetAudioOutput(output int) error { return nil }
func (device *FakeDevice) GetBalance() (float64, error) { return 0.0, nil }
func (device *FakeDevice) SetBalance(level float64) error { return nil }
func (device *FakeDevice) GetFadeEffects() (bool, error) { return false, nil }
func (device *FakeDevice) SetFadeEffects(on bool) error { return nil }
func (device *FakeDevice) GetFixedVolume() (bool, error) { return false, nil }
func (device *FakeDevice) SetFixedVolume(on bool) error { return nil }
func (device *FakeDevice) GetAudioInput() (int, error) { return 0, nil }
func (device *FakeDevice) GetAudioInputList() ([]string, error) { return nil, nil }
func (device *FakeDevice) SetAudioInput(input int) error { return nil }

func NewFakeDevice() *FakeDevice {
    return &FakeDevice{}
}
