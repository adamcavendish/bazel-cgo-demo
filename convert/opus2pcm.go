package convert

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/hraban/opus.v2"

	"github.com/adamcavendish/bazel-cgo-demo/buffer"
	"github.com/adamcavendish/bazel-cgo-demo/header"
)

const channels = 1

func Opus2Pcm(opusF *os.File, pcmF *os.File) error {
	if _, err := pcmF.Seek(44, io.SeekStart); err != nil {
		return fmt.Errorf("Seek() error: %v\n", err)
	}
	opusStream, err := opus.NewStream(opusF)
	if err != nil {
		return fmt.Errorf("opus.NewStream() error: %v\n", err)
	}
	defer opusStream.Close()
	pcmbuf := make([]float32, 16384)

	fileSize := 0
loop:
	for {
		n, err := opusStream.ReadFloat32(pcmbuf)
		switch err {
		case nil:
		case io.EOF:
			break loop
		default:
			return fmt.Errorf("stream Read() error: %v\n", err)
		}
		pcm := pcmbuf[:n*channels]
		nw, err := pcmF.Write(buffer.Fs2bs(pcm))
		if err != nil {
			return fmt.Errorf("Write() error: %v\n", err)
		}
		fileSize += nw
	}

	if _, err := pcmF.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("Seek() error: %v\n", err)
	}
	if _, err := pcmF.Write(header.WavHeader(int32(fileSize))); err != nil {
		return fmt.Errorf("Write() error: %v\n", err)
	}
	return nil
}
