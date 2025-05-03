package qrcode

import (
	"fmt"

	"github.com/pasca-l/wifi-qrcode-generator/utils"
)

type QRCode struct {
	Pattern Pattern
}

type Pattern [][]bool
type Coordinate struct {
	X int
	Y int
}

type Mask int

func NewPattern(size int) Pattern {
	pattern := make(Pattern, size)
	for i := range pattern {
		pattern[i] = make([]bool, size)
	}
	return pattern
}

func (p Pattern) DrawPattern(pat Pattern, coord Coordinate) (Pattern, error) {
	// check if pattern fits
	if coord.Y+len(pat) > len(p) || coord.X+len(pat) > len(p) {
		return nil, fmt.Errorf("invalid pattern to draw with size: %d at coordinate: %+v", len(pat), coord)
	}

	for y := 0; y < len(pat); y++ {
		for x := 0; x < len(pat); x++ {
			if pat[y][x] {
				// draws pattern from upper left corner
				p[coord.Y+y][coord.X+x] = true
			}
		}
	}
	return p, nil
}

func NewQRCode(msg utils.Bytes, spec QRCodeSpec) (QRCode, error) {
	pattern, err := GeneratePattern(msg, spec)
	if err != nil {
		return QRCode{}, err
	}
	return QRCode{
		Pattern: pattern,
	}, nil
}

func GeneratePattern(msg utils.Bytes, spec QRCodeSpec) (Pattern, error) {
	dim := calcSizeFromVersion(spec.version)
	pat := NewPattern(dim)

	_, err := pat.addFunctionalPattern(spec.version)
	if err != nil {
		return nil, err
	}

	return Pattern{}, nil
}

func calcSizeFromVersion(ver Version) int {
	// calculated size dimension from version
	return 21 + 4*(int(ver)-1)
}

func createFinderPattern() Pattern {
	pattern := NewPattern(7)
	for row := range 7 {
		for col := range 7 {
			if row == 0 || row == 6 || col == 0 || col == 6 || (row >= 2 && row <= 4 && col >= 2 && col <= 4) {
				pattern[row][col] = true
			}
		}
	}
	return pattern
}

func createAlignmentPattern() Pattern {
	pattern := NewPattern(5)
	for row := range 5 {
		for col := range 5 {
			if row == 0 || row == 4 || col == 0 || col == 4 || (row == 2 && col == 2) {
				pattern[row][col] = true
			}
		}
	}
	return pattern
}

var alignmentPatternCenterPosition = map[Version][]int{
	1:  {},
	2:  {18},
	3:  {22},
	4:  {26},
	5:  {30},
	6:  {34},
	7:  {6, 22, 38},
	8:  {6, 24, 42},
	9:  {6, 26, 46},
	10: {6, 28, 50},
	11: {6, 30, 54},
	12: {6, 32, 58},
	13: {6, 34, 62},
	14: {6, 26, 46, 66},
	15: {6, 26, 48, 70},
	16: {6, 26, 50, 74},
	17: {6, 30, 54, 78},
	18: {6, 30, 56, 82},
	19: {6, 30, 58, 86},
	20: {6, 34, 62, 90},
	21: {6, 28, 50, 72, 94},
	22: {6, 26, 50, 74, 98},
	23: {6, 30, 54, 78, 102},
	24: {6, 28, 54, 80, 106},
	25: {6, 32, 58, 84, 110},
	26: {6, 30, 58, 86, 114},
	27: {6, 34, 62, 90, 118},
	28: {6, 26, 50, 74, 98, 122},
	29: {6, 30, 54, 78, 102, 126},
	30: {6, 26, 52, 78, 104, 130},
	31: {6, 30, 56, 82, 108, 134},
	32: {6, 34, 60, 86, 112, 138},
	33: {6, 30, 58, 86, 114, 142},
	34: {6, 34, 62, 90, 118, 146},
	35: {6, 30, 54, 78, 102, 126, 150},
	36: {6, 24, 50, 76, 102, 128, 154},
	37: {6, 28, 54, 80, 106, 132, 158},
	38: {6, 32, 58, 84, 110, 136, 162},
	39: {6, 26, 54, 82, 110, 138, 166},
	40: {6, 30, 58, 86, 114, 142, 170},
}

func calcAlignmentPatternCoords(ver Version) ([]Coordinate, error) {
	centerPositions, exists := alignmentPatternCenterPosition[ver]
	if !exists {
		return nil, fmt.Errorf("alignment pattern center positions for version: %d does not exist", ver)
	}

	size := calcSizeFromVersion(ver)
	coords := make([]Coordinate, 0, len(centerPositions))
	for _, row := range centerPositions {
		for _, col := range centerPositions {
			// skip positions overlapping finder patterns
			if (row == 6 && col == 6) || (row == 6 && col == size-7) || (row == size-7 && col == 6) {
				continue
			}
			// displace coordinates representing upper left corner of the pattern
			coords = append(coords, Coordinate{X: col - 2, Y: row - 2})
		}
	}
	return coords, nil
}

func (p Pattern) addFunctionalPattern(ver Version) (Pattern, error) {
	// adds 7x7 finder patterns to corners
	finderPattern := createFinderPattern()
	p, err := p.DrawPattern(finderPattern, Coordinate{0, 0})
	if err != nil {
		return nil, err
	}
	p, err = p.DrawPattern(finderPattern, Coordinate{0, len(p) - 7})
	if err != nil {
		return nil, err
	}
	p, err = p.DrawPattern(finderPattern, Coordinate{len(p) - 7, 0})
	if err != nil {
		return nil, err
	}

	// adds timing patterns
	for i := 8; i < len(p)-8; i++ {
		p[6][i] = i%2 == 0 // horizontal pattern
		p[i][6] = i%2 == 0 // vertical pattern
	}

	// adds 5x5 alignment patterns
	alignmentPattern := createAlignmentPattern()
	coords, err := calcAlignmentPatternCoords(ver)
	if err != nil {
		return nil, err
	}
	for _, coord := range coords {
		p, err = p.DrawPattern(alignmentPattern, coord)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (p Pattern) applyData(msg utils.Bytes) Pattern {
	return Pattern{}
}

func (p Pattern) findBestMask() int {
	return 0
}

func (p Pattern) applyMask(mask int) Pattern {
	return Pattern{}
}
