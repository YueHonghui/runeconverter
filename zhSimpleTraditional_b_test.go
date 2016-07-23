package runeconverter

import (
	"bytes"
	"testing"
)

var testStr = []byte(`在歷史、文化方面，東亞各地受漢文化影響而形成漢字文化圈。在20世紀之前，以傳統漢字書寫的文言文，在中國、朝鮮半島、琉球、日本、越南等地被使用在幾乎所有正式的文書上。`)
var buf = bytes.NewBuffer(make([]byte, 1024))

func BenchmarkTrad2Simp(b *testing.B) {
	c := new(Traditional2Simple)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		Convert(testStr, buf, func(r rune, bf *bytes.Buffer) {
			bf.WriteRune(c.Convert(r))
		})
	}
}
