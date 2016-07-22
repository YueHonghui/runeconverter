package runeconverter

import (
	"bytes"
)

func Convert(input []byte, output *bytes.Buffer, cvt Converter) *bytes.Buffer {
	b := bytes.NewBuffer(input)
	var r rune
	var err error
	var n int
	for r, _, err = b.ReadRune(); err == nil; r, _, err = b.ReadRune() {
		output.WriteRune(cvt.Convert(r))
	}
	return output
}
