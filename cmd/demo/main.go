package main

import (
	"fmt"
	"os"

	"github.com/adamcavendish/bazel-cgo-demo/convert"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: demo file.opus file.wav. accepts only 48KHz sample rate, mono channel opus files.")
		return
	}
	opusFileName, pcmFileName := os.Args[1], os.Args[2]

	opusF, err := os.Open(opusFileName)
	if err != nil {
		fmt.Printf("Open(%s) error: %v\n", opusFileName, err)
		return
	}
	defer opusF.Close()
	pcmF, err := os.Create(pcmFileName)
	if err != nil {
		fmt.Printf("Create(%s) error: %v\n", pcmFileName, err)
		return
	}
	defer pcmF.Close()
	if err := convert.Opus2Pcm(opusF, pcmF); err != nil {
		fmt.Printf("Opus2Pcm() error: %v\n", err)
		return
	}
}
