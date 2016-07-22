package runeconverter

const (
	NormalizeZhLow2Big = iota
	NormalizeZhBig2Low
)

type NormalizeZhNum struct {
	mode   int
	mapuse map[rune]rune
}

var lower2bigger = map[rune]rune{
	'〇': '零',
	'一': '壹',
	'二': '贰',
	'三': '叁',
	'四': '肆',
	'五': '伍',
	'六': '陆',
	'七': '柒',
	'八': '捌',
	'九': '玖',
}

var bigger2lower = map[rune]rune{
	'零': '〇',
	'壹': '一',
	'贰': '二',
	'叁': '三',
	'肆': '四',
	'伍': '五',
	'陆': '六',
	'柒': '七',
	'捌': '八',
	'玖': '九',
}

func NewNormalizeZhNum(mode int) *NormalizeZhNum {
	if mode == NormalizeZhBig2Low {
		return &NormalizeZhNum{
			mode:   mode,
			mapuse: bigger2lower,
		}
	} else {
		return &NormalizeZhNum{
			mode:   mode,
			mapuse: lower2bigger,
		}
	}
}

func (nz *NormalizeZhNum) Convert(r rune) rune {
	if ret, ok := nz.mapuse[r]; ok {
		return ret
	}
	return r
}
