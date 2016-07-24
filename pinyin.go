package runeconverter

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 拼音风格
const (
	PinyinStyleNormal = iota
	PinyinStyleTone
)

type pinyinTable []pinyinItem

type pinyinItem struct {
	Code    uint32
	Cnt     int
	Pinyins [3][]byte
}

func (table pinyinTable) Less(i, j int) bool {
	return table[i].Code < table[j].Code
}

func (table pinyinTable) Len() int {
	return len(table)
}

func (table pinyinTable) Swap(i, j int) {
	table[i], table[j] = table[j], table[i]
}

type Pinyin struct {
	table [1024][]pinyinItem
	style int
}

func NewPinyin(dbpath string, style int) (*Pinyin, error) {
	f, err := os.OpenFile(dbpath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var code uint64
	var tableidx uint32
	var ret Pinyin
	ret.style = style
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		itms := strings.Split(scanner.Text(), ",")
		if len(itms) < 2 {
			return nil, errors.New("Invalid table file")
		}
		code, err = strconv.ParseUint(itms[0], 0, 32)
		if err != nil {
			return nil, err
		}
		pinyinitm := pinyinItem{
			Code: uint32(code),
			Cnt:  len(itms) - 1,
		}
		if pinyinitm.Cnt > 3 {
			pinyinitm.Cnt = 3
		}
		for i := 0; i < pinyinitm.Cnt; i++ {
			pinyinitm.Pinyins[i] = []byte(itms[i+1])
		}
		tableidx = pinyinitm.Code & (uint32(len(ret.table)) - 1)
		ret.table[tableidx] = append(ret.table[tableidx], pinyinitm)
	}
	for i := 0; i < len(ret.table); i++ {
		sort.Sort(pinyinTable(ret.table[i]))
	}
	return &ret, nil
}

func (py *Pinyin) Convert(r rune, buf *bytes.Buffer) {
	tableidx := int(r) & (len(py.table) - 1)
	tableentry := py.table[tableidx]
	idx := sort.Search(len(tableentry), func(i int) bool {
		return tableentry[i].Code >= uint32(r)
	})
	if idx < len(tableentry) && tableentry[idx].Code == uint32(r) {
		if py.style == PinyinStyleNormal {
			vlen := len(tableentry[idx].Pinyins[0])
			buf.Write(tableentry[idx].Pinyins[0][:vlen-1])
		} else {
			buf.Write(tableentry[idx].Pinyins[0])
		}
	} else {
		buf.WriteRune(r)
	}
}
