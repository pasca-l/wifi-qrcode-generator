package qrcode

import (
	"github.com/pasca-l/wifi-qrcode-generator/utils"
)

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

func (s QRCodeSpec) Encode() (utils.Bits, error) {
	codeword := utils.Bits{}

	indBits, err := getIndicatorBits(s.mode)
	if err != nil {
		return utils.Bits{}, err
	}
	codeword = append(codeword, indBits...)

	countIndicator, err := getCharacterCountIndicator(s.version, s.mode)
	if err != nil {
		return utils.Bits{}, err
	}
	srcCountBits, err := getSrcCountBits(s.src, countIndicator)
	if err != nil {
		return utils.Bits{}, err
	}
	codeword = append(codeword, srcCountBits...)

	srcBits, err := convertSrcToBits(s.mode, s.src)
	if err != nil {
		return utils.Bits{}, err
	}
	codeword = append(codeword, srcBits...)

	endBits, err := getTerminatorBits()
	if err != nil {
		return utils.Bits{}, err
	}
	codeword = append(codeword, endBits...)

	capacity, err := getVersionCapacity(s.version, s.ecl)
	if err != nil {
		return utils.Bits{}, err
	}
	codeword = codeword.AppendBitPadding()
	codeword = codeword.AppendBytePadding(capacity)

	return codeword, nil
}

func (s QRCodeSpec) GenerateQRCode() (QRCode, error) {
	return QRCode{}, nil
}
