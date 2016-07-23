package runeconverter

func dbc2Sbc(r rune) rune {
	if r == 12288 {
		return 32
	} else if r >= 65281 && r <= 65374 {
		return r - (65281 - 33)
	} else {
		return r
	}
}

func sbc2Dbc(r rune) rune {
	if r == 32 {
		return 12288
	} else if r >= 33 && r <= 126 {
		return r + (65281 - 33)
	} else {
		return r
	}
}

type BC struct {
	bcfunc func(rune) rune
}

const (
	BCModeDBC2SBC = iota
	BCModeSBC2DBC
)

func NewBC(mode int) *BC {
	if mode == BCModeDBC2SBC {
		return &BC{
			bcfunc: dbc2Sbc,
		}
	} else {
		return &BC{
			bcfunc: sbc2Dbc,
		}
	}
}

func (bc *BC) Convert(r rune) rune {
	return bc.bcfunc(r)
}
