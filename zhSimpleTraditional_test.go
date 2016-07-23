package runeconverter

import (
	"bytes"
	"testing"
)

var ttrad2simp = []byte(`東亞各地受漢文化影響而形成漢字文化圈`)
var ttrad2simpR = `东亚各地受汉文化影响而形成汉字文化圈`
var zstbuf = bytes.NewBuffer(make([]byte, 1024))

func TestTrad2Simp(t *testing.T) {
	c := new(Traditional2Simple)
	zstbuf.Reset()
	Convert(ttrad2simp, zstbuf, func(r rune, bf *bytes.Buffer) {
		bf.WriteRune(c.Convert(r))
	})
	if zstbuf.String() != ttrad2simpR {
		t.Errorf("expected: %s, actual: %s\n", ttrad2simpR, zstbuf.String())
	}
}
