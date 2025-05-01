package qrcode

import (
	"fmt"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
	"github.com/pasca-l/wifi-qrcode-generator/utils/math"
)

type QRCodeSpec struct {
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

func NewQRCodeSpec(src string, ecl ErrorCorrectionLevel) (QRCodeSpec, error) {
	mode := getEncodeMode(src)

	ver, err := getVersion(ecl, src)
	if err != nil {
		return QRCodeSpec{}, err
	}

	return QRCodeSpec{
		ecl:     ecl,
		mode:    mode,
		version: ver,
	}, nil
}

func EncodeSrc(src string, spec QRCodeSpec) (utils.Bytes, error) {
	msg := utils.Bits{}

	indBits, err := getIndicatorBits(spec.mode)
	if err != nil {
		return utils.Bytes{}, err
	}
	msg = append(msg, indBits...)

	srcBits, err := convertSrcToBits(spec.mode, src)
	if err != nil {
		return utils.Bytes{}, err
	}
	countIndicator, err := getLengthField(spec.version, spec.mode)
	if err != nil {
		return utils.Bytes{}, err
	}
	srcCountBits, err := getSrcCountBits(len(srcBits), countIndicator)
	if err != nil {
		return utils.Bytes{}, err
	}
	msg = append(msg, srcCountBits...)
	msg = append(msg, srcBits...)

	endBits, err := getTerminatorBits()
	if err != nil {
		return utils.Bytes{}, err
	}
	msg = append(msg, endBits...)

	capacity, err := getVersionCapacity(spec.version, spec.ecl)
	if err != nil {
		return utils.Bytes{}, err
	}
	msg = msg.AppendBitPadding()
	msg = msg.AppendBytePadding(capacity)

	msgBytes, err := msg.ToBytes()
	if err != nil {
		return utils.Bytes{}, err
	}

	return msgBytes, nil
}

func ApplyErrorCorrection(msg utils.Bytes, spec QRCodeSpec) (utils.Bytes, error) {
	blocks, exists := blockStructure[spec.version][spec.ecl]
	if !exists {
		return utils.Bytes{}, fmt.Errorf("unexpected block structure for version: %d, ecl: %s", spec.version, spec.ecl)
	}

	rs := math.ReedSolomon{}
	result := make(utils.Bytes, 0)
	for _, block := range blocks {
		subMsg := msg[:block.blockLength]
		msg = msg[block.blockLength:] // update message to its remaining part

		encoded, err := rs.Encode(subMsg, block.blockLength-block.codewordLength)
		if err != nil {
			return utils.Bytes{}, err
		}
		result = append(result, encoded...)
	}

	return result, nil
}

func (spec QRCodeSpec) GenerateQRCode(src string) (QRCode, error) {
	msg, err := EncodeSrc(src, spec)
	if err != nil {
		return QRCode{}, err
	}
	encoded, err := ApplyErrorCorrection(msg, spec)
	if err != nil {
		return QRCode{}, err
	}
	qrcode, err := NewQRCode(encoded, spec)
	if err != nil {
		return QRCode{}, err
	}

	return qrcode, nil
}
