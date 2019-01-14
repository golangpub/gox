package types

func CountBitsSize(i int64) uint {
	var bitsSize uint = 0
	for bitsSize < 64 {
		i = i >> 1
		if i == 0 {
			break
		}
		bitsSize++
	}
	return bitsSize
}

func KeepRightBits(i int64, bitSize uint) int64 {
	return ((i >> bitSize) << bitSize) ^ i
}

func LeftMultiRight(n int64) int64 {
	bitsSize := CountBitsSize(n)
	halfSize := bitsSize / 2
	left := n >> halfSize
	right := KeepRightBits(n, halfSize)
	if right == 0 {
		right = 1
	}
	return left * right
}
