package utils

type IntegerInfo struct {
	ByteSize int
	BitMask  uint64
}

var (
	INT16 = &IntegerInfo{ByteSize: 2, BitMask: 0x00FF}
	INT32 = &IntegerInfo{ByteSize: 4, BitMask: 0x00FF}
	INT64 = &IntegerInfo{ByteSize: 8, BitMask: 0x00FF}
)

func ConvertIntToByte(info *IntegerInfo, src uint64) ([]byte, error) {
	dst := make([]byte, info.ByteSize)

	for i := 0; i < info.ByteSize; i++ {
		dst[info.ByteSize-1-i] = byte(src & info.BitMask)
		src = src >> 8
	}
	return dst, nil
}
