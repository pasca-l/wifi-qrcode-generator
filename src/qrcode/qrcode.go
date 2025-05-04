package qrcode

import (
	"fmt"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
	"github.com/pasca-l/wifi-qrcode-generator/utils/math"
)

type QRCode struct {
	Pattern Pattern
}

type QRCodeSpec struct {
	mode    EncodeMode
	version Version
	ecl     ErrorCorrectionLevel
}

func NewQRCode(src string, spec QRCodeSpec) (QRCode, error) {
	msg, err := spec.EncodeSrc(src)
	if err != nil {
		return QRCode{}, err
	}
	encoded, err := spec.ApplyErrorCorrection(msg)
	if err != nil {
		return QRCode{}, err
	}

	pattern, err := GeneratePattern(encoded, spec)
	if err != nil {
		return QRCode{}, err
	}

	return QRCode{
		Pattern: pattern,
	}, nil
}

func NewQRCodeSpec(src string, ecl ErrorCorrectionLevel) (QRCodeSpec, error) {
	mode := getEncodeMode(src)

	ver, err := getVersion(ecl, src)
	if err != nil {
		return QRCodeSpec{}, err
	}

	return QRCodeSpec{
		mode:    mode,
		version: ver,
		ecl:     ecl,
	}, nil
}

func (spec QRCodeSpec) EncodeSrc(src string) (utils.Bytes, error) {
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

func (spec QRCodeSpec) ApplyErrorCorrection(msg utils.Bytes) (utils.Bytes, error) {
	blocks, exists := blockStructure[spec.version][spec.ecl]
	if !exists {
		return utils.Bytes{}, fmt.Errorf("unexpected block structure for version: %d, ecl: %s", spec.version, spec.ecl.ToString())
	}

	rs := math.ReedSolomon{}
	result := make(utils.Bytes, 0)
	for _, block := range blocks {
		subMsg := msg
		if len(msg) > block.blockLength {
			subMsg = msg[:block.blockLength]
			msg = msg[block.blockLength:] // update message to its remaining part
		}

		encoded, err := rs.Encode(subMsg, block.blockLength-block.codewordLength)
		if err != nil {
			return utils.Bytes{}, err
		}
		result = append(result, encoded...)
	}

	return result, nil
}
