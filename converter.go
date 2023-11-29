package caseconv

import "github.com/leaqved/splitter"

type Converter struct {
	Splitter   *splitter.Splitter
	Delimeter  rune
	Formatters []Format
}
