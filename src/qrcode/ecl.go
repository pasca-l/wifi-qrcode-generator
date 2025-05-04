package qrcode

type ErrorCorrectionLevel int

const (
	L ErrorCorrectionLevel = 1 // Low, 7% recoverable
	M ErrorCorrectionLevel = 0 // Medium, 15% recoverable
	Q ErrorCorrectionLevel = 3 // Quartile, 25% recoverable
	H ErrorCorrectionLevel = 2 // High, 30% recoverable
)

func (ecl ErrorCorrectionLevel) ToString() string {
	switch ecl {
	case L:
		return "L"
	case M:
		return "M"
	case Q:
		return "Q"
	case H:
		return "H"
	}
	return ""
}
