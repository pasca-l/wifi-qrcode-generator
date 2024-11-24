package qrcode

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

func (s QRCodeSpec) GenerateQRCode() (QRCode, error) {
	return QRCode{}, nil
}
