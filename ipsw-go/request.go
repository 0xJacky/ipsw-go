package ipsw_go

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetLatestFirmwareInfo(identifier string, firmwareType FirmwareType, lastTwoVer bool) (info []DeviceFirmware, err error) {

	url2 := fmt.Sprintf("%s/apple/firmwares/%s/latest?type=%s", "https://betahub.cn/api", identifier, firmwareType)

	u, _ := url.Parse(url2)

	if lastTwoVer {
		u.Query().Set("last_two_ver", "true")
	}

	resp, err := http.Get(u.String())

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
