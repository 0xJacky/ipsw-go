package ipsw_go

import (
	"time"
)

type FirmwareType string

const (
	ReleaseFirmware FirmwareType = "release"
	BetaFirmware                 = "beta"
)

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
	Version  string    `json:"version"`
	BuildID  string    `json:"build_id"`
	Size     int64     `json:"size"`
	UploadAt time.Time `json:"upload_at"`
	Url      string    `json:"url"`
	CowUrl   string    `json:"cow_url"`
	Filename string    `json:"filename"`
	Type     int       `json:"type"`
	Signing  int       `json:"signing"`
}
