package ipsw_go

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func GetLatestFirmwareInfo(identifier string) (info DeviceFirmware, err error) {

	url := fmt.Sprintf("%s/apple/firmwares/%s/latest", "https://betahub.cn/api", identifier)

	resp, err := http.Get(url)

	if err != nil {
		err = errors.Wrap(err, "error get latest firmware info http.Get")
		return
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		err = errors.Wrap(err, "error get latest firmware info ioutil.ReadAll")
		return
	}

	err = json.Unmarshal(bytes, &info)

	if err != nil {
		err = errors.Wrap(err, "error get latest firmware info json.Unmarshal")
		return
	}

	return
}
