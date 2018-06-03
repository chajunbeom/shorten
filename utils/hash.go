package utils

const PolyCRC32 = 0x04C11DB7

// bits reverse
func Reflect(in uint64, count uint) uint64 {
	ret := in
	for idx := uint(0); idx < count; idx++ {
		srcbit := uint64(1) << idx
		dstbit := uint64(1) << (count - idx - 1)
		if (in & srcbit) != 0 {
			ret |= dstbit
		} else {
			ret = ret & (^dstbit)
		}
	}
	return ret
}

func CRC32(src []byte) uint32 {
	topbit := uint64(1) << 31
	crc := uint64(topbit<<1 - 1)
	for i := 0; i < len(src); i++ {
		crc ^= (Reflect(uint64(src[i])&0x00FF, 8) << 24)
		for j := 0; j < 8; j++ {
			if (crc & topbit) != 0 {
				crc = (crc << 1) ^ uint64(PolyCRC32)
			} else {
				crc = crc << 1
			}
		}
	}
	crc = Reflect(crc, 32)
	return uint32(crc ^ (topbit<<1 - 1))
}
