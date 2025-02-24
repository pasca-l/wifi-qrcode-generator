package qrcode

import "fmt"

// codeblock is used to divide the data to not exceed 255 codewords in length,
// which is due to the restriction of using GF(2^8)
type Block struct {
	blockLength    int // total block length
	codewordLength int // codeword (data) length
}
type Blocks []Block

type BlockSpec struct {
	count int // count on repetition of blocks
	block Block
}

func NewBlockSpec(count int, blockLength int, codewordLength int) (BlockSpec, error) {
	if count < 1 || blockLength < 1 || codewordLength < 1 {
		return BlockSpec{}, fmt.Errorf("all given arguments must be larger than 0")
	}

	return BlockSpec{
		count: count,
		block: Block{
			blockLength:    blockLength,
			codewordLength: codewordLength,
		},
	}, nil
}

// generate block specs for constant specs
func genBlkSpec(count int, blockLength int, codewordLength int) BlockSpec {
	spec, _ := NewBlockSpec(count, blockLength, codewordLength)
	return spec
}

func (bs BlockSpec) toBlocks() Blocks {
	var blocks Blocks
	for _ = range bs.count {
		blocks = append(blocks, bs.block)
	}
	return blocks
}

func genBlks(spec []BlockSpec) Blocks {
	var blocks Blocks
	for _, s := range spec {
		blocks = append(blocks, s.toBlocks()...)
	}
	return blocks
}

// block structure used to encode with reed-solomon,
// where block values are represented in bytes
var blockStructure = map[Version]map[ErrorCorrectionLevel]Blocks{
	1: {
		L: genBlks([]BlockSpec{genBlkSpec(1, 26, 19)}),
		M: genBlks([]BlockSpec{genBlkSpec(1, 26, 16)}),
		Q: genBlks([]BlockSpec{genBlkSpec(1, 26, 13)}),
		H: genBlks([]BlockSpec{genBlkSpec(1, 26, 9)}),
	},
	2: {
		L: genBlks([]BlockSpec{genBlkSpec(1, 44, 34)}),
		M: genBlks([]BlockSpec{genBlkSpec(1, 44, 28)}),
		Q: genBlks([]BlockSpec{genBlkSpec(1, 44, 22)}),
		H: genBlks([]BlockSpec{genBlkSpec(1, 44, 16)}),
	},
	3: {
		L: genBlks([]BlockSpec{genBlkSpec(1, 70, 55)}),
		M: genBlks([]BlockSpec{genBlkSpec(1, 70, 44)}),
		Q: genBlks([]BlockSpec{genBlkSpec(2, 35, 17)}),
		H: genBlks([]BlockSpec{genBlkSpec(2, 35, 13)}),
	},
	4: {
		L: genBlks([]BlockSpec{genBlkSpec(1, 100, 80)}),
		M: genBlks([]BlockSpec{genBlkSpec(2, 50, 32)}),
		Q: genBlks([]BlockSpec{genBlkSpec(2, 50, 24)}),
		H: genBlks([]BlockSpec{genBlkSpec(4, 25, 9)}),
	},
	5: {
		L: genBlks([]BlockSpec{genBlkSpec(1, 134, 108)}),
		M: genBlks([]BlockSpec{genBlkSpec(2, 67, 43)}),
		Q: genBlks([]BlockSpec{genBlkSpec(2, 33, 15), genBlkSpec(2, 34, 16)}),
		H: genBlks([]BlockSpec{genBlkSpec(2, 33, 11), genBlkSpec(2, 34, 12)}),
	},
	6: {
		L: genBlks([]BlockSpec{genBlkSpec(2, 86, 68)}),
		M: genBlks([]BlockSpec{genBlkSpec(4, 43, 27)}),
		Q: genBlks([]BlockSpec{genBlkSpec(4, 43, 19)}),
		H: genBlks([]BlockSpec{genBlkSpec(4, 43, 15)}),
	},
	7: {
		L: genBlks([]BlockSpec{genBlkSpec(2, 98, 78)}),
		M: genBlks([]BlockSpec{genBlkSpec(4, 49, 31)}),
		Q: genBlks([]BlockSpec{genBlkSpec(2, 32, 14), genBlkSpec(4, 33, 15)}),
		H: genBlks([]BlockSpec{genBlkSpec(4, 39, 13), genBlkSpec(1, 40, 14)}),
	},
	8: {
		L: genBlks([]BlockSpec{genBlkSpec(2, 121, 97)}),
		M: genBlks([]BlockSpec{genBlkSpec(2, 60, 38), genBlkSpec(2, 61, 39)}),
		Q: genBlks([]BlockSpec{genBlkSpec(4, 40, 18), genBlkSpec(2, 41, 19)}),
		H: genBlks([]BlockSpec{genBlkSpec(4, 40, 14), genBlkSpec(2, 41, 15)}),
	},
	9: {
		L: genBlks([]BlockSpec{genBlkSpec(2, 146, 116)}),
		M: genBlks([]BlockSpec{genBlkSpec(3, 58, 36), genBlkSpec(2, 59, 37)}),
		Q: genBlks([]BlockSpec{genBlkSpec(4, 36, 16), genBlkSpec(4, 37, 17)}),
		H: genBlks([]BlockSpec{genBlkSpec(4, 36, 12), genBlkSpec(4, 37, 13)}),
	},
	10: {
		L: genBlks([]BlockSpec{genBlkSpec(2, 86, 68), genBlkSpec(2, 87, 69)}),
		M: genBlks([]BlockSpec{genBlkSpec(4, 69, 43), genBlkSpec(1, 70, 44)}),
		Q: genBlks([]BlockSpec{genBlkSpec(6, 43, 19), genBlkSpec(2, 44, 20)}),
		H: genBlks([]BlockSpec{genBlkSpec(6, 43, 15), genBlkSpec(2, 44, 16)}),
	},
	11: {
		L: genBlks([]BlockSpec{genBlkSpec(4, 101, 81)}),
		M: genBlks([]BlockSpec{genBlkSpec(1, 80, 50), genBlkSpec(4, 81, 51)}),
		Q: genBlks([]BlockSpec{genBlkSpec(4, 50, 22), genBlkSpec(4, 51, 23)}),
		H: genBlks([]BlockSpec{genBlkSpec(3, 36, 12), genBlkSpec(8, 37, 13)}),
	},
	12: {
		L: genBlks([]BlockSpec{genBlkSpec(2, 116, 92), genBlkSpec(2, 117, 93)}),
		M: genBlks([]BlockSpec{genBlkSpec(6, 58, 36), genBlkSpec(2, 59, 37)}),
		Q: genBlks([]BlockSpec{genBlkSpec(4, 46, 20), genBlkSpec(6, 47, 21)}),
		H: genBlks([]BlockSpec{genBlkSpec(7, 42, 14), genBlkSpec(4, 43, 15)}),
	},
	13: {
		L: genBlks([]BlockSpec{genBlkSpec(4, 133, 107)}),
		M: genBlks([]BlockSpec{genBlkSpec(8, 59, 37), genBlkSpec(1, 60, 38)}),
		Q: genBlks([]BlockSpec{genBlkSpec(8, 44, 20), genBlkSpec(4, 45, 21)}),
		H: genBlks([]BlockSpec{genBlkSpec(12, 33, 11), genBlkSpec(4, 34, 12)}),
	},
	14: {
		L: genBlks([]BlockSpec{genBlkSpec(3, 145, 115), genBlkSpec(1, 146, 116)}),
		M: genBlks([]BlockSpec{genBlkSpec(4, 64, 40), genBlkSpec(5, 65, 41)}),
		Q: genBlks([]BlockSpec{genBlkSpec(11, 36, 16), genBlkSpec(5, 37, 17)}),
		H: genBlks([]BlockSpec{genBlkSpec(11, 36, 12), genBlkSpec(5, 37, 13)}),
	},
	15: {
		L: genBlks([]BlockSpec{genBlkSpec(5, 109, 87), genBlkSpec(1, 110, 88)}),
		M: genBlks([]BlockSpec{genBlkSpec(5, 65, 41), genBlkSpec(5, 66, 42)}),
		Q: genBlks([]BlockSpec{genBlkSpec(5, 54, 24), genBlkSpec(7, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(11, 36, 12), genBlkSpec(7, 37, 13)}),
	},
	16: {
		L: genBlks([]BlockSpec{genBlkSpec(5, 122, 98), genBlkSpec(1, 123, 99)}),
		M: genBlks([]BlockSpec{genBlkSpec(7, 73, 45), genBlkSpec(3, 74, 46)}),
		Q: genBlks([]BlockSpec{genBlkSpec(15, 43, 19), genBlkSpec(2, 44, 20)}),
		H: genBlks([]BlockSpec{genBlkSpec(3, 45, 15), genBlkSpec(13, 46, 16)}),
	},
	17: {
		L: genBlks([]BlockSpec{genBlkSpec(1, 135, 107), genBlkSpec(5, 136, 108)}),
		M: genBlks([]BlockSpec{genBlkSpec(10, 74, 46), genBlkSpec(1, 75, 47)}),
		Q: genBlks([]BlockSpec{genBlkSpec(1, 50, 22), genBlkSpec(15, 51, 23)}),
		H: genBlks([]BlockSpec{genBlkSpec(2, 42, 14), genBlkSpec(17, 43, 15)}),
	},
	18: {
		L: genBlks([]BlockSpec{genBlkSpec(5, 150, 120), genBlkSpec(1, 151, 121)}),
		M: genBlks([]BlockSpec{genBlkSpec(9, 69, 43), genBlkSpec(4, 70, 44)}),
		Q: genBlks([]BlockSpec{genBlkSpec(17, 50, 22), genBlkSpec(1, 51, 23)}),
		H: genBlks([]BlockSpec{genBlkSpec(2, 42, 14), genBlkSpec(19, 43, 15)}),
	},
	19: {
		L: genBlks([]BlockSpec{genBlkSpec(3, 141, 113), genBlkSpec(4, 142, 114)}),
		M: genBlks([]BlockSpec{genBlkSpec(3, 70, 44), genBlkSpec(11, 71, 45)}),
		Q: genBlks([]BlockSpec{genBlkSpec(17, 47, 21), genBlkSpec(4, 48, 22)}),
		H: genBlks([]BlockSpec{genBlkSpec(9, 39, 13), genBlkSpec(16, 40, 14)}),
	},
	20: {
		L: genBlks([]BlockSpec{genBlkSpec(3, 135, 107), genBlkSpec(5, 136, 108)}),
		M: genBlks([]BlockSpec{genBlkSpec(3, 67, 41), genBlkSpec(13, 68, 42)}),
		Q: genBlks([]BlockSpec{genBlkSpec(15, 54, 24), genBlkSpec(5, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(15, 43, 15), genBlkSpec(10, 44, 16)}),
	},
	21: {
		L: genBlks([]BlockSpec{genBlkSpec(4, 144, 116), genBlkSpec(4, 145, 117)}),
		M: genBlks([]BlockSpec{genBlkSpec(17, 68, 42)}),
		Q: genBlks([]BlockSpec{genBlkSpec(17, 50, 22), genBlkSpec(6, 51, 23)}),
		H: genBlks([]BlockSpec{genBlkSpec(19, 46, 16), genBlkSpec(6, 47, 17)}),
	},
	22: {
		L: genBlks([]BlockSpec{genBlkSpec(2, 139, 111), genBlkSpec(7, 140, 112)}),
		M: genBlks([]BlockSpec{genBlkSpec(17, 74, 46)}),
		Q: genBlks([]BlockSpec{genBlkSpec(7, 54, 24), genBlkSpec(16, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(34, 37, 13)}),
	},
	23: {
		L: genBlks([]BlockSpec{genBlkSpec(4, 151, 121), genBlkSpec(5, 152, 122)}),
		M: genBlks([]BlockSpec{genBlkSpec(4, 75, 47), genBlkSpec(14, 76, 48)}),
		Q: genBlks([]BlockSpec{genBlkSpec(11, 54, 24), genBlkSpec(14, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(16, 45, 15), genBlkSpec(14, 46, 16)}),
	},
	24: {
		L: genBlks([]BlockSpec{genBlkSpec(6, 147, 117), genBlkSpec(4, 148, 118)}),
		M: genBlks([]BlockSpec{genBlkSpec(6, 73, 45), genBlkSpec(14, 74, 46)}),
		Q: genBlks([]BlockSpec{genBlkSpec(11, 54, 24), genBlkSpec(16, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(30, 46, 16), genBlkSpec(2, 47, 17)}),
	},
	25: {
		L: genBlks([]BlockSpec{genBlkSpec(8, 132, 106), genBlkSpec(4, 133, 107)}),
		M: genBlks([]BlockSpec{genBlkSpec(8, 75, 47), genBlkSpec(13, 76, 48)}),
		Q: genBlks([]BlockSpec{genBlkSpec(7, 54, 24), genBlkSpec(22, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(22, 45, 15), genBlkSpec(13, 46, 16)}),
	},
	26: {
		L: genBlks([]BlockSpec{genBlkSpec(10, 142, 114), genBlkSpec(2, 143, 115)}),
		M: genBlks([]BlockSpec{genBlkSpec(19, 74, 46), genBlkSpec(4, 75, 47)}),
		Q: genBlks([]BlockSpec{genBlkSpec(28, 50, 22), genBlkSpec(6, 51, 23)}),
		H: genBlks([]BlockSpec{genBlkSpec(33, 46, 16), genBlkSpec(4, 47, 17)}),
	},
	27: {
		L: genBlks([]BlockSpec{genBlkSpec(8, 152, 122), genBlkSpec(4, 153, 123)}),
		M: genBlks([]BlockSpec{genBlkSpec(22, 73, 45), genBlkSpec(3, 74, 46)}),
		Q: genBlks([]BlockSpec{genBlkSpec(8, 53, 23), genBlkSpec(26, 54, 24)}),
		H: genBlks([]BlockSpec{genBlkSpec(12, 45, 15), genBlkSpec(28, 46, 16)}),
	},
	28: {
		L: genBlks([]BlockSpec{genBlkSpec(3, 147, 117), genBlkSpec(10, 148, 118)}),
		M: genBlks([]BlockSpec{genBlkSpec(3, 73, 45), genBlkSpec(23, 74, 46)}),
		Q: genBlks([]BlockSpec{genBlkSpec(4, 54, 24), genBlkSpec(31, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(11, 45, 15), genBlkSpec(31, 46, 16)}),
	},
	29: {
		L: genBlks([]BlockSpec{genBlkSpec(7, 146, 116), genBlkSpec(7, 147, 117)}),
		M: genBlks([]BlockSpec{genBlkSpec(21, 73, 45), genBlkSpec(7, 74, 46)}),
		Q: genBlks([]BlockSpec{genBlkSpec(1, 53, 23), genBlkSpec(37, 54, 24)}),
		H: genBlks([]BlockSpec{genBlkSpec(19, 45, 15), genBlkSpec(26, 46, 16)}),
	},
	30: {
		L: genBlks([]BlockSpec{genBlkSpec(5, 145, 115), genBlkSpec(10, 146, 116)}),
		M: genBlks([]BlockSpec{genBlkSpec(19, 75, 47), genBlkSpec(10, 76, 48)}),
		Q: genBlks([]BlockSpec{genBlkSpec(15, 54, 24), genBlkSpec(25, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(23, 45, 15), genBlkSpec(25, 46, 16)}),
	},
	31: {
		L: genBlks([]BlockSpec{genBlkSpec(13, 145, 115), genBlkSpec(3, 146, 116)}),
		M: genBlks([]BlockSpec{genBlkSpec(2, 74, 46), genBlkSpec(29, 75, 47)}),
		Q: genBlks([]BlockSpec{genBlkSpec(42, 54, 24), genBlkSpec(1, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(23, 45, 15), genBlkSpec(28, 46, 16)}),
	},
	32: {
		L: genBlks([]BlockSpec{genBlkSpec(17, 145, 115)}),
		M: genBlks([]BlockSpec{genBlkSpec(10, 74, 46), genBlkSpec(23, 75, 47)}),
		Q: genBlks([]BlockSpec{genBlkSpec(10, 54, 24), genBlkSpec(35, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(19, 45, 15), genBlkSpec(35, 46, 16)}),
	},
	33: {
		L: genBlks([]BlockSpec{genBlkSpec(17, 145, 115), genBlkSpec(1, 146, 116)}),
		M: genBlks([]BlockSpec{genBlkSpec(14, 74, 46), genBlkSpec(21, 75, 47)}),
		Q: genBlks([]BlockSpec{genBlkSpec(29, 54, 24), genBlkSpec(19, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(11, 45, 15), genBlkSpec(46, 46, 16)}),
	},
	34: {
		L: genBlks([]BlockSpec{genBlkSpec(13, 145, 115), genBlkSpec(6, 146, 116)}),
		M: genBlks([]BlockSpec{genBlkSpec(14, 74, 46), genBlkSpec(23, 75, 47)}),
		Q: genBlks([]BlockSpec{genBlkSpec(44, 54, 24), genBlkSpec(7, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(59, 46, 16), genBlkSpec(1, 47, 17)}),
	},
	35: {
		L: genBlks([]BlockSpec{genBlkSpec(12, 151, 121), genBlkSpec(7, 152, 122)}),
		M: genBlks([]BlockSpec{genBlkSpec(12, 75, 47), genBlkSpec(26, 76, 48)}),
		Q: genBlks([]BlockSpec{genBlkSpec(39, 54, 24), genBlkSpec(14, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(22, 45, 15), genBlkSpec(41, 46, 16)}),
	},
	36: {
		L: genBlks([]BlockSpec{genBlkSpec(6, 151, 121), genBlkSpec(14, 152, 122)}),
		M: genBlks([]BlockSpec{genBlkSpec(6, 75, 47), genBlkSpec(34, 76, 48)}),
		Q: genBlks([]BlockSpec{genBlkSpec(46, 54, 24), genBlkSpec(10, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(2, 45, 15), genBlkSpec(64, 46, 16)}),
	},
	37: {
		L: genBlks([]BlockSpec{genBlkSpec(17, 152, 122), genBlkSpec(4, 153, 123)}),
		M: genBlks([]BlockSpec{genBlkSpec(29, 74, 46), genBlkSpec(14, 75, 47)}),
		Q: genBlks([]BlockSpec{genBlkSpec(49, 54, 24), genBlkSpec(10, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(24, 45, 15), genBlkSpec(46, 46, 16)}),
	},
	38: {
		L: genBlks([]BlockSpec{genBlkSpec(4, 152, 122), genBlkSpec(18, 153, 123)}),
		M: genBlks([]BlockSpec{genBlkSpec(13, 74, 46), genBlkSpec(32, 75, 47)}),
		Q: genBlks([]BlockSpec{genBlkSpec(48, 54, 24), genBlkSpec(14, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(42, 45, 15), genBlkSpec(32, 46, 16)}),
	},
	39: {
		L: genBlks([]BlockSpec{genBlkSpec(20, 147, 117), genBlkSpec(4, 148, 118)}),
		M: genBlks([]BlockSpec{genBlkSpec(40, 75, 47), genBlkSpec(7, 76, 48)}),
		Q: genBlks([]BlockSpec{genBlkSpec(43, 54, 24), genBlkSpec(22, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(10, 45, 15), genBlkSpec(67, 46, 16)}),
	},
	40: {
		L: genBlks([]BlockSpec{genBlkSpec(19, 148, 118), genBlkSpec(6, 149, 119)}),
		M: genBlks([]BlockSpec{genBlkSpec(18, 75, 47), genBlkSpec(31, 76, 48)}),
		Q: genBlks([]BlockSpec{genBlkSpec(34, 54, 24), genBlkSpec(34, 55, 25)}),
		H: genBlks([]BlockSpec{genBlkSpec(20, 45, 15), genBlkSpec(61, 46, 16)}),
	},
}
