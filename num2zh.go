package runeconverter

type Num2Zh int

var num2zh = [10]rune{'〇', '一', '二', '三', '四', '五', '六', '七', '八', '九'}

func (n2z *Num2Zh) Convert(r rune) rune {
	if r >= 48 && r <= 57 {
		return num2zh[r-48]
	}
	return r
}
