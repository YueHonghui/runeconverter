package runeconverter

import (
	"bytes"
)

func Convert(input []byte, output *bytes.Buffer, runefunc func(rune, *bytes.Buffer)) *bytes.Buffer {
	b := bytes.NewBuffer(input)
	var r rune
	var err error
	for r, _, err = b.ReadRune(); err == nil; r, _, err = b.ReadRune() {
		runefunc(r, output)
	}
	return output
}
