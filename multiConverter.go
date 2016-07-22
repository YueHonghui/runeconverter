package runeconverter

type Multi struct {
	Converters []Converter
}

func (m *Multi) Add(c Converter) *Multi {
	m.Converters = append(m.Converters, c)
	return m
}

func (m *Multi) Convert(r rune) rune {
	for _, c := range m.Converters {
		r = c.Convert(r)
	}
	return r
}
