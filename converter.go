package caseconv

import (
	"unicode"

	s "github.com/leaqved/splitter"
)

var DefaultSeparators = map[rune]struct{}{
	' ': {},
	'-': {},
	'_': {},
	'.': {},
	',': {},
	':': {},
	';': {},
}

type Converter struct {
	Splitter   *s.Splitter
	Delimiter  rune
	Formatters []Format
}

func New() *Converter {
	return &Converter{
		Splitter: DefaultSplitter(),
	}
}

func (c *Converter) Convert(str string) string {
	var res string
	words := c.Splitter.Split(str)
	for index := range words {
		for _, format := range c.Formatters {
			words[index] = format(index, words)
		}
		res += words[index]
		if index != len(words)-1 {
			res += string(c.Delimiter)
		}
	}
	return res
}

func DefaultSplitter() *s.Splitter {
	return s.New(
		s.AddSkip(
			s.Before(
				s.CharInSet(DefaultSeparators))),
		s.AddSplit(
			s.After(unicode.IsLetter).And(
				s.After(unicode.IsLower).And(
					s.Before(unicode.IsLetter).And(
						s.Before(unicode.IsUpper)))),
		),
		s.AddJoin(
			s.Everything()),
	)
}
