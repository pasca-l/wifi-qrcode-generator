package qrcode

import (
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
)

const imageSize int = 300
const black string = "#000000"

func DrawQRCode(w io.Writer, code QRCode) error {
	s := svg.New(w)
	s.Start(imageSize, imageSize)
	for y, row := range code.Pattern {
		for x, cell := range row {
			pixel := imageSize / len(code.Pattern)
			if cell {
				s.Square(x*pixel, y*pixel, pixel, fmt.Sprintf(`fill="%s"`, black))
			}
		}
	}
	s.End()

	return nil
}
