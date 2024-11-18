package qrcode

import (
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
)

const black string = "#000000"

func DrawQRCode(w io.Writer, code QRCode) error {
	s := svg.New(w)
	s.Start(300, 300)
	s.Square(0, 0, 300, fmt.Sprintf(`fill="%s"`, black))
	s.End()

	return nil
}
