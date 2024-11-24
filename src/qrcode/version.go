package qrcode

import "fmt"

type Version int

var versionCapacity = map[Version]map[ErrorCorrectionLevel]int{
	1:  {L: 152, M: 128, Q: 104, H: 72},
	2:  {L: 272, M: 224, Q: 176, H: 128},
	3:  {L: 440, M: 352, Q: 272, H: 208},
	4:  {L: 640, M: 512, Q: 384, H: 288},
	5:  {L: 864, M: 688, Q: 496, H: 368},
	6:  {L: 1088, M: 864, Q: 608, H: 480},
	7:  {L: 1248, M: 992, Q: 704, H: 528},
	8:  {L: 1552, M: 1232, Q: 880, H: 688},
	9:  {L: 1856, M: 1456, Q: 1056, H: 800},
	10: {L: 2192, M: 1728, Q: 1232, H: 976},
	11: {L: 2592, M: 2032, Q: 1440, H: 1120},
	12: {L: 2960, M: 2320, Q: 1648, H: 1264},
	13: {L: 3424, M: 2672, Q: 1952, H: 1440},
	14: {L: 3688, M: 2920, Q: 2088, H: 1576},
	15: {L: 4184, M: 3320, Q: 2360, H: 1784},
	16: {L: 4712, M: 3624, Q: 2600, H: 2024},
	17: {L: 5176, M: 4056, Q: 2936, H: 2264},
	18: {L: 5768, M: 4504, Q: 3176, H: 2504},
	19: {L: 6360, M: 5016, Q: 3560, H: 2728},
	20: {L: 6888, M: 5352, Q: 3880, H: 3080},
	21: {L: 7456, M: 5712, Q: 4096, H: 3248},
	22: {L: 8048, M: 6256, Q: 4544, H: 3536},
	23: {L: 8752, M: 6880, Q: 4912, H: 3712},
	24: {L: 9392, M: 7312, Q: 5312, H: 4112},
	25: {L: 10208, M: 8000, Q: 5744, H: 4304},
	26: {L: 10960, M: 8496, Q: 6032, H: 4768},
	27: {L: 11744, M: 9024, Q: 6464, H: 5024},
	28: {L: 12248, M: 9544, Q: 6968, H: 5288},
	29: {L: 13048, M: 10136, Q: 7288, H: 5608},
	30: {L: 13880, M: 10984, Q: 7880, H: 5960},
	31: {L: 14744, M: 11640, Q: 8264, H: 6344},
	32: {L: 15640, M: 12328, Q: 8920, H: 6760},
	33: {L: 16568, M: 13048, Q: 9368, H: 7208},
	34: {L: 17528, M: 13800, Q: 9848, H: 7688},
	35: {L: 18448, M: 14496, Q: 10288, H: 7888},
	36: {L: 19472, M: 15312, Q: 10832, H: 8432},
	37: {L: 20528, M: 15936, Q: 11408, H: 8768},
	38: {L: 21616, M: 16816, Q: 12016, H: 9136},
	39: {L: 22496, M: 17728, Q: 12656, H: 9776},
	40: {L: 23648, M: 18672, Q: 13328, H: 10208},
}

var characterCountIndicator = map[Version]map[EncodeMode]int{
	1:  {BinaryMode: 8, NumericMode: 10},
	2:  {BinaryMode: 8, NumericMode: 10},
	3:  {BinaryMode: 8, NumericMode: 10},
	4:  {BinaryMode: 8, NumericMode: 10},
	5:  {BinaryMode: 8, NumericMode: 10},
	6:  {BinaryMode: 8, NumericMode: 10},
	7:  {BinaryMode: 8, NumericMode: 10},
	8:  {BinaryMode: 8, NumericMode: 10},
	9:  {BinaryMode: 8, NumericMode: 10},
	10: {BinaryMode: 16, NumericMode: 12},
	11: {BinaryMode: 16, NumericMode: 12},
	12: {BinaryMode: 16, NumericMode: 12},
	13: {BinaryMode: 16, NumericMode: 12},
	14: {BinaryMode: 16, NumericMode: 12},
	15: {BinaryMode: 16, NumericMode: 12},
	16: {BinaryMode: 16, NumericMode: 12},
	17: {BinaryMode: 16, NumericMode: 12},
	18: {BinaryMode: 16, NumericMode: 12},
	19: {BinaryMode: 16, NumericMode: 12},
	20: {BinaryMode: 16, NumericMode: 12},
	21: {BinaryMode: 16, NumericMode: 12},
	22: {BinaryMode: 16, NumericMode: 12},
	23: {BinaryMode: 16, NumericMode: 12},
	24: {BinaryMode: 16, NumericMode: 12},
	25: {BinaryMode: 16, NumericMode: 12},
	26: {BinaryMode: 16, NumericMode: 12},
	27: {BinaryMode: 16, NumericMode: 14},
	28: {BinaryMode: 16, NumericMode: 14},
	29: {BinaryMode: 16, NumericMode: 14},
	30: {BinaryMode: 16, NumericMode: 14},
	31: {BinaryMode: 16, NumericMode: 14},
	32: {BinaryMode: 16, NumericMode: 14},
	33: {BinaryMode: 16, NumericMode: 14},
	34: {BinaryMode: 16, NumericMode: 14},
	35: {BinaryMode: 16, NumericMode: 14},
	36: {BinaryMode: 16, NumericMode: 14},
	37: {BinaryMode: 16, NumericMode: 14},
	38: {BinaryMode: 16, NumericMode: 14},
	39: {BinaryMode: 16, NumericMode: 14},
	40: {BinaryMode: 16, NumericMode: 14},
}

func getVersion(ecl ErrorCorrectionLevel, src string) (Version, error) {
	mode := getEncodeMode(src)
	srcLength := len(src)

	minVersion, err := findMinimumVersionToFit(ecl, mode, srcLength)
	if err != nil {
		return 0, err
	}

	return minVersion, nil
}

func findMinimumVersionToFit(ecl ErrorCorrectionLevel, mode EncodeMode, srcLength int) (Version, error) {
	modeIndicator := 4
	dataBits, err := func(mode EncodeMode, srcLength int) (int, error) {
		switch mode {
		case BinaryMode:
			return 8 * srcLength, nil

		case NumericMode:
			remainder := func(srcLength int) int {
				switch srcLength % 3 {
				case 1:
					return 4
				case 2:
					return 7
				default:
					return 0
				}
			}(srcLength)
			return 10*(srcLength/3) + remainder, nil

		default:
			return 0, fmt.Errorf("unexpected mode: %s", mode)
		}
	}(mode, srcLength)
	if err != nil {
		return 0, err
	}

	var requiredBits int
	for i := 1; i <= 40; i++ {
		v := Version(i)
		requiredBits = modeIndicator + characterCountIndicator[v][mode] + dataBits
		if versionCapacity[v][ecl] >= requiredBits {
			return v, nil
		}
	}
	return 0, fmt.Errorf("data is too large with size: %d bits", requiredBits)
}
