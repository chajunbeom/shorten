package utils

type base64unit struct {
	b       []byte
	padding int
}

const Base64MappingTable = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func (o *base64unit) base64Int32() int32 {
	for i := len(o.b); i < 3; i++ {
		o.b = append(o.b, 0)
		o.padding++
	}
	var ret int32
	ret = ret | int32(o.b[0])<<16
	ret = ret | int32(o.b[1])<<8
	ret = ret | int32(o.b[2])
	return ret
}

func (o *base64unit) base64() []byte {
	data := o.base64Int32()
	table := []byte(Base64MappingTable)
	ret := make([]byte, 4-o.padding)
	for i := 0; data > 0; i++ {
		if o.padding > 0 && byte(data)&0x3f == 0 {
			o.padding--
		} else {
			ret[3-o.padding-i] = table[byte(data)&0x3f]
		}
		data = data >> 6
	}
	return ret
}

func Base64Encode(src []byte) []byte {
	pad := (6 - ((len(src) * 8) % 6)) / 2
	size := (len(src) * 8) / 6
	if pad > 0 && pad < 3 {
		size++
	}
	dst := make([]byte, size)
	i := 0
	j := 0
	for i <= len(src) {
		temp := base64unit{b: src[i:]}
		copy(dst[j:], temp.base64())
		i += 3
		j += 4
	}
	return dst
}
