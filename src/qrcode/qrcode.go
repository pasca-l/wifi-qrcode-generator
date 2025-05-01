package qrcode

import "github.com/pasca-l/wifi-qrcode-generator/utils"

type QRCode struct {
	Pattern QRCodePattern
}
type QRCodePattern [][]bool

type Mask int

func NewQRCode(msg utils.Bytes, spec QRCodeSpec) (QRCode, error) {
	return QRCode{
		Pattern: GeneratePattern(msg, spec),
	}, nil
}

func NewPattern(ver Version) QRCodePattern {
	// calculated size dimension from version
	dim := 21 + 4*(int(ver)-1)

	pat := make(QRCodePattern, dim)
	for i := range pat {
		pat[i] = make([]bool, dim)
	}
	return pat
}

func GeneratePattern(msg utils.Bytes, spec QRCodeSpec) QRCodePattern {
	_ = NewPattern(spec.version)

	return QRCodePattern{}
}

func (p QRCodePattern) addFunctionalPattern() QRCodePattern {
	return QRCodePattern{}
}

func (p QRCodePattern) applyData(msg utils.Bytes) QRCodePattern {
	return QRCodePattern{}
}

func (p QRCodePattern) findBestMask() int {
	return 0
}

func (p QRCodePattern) applyMask(mask int) QRCodePattern {
	return QRCodePattern{}
}
