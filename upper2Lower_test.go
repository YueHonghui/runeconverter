package runeconverter

import (
	"bytes"
	"testing"
)

var mc = Multi{
	Converters: []Converter{
		NewBC(BCModeDBC2SBC),
		new(Traditional2Simple),
		new(NormalizeNum),
		new(Num2Zh),
		NewUpper2Lower(Lower),
	},
}

func TestTrad2Simp(t *testing.T) {
	tmpbuf := bytes.NewBuffer(make([]byte, 0))
	Convert([]byte("Hello World,優酷土豆，ａｂｃ"), tmpbuf, func(r rune, bf *bytes.Buffer) {
		bf.WriteRune(mc.Convert(r))
	})
	if "hello world,优酷土豆,abc" != tmpbuf.String() {
		t.Error("fail,result=", tmpbuf.String())
	}
}
