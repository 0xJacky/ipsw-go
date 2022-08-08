package ipsw_go

import (
	"github.com/cavaliergopher/grab/v3"
	"log"
	"os"
)

func buildReqs(identifier []string, firmwareType FirmwareType) (reqs []*grab.Request) {
	// create multiple download requests
	for _, v := range identifier {
		info, err := GetLatestFirmwareInfo(v, firmwareType)

		if err != nil {
			log.Println("[Error]", "GetLatestFirmwareInfo", v)
			continue
		}

		if info.Firmware.Url == "" {
			log.Println("[Error]", "GetLatestFirmwareInfo info.Firmware.Url is empty")
			continue
		}

		if info.Firmware.Version == "" {
			log.Println("[Error]", "GetLatestFirmwareInfo info.Firmware.Version is empty")
			continue
		}

		dir := "downloads/" + info.Firmware.Version

		if _, err = os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0777)

			if err != nil {
				log.Println("[Error]", "os.MkdirAll", dir, err)
				continue
			}
			err = os.Chmod(dir, 0777)
			if err != nil {
				log.Println("[Error]", "os.Chmod", dir, err)
				continue
			}
		}

		req, err := grab.NewRequest(dir+"/.", info.Firmware.Url)

		if err != nil {
			log.Println("[Error]", "GetLatestFirmwareInfo", v, err)
			continue
		}

		reqs = append(reqs, req)
	}

	return
}

func Worker(workers int, identifier []string) {
	reqs := buildReqs(identifier, ReleaseFirmware)

	reqs = append(reqs, buildReqs(identifier, BetaFirmware)...)

	client := grab.NewClient()

	respch := doBatch(client, workers, reqs)

	// check each response
	for resp := range respch {
		if err := resp.Err(); err != nil {
			log.Println("[Error]", "Do Batch", err)
		}
		log.Printf("[Info] Downloaded %s", resp.Filename)
	}

}
