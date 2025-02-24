package qrcode

import "github.com/pasca-l/wifi-qrcode-generator/utils"

type QRCode struct {
	pattern [][]bool
}

type QRCodeSpec struct {
	src     string
	ecl     ErrorCorrectionLevel
	mode    EncodeMode
	version Version
}

type ErrorCorrectionLevel string

const (
	L ErrorCorrectionLevel = "L" // Low, 7% recoverable
	M ErrorCorrectionLevel = "M" // Medium, 15% recoverable
	Q ErrorCorrectionLevel = "Q" // Quartile, 25% recoverable
	H ErrorCorrectionLevel = "H" // High, 30% recoverable
)

func NewQRCodeSpec(src string) (QRCodeSpec, error) {
	ecl := L
	mode := getEncodeMode(src)

	ver, err := getVersion(ecl, src)
	if err != nil {
		return QRCodeSpec{}, err
	}

	return QRCodeSpec{
		src:     src,
		ecl:     ecl,
		mode:    mode,
		version: ver,
	}, nil
}

func (s QRCodeSpec) EncodeSrc() (utils.Bits, error) {
	msg := utils.Bits{}

	indBits, err := getIndicatorBits(s.mode)
	if err != nil {
		return utils.Bits{}, err
	}
	msg = append(msg, indBits...)

	srcBits, err := convertSrcToBits(s.mode, s.src)
	if err != nil {
		return utils.Bits{}, err
	}
	countIndicator, err := getLengthField(s.version, s.mode)
	if err != nil {
		return utils.Bits{}, err
	}
	srcCountBits, err := getSrcCountBits(len(srcBits), countIndicator)
	if err != nil {
		return utils.Bits{}, err
	}
	msg = append(msg, srcCountBits...)
	msg = append(msg, srcBits...)

	endBits, err := getTerminatorBits()
	if err != nil {
		return utils.Bits{}, err
	}
	msg = append(msg, endBits...)

	capacity, err := getVersionCapacity(s.version, s.ecl)
	if err != nil {
		return utils.Bits{}, err
	}
	msg = msg.AppendBitPadding()
	msg = msg.AppendBytePadding(capacity)

	return msg, nil
}

func (s QRCodeSpec) GenerateQRCode() (QRCode, error) {
	return QRCode{}, nil
}
