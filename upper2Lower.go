package runeconverter

const (
	Upper = iota
	Lower
)

const (
	MaxASCII = '\u007F' // maximum ASCII value.
)

type Upper2Lower struct {
	mode int
}

func NewUpper2Lower(mode int) *Upper2Lower {
	return &Upper2Lower{
		mode: mode,
	}
}

func toUpper(r rune) rune {
	if r <= MaxASCII {
		if 'a' <= r && r <= 'z' {
			r -= 'a' - 'A'
		}
	}
	return r
}

func toLower(r rune) rune {
	if r <= MaxASCII {
		if 'A' <= r && r <= 'Z' {
			r += 'a' - 'A'
		}
	}
	return r
}

func (nz *Upper2Lower) Convert(r rune) rune {
	if nz.mode == Lower {
		return toLower(r)
	} else {
		return toUpper(r)
	}
}
