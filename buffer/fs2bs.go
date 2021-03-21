package buffer

func Fs2bs(data []float32) []byte {
	var buf []byte
	for _, vf := range data {
		v := int16(vf * 32768)
		if v > 32767 {
			v = 32767
		}
		if v < -32768 {
			v = -32768
		}
		lb, hb := byte(v&0xFF), byte(v>>8)
		buf = append(buf, lb, hb)
	}
	return buf
}
