package qrcode

import (
	"fmt"
	"net/url"
)

type WifiSpec struct {
	ssid       string
	password   string
	encryption Encryption
}

type Encryption string

const (
	NoPass Encryption = "nopass"
	WEP    Encryption = "WEP"
	WPA    Encryption = "WPA"
)

func NewWifiSpec(params url.Values) (WifiSpec, error) {
	enc, err := toEncryption(params.Get("encryption"))
	if err != nil {
		return WifiSpec{}, err
	}

	return WifiSpec{
		ssid:       params.Get("ssid"),
		password:   params.Get("password"),
		encryption: enc,
	}, nil
}

func toEncryption(param string) (Encryption, error) {
	switch param {
	case string(NoPass):
		return NoPass, nil
	case string(WEP):
		return WEP, nil
	case string(WPA):
		return WPA, nil
	default:
		return Encryption(""), fmt.Errorf(
			"cannot convert value '%s' to type Encryption", param,
		)
	}
}

func (s WifiSpec) Encode() string {
	// referenced: https://github.com/zxing/zxing/wiki/Barcode-Contents#wi-fi-network-config-android-ios-11
	return fmt.Sprintf(
		"WIFI:T:%s;S:\"%s\";P:\"%s\";;",
		s.encryption, s.ssid, s.password,
	)
}
