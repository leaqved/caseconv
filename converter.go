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
	Delimeter  rune
	Formatters []Format
}

func New() *Converter {
	return &Converter{
		Splitter: DefaultSplitter(),
	}
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
