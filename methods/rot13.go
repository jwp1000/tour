package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotty rot13Reader) Read(arr []byte) (int, error) {
	n, err := rotty.r.Read(arr)
	const input = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const output = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

	for i := range arr {
		inputIndex := strings.Index(input, string(arr[i]))
		if inputIndex < 0 {
			continue
		}
		arr[i] = output[inputIndex]
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
