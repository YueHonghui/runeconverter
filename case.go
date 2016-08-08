package runeconverter

import (
	"unicode"
)

type CaseLower int

type CaseUpper int

func (cl *CaseLower) Convert(r rune) rune {
	return unicode.ToLower(r)
}

func (cl *CaseUpper) Convert(r rune) rune {
	return unicode.ToUpper(r)
}
