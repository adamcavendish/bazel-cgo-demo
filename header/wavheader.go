package header

import "fmt"

func WavHeader(fileSize int32) []byte {
	// the fileSize is the raw PCM bytes length
	fileSize1 := fileSize + 36
	var buf []byte
	buf = append(buf, "RIFF"...)                 // const "RIFF"
	buf = append(buf, byte((fileSize1>>0)&0xFF)) // file size 1 (file size + the following header (36 bytes))
	buf = append(buf, byte((fileSize1>>2)&0xFF)) // file size 1 (file size + the following header (36 bytes))
	buf = append(buf, byte((fileSize1>>4)&0xFF)) // file size 1 (file size + the following header (36 bytes))
	buf = append(buf, byte((fileSize1>>6)&0xFF)) // file size 1 (file size + the following header (36 bytes))
	buf = append(buf, "WAVE"...)                 // const "WAVE"
	buf = append(buf, "fmt "...)                 // const "fmt "
	buf = append(buf, 0x10, 0x00, 0x00, 0x00)    // const 16
	buf = append(buf, 0x01, 0x00)                // PCM
	buf = append(buf, 0x01, 0x00)                // mono channel
	buf = append(buf, 0x80, 0xbb, 0x00, 0x00)    // 48kHz
	buf = append(buf, 0x00, 0x77, 0x01, 0x00)    // (Sample Rate * BitsPerSample * Channels) / 8
	buf = append(buf, 0x02, 0x00)                // (Bits per sample * Channels) / 8
	buf = append(buf, 0x10, 0x00)                // Bits per sample
	buf = append(buf, "data"...)                 // const "data"
	buf = append(buf, byte((fileSize>>0)&0xFF))  // file size
	buf = append(buf, byte((fileSize>>2)&0xFF))  // file size
	buf = append(buf, byte((fileSize>>4)&0xFF))  // file size
	buf = append(buf, byte((fileSize>>6)&0xFF))  // file size

	if len(buf) != 44 {
		panic(fmt.Sprintf("unexpected buf length: %d", len(buf)))
	}

	return buf
}
