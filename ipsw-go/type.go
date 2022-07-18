package ipsw_go

import (
	"github.com/spf13/cast"
	"time"
)

type FirmwareType int

const (
	ReleaseFirmware FirmwareType = iota + 1
	BetaFirmware
)

func (t FirmwareType) String() string {
	return cast.ToString(t)
}

type DeviceFirmware struct {
	DeviceID   uint     `json:"device_id"`
	Device     Device   `json:"device"`
	FirmwareID uint     `json:"firmware_id"`
	Firmware   Firmware `json:"firmware"`
}

type Device struct {
	Name        string     `json:"name"`
	Identifier  string     `json:"identifier"`
	ReleaseDate *time.Time `json:"release_date"`
}

type Firmware struct {
	Version  string       `json:"version"`
	BuildID  string       `json:"build_id"`
	Size     int64        `json:"size"`
	UploadAt time.Time    `json:"upload_at"`
	Url      string       `json:"url"`
	CowUrl   string       `json:"cow_url"`
	Filename string       `json:"filename"`
	Type     FirmwareType `json:"type"`
	Signing  int          `json:"signing"`
}
