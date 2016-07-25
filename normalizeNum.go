package runeconverter

type NormalizeNum int

func (tn *NormalizeNum) Convert(r rune) rune {
	//①(9312)~⑨(9320) ----> 1(49)~9(57)
	if r >= 9312 && r <= 9320 {
		return r - (9312 - 49)
	} else if r >= 9352 && r <= 9360 { //⒈(9352)~⒐(9360) ---->1(49)~9(57)
		return r - (9352 - 49)
	} else {
		return r
	}
}
